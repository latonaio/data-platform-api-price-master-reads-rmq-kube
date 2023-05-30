package dpfm_api_output_formatter

import (
	"data-platform-api-price-master-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-price-master-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToPriceMaster(sdc *api_input_reader.SDC, rows *sql.Rows) ([]PriceMaster, error) {
	priceMaster := make([]PriceMaster, 0)
	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		pm := &requests.PriceMaster{}
		err := rows.Scan(
			&pm.SupplyChainRelationshipID,
			&pm.Buyer,
			&pm.Seller,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionValidityEndDate,
			&pm.ConditionValidityStartDate,
			&pm.Product,
			&pm.ConditionType,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.ConditionRateValue,
			&pm.ConditionRateValueUnit,
			&pm.ConditionScaleQuantity,
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
		data := pm

		priceMaster = append(priceMaster, PriceMaster{
			SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
			Buyer:                      data.Buyer,
			Seller:                     data.Seller,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionValidityEndDate:   data.ConditionValidityEndDate,
			ConditionValidityStartDate: data.ConditionValidityStartDate,
			Product:                    data.Product,
			ConditionType:              data.ConditionType,
			CreationDate:               data.CreationDate,
			LastChangeDate:             data.LastChangeDate,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionRateValueUnit:     data.ConditionRateValueUnit,
			ConditionScaleQuantity:     data.ConditionScaleQuantity,
			ConditionRateRatio:         data.ConditionRateRatio,
			ConditionRateRatioUnit:     data.ConditionRateRatioUnit,
			ConditionCurrency:          data.ConditionCurrency,
			BaseUnit:                   data.BaseUnit,
			ConditionIsDeleted:         data.ConditionIsDeleted,
		})

	}

	return priceMaster, nil
}
