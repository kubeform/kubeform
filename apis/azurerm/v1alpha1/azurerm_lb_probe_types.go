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

type AzurermLbProbe struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLbProbeSpec   `json:"spec,omitempty"`
	Status            AzurermLbProbeStatus `json:"status,omitempty"`
}

type AzurermLbProbeSpec struct {
	Name              string   `json:"name"`
	Location          string   `json:"location"`
	Protocol          string   `json:"protocol"`
	IntervalInSeconds int      `json:"interval_in_seconds"`
	NumberOfProbes    int      `json:"number_of_probes"`
	ResourceGroupName string   `json:"resource_group_name"`
	LoadbalancerId    string   `json:"loadbalancer_id"`
	Port              int      `json:"port"`
	RequestPath       string   `json:"request_path"`
	LoadBalancerRules []string `json:"load_balancer_rules"`
}

type AzurermLbProbeStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermLbProbeList is a list of AzurermLbProbes
type AzurermLbProbeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLbProbe CRD objects
	Items []AzurermLbProbe `json:"items,omitempty"`
}
