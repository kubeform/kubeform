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

type StoragegatewayUploadBuffer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewayUploadBufferSpec   `json:"spec,omitempty"`
	Status            StoragegatewayUploadBufferStatus `json:"status,omitempty"`
}

type StoragegatewayUploadBufferSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	DiskID     string `json:"diskID" tf:"disk_id"`
	GatewayArn string `json:"gatewayArn" tf:"gateway_arn"`
}

type StoragegatewayUploadBufferStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StoragegatewayUploadBufferList is a list of StoragegatewayUploadBuffers
type StoragegatewayUploadBufferList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StoragegatewayUploadBuffer CRD objects
	Items []StoragegatewayUploadBuffer `json:"items,omitempty"`
}
