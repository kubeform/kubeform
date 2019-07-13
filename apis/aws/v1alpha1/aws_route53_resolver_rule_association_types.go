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

type AwsRoute53ResolverRuleAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRoute53ResolverRuleAssociationSpec   `json:"spec,omitempty"`
	Status            AwsRoute53ResolverRuleAssociationStatus `json:"status,omitempty"`
}

type AwsRoute53ResolverRuleAssociationSpec struct {
	ResolverRuleId string `json:"resolver_rule_id"`
	VpcId          string `json:"vpc_id"`
	Name           string `json:"name"`
}



type AwsRoute53ResolverRuleAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRoute53ResolverRuleAssociationList is a list of AwsRoute53ResolverRuleAssociations
type AwsRoute53ResolverRuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRoute53ResolverRuleAssociation CRD objects
	Items []AwsRoute53ResolverRuleAssociation `json:"items,omitempty"`
}