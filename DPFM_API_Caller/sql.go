package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-payment-terms-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-payment-terms-reads-rmq-kube/DPFM_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var paymentTerms *[]dpfm_api_output_formatter.PaymentTerms
	var paymentTermsText *[]dpfm_api_output_formatter.PaymentTermsText
	for _, fn := range accepter {
		switch fn {
		case "PaymentTerms":
			func() {
				paymentTerms = c.PaymentTerms(mtx, input, output, errs, log)
			}()
		case "PaymentTermsText":
			func() {
				paymentTermsText = c.PaymentTermsText(mtx, input, output, errs, log)
			}()
		case "PaymentTermsTexts":
			func() {
				paymentTermsText = c.PaymentTermsTexts(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		PaymentTerms:     paymentTerms,
		PaymentTermsText: paymentTermsText,
	}

	return data
}

func (c *DPFMAPICaller) PaymentTerms(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PaymentTerms {
	paymentTerms := input.PaymentTerms.PaymentTerms
	baseDate := input.PaymentTerms.BaseDate

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_terms_payment_terms_data
		WHERE (PaymentTerms, BaseDate) = (?, ?);`, paymentTerms, baseDate,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPaymentTerms(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PaymentTermsText(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PaymentTermsText {
	var args []interface{}
	paymentTerms := input.PaymentTerms.PaymentTerms
	paymentTermsText := input.PaymentTerms.PaymentTermsText

	cnt := 0
	for _, v := range paymentTermsText {
		args = append(args, paymentTerms, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT * 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_terms_payment_terms_text_data
		WHERE (PaymentTerms, Language) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPaymentTermsText(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PaymentTermsTexts(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PaymentTermsText {
	var args []interface{}
	paymentTermsText := input.PaymentTerms.PaymentTermsText

	cnt := 0
	for _, v := range paymentTermsText {
		args = append(args, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?),", cnt-1) + "(?)"
	rows, err := c.db.Query(
		`SELECT * 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_terms_payment_terms_text_data
		WHERE Language IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	//
	data, err := dpfm_api_output_formatter.ConvertToPaymentTermsTexts(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
