package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermStreamAnalyticsStreamInputBlob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStreamAnalyticsStreamInputBlobSpec   `json:"spec,omitempty"`
	Status            AzurermStreamAnalyticsStreamInputBlobStatus `json:"status,omitempty"`
}

type AzurermStreamAnalyticsStreamInputBlobSpecSerialization struct {
	Encoding       string `json:"encoding"`
	Type           string `json:"type"`
	FieldDelimiter string `json:"field_delimiter"`
}

type AzurermStreamAnalyticsStreamInputBlobSpec struct {
	DateFormat             string                                      `json:"date_format"`
	StorageAccountKey      string                                      `json:"storage_account_key"`
	StorageContainerName   string                                      `json:"storage_container_name"`
	Serialization          []AzurermStreamAnalyticsStreamInputBlobSpec `json:"serialization"`
	Name                   string                                      `json:"name"`
	ResourceGroupName      string                                      `json:"resource_group_name"`
	StorageAccountName     string                                      `json:"storage_account_name"`
	TimeFormat             string                                      `json:"time_format"`
	StreamAnalyticsJobName string                                      `json:"stream_analytics_job_name"`
	PathPattern            string                                      `json:"path_pattern"`
}



type AzurermStreamAnalyticsStreamInputBlobStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermStreamAnalyticsStreamInputBlobList is a list of AzurermStreamAnalyticsStreamInputBlobs
type AzurermStreamAnalyticsStreamInputBlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStreamAnalyticsStreamInputBlob CRD objects
	Items []AzurermStreamAnalyticsStreamInputBlob `json:"items,omitempty"`
}