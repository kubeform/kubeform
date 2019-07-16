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

type EcrRepositoryPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EcrRepositoryPolicySpec   `json:"spec,omitempty"`
	Status            EcrRepositoryPolicyStatus `json:"status,omitempty"`
}

type EcrRepositoryPolicySpec struct {
	Policy     string `json:"policy"`
	Repository string `json:"repository"`
}

type EcrRepositoryPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EcrRepositoryPolicyList is a list of EcrRepositoryPolicys
type EcrRepositoryPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EcrRepositoryPolicy CRD objects
	Items []EcrRepositoryPolicy `json:"items,omitempty"`
}
