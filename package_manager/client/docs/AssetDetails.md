# AssetDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** | The asset&#x27;s unique identifier. | [optional] [default to null]
**Type_** | **string** | The asset&#x27;s type, can be \&quot;addon\&quot; or \&quot;project\&quot;. | [optional] [default to null]
**Author** | **string** | The author&#x27;s username. | [optional] [default to null]
**AuthorId** | **string** | The author&#x27;s unique identifier. | [optional] [default to null]
**Category** | **string** | The category the asset belongs to. | [optional] [default to null]
**CategoryId** | **string** | The unique identifier of the category the asset belongs to. | [optional] [default to null]
**DownloadProvider** | **string** |  | [optional] [default to null]
**DownloadCommit** | **string** |  | [optional] [default to null]
**DownloadHash** | **string** | The asset&#x27;s SHA-256 hash for the latest version. **Note:** This is currently always an empty string as asset versions&#x27; hashes aren&#x27;t computed and stored yet.  | [optional] 
**Cost** | **string** | The asset&#x27;s license as a [SPDX license identifier](https://spdx.org/licenses/). For compatibility reasons, this field is called &#x60;cost&#x60; instead of &#x60;license&#x60;.  | [optional] [default to null]
**GodotVersion** | **string** | The Godot version the asset&#x27;s latest version is intended for (in &#x60;major.minor&#x60; format).&lt;br&gt; This field is present for compatibility reasons with the Godot editor. See also the &#x60;versions&#x60; array.  | [optional] [default to null]
**IconUrl** | **string** | The asset&#x27;s icon URL (should always be a PNG image). | [optional] [default to null]
**IsArchived** | **bool** | If &#x60;true&#x60;, the asset is marked as archived by its author. When archived, it can&#x27;t receive any further reviews but can still be unarchived at any time by the author.  | [optional] [default to null]
**IssuesUrl** | **string** | The asset&#x27;s issue reporting URL (typically associated with the Git repository specified in &#x60;browse_url&#x60;).  | [optional] [default to null]
**ModifyDate** | [**time.Time**](time.Time.md) | The date on which the asset entry was last updated. Note that entries can be edited independently of new asset versions being released.  | [optional] [default to null]
**Rating** | **string** | The asset&#x27;s rating (unused). For compatibility reasons, a value of 0 is always returned. You most likely want &#x60;score&#x60; instead.  | [optional] [default to null]
**SupportLevel** | **string** | The asset&#x27;s support level. | [optional] [default to null]
**Title** | **string** | The asset&#x27;s title (usually less than 50 characters). | [optional] [default to null]
**Version** | **string** | The asset revision string (starting from 1).&lt;br&gt; Every time the asset is edited (for anyone and for any reason), this string is incremented by 1.  | [optional] [default to null]
**VersionString** | **string** | The version string of the latest version (free-form, but usually &#x60;major.minor&#x60; or &#x60;major.minor.patch&#x60;).&lt;br&gt; This field is present for compatibility reasons with the Godot editor. See also the &#x60;versions&#x60; array.  | [optional] [default to null]
**Searchable** | **string** |  | [optional] [default to null]
**Previews** | [**[]AssetPreview**](AssetPreview.md) |  | [optional] [default to null]
**BrowseUrl** | **string** | The asset&#x27;s browsable repository URL. | [optional] [default to null]
**Description** | **string** | The asset&#x27;s full description. | [optional] [default to null]
**DownloadUrl** | **string** | The download link of the asset&#x27;s latest version (should always point to a ZIP archive).&lt;br&gt; This field is present for compatibility reasons with the Godot editor. See also the &#x60;versions&#x60; array.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

