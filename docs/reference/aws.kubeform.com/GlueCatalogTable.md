## GlueCatalogTable
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlueCatalogTable` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlueCatalogTableSpec](#GlueCatalogTableSpec)***||
| `status` | ***[GlueCatalogTableStatus](#GlueCatalogTableStatus)***||
## GlueCatalogTableSpec
##### (Appears on:[GlueCatalogTable](#GlueCatalogTable), [GlueCatalogTableStatus](#GlueCatalogTableStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `catalogID` | ***string***| ***(Optional)*** |
| `databaseName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `owner` | ***string***| ***(Optional)*** |
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `partitionKeys` | ***[[]GlueCatalogTableSpecPartitionKeys](#GlueCatalogTableSpecPartitionKeys)***| ***(Optional)*** |
| `retention` | ***int***| ***(Optional)*** |
| `storageDescriptor` | ***[[]GlueCatalogTableSpecStorageDescriptor](#GlueCatalogTableSpecStorageDescriptor)***| ***(Optional)*** |
| `tableType` | ***string***| ***(Optional)*** |
| `viewExpandedText` | ***string***| ***(Optional)*** |
| `viewOriginalText` | ***string***| ***(Optional)*** |
## GlueCatalogTableSpecPartitionKeys
##### (Appears on:[GlueCatalogTableSpec](#GlueCatalogTableSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `comment` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `type` | ***string***| ***(Optional)*** |
## GlueCatalogTableSpecStorageDescriptor
##### (Appears on:[GlueCatalogTableSpec](#GlueCatalogTableSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketColumns` | ***[]string***| ***(Optional)*** |
| `columns` | ***[[]GlueCatalogTableSpecStorageDescriptorColumns](#GlueCatalogTableSpecStorageDescriptorColumns)***| ***(Optional)*** |
| `compressed` | ***bool***| ***(Optional)*** |
| `inputFormat` | ***string***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `numberOfBuckets` | ***int***| ***(Optional)*** |
| `outputFormat` | ***string***| ***(Optional)*** |
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `serDeInfo` | ***[[]GlueCatalogTableSpecStorageDescriptorSerDeInfo](#GlueCatalogTableSpecStorageDescriptorSerDeInfo)***| ***(Optional)*** |
| `skewedInfo` | ***[[]GlueCatalogTableSpecStorageDescriptorSkewedInfo](#GlueCatalogTableSpecStorageDescriptorSkewedInfo)***| ***(Optional)*** |
| `sortColumns` | ***[[]GlueCatalogTableSpecStorageDescriptorSortColumns](#GlueCatalogTableSpecStorageDescriptorSortColumns)***| ***(Optional)*** |
| `storedAsSubDirectories` | ***bool***| ***(Optional)*** |
## GlueCatalogTableSpecStorageDescriptorColumns
##### (Appears on:[GlueCatalogTableSpecStorageDescriptor](#GlueCatalogTableSpecStorageDescriptor))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `comment` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `type` | ***string***| ***(Optional)*** |
## GlueCatalogTableSpecStorageDescriptorSerDeInfo
##### (Appears on:[GlueCatalogTableSpecStorageDescriptor](#GlueCatalogTableSpecStorageDescriptor))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `serializationLibrary` | ***string***| ***(Optional)*** |
## GlueCatalogTableSpecStorageDescriptorSkewedInfo
##### (Appears on:[GlueCatalogTableSpecStorageDescriptor](#GlueCatalogTableSpecStorageDescriptor))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `skewedColumnNames` | ***[]string***| ***(Optional)*** |
| `skewedColumnValueLocationMaps` | ***map[string]string***| ***(Optional)*** |
| `skewedColumnValues` | ***[]string***| ***(Optional)*** |
## GlueCatalogTableSpecStorageDescriptorSortColumns
##### (Appears on:[GlueCatalogTableSpecStorageDescriptor](#GlueCatalogTableSpecStorageDescriptor))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `column` | ***string***||
| `sortOrder` | ***int***||
## GlueCatalogTableStatus
##### (Appears on:[GlueCatalogTable](#GlueCatalogTable))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlueCatalogTableSpec](#GlueCatalogTableSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
