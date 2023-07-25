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
		case "SinglePaymentTerms":
			func() {
				paymentTerms = c.SinglePaymentTerms(mtx, input, output, errs, log)
			}()
		case "MultiplePaymentTerms":
			func() {
				paymentTerms = c.MultiplePaymentTerms(mtx, input, output, errs, log)
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

func (c *DPFMAPICaller) SinglePaymentTerms(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PaymentTerms {
	where := fmt.Sprintf("WHERE PaymentTerms = '%s'", input.PaymentTerms.PaymentTerms)

	if input.PaymentTerms.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.PaymentTerms.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_terms_payment_terms_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, PaymentMethod DESC, BaseDate DESC;`,
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

func (c *DPFMAPICaller) MultiplePaymentTerms(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PaymentTerms {

	if input.PaymentTerms.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.PaymentTerms.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_terms_payment_terms_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, PaymentMethod DESC, BaseDate DESC;`,
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
