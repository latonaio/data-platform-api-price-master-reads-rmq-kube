package dpfm_api_output_formatter

import (
	"data-platform-api-price-master-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-price-master-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToPriceMaster(sdc *api_input_reader.SDC, rows *sql.Rows) ([]Header, error) {
	priceMaster := make([]Header, 0)
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
			&pm.Product,
			&pm.ConditionValidityStartDate,
			&pm.ConditionValidityEndDate,
			&pm.ConditionType,
			&pm.ConditionRateValue,
			&pm.ConditionRateValueUnit,
			&pm.ConditionScaleQuantity,
			&pm.ConditionCurrency,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
		data := pm

		priceMaster = append(priceMaster, Header{
			SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
			Buyer:                      data.Buyer,
			Seller:                     data.Seller,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			Product:                    data.Product,
			ConditionValidityStartDate: data.ConditionValidityStartDate,
			ConditionValidityEndDate:   data.ConditionValidityEndDate,
			ConditionType:              data.ConditionType,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionRateValueUnit:     data.ConditionRateValueUnit,
			ConditionScaleQuantity:     data.ConditionScaleQuantity,
			ConditionCurrency:          data.ConditionCurrency,
			CreationDate:               data.CreationDate,
			LastChangeDate:             data.LastChangeDate,
			IsMarkedForDeletion:        data.IsMarkedForDeletion,
		})

	}

	return priceMaster, nil
}
