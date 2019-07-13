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

type AzurermApiManagementProperty struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementPropertySpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementPropertyStatus `json:"status,omitempty"`
}

type AzurermApiManagementPropertySpec struct {
	ApiManagementName string   `json:"api_management_name"`
	DisplayName       string   `json:"display_name"`
	Value             string   `json:"value"`
	Secret            bool     `json:"secret"`
	Tags              []string `json:"tags"`
	Name              string   `json:"name"`
	ResourceGroupName string   `json:"resource_group_name"`
}



type AzurermApiManagementPropertyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApiManagementPropertyList is a list of AzurermApiManagementPropertys
type AzurermApiManagementPropertyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementProperty CRD objects
	Items []AzurermApiManagementProperty `json:"items,omitempty"`
}