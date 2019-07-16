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

type ServiceAccountIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountIamPolicySpec   `json:"spec,omitempty"`
	Status            ServiceAccountIamPolicyStatus `json:"status,omitempty"`
}

type ServiceAccountIamPolicySpec struct {
	PolicyData       string `json:"policy_data"`
	ServiceAccountId string `json:"service_account_id"`
}

type ServiceAccountIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServiceAccountIamPolicyList is a list of ServiceAccountIamPolicys
type ServiceAccountIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceAccountIamPolicy CRD objects
	Items []ServiceAccountIamPolicy `json:"items,omitempty"`
}
