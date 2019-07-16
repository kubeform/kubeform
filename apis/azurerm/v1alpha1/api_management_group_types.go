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

type ApiManagementGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementGroupSpec   `json:"spec,omitempty"`
	Status            ApiManagementGroupStatus `json:"status,omitempty"`
}

type ApiManagementGroupSpec struct {
	ApiManagementName string `json:"api_management_name"`
	// +optional
	Description string `json:"description,omitempty"`
	DisplayName string `json:"display_name"`
	// +optional
	ExternalId        string `json:"external_id,omitempty"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	Type string `json:"type,omitempty"`
}

type ApiManagementGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementGroupList is a list of ApiManagementGroups
type ApiManagementGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementGroup CRD objects
	Items []ApiManagementGroup `json:"items,omitempty"`
}
