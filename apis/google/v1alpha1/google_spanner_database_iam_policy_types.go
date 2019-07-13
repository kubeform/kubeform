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

type GoogleSpannerDatabaseIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSpannerDatabaseIamPolicySpec   `json:"spec,omitempty"`
	Status            GoogleSpannerDatabaseIamPolicyStatus `json:"status,omitempty"`
}

type GoogleSpannerDatabaseIamPolicySpec struct {
	PolicyData string `json:"policy_data"`
	Etag       string `json:"etag"`
	Database   string `json:"database"`
	Project    string `json:"project"`
	Instance   string `json:"instance"`
}



type GoogleSpannerDatabaseIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleSpannerDatabaseIamPolicyList is a list of GoogleSpannerDatabaseIamPolicys
type GoogleSpannerDatabaseIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSpannerDatabaseIamPolicy CRD objects
	Items []GoogleSpannerDatabaseIamPolicy `json:"items,omitempty"`
}