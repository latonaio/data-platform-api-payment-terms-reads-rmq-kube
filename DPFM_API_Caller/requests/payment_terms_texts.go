package requests

type PaymentTermsTexts struct {
	PaymentTerms     string  `json:"PaymentTerms"`
	Language         string  `json:"Language"`
	PaymentTermsName *string `json:"PaymentTermsName"`
}
