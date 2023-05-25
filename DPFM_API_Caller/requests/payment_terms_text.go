package requests

type PaymentTermsText struct {
	PaymentTerms     string  `json:"PaymentTerms"`
	Language         string  `json:"Language"`
	PaymentTermsName *string `json:"PaymentTermsName"`
}
