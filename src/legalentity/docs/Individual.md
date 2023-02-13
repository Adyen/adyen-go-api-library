# Individual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BirthData** | Pointer to [**BirthData**](BirthData.md) |  | [optional] 
**Email** | Pointer to **string** | The email address of the legal entity. | [optional] 
**IdentificationData** | Pointer to [**IdentificationData**](IdentificationData.md) |  | [optional] 
**Name** | [**Name**](Name.md) |  | 
**Nationality** | Pointer to **string** | The individual&#39;s nationality. | [optional] 
**Phone** | Pointer to [**PhoneNumber**](PhoneNumber.md) |  | [optional] 
**ResidentialAddress** | [**Address**](Address.md) |  | 
**TaxInformation** | Pointer to [**[]TaxInformation**](TaxInformation.md) | The tax information of the individual. | [optional] 
**WebData** | Pointer to [**WebData**](WebData.md) |  | [optional] 

## Methods

### NewIndividual

`func NewIndividual(name Name, residentialAddress Address, ) *Individual`

NewIndividual instantiates a new Individual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualWithDefaults

`func NewIndividualWithDefaults() *Individual`

NewIndividualWithDefaults instantiates a new Individual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBirthData

`func (o *Individual) GetBirthData() BirthData`

GetBirthData returns the BirthData field if non-nil, zero value otherwise.

### GetBirthDataOk

`func (o *Individual) GetBirthDataOk() (*BirthData, bool)`

GetBirthDataOk returns a tuple with the BirthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthData

`func (o *Individual) SetBirthData(v BirthData)`

SetBirthData sets BirthData field to given value.

### HasBirthData

`func (o *Individual) HasBirthData() bool`

HasBirthData returns a boolean if a field has been set.

### GetEmail

`func (o *Individual) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Individual) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Individual) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Individual) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIdentificationData

`func (o *Individual) GetIdentificationData() IdentificationData`

GetIdentificationData returns the IdentificationData field if non-nil, zero value otherwise.

### GetIdentificationDataOk

`func (o *Individual) GetIdentificationDataOk() (*IdentificationData, bool)`

GetIdentificationDataOk returns a tuple with the IdentificationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationData

`func (o *Individual) SetIdentificationData(v IdentificationData)`

SetIdentificationData sets IdentificationData field to given value.

### HasIdentificationData

`func (o *Individual) HasIdentificationData() bool`

HasIdentificationData returns a boolean if a field has been set.

### GetName

`func (o *Individual) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Individual) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Individual) SetName(v Name)`

SetName sets Name field to given value.


### GetNationality

`func (o *Individual) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *Individual) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *Individual) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *Individual) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetPhone

`func (o *Individual) GetPhone() PhoneNumber`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Individual) GetPhoneOk() (*PhoneNumber, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Individual) SetPhone(v PhoneNumber)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Individual) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetResidentialAddress

`func (o *Individual) GetResidentialAddress() Address`

GetResidentialAddress returns the ResidentialAddress field if non-nil, zero value otherwise.

### GetResidentialAddressOk

`func (o *Individual) GetResidentialAddressOk() (*Address, bool)`

GetResidentialAddressOk returns a tuple with the ResidentialAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentialAddress

`func (o *Individual) SetResidentialAddress(v Address)`

SetResidentialAddress sets ResidentialAddress field to given value.


### GetTaxInformation

`func (o *Individual) GetTaxInformation() []TaxInformation`

GetTaxInformation returns the TaxInformation field if non-nil, zero value otherwise.

### GetTaxInformationOk

`func (o *Individual) GetTaxInformationOk() (*[]TaxInformation, bool)`

GetTaxInformationOk returns a tuple with the TaxInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxInformation

`func (o *Individual) SetTaxInformation(v []TaxInformation)`

SetTaxInformation sets TaxInformation field to given value.

### HasTaxInformation

`func (o *Individual) HasTaxInformation() bool`

HasTaxInformation returns a boolean if a field has been set.

### GetWebData

`func (o *Individual) GetWebData() WebData`

GetWebData returns the WebData field if non-nil, zero value otherwise.

### GetWebDataOk

`func (o *Individual) GetWebDataOk() (*WebData, bool)`

GetWebDataOk returns a tuple with the WebData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebData

`func (o *Individual) SetWebData(v WebData)`

SetWebData sets WebData field to given value.

### HasWebData

`func (o *Individual) HasWebData() bool`

HasWebData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


