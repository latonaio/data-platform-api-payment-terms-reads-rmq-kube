package dpfm_api_output_formatter

import (
	"data-platform-api-payment-terms-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToPaymentTerms(rows *sql.Rows) (*[]PaymentTerms, error) {
	defer rows.Close()
	paymentTerms := make([]PaymentTerms, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PaymentTerms{}

		err := rows.Scan(
			&pm.PaymentTerms,
			&pm.BaseDate,
			&pm.BaseDateCalcAddMonth,
			&pm.BaseDateCalcFixedDate,
			&pm.PaymentDueDateCalcAddMonth,
			&pm.PaymentDueDateCalcFixedDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &paymentTerms, nil
		}

		data := pm
		paymentTerms = append(paymentTerms, PaymentTerms{
			PaymentTerms:                data.PaymentTerms,
			BaseDate:                    data.BaseDate,
			BaseDateCalcAddMonth:        data.BaseDateCalcAddMonth,
			BaseDateCalcFixedDate:       data.BaseDateCalcFixedDate,
			PaymentDueDateCalcAddMonth:  data.PaymentDueDateCalcAddMonth,
			PaymentDueDateCalcFixedDate: data.PaymentDueDateCalcFixedDate,
			CreationDate:				 data.CreationDate,
			LastChangeDate:				 data.LastChangeDate,
			IsMarkedForDeletion:		 data.IsMarkedForDeletion,
		})
	}

	return &paymentTerms, nil
}

func ConvertToPaymentTermsText(rows *sql.Rows) (*[]PaymentTermsText, error) {
	defer rows.Close()
	paymentTermsText := make([]PaymentTermsText, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PaymentTermsText{}

		err := rows.Scan(
			&pm.PaymentTerms,
			&pm.Language,
			&pm.PaymentTermsName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &paymentTermsText, err
		}

		data := pm
		paymentTermsText = append(paymentTermsText, PaymentTermsText{
			PaymentTerms:			data.PaymentTerms,
			Language:				data.Language,
			PaymentTermsName:		data.PaymentTermsName,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &paymentTermsText, nil
}
