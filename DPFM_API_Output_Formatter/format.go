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
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &paymentTermsText, err
		}

		data := pm
		paymentTermsText = append(paymentTermsText, PaymentTermsText{
			PaymentTerms:     data.PaymentTerms,
			Language:         data.Language,
			PaymentTermsName: data.PaymentTermsName,
		})
	}

	return &paymentTermsText, nil
}

func ConvertToPaymentTermsTexts(rows *sql.Rows) (*[]PaymentTermsText, error) {
	defer rows.Close()
	paymentTermsText := make([]PaymentTermsText, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PaymentTermsTexts{}

		err := rows.Scan(
			&pm.PaymentTerms,
			&pm.Language,
			&pm.PaymentTermsName,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &paymentTermsText, err
		}

		data := pm
		paymentTermsText = append(paymentTermsText, PaymentTermsText{
			PaymentTerms:     data.PaymentTerms,
			Language:         data.Language,
			PaymentTermsName: data.PaymentTermsName,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &paymentTermsText, nil
	}

	return &paymentTermsText, nil
}
