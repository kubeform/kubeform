package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ElasticsearchDomainPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticsearchDomainPolicySpec   `json:"spec,omitempty"`
	Status            ElasticsearchDomainPolicyStatus `json:"status,omitempty"`
}

type ElasticsearchDomainPolicySpec struct {
	AccessPolicies string                    `json:"accessPolicies" tf:"access_policies"`
	DomainName     string                    `json:"domainName" tf:"domain_name"`
	ProviderRef    core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ElasticsearchDomainPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticsearchDomainPolicyList is a list of ElasticsearchDomainPolicys
type ElasticsearchDomainPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticsearchDomainPolicy CRD objects
	Items []ElasticsearchDomainPolicy `json:"items,omitempty"`
}
