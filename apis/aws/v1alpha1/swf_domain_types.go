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

type SwfDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SwfDomainSpec   `json:"spec,omitempty"`
	Status            SwfDomainStatus `json:"status,omitempty"`
}

type SwfDomainSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	NamePrefix                             string `json:"name_prefix,omitempty"`
	WorkflowExecutionRetentionPeriodInDays string `json:"workflow_execution_retention_period_in_days"`
}

type SwfDomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
