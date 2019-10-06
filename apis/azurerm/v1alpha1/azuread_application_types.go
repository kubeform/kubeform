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

type AzureadApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzureadApplicationSpec   `json:"spec,omitempty"`
	Status            AzureadApplicationStatus `json:"status,omitempty"`
}

type AzureadApplicationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApplicationID string `json:"applicationID,omitempty" tf:"application_id,omitempty"`
	// +optional
	AvailableToOtherTenants bool `json:"availableToOtherTenants,omitempty" tf:"available_to_other_tenants,omitempty"`
	// +optional
	Homepage string `json:"homepage,omitempty" tf:"homepage,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	IdentifierUris []string `json:"identifierUris,omitempty" tf:"identifier_uris,omitempty"`
	Name           string   `json:"name" tf:"name"`
	// +optional
	Oauth2AllowImplicitFlow bool `json:"oauth2AllowImplicitFlow,omitempty" tf:"oauth2_allow_implicit_flow,omitempty"`
	// +optional
	ReplyUrls []string `json:"replyUrls,omitempty" tf:"reply_urls,omitempty"`
}

type AzureadApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AzureadApplicationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzureadApplicationList is a list of AzureadApplications
type AzureadApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzureadApplication CRD objects
	Items []AzureadApplication `json:"items,omitempty"`
}
