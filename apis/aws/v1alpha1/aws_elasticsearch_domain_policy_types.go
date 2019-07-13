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

type AwsElasticsearchDomainPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElasticsearchDomainPolicySpec   `json:"spec,omitempty"`
	Status            AwsElasticsearchDomainPolicyStatus `json:"status,omitempty"`
}

type AwsElasticsearchDomainPolicySpec struct {
	DomainName     string `json:"domain_name"`
	AccessPolicies string `json:"access_policies"`
}



type AwsElasticsearchDomainPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsElasticsearchDomainPolicyList is a list of AwsElasticsearchDomainPolicys
type AwsElasticsearchDomainPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticsearchDomainPolicy CRD objects
	Items []AwsElasticsearchDomainPolicy `json:"items,omitempty"`
}