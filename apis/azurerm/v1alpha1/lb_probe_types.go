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

type LbProbe struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbProbeSpec   `json:"spec,omitempty"`
	Status            LbProbeStatus `json:"status,omitempty"`
}

type LbProbeSpec struct {
	// +optional
	IntervalInSeconds int    `json:"interval_in_seconds,omitempty"`
	LoadbalancerId    string `json:"loadbalancer_id"`
	// +optional
	// Deprecated
	Location string `json:"location,omitempty"`
	Name     string `json:"name"`
	// +optional
	NumberOfProbes int `json:"number_of_probes,omitempty"`
	Port           int `json:"port"`
	// +optional
	RequestPath       string `json:"request_path,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
}

type LbProbeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
