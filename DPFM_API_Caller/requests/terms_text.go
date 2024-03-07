package requests

type Text struct {
	PaymentTerms		string  `json:"PaymentTerms"`
	Language			string  `json:"Language"`
	PaymentTermsName	string	`json:"PaymentTermsName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
