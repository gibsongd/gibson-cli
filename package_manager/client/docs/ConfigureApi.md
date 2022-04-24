# {{classname}}

All URIs are relative to *https://godotengine.org/asset-library/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureGet**](ConfigureApi.md#ConfigureGet) | **Get** /configure | Fetch categories

# **ConfigureGet**
> Category ConfigureGet(ctx, optional)
Fetch categories

Returns category names and IDs (used for editor integration).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigureApiConfigureGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigureApiConfigureGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**|  | [default to any]
 **session** | **optional.Bool**|  | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

