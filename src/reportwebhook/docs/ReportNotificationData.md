# reportwebhookData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolder** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**BalanceAccount** | Pointer to [**ResourceReference**](ResourceReference.md) |  | [optional] 
**BalancePlatform** | Pointer to **string** | The unique identifier of the balance platform. | [optional] 
**CreationDate** | Pointer to **time.Time** | The date and time when the event was triggered, in ISO 8601 extended format. For example, **2020-12-18T10:15:30+01:00**. | [optional] 
**DownloadUrl** | **string** | The URL at which you can download the report. To download, you must authenticate your GET request with your [API credentials](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/overview). | 
**FileName** | **string** | The filename of the report. | 
**ReportType** | **string** | Type of report. | 

## Methods

### NewreportwebhookData

`func NewreportwebhookData(downloadUrl string, fileName string, reportType string, ) *reportwebhookData`

NewreportwebhookData instantiates a new reportwebhookData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewreportwebhookDataWithDefaults

`func NewreportwebhookDataWithDefaults() *reportwebhookData`

NewreportwebhookDataWithDefaults instantiates a new reportwebhookData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolder

`func (o *reportwebhookData) GetAccountHolder() ResourceReference`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *reportwebhookData) GetAccountHolderOk() (*ResourceReference, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *reportwebhookData) SetAccountHolder(v ResourceReference)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *reportwebhookData) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.

### GetBalanceAccount

`func (o *reportwebhookData) GetBalanceAccount() ResourceReference`

GetBalanceAccount returns the BalanceAccount field if non-nil, zero value otherwise.

### GetBalanceAccountOk

`func (o *reportwebhookData) GetBalanceAccountOk() (*ResourceReference, bool)`

GetBalanceAccountOk returns a tuple with the BalanceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAccount

`func (o *reportwebhookData) SetBalanceAccount(v ResourceReference)`

SetBalanceAccount sets BalanceAccount field to given value.

### HasBalanceAccount

`func (o *reportwebhookData) HasBalanceAccount() bool`

HasBalanceAccount returns a boolean if a field has been set.

### GetBalancePlatform

`func (o *reportwebhookData) GetBalancePlatform() string`

GetBalancePlatform returns the BalancePlatform field if non-nil, zero value otherwise.

### GetBalancePlatformOk

`func (o *reportwebhookData) GetBalancePlatformOk() (*string, bool)`

GetBalancePlatformOk returns a tuple with the BalancePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalancePlatform

`func (o *reportwebhookData) SetBalancePlatform(v string)`

SetBalancePlatform sets BalancePlatform field to given value.

### HasBalancePlatform

`func (o *reportwebhookData) HasBalancePlatform() bool`

HasBalancePlatform returns a boolean if a field has been set.

### GetCreationDate

`func (o *reportwebhookData) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *reportwebhookData) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *reportwebhookData) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *reportwebhookData) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *reportwebhookData) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *reportwebhookData) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *reportwebhookData) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### GetFileName

`func (o *reportwebhookData) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *reportwebhookData) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *reportwebhookData) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetReportType

`func (o *reportwebhookData) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *reportwebhookData) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *reportwebhookData) SetReportType(v string)`

SetReportType sets ReportType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


