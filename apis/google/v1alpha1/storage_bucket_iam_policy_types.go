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

type StorageBucketIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageBucketIamPolicySpec   `json:"spec,omitempty"`
	Status            StorageBucketIamPolicyStatus `json:"status,omitempty"`
}

type StorageBucketIamPolicySpec struct {
	Bucket     string `json:"bucket"`
	PolicyData string `json:"policy_data"`
}

type StorageBucketIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageBucketIamPolicyList is a list of StorageBucketIamPolicys
type StorageBucketIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageBucketIamPolicy CRD objects
	Items []StorageBucketIamPolicy `json:"items,omitempty"`
}
