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

type OpsworksPermission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksPermissionSpec   `json:"spec,omitempty"`
	Status            OpsworksPermissionStatus `json:"status,omitempty"`
}

type OpsworksPermissionSpec struct {
	// +optional
	AllowSSH bool `json:"allowSSH,omitempty" tf:"allow_ssh,omitempty"`
	// +optional
	AllowSudo bool `json:"allowSudo,omitempty" tf:"allow_sudo,omitempty"`
	// +optional
	Level string `json:"level,omitempty" tf:"level,omitempty"`
	// +optional
	StackID     string                    `json:"stackID,omitempty" tf:"stack_id,omitempty"`
	UserArn     string                    `json:"userArn" tf:"user_arn"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type OpsworksPermissionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksPermissionList is a list of OpsworksPermissions
type OpsworksPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksPermission CRD objects
	Items []OpsworksPermission `json:"items,omitempty"`
}
