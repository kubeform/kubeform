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

type LbNatPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbNatPoolSpec   `json:"spec,omitempty"`
	Status            LbNatPoolStatus `json:"status,omitempty"`
}

type LbNatPoolSpec struct {
	BackendPort                 int    `json:"backend_port"`
	FrontendIpConfigurationName string `json:"frontend_ip_configuration_name"`
	FrontendPortEnd             int    `json:"frontend_port_end"`
	FrontendPortStart           int    `json:"frontend_port_start"`
	LoadbalancerId              string `json:"loadbalancer_id"`
	// +optional
	// Deprecated
	Location          string `json:"location,omitempty"`
	Name              string `json:"name"`
	Protocol          string `json:"protocol"`
	ResourceGroupName string `json:"resource_group_name"`
}

type LbNatPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbNatPoolList is a list of LbNatPools
type LbNatPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbNatPool CRD objects
	Items []LbNatPool `json:"items,omitempty"`
}
