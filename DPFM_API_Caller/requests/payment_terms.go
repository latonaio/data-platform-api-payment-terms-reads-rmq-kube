package requests

type PaymentTerms struct {
	PaymentTerms                string	`json:"PaymentTerms"`
	BaseDate                    int		`json:"BaseDate"`
	BaseDateCalcAddMonth        int		`json:"BaseDateCalcAddMonth"`
	BaseDateCalcFixedDate       int		`json:"BaseDateCalcFixedDate"`
	PaymentDueDateCalcAddMonth  int		`json:"PaymentDueDateCalcAddMonth"`
	PaymentDueDateCalcFixedDate int		`json:"PaymentDueDateCalcFixedDate"`
	CreationDate				string	`json:"CreationDate"`
	LastChangeDate				string	`json:"LastChangeDate"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}
