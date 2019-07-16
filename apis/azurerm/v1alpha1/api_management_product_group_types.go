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

type ApiManagementProductGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementProductGroupSpec   `json:"spec,omitempty"`
	Status            ApiManagementProductGroupStatus `json:"status,omitempty"`
}

type ApiManagementProductGroupSpec struct {
	ApiManagementName string `json:"api_management_name"`
	GroupName         string `json:"group_name"`
	ProductId         string `json:"product_id"`
	ResourceGroupName string `json:"resource_group_name"`
}

type ApiManagementProductGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementProductGroupList is a list of ApiManagementProductGroups
type ApiManagementProductGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementProductGroup CRD objects
	Items []ApiManagementProductGroup `json:"items,omitempty"`
}
