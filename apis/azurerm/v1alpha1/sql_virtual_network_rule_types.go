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

type SqlVirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlVirtualNetworkRuleSpec   `json:"spec,omitempty"`
	Status            SqlVirtualNetworkRuleStatus `json:"status,omitempty"`
}

type SqlVirtualNetworkRuleSpec struct {
	// +optional
	IgnoreMissingVnetServiceEndpoint bool   `json:"ignore_missing_vnet_service_endpoint,omitempty"`
	Name                             string `json:"name"`
	ResourceGroupName                string `json:"resource_group_name"`
	ServerName                       string `json:"server_name"`
	SubnetId                         string `json:"subnet_id"`
}

type SqlVirtualNetworkRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlVirtualNetworkRuleList is a list of SqlVirtualNetworkRules
type SqlVirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlVirtualNetworkRule CRD objects
	Items []SqlVirtualNetworkRule `json:"items,omitempty"`
}
