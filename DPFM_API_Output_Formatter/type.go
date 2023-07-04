package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	PriceMaster []PriceMaster `json:"PriceMaster"`
}

type PriceMaster struct {
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID"`
	Buyer                      int      `json:"Buyer"`
	Seller                     int      `json:"Seller"`
	ConditionRecord            int      `json:"ConditionRecord"`
	ConditionSequentialNumber  int      `json:"ConditionSequentialNumber"`
	Product                    string   `json:"Product"`
	ConditionValidityStartDate string   `json:"ConditionValidityStartDate"`
	ConditionValidityEndDate   string   `json:"ConditionValidityEndDate"`
	ConditionType              string   `json:"ConditionType"`
	ConditionRateValue         float32  `json:"ConditionRateValue"`
	ConditionRateValueUnit     int      `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity     int      `json:"ConditionScaleQuantity"`
	ConditionCurrency          string   `json:"ConditionCurrency"`
	CreationDate               string   `json:"CreationDate"`
	LastChangeDate             string   `json:"LastChangeDate"`
	IsMarkedForDeletion        *bool    `json:"IsMarkedForDeletion"`
}
