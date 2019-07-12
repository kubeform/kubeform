package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermLogicAppActionCustom struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLogicAppActionCustomSpec   `json:"spec,omitempty"`
	Status            AzurermLogicAppActionCustomStatus `json:"status,omitempty"`
}

type AzurermLogicAppActionCustomSpec struct {
	Body       string `json:"body"`
	Name       string `json:"name"`
	LogicAppId string `json:"logic_app_id"`
}



type AzurermLogicAppActionCustomStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermLogicAppActionCustomList is a list of AzurermLogicAppActionCustoms
type AzurermLogicAppActionCustomList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLogicAppActionCustom CRD objects
	Items []AzurermLogicAppActionCustom `json:"items,omitempty"`
}