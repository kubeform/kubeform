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

type AzurermDevTestPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDevTestPolicySpec   `json:"spec,omitempty"`
	Status            AzurermDevTestPolicyStatus `json:"status,omitempty"`
}

type AzurermDevTestPolicySpec struct {
	ResourceGroupName string            `json:"resource_group_name"`
	FactData          string            `json:"fact_data"`
	Tags              map[string]string `json:"tags"`
	EvaluatorType     string            `json:"evaluator_type"`
	Description       string            `json:"description"`
	Name              string            `json:"name"`
	PolicySetName     string            `json:"policy_set_name"`
	LabName           string            `json:"lab_name"`
	Threshold         string            `json:"threshold"`
}



type AzurermDevTestPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDevTestPolicyList is a list of AzurermDevTestPolicys
type AzurermDevTestPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDevTestPolicy CRD objects
	Items []AzurermDevTestPolicy `json:"items,omitempty"`
}