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

type DevTestPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevTestPolicySpec   `json:"spec,omitempty"`
	Status            DevTestPolicyStatus `json:"status,omitempty"`
}

type DevTestPolicySpec struct {
	// +optional
	Description   string `json:"description,omitempty"`
	EvaluatorType string `json:"evaluator_type"`
	// +optional
	FactData          string `json:"fact_data,omitempty"`
	LabName           string `json:"lab_name"`
	Name              string `json:"name"`
	PolicySetName     string `json:"policy_set_name"`
	ResourceGroupName string `json:"resource_group_name"`
	Threshold         string `json:"threshold"`
}

type DevTestPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DevTestPolicyList is a list of DevTestPolicys
type DevTestPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DevTestPolicy CRD objects
	Items []DevTestPolicy `json:"items,omitempty"`
}
