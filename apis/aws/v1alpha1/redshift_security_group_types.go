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

type RedshiftSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedshiftSecurityGroupSpec   `json:"spec,omitempty"`
	Status            RedshiftSecurityGroupStatus `json:"status,omitempty"`
}

type RedshiftSecurityGroupSpecIngress struct {
	// +optional
	Cidr string `json:"cidr,omitempty" tf:"cidr,omitempty"`
	// +optional
	SecurityGroupName string `json:"securityGroupName,omitempty" tf:"security_group_name,omitempty"`
	// +optional
	SecurityGroupOwnerID string `json:"securityGroupOwnerID,omitempty" tf:"security_group_owner_id,omitempty"`
}

type RedshiftSecurityGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description string                             `json:"description,omitempty" tf:"description,omitempty"`
	Ingress     []RedshiftSecurityGroupSpecIngress `json:"ingress" tf:"ingress"`
	Name        string                             `json:"name" tf:"name"`
}

type RedshiftSecurityGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *RedshiftSecurityGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedshiftSecurityGroupList is a list of RedshiftSecurityGroups
type RedshiftSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RedshiftSecurityGroup CRD objects
	Items []RedshiftSecurityGroup `json:"items,omitempty"`
}
