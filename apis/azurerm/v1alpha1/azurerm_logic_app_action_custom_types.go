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

type AzurermLogicAppActionCustom struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLogicAppActionCustomSpec   `json:"spec,omitempty"`
	Status            AzurermLogicAppActionCustomStatus `json:"status,omitempty"`
}

type AzurermLogicAppActionCustomSpec struct {
	Name       string `json:"name"`
	LogicAppId string `json:"logic_app_id"`
	Body       string `json:"body"`
}



type AzurermLogicAppActionCustomStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermLogicAppActionCustomList is a list of AzurermLogicAppActionCustoms
type AzurermLogicAppActionCustomList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLogicAppActionCustom CRD objects
	Items []AzurermLogicAppActionCustom `json:"items,omitempty"`
}