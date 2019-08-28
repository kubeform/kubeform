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

type Route53ResolverRuleAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverRuleAssociationSpec   `json:"spec,omitempty"`
	Status            Route53ResolverRuleAssociationStatus `json:"status,omitempty"`
}

type Route53ResolverRuleAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Name           string `json:"name,omitempty" tf:"name,omitempty"`
	ResolverRuleID string `json:"resolverRuleID" tf:"resolver_rule_id"`
	VpcID          string `json:"vpcID" tf:"vpc_id"`
}



type Route53ResolverRuleAssociationStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *Route53ResolverRuleAssociationSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Route53ResolverRuleAssociationList is a list of Route53ResolverRuleAssociations
type Route53ResolverRuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route53ResolverRuleAssociation CRD objects
	Items []Route53ResolverRuleAssociation `json:"items,omitempty"`
}