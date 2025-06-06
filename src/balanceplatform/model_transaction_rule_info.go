/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the TransactionRuleInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransactionRuleInfo{}

// TransactionRuleInfo struct for TransactionRuleInfo
type TransactionRuleInfo struct {
	// The level at which data must be accumulated, used in rules with `type` **velocity** or **maxUsage**. The level must be the [same or lower in hierarchy](https://docs.adyen.com/issuing/transaction-rules#accumulate-data) than the `entityKey`.  If not provided, by default, the rule will accumulate data at the **paymentInstrument** level.  Possible values: **paymentInstrument**, **paymentInstrumentGroup**, **balanceAccount**, **accountHolder**, **balancePlatform**.
	AggregationLevel *string `json:"aggregationLevel,omitempty"`
	// Your description for the transaction rule.
	Description string `json:"description"`
	// The date when the rule will stop being evaluated, in ISO 8601 extended offset date-time format. For example, **2020-12-18T10:15:30+01:00**.  If not provided, the rule will be evaluated until the rule status is set to **inactive**.
	EndDate   *string                  `json:"endDate,omitempty"`
	EntityKey TransactionRuleEntityKey `json:"entityKey"`
	Interval  TransactionRuleInterval  `json:"interval"`
	// The [outcome](https://docs.adyen.com/issuing/transaction-rules#outcome) that will be applied when a transaction meets the conditions of the rule.  Possible values: * **hardBlock**: the transaction is declined. * **scoreBased**: the transaction is assigned the `score` you specified. Adyen calculates the total score and if it exceeds 100, the transaction is declined.  Default value: **hardBlock**.  > **scoreBased** is not allowed when `requestType` is **bankTransfer**.
	OutcomeType *string `json:"outcomeType,omitempty"`
	// Your reference for the transaction rule.
	Reference string `json:"reference"`
	// Indicates the type of request to which the rule applies. If not provided, by default, this is set to **authorization**.  Possible values: **authorization**, **authentication**, **tokenization**, **bankTransfer**.
	RequestType      *string                     `json:"requestType,omitempty"`
	RuleRestrictions TransactionRuleRestrictions `json:"ruleRestrictions"`
	// A positive or negative score applied to the transaction if it meets the conditions of the rule. Required when `outcomeType` is **scoreBased**.  The value must be between **-100** and **100**.
	Score *int32 `json:"score,omitempty"`
	// The date when the rule will start to be evaluated, in ISO 8601 extended offset date-time format. For example, **2020-12-18T10:15:30+01:00**.  If not provided when creating a transaction rule, the `startDate` is set to the date when the rule status is set to **active**.
	StartDate *string `json:"startDate,omitempty"`
	// The status of the transaction rule. If you provide a `startDate` in the request, the rule is automatically created  with an **active** status.   Possible values: **active**, **inactive**.
	Status *string `json:"status,omitempty"`
	// The [type of rule](https://docs.adyen.com/issuing/transaction-rules#rule-types), which defines if a rule blocks transactions based on individual characteristics or accumulates data.  Possible values:  * **blockList**: decline a transaction when the conditions are met.  * **maxUsage**: add the amount or number of transactions for the lifetime of a payment instrument, and then decline a transaction when the specified limits are met.  * **velocity**: add the amount or number of transactions based on a specified time interval, and then decline a transaction when the specified limits are met.
	Type string `json:"type"`
}

// NewTransactionRuleInfo instantiates a new TransactionRuleInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionRuleInfo(description string, entityKey TransactionRuleEntityKey, interval TransactionRuleInterval, reference string, ruleRestrictions TransactionRuleRestrictions, type_ string) *TransactionRuleInfo {
	this := TransactionRuleInfo{}
	this.Description = description
	this.EntityKey = entityKey
	this.Interval = interval
	this.Reference = reference
	this.RuleRestrictions = ruleRestrictions
	this.Type = type_
	return &this
}

// NewTransactionRuleInfoWithDefaults instantiates a new TransactionRuleInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionRuleInfoWithDefaults() *TransactionRuleInfo {
	this := TransactionRuleInfo{}
	return &this
}

// GetAggregationLevel returns the AggregationLevel field value if set, zero value otherwise.
func (o *TransactionRuleInfo) GetAggregationLevel() string {
	if o == nil || common.IsNil(o.AggregationLevel) {
		var ret string
		return ret
	}
	return *o.AggregationLevel
}

// GetAggregationLevelOk returns a tuple with the AggregationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleInfo) GetAggregationLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.AggregationLevel) {
		return nil, false
	}
	return o.AggregationLevel, true
}

// HasAggregationLevel returns a boolean if a field has been set.
func (o *TransactionRuleInfo) HasAggregationLevel() bool {
	if o != nil && !common.IsNil(o.AggregationLevel) {
		return true
	}

	return false
}

// SetAggregationLevel gets a reference to the given string and assigns it to the AggregationLevel field.
func (o *TransactionRuleInfo) SetAggregationLevel(v string) {
	o.AggregationLevel = &v
}

// GetDescription returns the Description field value
func (o *TransactionRuleInfo) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TransactionRuleInfo) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TransactionRuleInfo) SetDescription(v string) {
	o.Description = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *TransactionRuleInfo) GetEndDate() string {
	if o == nil || common.IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleInfo) GetEndDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *TransactionRuleInfo) HasEndDate() bool {
	if o != nil && !common.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *TransactionRuleInfo) SetEndDate(v string) {
	o.EndDate = &v
}

// GetEntityKey returns the EntityKey field value
func (o *TransactionRuleInfo) GetEntityKey() TransactionRuleEntityKey {
	if o == nil {
		var ret TransactionRuleEntityKey
		return ret
	}

	return o.EntityKey
}

// GetEntityKeyOk returns a tuple with the EntityKey field value
// and a boolean to check if the value has been set.
func (o *TransactionRuleInfo) GetEntityKeyOk() (*TransactionRuleEntityKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityKey, true
}

// SetEntityKey sets field value
func (o *TransactionRuleInfo) SetEntityKey(v TransactionRuleEntityKey) {
	o.EntityKey = v
}

// GetInterval returns the Interval field value
func (o *TransactionRuleInfo) GetInterval() TransactionRuleInterval {
	if o == nil {
		var ret TransactionRuleInterval
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *TransactionRuleInfo) GetIntervalOk() (*TransactionRuleInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *TransactionRuleInfo) SetInterval(v TransactionRuleInterval) {
	o.Interval = v
}

// GetOutcomeType returns the OutcomeType field value if set, zero value otherwise.
func (o *TransactionRuleInfo) GetOutcomeType() string {
	if o == nil || common.IsNil(o.OutcomeType) {
		var ret string
		return ret
	}
	return *o.OutcomeType
}

// GetOutcomeTypeOk returns a tuple with the OutcomeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleInfo) GetOutcomeTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OutcomeType) {
		return nil, false
	}
	return o.OutcomeType, true
}

// HasOutcomeType returns a boolean if a field has been set.
func (o *TransactionRuleInfo) HasOutcomeType() bool {
	if o != nil && !common.IsNil(o.OutcomeType) {
		return true
	}

	return false
}

// SetOutcomeType gets a reference to the given string and assigns it to the OutcomeType field.
func (o *TransactionRuleInfo) SetOutcomeType(v string) {
	o.OutcomeType = &v
}

// GetReference returns the Reference field value
func (o *TransactionRuleInfo) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *TransactionRuleInfo) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *TransactionRuleInfo) SetReference(v string) {
	o.Reference = v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *TransactionRuleInfo) GetRequestType() string {
	if o == nil || common.IsNil(o.RequestType) {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleInfo) GetRequestTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestType) {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *TransactionRuleInfo) HasRequestType() bool {
	if o != nil && !common.IsNil(o.RequestType) {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *TransactionRuleInfo) SetRequestType(v string) {
	o.RequestType = &v
}

// GetRuleRestrictions returns the RuleRestrictions field value
func (o *TransactionRuleInfo) GetRuleRestrictions() TransactionRuleRestrictions {
	if o == nil {
		var ret TransactionRuleRestrictions
		return ret
	}

	return o.RuleRestrictions
}

// GetRuleRestrictionsOk returns a tuple with the RuleRestrictions field value
// and a boolean to check if the value has been set.
func (o *TransactionRuleInfo) GetRuleRestrictionsOk() (*TransactionRuleRestrictions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleRestrictions, true
}

// SetRuleRestrictions sets field value
func (o *TransactionRuleInfo) SetRuleRestrictions(v TransactionRuleRestrictions) {
	o.RuleRestrictions = v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *TransactionRuleInfo) GetScore() int32 {
	if o == nil || common.IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleInfo) GetScoreOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *TransactionRuleInfo) HasScore() bool {
	if o != nil && !common.IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *TransactionRuleInfo) SetScore(v int32) {
	o.Score = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *TransactionRuleInfo) GetStartDate() string {
	if o == nil || common.IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleInfo) GetStartDateOk() (*string, bool) {
	if o == nil || common.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *TransactionRuleInfo) HasStartDate() bool {
	if o != nil && !common.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *TransactionRuleInfo) SetStartDate(v string) {
	o.StartDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransactionRuleInfo) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionRuleInfo) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransactionRuleInfo) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TransactionRuleInfo) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value
func (o *TransactionRuleInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionRuleInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionRuleInfo) SetType(v string) {
	o.Type = v
}

func (o TransactionRuleInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionRuleInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AggregationLevel) {
		toSerialize["aggregationLevel"] = o.AggregationLevel
	}
	toSerialize["description"] = o.Description
	if !common.IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["entityKey"] = o.EntityKey
	toSerialize["interval"] = o.Interval
	if !common.IsNil(o.OutcomeType) {
		toSerialize["outcomeType"] = o.OutcomeType
	}
	toSerialize["reference"] = o.Reference
	if !common.IsNil(o.RequestType) {
		toSerialize["requestType"] = o.RequestType
	}
	toSerialize["ruleRestrictions"] = o.RuleRestrictions
	if !common.IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !common.IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTransactionRuleInfo struct {
	value *TransactionRuleInfo
	isSet bool
}

func (v NullableTransactionRuleInfo) Get() *TransactionRuleInfo {
	return v.value
}

func (v *NullableTransactionRuleInfo) Set(val *TransactionRuleInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionRuleInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionRuleInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionRuleInfo(val *TransactionRuleInfo) *NullableTransactionRuleInfo {
	return &NullableTransactionRuleInfo{value: val, isSet: true}
}

func (v NullableTransactionRuleInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionRuleInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TransactionRuleInfo) isValidOutcomeType() bool {
	var allowedEnumValues = []string{"enforceSCA", "hardBlock", "scoreBased", "timedBlock"}
	for _, allowed := range allowedEnumValues {
		if o.GetOutcomeType() == allowed {
			return true
		}
	}
	return false
}
func (o *TransactionRuleInfo) isValidRequestType() bool {
	var allowedEnumValues = []string{"authentication", "authorization", "bankTransfer", "tokenization"}
	for _, allowed := range allowedEnumValues {
		if o.GetRequestType() == allowed {
			return true
		}
	}
	return false
}
func (o *TransactionRuleInfo) isValidStatus() bool {
	var allowedEnumValues = []string{"active", "inactive"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
func (o *TransactionRuleInfo) isValidType() bool {
	var allowedEnumValues = []string{"allowList", "blockList", "maxUsage", "velocity"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
