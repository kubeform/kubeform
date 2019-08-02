package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type LogicAppTriggerHTTPRequest struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogicAppTriggerHTTPRequestSpec   `json:"spec,omitempty"`
	Status            LogicAppTriggerHTTPRequestStatus `json:"status,omitempty"`
}

type LogicAppTriggerHTTPRequestSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	LogicAppID string `json:"logicAppID" tf:"logic_app_id"`
	// +optional
	Method string `json:"method,omitempty" tf:"method,omitempty"`
	Name   string `json:"name" tf:"name"`
	// +optional
	RelativePath string `json:"relativePath,omitempty" tf:"relative_path,omitempty"`
	Schema       string `json:"schema" tf:"schema"`
}

type LogicAppTriggerHTTPRequestStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *LogicAppTriggerHTTPRequestSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogicAppTriggerHTTPRequestList is a list of LogicAppTriggerHTTPRequests
type LogicAppTriggerHTTPRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogicAppTriggerHTTPRequest CRD objects
	Items []LogicAppTriggerHTTPRequest `json:"items,omitempty"`
}
