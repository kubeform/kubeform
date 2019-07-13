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

type AzurermPublicIpPrefix struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPublicIpPrefixSpec   `json:"spec,omitempty"`
	Status            AzurermPublicIpPrefixStatus `json:"status,omitempty"`
}

type AzurermPublicIpPrefixSpec struct {
	IpPrefix          string            `json:"ip_prefix"`
	Zones             []string          `json:"zones"`
	Tags              map[string]string `json:"tags"`
	Name              string            `json:"name"`
	Location          string            `json:"location"`
	ResourceGroupName string            `json:"resource_group_name"`
	Sku               string            `json:"sku"`
	PrefixLength      int               `json:"prefix_length"`
}



type AzurermPublicIpPrefixStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermPublicIpPrefixList is a list of AzurermPublicIpPrefixs
type AzurermPublicIpPrefixList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPublicIpPrefix CRD objects
	Items []AzurermPublicIpPrefix `json:"items,omitempty"`
}