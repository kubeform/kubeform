package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AvailabilitySet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AvailabilitySetSpec   `json:"spec,omitempty"`
	Status            AvailabilitySetStatus `json:"status,omitempty"`
}

type AvailabilitySetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Location string `json:"location" tf:"location"`
	// +optional
	Managed bool   `json:"managed,omitempty" tf:"managed,omitempty"`
	Name    string `json:"name" tf:"name"`
	// +optional
	PlatformFaultDomainCount int `json:"platformFaultDomainCount,omitempty" tf:"platform_fault_domain_count,omitempty"`
	// +optional
	PlatformUpdateDomainCount int    `json:"platformUpdateDomainCount,omitempty" tf:"platform_update_domain_count,omitempty"`
	ResourceGroupName         string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}



type AvailabilitySetStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *AvailabilitySetSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AvailabilitySetList is a list of AvailabilitySets
type AvailabilitySetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AvailabilitySet CRD objects
	Items []AvailabilitySet `json:"items,omitempty"`
}