# StockData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketIdentifier** | Pointer to **string** | The four-digit [Market Identifier Code](https://en.wikipedia.org/wiki/Market_Identifier_Code) of the stock market where the organization&#39;s stocks are traded. | [optional] 
**StockNumber** | Pointer to **string** | The 12-digit International Securities Identification Number (ISIN) of the company, without dashes (-). | [optional] 
**TickerSymbol** | Pointer to **string** | The stock ticker symbol. | [optional] 

## Methods

### NewStockData

`func NewStockData() *StockData`

NewStockData instantiates a new StockData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockDataWithDefaults

`func NewStockDataWithDefaults() *StockData`

NewStockDataWithDefaults instantiates a new StockData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketIdentifier

`func (o *StockData) GetMarketIdentifier() string`

GetMarketIdentifier returns the MarketIdentifier field if non-nil, zero value otherwise.

### GetMarketIdentifierOk

`func (o *StockData) GetMarketIdentifierOk() (*string, bool)`

GetMarketIdentifierOk returns a tuple with the MarketIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketIdentifier

`func (o *StockData) SetMarketIdentifier(v string)`

SetMarketIdentifier sets MarketIdentifier field to given value.

### HasMarketIdentifier

`func (o *StockData) HasMarketIdentifier() bool`

HasMarketIdentifier returns a boolean if a field has been set.

### GetStockNumber

`func (o *StockData) GetStockNumber() string`

GetStockNumber returns the StockNumber field if non-nil, zero value otherwise.

### GetStockNumberOk

`func (o *StockData) GetStockNumberOk() (*string, bool)`

GetStockNumberOk returns a tuple with the StockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockNumber

`func (o *StockData) SetStockNumber(v string)`

SetStockNumber sets StockNumber field to given value.

### HasStockNumber

`func (o *StockData) HasStockNumber() bool`

HasStockNumber returns a boolean if a field has been set.

### GetTickerSymbol

`func (o *StockData) GetTickerSymbol() string`

GetTickerSymbol returns the TickerSymbol field if non-nil, zero value otherwise.

### GetTickerSymbolOk

`func (o *StockData) GetTickerSymbolOk() (*string, bool)`

GetTickerSymbolOk returns a tuple with the TickerSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerSymbol

`func (o *StockData) SetTickerSymbol(v string)`

SetTickerSymbol sets TickerSymbol field to given value.

### HasTickerSymbol

`func (o *StockData) HasTickerSymbol() bool`

HasTickerSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


