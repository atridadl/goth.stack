package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/uptrace/bunrouter"
)

type rateLimiter struct {
	visitors map[string]*visitor
	mu       sync.Mutex
	rps      int
}

type visitor struct {
	firstSeen time.Time
	requests  int
}

func NewRateLimiter(rps int) *rateLimiter {
	return &rateLimiter{
		visitors: make(map[string]*visitor),
		rps:      rps,
	}
}

func (rl *rateLimiter) RateLimit(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		rl.mu.Lock()
		defer rl.mu.Unlock()

		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			// handle error, e.g., return an HTTP 500 error
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return nil
		}

		v, exists := rl.visitors[ip]
		if !exists || time.Since(v.firstSeen) > 1*time.Minute {
			v = &visitor{
				firstSeen: time.Now(),
			}
			rl.visitors[ip] = v
		}

		v.requests++

		// Limit each IP to rps requests per minute
		if v.requests > rl.rps {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return nil
		}

		return next(w, req)
	}
}
