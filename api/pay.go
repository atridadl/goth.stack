package api

import (
	"net/http"

	"atri.dad/lib"
	"github.com/labstack/echo/v4"
)

type PayRequest struct {
	SuccessUrl string `json:"successUrl"`
	CancelUrl  string `json:"cancelUrl"`
	PriceId    string `json:"priceId"`
}

func Pay(c echo.Context) error {
	payReq := new(PayRequest)

	if err := c.Bind(payReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	lib.CreateCheckoutSession(c.Response().Writer, c.Request(), payReq.SuccessUrl, payReq.CancelUrl, payReq.PriceId)

	return c.String(http.StatusOK, "Checkout session created")
}
