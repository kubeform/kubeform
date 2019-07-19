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

type TransferUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransferUserSpec   `json:"spec,omitempty"`
	Status            TransferUserStatus `json:"status,omitempty"`
}

type TransferUserSpec struct {
	// +optional
	HomeDirectory string `json:"homeDirectory,omitempty" tf:"home_directory,omitempty"`
	// +optional
	Policy   string `json:"policy,omitempty" tf:"policy,omitempty"`
	Role     string `json:"role" tf:"role"`
	ServerID string `json:"serverID" tf:"server_id"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	UserName    string                    `json:"userName" tf:"user_name"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type TransferUserStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TransferUserList is a list of TransferUsers
type TransferUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TransferUser CRD objects
	Items []TransferUser `json:"items,omitempty"`
}
