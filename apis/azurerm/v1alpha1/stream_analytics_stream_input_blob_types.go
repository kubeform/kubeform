package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	Encoding string `json:"encoding,omitempty" tf:"encoding,omitempty"`
	// +optional
	FieldDelimiter string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`
	Type           string `json:"type" tf:"type"`
}

type StreamAnalyticsStreamInputBlobSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	DateFormat        string `json:"dateFormat" tf:"date_format"`
	Name              string `json:"name" tf:"name"`
	PathPattern       string `json:"pathPattern" tf:"path_pattern"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Serialization          []StreamAnalyticsStreamInputBlobSpecSerialization `json:"serialization" tf:"serialization"`
	StorageAccountName     string                                            `json:"storageAccountName" tf:"storage_account_name"`
	StorageContainerName   string                                            `json:"storageContainerName" tf:"storage_container_name"`
	StreamAnalyticsJobName string                                            `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name"`
	TimeFormat             string                                            `json:"timeFormat" tf:"time_format"`
}

type StreamAnalyticsStreamInputBlobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
