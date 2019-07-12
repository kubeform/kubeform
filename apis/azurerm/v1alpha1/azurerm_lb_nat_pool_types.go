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

type AzurermLbNatPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLbNatPoolSpec   `json:"spec,omitempty"`
	Status            AzurermLbNatPoolStatus `json:"status,omitempty"`
}

type AzurermLbNatPoolSpec struct {
	Protocol                    string `json:"protocol"`
	FrontendPortStart           int    `json:"frontend_port_start"`
	FrontendIpConfigurationName string `json:"frontend_ip_configuration_name"`
	FrontendIpConfigurationId   string `json:"frontend_ip_configuration_id"`
	Name                        string `json:"name"`
	Location                    string `json:"location"`
	FrontendPortEnd             int    `json:"frontend_port_end"`
	BackendPort                 int    `json:"backend_port"`
	ResourceGroupName           string `json:"resource_group_name"`
	LoadbalancerId              string `json:"loadbalancer_id"`
}

type AzurermLbNatPoolStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermLbNatPoolList is a list of AzurermLbNatPools
type AzurermLbNatPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLbNatPool CRD objects
	Items []AzurermLbNatPool `json:"items,omitempty"`
}
