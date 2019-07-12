package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServiceDiscoveryPrivateDnsNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsServiceDiscoveryPrivateDnsNamespaceSpec   `json:"spec,omitempty"`
	Status            AwsServiceDiscoveryPrivateDnsNamespaceStatus `json:"status,omitempty"`
}

type AwsServiceDiscoveryPrivateDnsNamespaceSpec struct {
	Description string `json:"description"`
	Vpc         string `json:"vpc"`
	Arn         string `json:"arn"`
	HostedZone  string `json:"hosted_zone"`
	Name        string `json:"name"`
}



type AwsServiceDiscoveryPrivateDnsNamespaceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsServiceDiscoveryPrivateDnsNamespaceList is a list of AwsServiceDiscoveryPrivateDnsNamespaces
type AwsServiceDiscoveryPrivateDnsNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsServiceDiscoveryPrivateDnsNamespace CRD objects
	Items []AwsServiceDiscoveryPrivateDnsNamespace `json:"items,omitempty"`
}