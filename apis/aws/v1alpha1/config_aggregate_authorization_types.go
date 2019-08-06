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

type ConfigAggregateAuthorization struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigAggregateAuthorizationSpec   `json:"spec,omitempty"`
	Status            ConfigAggregateAuthorizationStatus `json:"status,omitempty"`
}

type ConfigAggregateAuthorizationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AccountID string `json:"accountID" tf:"account_id"`
	// +optional
	Arn    string `json:"arn,omitempty" tf:"arn,omitempty"`
	Region string `json:"region" tf:"region"`
}

type ConfigAggregateAuthorizationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ConfigAggregateAuthorizationSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ConfigAggregateAuthorizationList is a list of ConfigAggregateAuthorizations
type ConfigAggregateAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigAggregateAuthorization CRD objects
	Items []ConfigAggregateAuthorization `json:"items,omitempty"`
}
