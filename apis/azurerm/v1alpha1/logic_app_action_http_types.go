package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type LogicAppActionHTTP struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogicAppActionHTTPSpec   `json:"spec,omitempty"`
	Status            LogicAppActionHTTPStatus `json:"status,omitempty"`
}

type LogicAppActionHTTPSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Body string `json:"body,omitempty" tf:"body,omitempty"`
	// +optional
	Headers    map[string]string `json:"headers,omitempty" tf:"headers,omitempty"`
	LogicAppID string            `json:"logicAppID" tf:"logic_app_id"`
	Method     string            `json:"method" tf:"method"`
	Name       string            `json:"name" tf:"name"`
	Uri        string            `json:"uri" tf:"uri"`
}

type LogicAppActionHTTPStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LogicAppActionHTTPSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogicAppActionHTTPList is a list of LogicAppActionHTTPs
type LogicAppActionHTTPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogicAppActionHTTP CRD objects
	Items []LogicAppActionHTTP `json:"items,omitempty"`
}
