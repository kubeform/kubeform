package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SfnStateMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SfnStateMachineSpec   `json:"spec,omitempty"`
	Status            SfnStateMachineStatus `json:"status,omitempty"`
}

type SfnStateMachineSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Definition string `json:"definition" tf:"definition"`
	Name       string `json:"name" tf:"name"`
	RoleArn    string `json:"roleArn" tf:"role_arn"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SfnStateMachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *SfnStateMachineSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SfnStateMachineList is a list of SfnStateMachines
type SfnStateMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SfnStateMachine CRD objects
	Items []SfnStateMachine `json:"items,omitempty"`
}
