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

type AwsIotPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIotPolicySpec   `json:"spec,omitempty"`
	Status            AwsIotPolicyStatus `json:"status,omitempty"`
}

type AwsIotPolicySpec struct {
	DefaultVersionId string `json:"default_version_id"`
	Name             string `json:"name"`
	Policy           string `json:"policy"`
	Arn              string `json:"arn"`
}



type AwsIotPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIotPolicyList is a list of AwsIotPolicys
type AwsIotPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIotPolicy CRD objects
	Items []AwsIotPolicy `json:"items,omitempty"`
}