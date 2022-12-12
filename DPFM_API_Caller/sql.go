package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-price-master-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-price-master-reads-rmq-kube/DPFM_API_Output_Formatter"
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
	var priceMaster *dpfm_api_output_formatter.PriceMaster
	for _, fn := range accepter {
		switch fn {
		case "PriceMaster":
			func() {
				priceMaster = c.PriceMaster(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		PriceMaster: priceMaster,
	}

	return data
}

func (c *DPFMAPICaller) PriceMaster(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.PriceMaster {
	businessPartner := input.PriceMaster.BusinessPartner
	conditionRecordCategory := input.PriceMaster.ConditionRecordCategory
	conditionRecord := input.PriceMaster.ConditionRecord
	conditionSequentialNumber := input.PriceMaster.ConditionSequentialNumber
	conditionType := input.PriceMaster.ConditionType
	conditionValidityEndDate := input.PriceMaster.ConditionValidityEndDate
	conditionValidityStartDate := input.PriceMaster.ConditionValidityStartDate
	product := input.PriceMaster.Product

	rows, err := c.db.Query(
		`SELECT BusinessPartner, ConditionRecordCategory, ConditionRecord, ConditionSequentialNumber, 
		ConditionType, ConditionValidityEndDate, ConditionValidityStartDate, Product, Customer, Supplier, 
		CreationDate, ConditionRateValue, ConditionRateValueUnit, ConditionRateRatio, ConditionRateRatioUnit, 
		ConditionCurrency, BaseUnit, ConditionIsDeleted
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_price_master_price_master_data
		WHERE (BusinessPartner, ConditionRecordCategory, ConditionRecord, ConditionSequentialNumber, 
		ConditionType, ConditionValidityEndDate, ConditionValidityStartDate, Product) = (?, ?, ?, ?, ?, ?, ?, ?);`, businessPartner, conditionRecordCategory, conditionRecord, conditionSequentialNumber, conditionType, conditionValidityEndDate, conditionValidityStartDate, product,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToPriceMaster(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
