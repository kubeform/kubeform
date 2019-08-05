package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CodecommitTrigger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodecommitTriggerSpec   `json:"spec,omitempty"`
	Status            CodecommitTriggerStatus `json:"status,omitempty"`
}

type CodecommitTriggerSpecTrigger struct {
	// +optional
	Branches []string `json:"branches,omitempty" tf:"branches,omitempty"`
	// +optional
	CustomData     string   `json:"customData,omitempty" tf:"custom_data,omitempty"`
	DestinationArn string   `json:"destinationArn" tf:"destination_arn"`
	Events         []string `json:"events" tf:"events"`
	Name           string   `json:"name" tf:"name"`
}

type CodecommitTriggerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	RepositoryName string `json:"repositoryName" tf:"repository_name"`
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:UniqueItems=true
	Trigger []CodecommitTriggerSpecTrigger `json:"trigger" tf:"trigger"`
}

type CodecommitTriggerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CodecommitTriggerSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CodecommitTriggerList is a list of CodecommitTriggers
type CodecommitTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CodecommitTrigger CRD objects
	Items []CodecommitTrigger `json:"items,omitempty"`
}
