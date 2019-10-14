package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type LbProbe struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbProbeSpec   `json:"spec,omitempty"`
	Status            LbProbeStatus `json:"status,omitempty"`
}

type LbProbeSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	IntervalInSeconds int `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`
	// +optional
	LoadBalancerRules []string `json:"loadBalancerRules,omitempty" tf:"load_balancer_rules,omitempty"`
	LoadbalancerID    string   `json:"loadbalancerID" tf:"loadbalancer_id"`
	// +optional
	// Deprecated
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	Name     string `json:"name" tf:"name"`
	// +optional
	NumberOfProbes int `json:"numberOfProbes,omitempty" tf:"number_of_probes,omitempty"`
	Port           int `json:"port" tf:"port"`
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// +optional
	RequestPath       string `json:"requestPath,omitempty" tf:"request_path,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

type LbProbeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LbProbeSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbProbeList is a list of LbProbes
type LbProbeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbProbe CRD objects
	Items []LbProbe `json:"items,omitempty"`
}
