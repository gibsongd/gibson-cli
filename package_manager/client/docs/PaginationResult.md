# PaginationResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **string** | The requested page string. | [optional] [default to null]
**PageLength** | **string** | The requested page length.&lt;br&gt; **Note:** This can be higher than the total amount of items returned.  | [optional] [default to null]
**Pages** | **string** | The total string of pages available.&lt;br&gt; **Note:** If requesting a page higher than this value, a successful response will be returned (status code 200) but no items will be listed.  | [optional] [default to null]
**TotalItems** | **string** | The total string of items available. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

