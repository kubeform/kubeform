package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermPublicIpPrefix struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPublicIpPrefixSpec   `json:"spec,omitempty"`
	Status            AzurermPublicIpPrefixStatus `json:"status,omitempty"`
}

type AzurermPublicIpPrefixSpec struct {
	Location          string            `json:"location"`
	ResourceGroupName string            `json:"resource_group_name"`
	Sku               string            `json:"sku"`
	PrefixLength      int               `json:"prefix_length"`
	IpPrefix          string            `json:"ip_prefix"`
	Zones             []string          `json:"zones"`
	Tags              map[string]string `json:"tags"`
	Name              string            `json:"name"`
}

type AzurermPublicIpPrefixStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermPublicIpPrefixList is a list of AzurermPublicIpPrefixs
type AzurermPublicIpPrefixList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPublicIpPrefix CRD objects
	Items []AzurermPublicIpPrefix `json:"items,omitempty"`
}
