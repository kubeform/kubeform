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

type AzurermLogicAppTriggerHttpRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLogicAppTriggerHttpRequestSpec   `json:"spec,omitempty"`
	Status            AzurermLogicAppTriggerHttpRequestStatus `json:"status,omitempty"`
}

type AzurermLogicAppTriggerHttpRequestSpec struct {
	Name         string `json:"name"`
	LogicAppId   string `json:"logic_app_id"`
	Schema       string `json:"schema"`
	Method       string `json:"method"`
	RelativePath string `json:"relative_path"`
}



type AzurermLogicAppTriggerHttpRequestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermLogicAppTriggerHttpRequestList is a list of AzurermLogicAppTriggerHttpRequests
type AzurermLogicAppTriggerHttpRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLogicAppTriggerHttpRequest CRD objects
	Items []AzurermLogicAppTriggerHttpRequest `json:"items,omitempty"`
}