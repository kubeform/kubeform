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

type AzurermLogicAppTriggerCustom struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLogicAppTriggerCustomSpec   `json:"spec,omitempty"`
	Status            AzurermLogicAppTriggerCustomStatus `json:"status,omitempty"`
}

type AzurermLogicAppTriggerCustomSpec struct {
	Name       string `json:"name"`
	LogicAppId string `json:"logic_app_id"`
	Body       string `json:"body"`
}



type AzurermLogicAppTriggerCustomStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermLogicAppTriggerCustomList is a list of AzurermLogicAppTriggerCustoms
type AzurermLogicAppTriggerCustomList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLogicAppTriggerCustom CRD objects
	Items []AzurermLogicAppTriggerCustom `json:"items,omitempty"`
}