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

type GoogleSpannerInstanceIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSpannerInstanceIamPolicySpec   `json:"spec,omitempty"`
	Status            GoogleSpannerInstanceIamPolicyStatus `json:"status,omitempty"`
}

type GoogleSpannerInstanceIamPolicySpec struct {
	Project    string `json:"project"`
	PolicyData string `json:"policy_data"`
	Etag       string `json:"etag"`
	Instance   string `json:"instance"`
}



type GoogleSpannerInstanceIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleSpannerInstanceIamPolicyList is a list of GoogleSpannerInstanceIamPolicys
type GoogleSpannerInstanceIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSpannerInstanceIamPolicy CRD objects
	Items []GoogleSpannerInstanceIamPolicy `json:"items,omitempty"`
}