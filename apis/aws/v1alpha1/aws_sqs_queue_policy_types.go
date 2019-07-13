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

type AwsSqsQueuePolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSqsQueuePolicySpec   `json:"spec,omitempty"`
	Status            AwsSqsQueuePolicyStatus `json:"status,omitempty"`
}

type AwsSqsQueuePolicySpec struct {
	QueueUrl string `json:"queue_url"`
	Policy   string `json:"policy"`
}



type AwsSqsQueuePolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSqsQueuePolicyList is a list of AwsSqsQueuePolicys
type AwsSqsQueuePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSqsQueuePolicy CRD objects
	Items []AwsSqsQueuePolicy `json:"items,omitempty"`
}