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

type DefaultVpc struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultVpcSpec   `json:"spec,omitempty"`
	Status            DefaultVpcStatus `json:"status,omitempty"`
}

type DefaultVpcSpec struct {
	// +optional
	EnableDnsSupport bool `json:"enable_dns_support,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DefaultVpcStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DefaultVpcList is a list of DefaultVpcs
type DefaultVpcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DefaultVpc CRD objects
	Items []DefaultVpc `json:"items,omitempty"`
}
