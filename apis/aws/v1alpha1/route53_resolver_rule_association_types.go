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

type Route53ResolverRuleAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Route53ResolverRuleAssociationSpec   `json:"spec,omitempty"`
	Status            Route53ResolverRuleAssociationStatus `json:"status,omitempty"`
}

type Route53ResolverRuleAssociationSpec struct {
	// +optional
	Name           string                    `json:"name,omitempty" tf:"name,omitempty"`
	ResolverRuleID string                    `json:"resolverRuleID" tf:"resolver_rule_id"`
	VpcID          string                    `json:"vpcID" tf:"vpc_id"`
	ProviderRef    core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type Route53ResolverRuleAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
