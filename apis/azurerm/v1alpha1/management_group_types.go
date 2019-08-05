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

type ManagementGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementGroupSpec   `json:"spec,omitempty"`
	Status            ManagementGroupStatus `json:"status,omitempty"`
}

type ManagementGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty"`
	// +optional
	GroupID string `json:"groupID,omitempty" tf:"group_id,omitempty"`
	// +optional
	ParentManagementGroupID string `json:"parentManagementGroupID,omitempty" tf:"parent_management_group_id,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubscriptionIDS []string `json:"subscriptionIDS,omitempty" tf:"subscription_ids,omitempty"`
}

type ManagementGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ManagementGroupSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ManagementGroupList is a list of ManagementGroups
type ManagementGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagementGroup CRD objects
	Items []ManagementGroup `json:"items,omitempty"`
}
