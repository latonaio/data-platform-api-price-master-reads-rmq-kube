package dpfm_api_input_reader

import (
	"data-platform-api-price-master-reads-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToPriceMaster() *requests.PriceMaster {
	data := sdc.PriceMaster
	return &requests.PriceMaster{
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
}
