package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesReceiptRuleSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesReceiptRuleSetSpec   `json:"spec,omitempty"`
	Status            AwsSesReceiptRuleSetStatus `json:"status,omitempty"`
}

type AwsSesReceiptRuleSetSpec struct {
	RuleSetName string `json:"rule_set_name"`
}



type AwsSesReceiptRuleSetStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesReceiptRuleSetList is a list of AwsSesReceiptRuleSets
type AwsSesReceiptRuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesReceiptRuleSet CRD objects
	Items []AwsSesReceiptRuleSet `json:"items,omitempty"`
}