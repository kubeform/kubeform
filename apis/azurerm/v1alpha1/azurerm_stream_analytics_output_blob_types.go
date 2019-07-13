package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AzurermStreamAnalyticsOutputBlob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStreamAnalyticsOutputBlobSpec   `json:"spec,omitempty"`
	Status            AzurermStreamAnalyticsOutputBlobStatus `json:"status,omitempty"`
}

type AzurermStreamAnalyticsOutputBlobSpecSerialization struct {
	FieldDelimiter string `json:"field_delimiter"`
	Encoding       string `json:"encoding"`
	Format         string `json:"format"`
	Type           string `json:"type"`
}

type AzurermStreamAnalyticsOutputBlobSpec struct {
	StorageAccountName     string                                 `json:"storage_account_name"`
	StorageContainerName   string                                 `json:"storage_container_name"`
	TimeFormat             string                                 `json:"time_format"`
	Serialization          []AzurermStreamAnalyticsOutputBlobSpec `json:"serialization"`
	StreamAnalyticsJobName string                                 `json:"stream_analytics_job_name"`
	ResourceGroupName      string                                 `json:"resource_group_name"`
	DateFormat             string                                 `json:"date_format"`
	PathPattern            string                                 `json:"path_pattern"`
	StorageAccountKey      string                                 `json:"storage_account_key"`
	Name                   string                                 `json:"name"`
}



type AzurermStreamAnalyticsOutputBlobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermStreamAnalyticsOutputBlobList is a list of AzurermStreamAnalyticsOutputBlobs
type AzurermStreamAnalyticsOutputBlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStreamAnalyticsOutputBlob CRD objects
	Items []AzurermStreamAnalyticsOutputBlob `json:"items,omitempty"`
}