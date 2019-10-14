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

type IamOpenidConnectProvider struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamOpenidConnectProviderSpec   `json:"spec,omitempty"`
	Status            IamOpenidConnectProviderStatus `json:"status,omitempty"`
}

type IamOpenidConnectProviderSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn            string   `json:"arn,omitempty" tf:"arn,omitempty"`
	ClientIDList   []string `json:"clientIDList" tf:"client_id_list"`
	ThumbprintList []string `json:"thumbprintList" tf:"thumbprint_list"`
	Url            string   `json:"url" tf:"url"`
}

type IamOpenidConnectProviderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IamOpenidConnectProviderSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamOpenidConnectProviderList is a list of IamOpenidConnectProviders
type IamOpenidConnectProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamOpenidConnectProvider CRD objects
	Items []IamOpenidConnectProvider `json:"items,omitempty"`
}
