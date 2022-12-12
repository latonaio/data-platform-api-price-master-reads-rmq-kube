package requests

type PriceMaster struct {
	BusinessPartner            *int     `json:"BusinessPartner"`
	ConditionRecordCategory    string   `json:"ConditionRecordCategory"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionType              string   `json:"ConditionType"`
	ConditionValidityEndDate   *string  `json:"ConditionValidityEndDate"`
	ConditionValidityStartDate *string  `json:"ConditionValidityStartDate"`
	Product                    string   `json:"Product"`
	Customer                   *int     `json:"Customer"`
	Supplier                   *int     `json:"Supplier"`
	CreationDate               *string  `json:"CreationDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit     string   `json:"ConditionRateValueUnit"`
	ConditionRateRatio         *float32 `json:"ConditionRateRatio"`
	ConditionRateRatioUnit     string   `json:"ConditionRateRatioUnit"`
	ConditionCurrency          string   `json:"ConditionCurrency"`
	BaseUnit                   string   `json:"BaseUnit"`
	ConditionIsDeleted         *bool    `json:"ConditionIsDeleted"`
}
