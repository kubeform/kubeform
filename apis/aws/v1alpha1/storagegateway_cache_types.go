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

type StoragegatewayCache struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewayCacheSpec   `json:"spec,omitempty"`
	Status            StoragegatewayCacheStatus `json:"status,omitempty"`
}

type StoragegatewayCacheSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	DiskID     string `json:"diskID" tf:"disk_id"`
	GatewayArn string `json:"gatewayArn" tf:"gateway_arn"`
}

type StoragegatewayCacheStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StoragegatewayCacheList is a list of StoragegatewayCaches
type StoragegatewayCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StoragegatewayCache CRD objects
	Items []StoragegatewayCache `json:"items,omitempty"`
}
