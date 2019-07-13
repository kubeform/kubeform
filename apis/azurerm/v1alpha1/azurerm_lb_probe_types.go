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
	IntervalInSeconds int      `json:"interval_in_seconds"`
	Location          string   `json:"location"`
	ResourceGroupName string   `json:"resource_group_name"`
	Port              int      `json:"port"`
	RequestPath       string   `json:"request_path"`
	NumberOfProbes    int      `json:"number_of_probes"`
	LoadBalancerRules []string `json:"load_balancer_rules"`
	Name              string   `json:"name"`
	LoadbalancerId    string   `json:"loadbalancer_id"`
	Protocol          string   `json:"protocol"`
}



type AzurermLbProbeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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