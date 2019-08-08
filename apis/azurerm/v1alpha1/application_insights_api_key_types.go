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

type ApplicationInsightsAPIKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationInsightsAPIKeySpec   `json:"spec,omitempty"`
	Status            ApplicationInsightsAPIKeyStatus `json:"status,omitempty"`
}

type ApplicationInsightsAPIKeySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	ApiKey                string `json:"-" sensitive:"true" tf:"api_key,omitempty"`
	ApplicationInsightsID string `json:"applicationInsightsID" tf:"application_insights_id"`
	Name                  string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ReadPermissions []string `json:"readPermissions,omitempty" tf:"read_permissions,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	WritePermissions []string `json:"writePermissions,omitempty" tf:"write_permissions,omitempty"`
}

type ApplicationInsightsAPIKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApplicationInsightsAPIKeySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApplicationInsightsAPIKeyList is a list of ApplicationInsightsAPIKeys
type ApplicationInsightsAPIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApplicationInsightsAPIKey CRD objects
	Items []ApplicationInsightsAPIKey `json:"items,omitempty"`
}
