package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryPublicDnsNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsServiceDiscoveryPublicDnsNamespaceSpec   `json:"spec,omitempty"`
	Status            AwsServiceDiscoveryPublicDnsNamespaceStatus `json:"status,omitempty"`
}

type AwsServiceDiscoveryPublicDnsNamespaceSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Arn         string `json:"arn"`
	HostedZone  string `json:"hosted_zone"`
}



type AwsServiceDiscoveryPublicDnsNamespaceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsServiceDiscoveryPublicDnsNamespaceList is a list of AwsServiceDiscoveryPublicDnsNamespaces
type AwsServiceDiscoveryPublicDnsNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsServiceDiscoveryPublicDnsNamespace CRD objects
	Items []AwsServiceDiscoveryPublicDnsNamespace `json:"items,omitempty"`
}