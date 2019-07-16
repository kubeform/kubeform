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

type ApiManagementApiOperationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementApiOperationPolicySpec   `json:"spec,omitempty"`
	Status            ApiManagementApiOperationPolicyStatus `json:"status,omitempty"`
}

type ApiManagementApiOperationPolicySpec struct {
	ApiManagementName string `json:"api_management_name"`
	ApiName           string `json:"api_name"`
	OperationId       string `json:"operation_id"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	XmlLink string `json:"xml_link,omitempty"`
}

type ApiManagementApiOperationPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementApiOperationPolicyList is a list of ApiManagementApiOperationPolicys
type ApiManagementApiOperationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementApiOperationPolicy CRD objects
	Items []ApiManagementApiOperationPolicy `json:"items,omitempty"`
}
