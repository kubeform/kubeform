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

type AzurermSecurityCenterWorkspace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSecurityCenterWorkspaceSpec   `json:"spec,omitempty"`
	Status            AzurermSecurityCenterWorkspaceStatus `json:"status,omitempty"`
}

type AzurermSecurityCenterWorkspaceSpec struct {
	Scope       string `json:"scope"`
	WorkspaceId string `json:"workspace_id"`
}



type AzurermSecurityCenterWorkspaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermSecurityCenterWorkspaceList is a list of AzurermSecurityCenterWorkspaces
type AzurermSecurityCenterWorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSecurityCenterWorkspace CRD objects
	Items []AzurermSecurityCenterWorkspace `json:"items,omitempty"`
}