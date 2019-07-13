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

type AwsSesActiveReceiptRuleSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesActiveReceiptRuleSetSpec   `json:"spec,omitempty"`
	Status            AwsSesActiveReceiptRuleSetStatus `json:"status,omitempty"`
}

type AwsSesActiveReceiptRuleSetSpec struct {
	RuleSetName string `json:"rule_set_name"`
}



type AwsSesActiveReceiptRuleSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSesActiveReceiptRuleSetList is a list of AwsSesActiveReceiptRuleSets
type AwsSesActiveReceiptRuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesActiveReceiptRuleSet CRD objects
	Items []AwsSesActiveReceiptRuleSet `json:"items,omitempty"`
}