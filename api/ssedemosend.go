package api

import (
	"encoding/json"
	"net/http"

	"github.com/uptrace/bunrouter"
	"goth.stack/lib"
)

func SSEDemoSend(w http.ResponseWriter, req bunrouter.Request) error {
	// Get query parameters
	queryParams := req.URL.Query()

	// Get channel from query parameters
	channel := queryParams.Get("channel")
	if channel == "" {
		channel = "default"
	}

	// Get message from query parameters, form value, or request body
	message := queryParams.Get("message")
	if message == "" {
		message = req.PostFormValue("message")
		if message == "" {
			var body map[string]string
			err := json.NewDecoder(req.Body).Decode(&body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return err
			}
			message = body["message"]
		}
	}

	if message == "" {
		errMsg := map[string]string{"error": "message parameter is required"}
		errMsgBytes, _ := json.Marshal(errMsg)
		http.Error(w, string(errMsgBytes), http.StatusBadRequest)
		return nil
	}

	// Send message
	lib.SendSSE("default", message)

	statusMsg := map[string]string{"status": "message sent"}
	statusMsgBytes, _ := json.Marshal(statusMsg)
	w.Write(statusMsgBytes)

	return nil
}
