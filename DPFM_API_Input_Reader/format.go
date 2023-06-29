package dpfm_api_input_reader

import (
	"data-platform-api-price-master-reads-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToPriceMaster() *requests.PriceMaster {
	data := sdc.PriceMaster
	return &requests.PriceMaster{
		SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
		Buyer:                      data.Buyer,
		Seller:                     data.Seller,
		ConditionRecord:            data.ConditionRecord,
		ConditionSequentialNumber:  data.ConditionSequentialNumber,
		ConditionValidityStartDate: data.ConditionValidityStartDate,
		ConditionValidityEndDate:   data.ConditionValidityEndDate,
		Product:                    data.Product,
		ConditionType:              data.ConditionType,
		CreationDate:               data.CreationDate,
		ConditionRateValue:         data.ConditionRateValue,
		ConditionRateValueUnit:     data.ConditionRateValueUnit,
		ConditionCurrency:          data.ConditionCurrency,
		IsMarkedForDeletion:        data.IsMarkedForDeletion,
	}
}
