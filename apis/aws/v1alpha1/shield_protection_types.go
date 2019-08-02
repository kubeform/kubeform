package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ShieldProtection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ShieldProtectionSpec   `json:"spec,omitempty"`
	Status            ShieldProtectionStatus `json:"status,omitempty"`
}

type ShieldProtectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Name        string `json:"name" tf:"name"`
	ResourceArn string `json:"resourceArn" tf:"resource_arn"`
}

type ShieldProtectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *ShieldProtectionSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ShieldProtectionList is a list of ShieldProtections
type ShieldProtectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ShieldProtection CRD objects
	Items []ShieldProtection `json:"items,omitempty"`
}
