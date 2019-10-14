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

type SsmActivation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmActivationSpec   `json:"spec,omitempty"`
	Status            SsmActivationStatus `json:"status,omitempty"`
}

type SsmActivationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ActivationCode string `json:"activationCode,omitempty" tf:"activation_code,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	ExpirationDate string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`
	// +optional
	Expired string `json:"expired,omitempty" tf:"expired,omitempty"`
	IamRole string `json:"iamRole" tf:"iam_role"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	RegistrationCount int `json:"registrationCount,omitempty" tf:"registration_count,omitempty"`
	// +optional
	RegistrationLimit int `json:"registrationLimit,omitempty" tf:"registration_limit,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SsmActivationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SsmActivationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmActivationList is a list of SsmActivations
type SsmActivationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmActivation CRD objects
	Items []SsmActivation `json:"items,omitempty"`
}
