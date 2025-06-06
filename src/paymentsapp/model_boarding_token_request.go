/*
Payments App API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paymentsapp

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v21/src/common"
)

// checks if the BoardingTokenRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BoardingTokenRequest{}

// BoardingTokenRequest struct for BoardingTokenRequest
type BoardingTokenRequest struct {
	// The boardingToken request token.
	BoardingRequestToken string `json:"boardingRequestToken"`
}

// NewBoardingTokenRequest instantiates a new BoardingTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoardingTokenRequest(boardingRequestToken string) *BoardingTokenRequest {
	this := BoardingTokenRequest{}
	this.BoardingRequestToken = boardingRequestToken
	return &this
}

// NewBoardingTokenRequestWithDefaults instantiates a new BoardingTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoardingTokenRequestWithDefaults() *BoardingTokenRequest {
	this := BoardingTokenRequest{}
	return &this
}

// GetBoardingRequestToken returns the BoardingRequestToken field value
func (o *BoardingTokenRequest) GetBoardingRequestToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BoardingRequestToken
}

// GetBoardingRequestTokenOk returns a tuple with the BoardingRequestToken field value
// and a boolean to check if the value has been set.
func (o *BoardingTokenRequest) GetBoardingRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoardingRequestToken, true
}

// SetBoardingRequestToken sets field value
func (o *BoardingTokenRequest) SetBoardingRequestToken(v string) {
	o.BoardingRequestToken = v
}

func (o BoardingTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoardingTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boardingRequestToken"] = o.BoardingRequestToken
	return toSerialize, nil
}

type NullableBoardingTokenRequest struct {
	value *BoardingTokenRequest
	isSet bool
}

func (v NullableBoardingTokenRequest) Get() *BoardingTokenRequest {
	return v.value
}

func (v *NullableBoardingTokenRequest) Set(val *BoardingTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBoardingTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBoardingTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoardingTokenRequest(val *BoardingTokenRequest) *NullableBoardingTokenRequest {
	return &NullableBoardingTokenRequest{value: val, isSet: true}
}

func (v NullableBoardingTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoardingTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
