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

type DbSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbSecurityGroupSpec   `json:"spec,omitempty"`
	Status            DbSecurityGroupStatus `json:"status,omitempty"`
}

type DbSecurityGroupSpecIngress struct {
	// +optional
	Cidr string `json:"cidr,omitempty" tf:"cidr,omitempty"`
	// +optional
	SecurityGroupID string `json:"securityGroupID,omitempty" tf:"security_group_id,omitempty"`
	// +optional
	SecurityGroupName string `json:"securityGroupName,omitempty" tf:"security_group_name,omitempty"`
	// +optional
	SecurityGroupOwnerID string `json:"securityGroupOwnerID,omitempty" tf:"security_group_owner_id,omitempty"`
}

type DbSecurityGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Description string                       `json:"description,omitempty" tf:"description,omitempty"`
	Ingress     []DbSecurityGroupSpecIngress `json:"ingress" tf:"ingress"`
	Name        string                       `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DbSecurityGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DbSecurityGroupSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DbSecurityGroupList is a list of DbSecurityGroups
type DbSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DbSecurityGroup CRD objects
	Items []DbSecurityGroup `json:"items,omitempty"`
}
