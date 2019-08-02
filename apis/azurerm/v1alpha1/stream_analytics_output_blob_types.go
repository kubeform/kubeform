package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	Encoding string `json:"encoding,omitempty" tf:"encoding,omitempty"`
	// +optional
	FieldDelimiter string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`
	// +optional
	Format string `json:"format,omitempty" tf:"format,omitempty"`
	Type   string `json:"type" tf:"type"`
}

type StreamAnalyticsOutputBlobSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	DateFormat        string `json:"dateFormat" tf:"date_format"`
	Name              string `json:"name" tf:"name"`
	PathPattern       string `json:"pathPattern" tf:"path_pattern"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Serialization          []StreamAnalyticsOutputBlobSpecSerialization `json:"serialization" tf:"serialization"`
	StorageAccountKey      string                                       `json:"-" sensitive:"true" tf:"storage_account_key"`
	StorageAccountName     string                                       `json:"storageAccountName" tf:"storage_account_name"`
	StorageContainerName   string                                       `json:"storageContainerName" tf:"storage_container_name"`
	StreamAnalyticsJobName string                                       `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name"`
	TimeFormat             string                                       `json:"timeFormat" tf:"time_format"`
}

type StreamAnalyticsOutputBlobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *StreamAnalyticsOutputBlobSpec `json:"state,omitempty"`
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
