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

type StreamAnalyticsOutputBlob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsOutputBlobSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsOutputBlobStatus `json:"status,omitempty"`
}

type StreamAnalyticsOutputBlobSpecSerialization struct {
	// +optional
	Encoding string `json:"encoding,omitempty"`
	// +optional
	FieldDelimiter string `json:"field_delimiter,omitempty"`
	// +optional
	Format string `json:"format,omitempty"`
	Type   string `json:"type"`
}

type StreamAnalyticsOutputBlobSpec struct {
	DateFormat        string `json:"date_format"`
	Name              string `json:"name"`
	PathPattern       string `json:"path_pattern"`
	ResourceGroupName string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Serialization          []StreamAnalyticsOutputBlobSpec `json:"serialization"`
	StorageAccountKey      string                          `json:"storage_account_key"`
	StorageAccountName     string                          `json:"storage_account_name"`
	StorageContainerName   string                          `json:"storage_container_name"`
	StreamAnalyticsJobName string                          `json:"stream_analytics_job_name"`
	TimeFormat             string                          `json:"time_format"`
}

type StreamAnalyticsOutputBlobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StreamAnalyticsOutputBlobList is a list of StreamAnalyticsOutputBlobs
type StreamAnalyticsOutputBlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamAnalyticsOutputBlob CRD objects
	Items []StreamAnalyticsOutputBlob `json:"items,omitempty"`
}
