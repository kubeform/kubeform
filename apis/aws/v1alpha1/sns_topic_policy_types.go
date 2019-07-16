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

type SnsTopicPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnsTopicPolicySpec   `json:"spec,omitempty"`
	Status            SnsTopicPolicyStatus `json:"status,omitempty"`
}

type SnsTopicPolicySpec struct {
	Arn    string `json:"arn"`
	Policy string `json:"policy"`
}

type SnsTopicPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SnsTopicPolicyList is a list of SnsTopicPolicys
type SnsTopicPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SnsTopicPolicy CRD objects
	Items []SnsTopicPolicy `json:"items,omitempty"`
}
