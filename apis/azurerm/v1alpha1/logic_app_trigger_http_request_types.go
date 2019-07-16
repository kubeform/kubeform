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

type LogicAppTriggerHttpRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogicAppTriggerHttpRequestSpec   `json:"spec,omitempty"`
	Status            LogicAppTriggerHttpRequestStatus `json:"status,omitempty"`
}

type LogicAppTriggerHttpRequestSpec struct {
	LogicAppId string `json:"logic_app_id"`
	// +optional
	Method string `json:"method,omitempty"`
	Name   string `json:"name"`
	// +optional
	RelativePath string `json:"relative_path,omitempty"`
	Schema       string `json:"schema"`
}

type LogicAppTriggerHttpRequestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogicAppTriggerHttpRequestList is a list of LogicAppTriggerHttpRequests
type LogicAppTriggerHttpRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogicAppTriggerHttpRequest CRD objects
	Items []LogicAppTriggerHttpRequest `json:"items,omitempty"`
}
