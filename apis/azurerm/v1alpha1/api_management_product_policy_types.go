package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ApiManagementProductPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementProductPolicySpec   `json:"spec,omitempty"`
	Status            ApiManagementProductPolicyStatus `json:"status,omitempty"`
}

type ApiManagementProductPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	ProductID         string `json:"productID" tf:"product_id"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	XmlContent string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`
	// +optional
	XmlLink string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

type ApiManagementProductPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementProductPolicyList is a list of ApiManagementProductPolicys
type ApiManagementProductPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementProductPolicy CRD objects
	Items []ApiManagementProductPolicy `json:"items,omitempty"`
}
