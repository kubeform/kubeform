package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermResourceGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermResourceGroupSpec   `json:"spec,omitempty"`
	Status            AzurermResourceGroupStatus `json:"status,omitempty"`
}

type AzurermResourceGroupSpec struct {
	Name     string            `json:"name"`
	Location string            `json:"location"`
	Tags     map[string]string `json:"tags"`
}



type AzurermResourceGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermResourceGroupList is a list of AzurermResourceGroups
type AzurermResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermResourceGroup CRD objects
	Items []AzurermResourceGroup `json:"items,omitempty"`
}