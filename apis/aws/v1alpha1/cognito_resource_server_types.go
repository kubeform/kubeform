package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CognitoResourceServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoResourceServerSpec   `json:"spec,omitempty"`
	Status            CognitoResourceServerStatus `json:"status,omitempty"`
}

type CognitoResourceServerSpecScope struct {
	ScopeDescription string `json:"scopeDescription" tf:"scope_description"`
	ScopeName        string `json:"scopeName" tf:"scope_name"`
}

type CognitoResourceServerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	Identifier  string                    `json:"identifier" tf:"identifier"`
	Name        string                    `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=25
	// +kubebuilder:validation:UniqueItems=true
	Scope      []CognitoResourceServerSpecScope `json:"scope,omitempty" tf:"scope,omitempty"`
	UserPoolID string                           `json:"userPoolID" tf:"user_pool_id"`
}

type CognitoResourceServerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CognitoResourceServerList is a list of CognitoResourceServers
type CognitoResourceServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CognitoResourceServer CRD objects
	Items []CognitoResourceServer `json:"items,omitempty"`
}
