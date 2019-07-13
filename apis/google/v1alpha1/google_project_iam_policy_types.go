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

type GoogleProjectIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleProjectIamPolicySpec   `json:"spec,omitempty"`
	Status            GoogleProjectIamPolicyStatus `json:"status,omitempty"`
}

type GoogleProjectIamPolicySpec struct {
	Authoritative  bool   `json:"authoritative"`
	Etag           string `json:"etag"`
	RestorePolicy  string `json:"restore_policy"`
	DisableProject bool   `json:"disable_project"`
	Project        string `json:"project"`
	PolicyData     string `json:"policy_data"`
}



type GoogleProjectIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleProjectIamPolicyList is a list of GoogleProjectIamPolicys
type GoogleProjectIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleProjectIamPolicy CRD objects
	Items []GoogleProjectIamPolicy `json:"items,omitempty"`
}