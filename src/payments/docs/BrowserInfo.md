# BrowserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptHeader** | **string** | The accept header value of the shopper&#39;s browser. | 
**ColorDepth** | **int32** | The color depth of the shopper&#39;s browser in bits per pixel. This should be obtained by using the browser&#39;s &#x60;screen.colorDepth&#x60; property. Accepted values: 1, 4, 8, 15, 16, 24, 30, 32 or 48 bit color depth. | 
**JavaEnabled** | **bool** | Boolean value indicating if the shopper&#39;s browser is able to execute Java. | 
**JavaScriptEnabled** | Pointer to **bool** | Boolean value indicating if the shopper&#39;s browser is able to execute JavaScript. A default &#39;true&#39; value is assumed if the field is not present. | [optional] [default to true]
**Language** | **string** | The &#x60;navigator.language&#x60; value of the shopper&#39;s browser (as defined in IETF BCP 47). | 
**ScreenHeight** | **int32** | The total height of the shopper&#39;s device screen in pixels. | 
**ScreenWidth** | **int32** | The total width of the shopper&#39;s device screen in pixels. | 
**TimeZoneOffset** | **int32** | Time difference between UTC time and the shopper&#39;s browser local time, in minutes. | 
**UserAgent** | **string** | The user agent value of the shopper&#39;s browser. | 

## Methods

### NewBrowserInfo

`func NewBrowserInfo(acceptHeader string, colorDepth int32, javaEnabled bool, language string, screenHeight int32, screenWidth int32, timeZoneOffset int32, userAgent string, ) *BrowserInfo`

NewBrowserInfo instantiates a new BrowserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserInfoWithDefaults

`func NewBrowserInfoWithDefaults() *BrowserInfo`

NewBrowserInfoWithDefaults instantiates a new BrowserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptHeader

`func (o *BrowserInfo) GetAcceptHeader() string`

GetAcceptHeader returns the AcceptHeader field if non-nil, zero value otherwise.

### GetAcceptHeaderOk

`func (o *BrowserInfo) GetAcceptHeaderOk() (*string, bool)`

GetAcceptHeaderOk returns a tuple with the AcceptHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptHeader

`func (o *BrowserInfo) SetAcceptHeader(v string)`

SetAcceptHeader sets AcceptHeader field to given value.


### GetColorDepth

`func (o *BrowserInfo) GetColorDepth() int32`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *BrowserInfo) GetColorDepthOk() (*int32, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *BrowserInfo) SetColorDepth(v int32)`

SetColorDepth sets ColorDepth field to given value.


### GetJavaEnabled

`func (o *BrowserInfo) GetJavaEnabled() bool`

GetJavaEnabled returns the JavaEnabled field if non-nil, zero value otherwise.

### GetJavaEnabledOk

`func (o *BrowserInfo) GetJavaEnabledOk() (*bool, bool)`

GetJavaEnabledOk returns a tuple with the JavaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaEnabled

`func (o *BrowserInfo) SetJavaEnabled(v bool)`

SetJavaEnabled sets JavaEnabled field to given value.


### GetJavaScriptEnabled

`func (o *BrowserInfo) GetJavaScriptEnabled() bool`

GetJavaScriptEnabled returns the JavaScriptEnabled field if non-nil, zero value otherwise.

### GetJavaScriptEnabledOk

`func (o *BrowserInfo) GetJavaScriptEnabledOk() (*bool, bool)`

GetJavaScriptEnabledOk returns a tuple with the JavaScriptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaScriptEnabled

`func (o *BrowserInfo) SetJavaScriptEnabled(v bool)`

SetJavaScriptEnabled sets JavaScriptEnabled field to given value.

### HasJavaScriptEnabled

`func (o *BrowserInfo) HasJavaScriptEnabled() bool`

HasJavaScriptEnabled returns a boolean if a field has been set.

### GetLanguage

`func (o *BrowserInfo) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BrowserInfo) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BrowserInfo) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetScreenHeight

`func (o *BrowserInfo) GetScreenHeight() int32`

GetScreenHeight returns the ScreenHeight field if non-nil, zero value otherwise.

### GetScreenHeightOk

`func (o *BrowserInfo) GetScreenHeightOk() (*int32, bool)`

GetScreenHeightOk returns a tuple with the ScreenHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenHeight

`func (o *BrowserInfo) SetScreenHeight(v int32)`

SetScreenHeight sets ScreenHeight field to given value.


### GetScreenWidth

`func (o *BrowserInfo) GetScreenWidth() int32`

GetScreenWidth returns the ScreenWidth field if non-nil, zero value otherwise.

### GetScreenWidthOk

`func (o *BrowserInfo) GetScreenWidthOk() (*int32, bool)`

GetScreenWidthOk returns a tuple with the ScreenWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenWidth

`func (o *BrowserInfo) SetScreenWidth(v int32)`

SetScreenWidth sets ScreenWidth field to given value.


### GetTimeZoneOffset

`func (o *BrowserInfo) GetTimeZoneOffset() int32`

GetTimeZoneOffset returns the TimeZoneOffset field if non-nil, zero value otherwise.

### GetTimeZoneOffsetOk

`func (o *BrowserInfo) GetTimeZoneOffsetOk() (*int32, bool)`

GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneOffset

`func (o *BrowserInfo) SetTimeZoneOffset(v int32)`

SetTimeZoneOffset sets TimeZoneOffset field to given value.


### GetUserAgent

`func (o *BrowserInfo) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *BrowserInfo) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *BrowserInfo) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


