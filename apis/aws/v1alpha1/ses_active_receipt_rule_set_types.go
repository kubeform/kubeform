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

type SesActiveReceiptRuleSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesActiveReceiptRuleSetSpec   `json:"spec,omitempty"`
	Status            SesActiveReceiptRuleSetStatus `json:"status,omitempty"`
}

type SesActiveReceiptRuleSetSpec struct {
	RuleSetName string                    `json:"ruleSetName" tf:"rule_set_name"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SesActiveReceiptRuleSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SesActiveReceiptRuleSetList is a list of SesActiveReceiptRuleSets
type SesActiveReceiptRuleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SesActiveReceiptRuleSet CRD objects
	Items []SesActiveReceiptRuleSet `json:"items,omitempty"`
}
