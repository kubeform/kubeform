package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SecurityCenterWorkspace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityCenterWorkspaceSpec   `json:"spec,omitempty"`
	Status            SecurityCenterWorkspaceStatus `json:"status,omitempty"`
}

type SecurityCenterWorkspaceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Scope       string `json:"scope" tf:"scope"`
	WorkspaceID string `json:"workspaceID" tf:"workspace_id"`
}

type SecurityCenterWorkspaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SecurityCenterWorkspaceList is a list of SecurityCenterWorkspaces
type SecurityCenterWorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SecurityCenterWorkspace CRD objects
	Items []SecurityCenterWorkspace `json:"items,omitempty"`
}
