package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CognitiveAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitiveAccountSpec   `json:"spec,omitempty"`
	Status            CognitiveAccountStatus `json:"status,omitempty"`
}

type CognitiveAccountSpecSku struct {
	Name string `json:"name" tf:"name"`
	Tier string `json:"tier" tf:"tier"`
}

type CognitiveAccountSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Endpoint string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
	Kind     string `json:"kind" tf:"kind"`
	Location string `json:"location" tf:"location"`
	Name     string `json:"name" tf:"name"`
	// +optional
	PrimaryAccessKey  string `json:"-" sensitive:"true" tf:"primary_access_key,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAccessKey string `json:"-" sensitive:"true" tf:"secondary_access_key,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Sku []CognitiveAccountSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}



type CognitiveAccountStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *CognitiveAccountSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CognitiveAccountList is a list of CognitiveAccounts
type CognitiveAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CognitiveAccount CRD objects
	Items []CognitiveAccount `json:"items,omitempty"`
}