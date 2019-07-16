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

type NetworkConnectionMonitor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkConnectionMonitorSpec   `json:"spec,omitempty"`
	Status            NetworkConnectionMonitorStatus `json:"status,omitempty"`
}

type NetworkConnectionMonitorSpecDestination struct {
	// +optional
	Address string `json:"address,omitempty"`
	Port    int    `json:"port"`
	// +optional
	VirtualMachineId string `json:"virtual_machine_id,omitempty"`
}

type NetworkConnectionMonitorSpecSource struct {
	// +optional
	Port             int    `json:"port,omitempty"`
	VirtualMachineId string `json:"virtual_machine_id"`
}

type NetworkConnectionMonitorSpec struct {
	// +optional
	AutoStart bool `json:"auto_start,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Destination []NetworkConnectionMonitorSpec `json:"destination"`
	// +optional
	IntervalInSeconds  int    `json:"interval_in_seconds,omitempty"`
	Location           string `json:"location"`
	Name               string `json:"name"`
	NetworkWatcherName string `json:"network_watcher_name"`
	ResourceGroupName  string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Source []NetworkConnectionMonitorSpec `json:"source"`
}

type NetworkConnectionMonitorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
