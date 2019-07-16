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

type ComputeSubnetworkIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSubnetworkIamPolicySpec   `json:"spec,omitempty"`
	Status            ComputeSubnetworkIamPolicyStatus `json:"status,omitempty"`
}

type ComputeSubnetworkIamPolicySpec struct {
	PolicyData string `json:"policy_data"`
	// Deprecated
	Subnetwork string `json:"subnetwork"`
}

type ComputeSubnetworkIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSubnetworkIamPolicyList is a list of ComputeSubnetworkIamPolicys
type ComputeSubnetworkIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeSubnetworkIamPolicy CRD objects
	Items []ComputeSubnetworkIamPolicy `json:"items,omitempty"`
}
