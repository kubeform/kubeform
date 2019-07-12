package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermLbBackendAddressPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLbBackendAddressPoolSpec   `json:"spec,omitempty"`
	Status            AzurermLbBackendAddressPoolStatus `json:"status,omitempty"`
}

type AzurermLbBackendAddressPoolSpec struct {
	Name                    string   `json:"name"`
	Location                string   `json:"location"`
	ResourceGroupName       string   `json:"resource_group_name"`
	LoadbalancerId          string   `json:"loadbalancer_id"`
	BackendIpConfigurations []string `json:"backend_ip_configurations"`
	LoadBalancingRules      []string `json:"load_balancing_rules"`
}



type AzurermLbBackendAddressPoolStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermLbBackendAddressPoolList is a list of AzurermLbBackendAddressPools
type AzurermLbBackendAddressPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLbBackendAddressPool CRD objects
	Items []AzurermLbBackendAddressPool `json:"items,omitempty"`
}