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

type StreamAnalyticsStreamInputBlob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsStreamInputBlobSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsStreamInputBlobStatus `json:"status,omitempty"`
}

type StreamAnalyticsStreamInputBlobSpecSerialization struct {
	// +optional
	Encoding string `json:"encoding,omitempty"`
	// +optional
	FieldDelimiter string `json:"field_delimiter,omitempty"`
	Type           string `json:"type"`
}

type StreamAnalyticsStreamInputBlobSpec struct {
	DateFormat        string `json:"date_format"`
	Name              string `json:"name"`
	PathPattern       string `json:"path_pattern"`
	ResourceGroupName string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Serialization          []StreamAnalyticsStreamInputBlobSpec `json:"serialization"`
	StorageAccountKey      string                               `json:"storage_account_key"`
	StorageAccountName     string                               `json:"storage_account_name"`
	StorageContainerName   string                               `json:"storage_container_name"`
	StreamAnalyticsJobName string                               `json:"stream_analytics_job_name"`
	TimeFormat             string                               `json:"time_format"`
}

type StreamAnalyticsStreamInputBlobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StreamAnalyticsStreamInputBlobList is a list of StreamAnalyticsStreamInputBlobs
type StreamAnalyticsStreamInputBlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamAnalyticsStreamInputBlob CRD objects
	Items []StreamAnalyticsStreamInputBlob `json:"items,omitempty"`
}
