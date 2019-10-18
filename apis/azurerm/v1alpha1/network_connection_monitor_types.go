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
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type NetworkConnectionMonitor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkConnectionMonitorSpec   `json:"spec,omitempty"`
	Status            NetworkConnectionMonitorStatus `json:"status,omitempty"`
}

type NetworkConnectionMonitorSpecDestination struct {
	// +optional
	Address string `json:"address,omitempty" tf:"address,omitempty"`
	Port    int    `json:"port" tf:"port"`
	// +optional
	VirtualMachineID string `json:"virtualMachineID,omitempty" tf:"virtual_machine_id,omitempty"`
}

type NetworkConnectionMonitorSpecSource struct {
	// +optional
	Port             int    `json:"port,omitempty" tf:"port,omitempty"`
	VirtualMachineID string `json:"virtualMachineID" tf:"virtual_machine_id"`
}

type NetworkConnectionMonitorSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoStart bool `json:"autoStart,omitempty" tf:"auto_start,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Destination []NetworkConnectionMonitorSpecDestination `json:"destination" tf:"destination"`
	// +optional
	IntervalInSeconds  int    `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`
	Location           string `json:"location" tf:"location"`
	Name               string `json:"name" tf:"name"`
	NetworkWatcherName string `json:"networkWatcherName" tf:"network_watcher_name"`
	ResourceGroupName  string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Source []NetworkConnectionMonitorSpecSource `json:"source" tf:"source"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NetworkConnectionMonitorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NetworkConnectionMonitorSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkConnectionMonitorList is a list of NetworkConnectionMonitors
type NetworkConnectionMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkConnectionMonitor CRD objects
	Items []NetworkConnectionMonitor `json:"items,omitempty"`
}
