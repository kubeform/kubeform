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

type SsmParameter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmParameterSpec   `json:"spec,omitempty"`
	Status            SsmParameterStatus `json:"status,omitempty"`
}

type SsmParameterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AllowedPattern string `json:"allowedPattern,omitempty" tf:"allowed_pattern,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	KeyID string `json:"keyID,omitempty" tf:"key_id,omitempty"`
	Name  string `json:"name" tf:"name"`
	// +optional
	Overwrite bool `json:"overwrite,omitempty" tf:"overwrite,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Tier  string `json:"tier,omitempty" tf:"tier,omitempty"`
	Type  string `json:"type" tf:"type"`
	Value string `json:"-" sensitive:"true" tf:"value"`
}

type SsmParameterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SsmParameterSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmParameterList is a list of SsmParameters
type SsmParameterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmParameter CRD objects
	Items []SsmParameter `json:"items,omitempty"`
}
