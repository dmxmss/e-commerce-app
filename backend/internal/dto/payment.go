package dto

type CreatePaymentRequest struct {
	ProductIds []int `json:"product_ids"`
	Currency string `json:"currency"`
}

type CreatePaymentResponse struct {
	ClientSecret string `json:"client_secret"`
	PaymentId string `json:"payment_id"`
}

type GetPaymentRequest struct {
	ID string `json:"id"`
}

type GetPaymentResponse struct {
	ClientSecret string `json:"client_secret"`
	PaymentId string `json:"payment_id"`
}
