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

type EcsCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EcsClusterSpec   `json:"spec,omitempty"`
	Status            EcsClusterStatus `json:"status,omitempty"`
}

type EcsClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	Name        string                    `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EcsClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EcsClusterList is a list of EcsClusters
type EcsClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EcsCluster CRD objects
	Items []EcsCluster `json:"items,omitempty"`
}
