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

type DevTestPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevTestPolicySpec   `json:"spec,omitempty"`
	Status            DevTestPolicyStatus `json:"status,omitempty"`
}

type DevTestPolicySpec struct {
	// +optional
	Description   string `json:"description,omitempty" tf:"description,omitempty"`
	EvaluatorType string `json:"evaluatorType" tf:"evaluator_type"`
	// +optional
	FactData          string `json:"factData,omitempty" tf:"fact_data,omitempty"`
	LabName           string `json:"labName" tf:"lab_name"`
	Name              string `json:"name" tf:"name"`
	PolicySetName     string `json:"policySetName" tf:"policy_set_name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	Threshold   string                    `json:"threshold" tf:"threshold"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DevTestPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
