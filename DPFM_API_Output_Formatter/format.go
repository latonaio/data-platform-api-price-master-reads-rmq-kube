package dpfm_api_output_formatter

import (
	"data-platform-api-price-master-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-price-master-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToPriceMaster(sdc *api_input_reader.SDC, rows *sql.Rows) (*PriceMaster, error) {
	pm := &requests.PriceMaster{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.ConditionRecordCategory,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.ConditionValidityEndDate,
			&pm.ConditionValidityStartDate,
			&pm.Product,
			&pm.Customer,
			&pm.Supplier,
			&pm.CreationDate,
			&pm.ConditionRateValue,
			&pm.ConditionRateValueUnit,
			&pm.ConditionRateRatio,
			&pm.ConditionRateRatioUnit,
			&pm.ConditionCurrency,
			&pm.BaseUnit,
			&pm.ConditionIsDeleted,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	priceMaster := &PriceMaster{
		BusinessPartner:            data.BusinessPartner,
		ConditionRecordCategory:    data.ConditionRecordCategory,
		ConditionRecord:            data.ConditionRecord,
		ConditionSequentialNumber:  data.ConditionSequentialNumber,
		ConditionType:              data.ConditionType,
		ConditionValidityEndDate:   data.ConditionValidityEndDate,
		ConditionValidityStartDate: data.ConditionValidityStartDate,
		Product:                    data.Product,
		Customer:                   data.Customer,
		Supplier:                   data.Supplier,
		CreationDate:               data.CreationDate,
		ConditionRateValue:         data.ConditionRateValue,
		ConditionRateValueUnit:     data.ConditionRateValueUnit,
		ConditionRateRatio:         data.ConditionRateRatio,
		ConditionRateRatioUnit:     data.ConditionRateRatioUnit,
		ConditionCurrency:          data.ConditionCurrency,
		BaseUnit:                   data.BaseUnit,
		ConditionIsDeleted:         data.ConditionIsDeleted,
	}
	return priceMaster, nil
}
