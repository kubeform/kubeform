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

type ApiManagementAPIPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementAPIPolicySpec   `json:"spec,omitempty"`
	Status            ApiManagementAPIPolicyStatus `json:"status,omitempty"`
}

type ApiManagementAPIPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	ApiName           string `json:"apiName" tf:"api_name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	XmlContent string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`
	// +optional
	XmlLink string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

type ApiManagementAPIPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiManagementAPIPolicySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementAPIPolicyList is a list of ApiManagementAPIPolicys
type ApiManagementAPIPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementAPIPolicy CRD objects
	Items []ApiManagementAPIPolicy `json:"items,omitempty"`
}
