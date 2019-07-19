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

type CodedeployApp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodedeployAppSpec   `json:"spec,omitempty"`
	Status            CodedeployAppStatus `json:"status,omitempty"`
}

type CodedeployAppSpec struct {
	// +optional
	ComputePlatform string `json:"computePlatform,omitempty" tf:"compute_platform,omitempty"`
	Name            string `json:"name" tf:"name"`
	// +optional
	UniqueID    string                    `json:"uniqueID,omitempty" tf:"unique_id,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CodedeployAppStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CodedeployAppList is a list of CodedeployApps
type CodedeployAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CodedeployApp CRD objects
	Items []CodedeployApp `json:"items,omitempty"`
}
