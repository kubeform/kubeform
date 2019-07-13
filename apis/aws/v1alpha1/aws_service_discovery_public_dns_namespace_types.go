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

type AwsServiceDiscoveryPublicDnsNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsServiceDiscoveryPublicDnsNamespaceSpec   `json:"spec,omitempty"`
	Status            AwsServiceDiscoveryPublicDnsNamespaceStatus `json:"status,omitempty"`
}

type AwsServiceDiscoveryPublicDnsNamespaceSpec struct {
	Arn         string `json:"arn"`
	HostedZone  string `json:"hosted_zone"`
	Name        string `json:"name"`
	Description string `json:"description"`
}



type AwsServiceDiscoveryPublicDnsNamespaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsServiceDiscoveryPublicDnsNamespaceList is a list of AwsServiceDiscoveryPublicDnsNamespaces
type AwsServiceDiscoveryPublicDnsNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsServiceDiscoveryPublicDnsNamespace CRD objects
	Items []AwsServiceDiscoveryPublicDnsNamespace `json:"items,omitempty"`
}