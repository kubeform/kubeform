package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
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
	ApiManagementName string `json:"api_management_name"`
	ClientId          string `json:"client_id"`
	ClientSecret      string `json:"client_secret"`
	// +optional
	Description       string `json:"description,omitempty"`
	DisplayName       string `json:"display_name"`
	MetadataEndpoint  string `json:"metadata_endpoint"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
}

type ApiManagementOpenidConnectProviderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
