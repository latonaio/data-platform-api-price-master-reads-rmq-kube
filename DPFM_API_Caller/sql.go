package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-price-master-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-price-master-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
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
	var priceMaster []dpfm_api_output_formatter.PriceMaster
	for _, fn := range accepter {
		switch fn {
		case "PriceMaster":
			func() {
				priceMaster = c.PriceMaster(mtx, input, output, errs, log)
			}()
		case "PriceMasters":
			func() {
				priceMaster = c.PriceMasters(mtx, input, output, errs, log)
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
) []dpfm_api_output_formatter.PriceMaster {
	supplyChainRelationshipID := input.PriceMaster.SupplyChainRelationshipID
	buyer := input.PriceMaster.Buyer
	seller := input.PriceMaster.Seller
	conditionRecord := input.PriceMaster.ConditionRecord
	conditionSequentialNumber := input.PriceMaster.ConditionSequentialNumber
	product := input.PriceMaster.Product
	conditionValidityStartDate := input.PriceMaster.ConditionValidityStartDate
	conditionValidityEndDate := input.PriceMaster.ConditionValidityEndDate

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_price_master_price_master_data
		WHERE (
		       SupplyChainRelationshipID,
		       Buyer,
		       Seller,
		       ConditionRecord,
		       ConditionSequentialNumber,
		       Product,
		       ConditionValidityStartDate,
		       ConditionValidityEndDate
		) = (?, ?, ?, ?, ?, ?, ?);`,
		supplyChainRelationshipID,
		buyer,
		seller,
		conditionRecord,
		conditionSequentialNumber,
		product,
		conditionValidityStartDate,
		conditionValidityEndDate,
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

func (c *DPFMAPICaller) PriceMasters(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) []dpfm_api_output_formatter.PriceMaster {
	priceMster := input.PriceMaster
	where := fmt.Sprintf("WHERE (Buyer = %d OR Seller = %d)", priceMster.Buyer, priceMster.Seller)

	if input.PriceMaster.SupplyChainRelationshipID > 0 {
		where = fmt.Sprintf("%s\nAND SupplyChainRelationshipID = %v", where, input.PriceMaster.SupplyChainRelationshipID)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_price_master_price_master_data
		` + where + ` ;`,
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
