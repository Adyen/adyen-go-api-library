# Company

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CompanyLinks**](CompanyLinks.md) |  | [optional] 
**DataCenters** | Pointer to [**[]DataCenter**](DataCenter.md) | List of available data centers.  Adyen has several data centers around the world.In the URL that you use for making API requests, we recommend you use the live URL prefix from the data center closest to your shoppers. | [optional] 
**Description** | Pointer to **string** | Your description for the company account, maximum 300 characters | [optional] 
**Id** | Pointer to **string** | The unique identifier of the company account. | [optional] 
**Name** | Pointer to **string** | The legal or trading name of the company. | [optional] 
**Reference** | Pointer to **string** | Your reference to the account | [optional] 
**Status** | Pointer to **string** | The status of the company account.  Possible values:  * **Active**: Users can log in. Processing and payout capabilities depend on the status of the merchant account. * **Inactive**: Users can log in. Payment processing and payouts are disabled. * **Closed**: The company account is closed and this cannot be reversed. Users cannot log in. Payment processing and payouts are disabled. | [optional] 

## Methods

### NewCompany

`func NewCompany() *Company`

NewCompany instantiates a new Company object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWithDefaults

`func NewCompanyWithDefaults() *Company`

NewCompanyWithDefaults instantiates a new Company object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Company) GetLinks() CompanyLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Company) GetLinksOk() (*CompanyLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Company) SetLinks(v CompanyLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Company) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDataCenters

`func (o *Company) GetDataCenters() []DataCenter`

GetDataCenters returns the DataCenters field if non-nil, zero value otherwise.

### GetDataCentersOk

`func (o *Company) GetDataCentersOk() (*[]DataCenter, bool)`

GetDataCentersOk returns a tuple with the DataCenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCenters

`func (o *Company) SetDataCenters(v []DataCenter)`

SetDataCenters sets DataCenters field to given value.

### HasDataCenters

`func (o *Company) HasDataCenters() bool`

HasDataCenters returns a boolean if a field has been set.

### GetDescription

`func (o *Company) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Company) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Company) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Company) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Company) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Company) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Company) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Company) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Company) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Company) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Company) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Company) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReference

`func (o *Company) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Company) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Company) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Company) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetStatus

`func (o *Company) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Company) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Company) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Company) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


