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
	FrontendPortEnd             int    `json:"frontend_port_end"`
	BackendPort                 int    `json:"backend_port"`
	FrontendIpConfigurationName string `json:"frontend_ip_configuration_name"`
	Name                        string `json:"name"`
	ResourceGroupName           string `json:"resource_group_name"`
	LoadbalancerId              string `json:"loadbalancer_id"`
	FrontendIpConfigurationId   string `json:"frontend_ip_configuration_id"`
	Location                    string `json:"location"`
	Protocol                    string `json:"protocol"`
	FrontendPortStart           int    `json:"frontend_port_start"`
}



type AzurermLbNatPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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