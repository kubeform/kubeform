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

type SwfDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SwfDomainSpec   `json:"spec,omitempty"`
	Status            SwfDomainStatus `json:"status,omitempty"`
}

type SwfDomainSpec struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix                             string                    `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	WorkflowExecutionRetentionPeriodInDays string                    `json:"workflowExecutionRetentionPeriodInDays" tf:"workflow_execution_retention_period_in_days"`
	ProviderRef                            core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SwfDomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SwfDomainList is a list of SwfDomains
type SwfDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SwfDomain CRD objects
	Items []SwfDomain `json:"items,omitempty"`
}
