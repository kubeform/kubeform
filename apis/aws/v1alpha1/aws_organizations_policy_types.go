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

type AwsOrganizationsPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOrganizationsPolicySpec   `json:"spec,omitempty"`
	Status            AwsOrganizationsPolicyStatus `json:"status,omitempty"`
}

type AwsOrganizationsPolicySpec struct {
	Arn         string `json:"arn"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}



type AwsOrganizationsPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOrganizationsPolicyList is a list of AwsOrganizationsPolicys
type AwsOrganizationsPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOrganizationsPolicy CRD objects
	Items []AwsOrganizationsPolicy `json:"items,omitempty"`
}