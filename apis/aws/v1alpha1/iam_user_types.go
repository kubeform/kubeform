package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type IamUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamUserSpec   `json:"spec,omitempty"`
	Status            IamUserStatus `json:"status,omitempty"`
}

type IamUserSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// Delete user even if it has non-Terraform-managed IAM access keys, login profile or MFA devices
	// +optional
	ForceDestroy bool   `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`
	Name         string `json:"name" tf:"name"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	PermissionsBoundary string `json:"permissionsBoundary,omitempty" tf:"permissions_boundary,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UniqueID string `json:"uniqueID,omitempty" tf:"unique_id,omitempty"`
}



type IamUserStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *IamUserSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamUserList is a list of IamUsers
type IamUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamUser CRD objects
	Items []IamUser `json:"items,omitempty"`
}