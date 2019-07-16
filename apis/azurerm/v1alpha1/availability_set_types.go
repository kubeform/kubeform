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

type AvailabilitySet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AvailabilitySetSpec   `json:"spec,omitempty"`
	Status            AvailabilitySetStatus `json:"status,omitempty"`
}

type AvailabilitySetSpec struct {
	Location string `json:"location"`
	// +optional
	Managed bool   `json:"managed,omitempty"`
	Name    string `json:"name"`
	// +optional
	PlatformFaultDomainCount int `json:"platform_fault_domain_count,omitempty"`
	// +optional
	PlatformUpdateDomainCount int    `json:"platform_update_domain_count,omitempty"`
	ResourceGroupName         string `json:"resource_group_name"`
}

type AvailabilitySetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
