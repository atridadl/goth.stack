package lib

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

// init function
func init() {
	godotenv.Load(".env")
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}

func CreateCheckoutSession(w http.ResponseWriter, r *http.Request, successUrl string, cancelUrl string, priceId string) {
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    stripe.String(priceId),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:         stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:   stripe.String(successUrl),
		CancelURL:    stripe.String(cancelUrl),
		AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(true)},
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}
