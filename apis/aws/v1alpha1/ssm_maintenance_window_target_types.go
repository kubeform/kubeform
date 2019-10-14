package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type SsmMaintenanceWindowTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmMaintenanceWindowTargetSpec   `json:"spec,omitempty"`
	Status            SsmMaintenanceWindowTargetStatus `json:"status,omitempty"`
}

type SsmMaintenanceWindowTargetSpecTargets struct {
	Key    string   `json:"key" tf:"key"`
	Values []string `json:"values" tf:"values"`
}

type SsmMaintenanceWindowTargetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	OwnerInformation string `json:"ownerInformation,omitempty" tf:"owner_information,omitempty"`
	ResourceType     string `json:"resourceType" tf:"resource_type"`
	// +kubebuilder:validation:MaxItems=5
	Targets  []SsmMaintenanceWindowTargetSpecTargets `json:"targets" tf:"targets"`
	WindowID string                                  `json:"windowID" tf:"window_id"`
}

type SsmMaintenanceWindowTargetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SsmMaintenanceWindowTargetSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmMaintenanceWindowTargetList is a list of SsmMaintenanceWindowTargets
type SsmMaintenanceWindowTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmMaintenanceWindowTarget CRD objects
	Items []SsmMaintenanceWindowTarget `json:"items,omitempty"`
}
