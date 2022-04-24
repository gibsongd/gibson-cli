# {{classname}}

All URIs are relative to *https://godotengine.org/asset-library/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetGet**](AssetsApi.md#AssetGet) | **Get** /asset | List assets
[**AssetIdGet**](AssetsApi.md#AssetIdGet) | **Get** /asset/{id} | Get information about an asset

# **AssetGet**
> PaginatedAssetList AssetGet(ctx, optional)
List assets

Return a paginated list of assets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AssetsApiAssetGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiAssetGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**|  | [default to any]
 **category** | **optional.String**|  | 
 **support** | **optional.String**|  | 
 **filter** | **optional.String**|  | 
 **user** | **optional.String**|  | 
 **godotVersion** | **optional.String**|  | 
 **maxResults** | **optional.String**|  | 
 **page** | **optional.String**|  | 
 **offset** | **optional.String**|  | 
 **sort** | **optional.String**|  | 
 **reverse** | **optional.Bool**|  | 

### Return type

[**PaginatedAssetList**](PaginatedAssetList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetIdGet**
> AssetDetails AssetIdGet(ctx, id)
Get information about an asset

Get information about a single asset.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The asset&#x27;s unique identifier. | 

### Return type

[**AssetDetails**](AssetDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

