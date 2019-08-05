package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type IamServiceLinkedRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamServiceLinkedRoleSpec   `json:"spec,omitempty"`
	Status            IamServiceLinkedRoleStatus `json:"status,omitempty"`
}

type IamServiceLinkedRoleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Arn            string `json:"arn" tf:"arn"`
	AwsServiceName string `json:"awsServiceName" tf:"aws_service_name"`
	CreateDate     string `json:"createDate" tf:"create_date"`
	// +optional
	CustomSuffix string `json:"customSuffix,omitempty" tf:"custom_suffix,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	Path        string `json:"path" tf:"path"`
	UniqueID    string `json:"uniqueID" tf:"unique_id"`
}

type IamServiceLinkedRoleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IamServiceLinkedRoleSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamServiceLinkedRoleList is a list of IamServiceLinkedRoles
type IamServiceLinkedRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamServiceLinkedRole CRD objects
	Items []IamServiceLinkedRole `json:"items,omitempty"`
}
