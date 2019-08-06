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

type ApiManagementOpenidConnectProvider struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementOpenidConnectProviderSpec   `json:"spec,omitempty"`
	Status            ApiManagementOpenidConnectProviderStatus `json:"status,omitempty"`
}

type ApiManagementOpenidConnectProviderSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	ClientID          string `json:"-" sensitive:"true" tf:"client_id"`
	ClientSecret      string `json:"-" sensitive:"true" tf:"client_secret"`
	// +optional
	Description       string `json:"description,omitempty" tf:"description,omitempty"`
	DisplayName       string `json:"displayName" tf:"display_name"`
	MetadataEndpoint  string `json:"metadataEndpoint" tf:"metadata_endpoint"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

type ApiManagementOpenidConnectProviderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiManagementOpenidConnectProviderSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementOpenidConnectProviderList is a list of ApiManagementOpenidConnectProviders
type ApiManagementOpenidConnectProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementOpenidConnectProvider CRD objects
	Items []ApiManagementOpenidConnectProvider `json:"items,omitempty"`
}
