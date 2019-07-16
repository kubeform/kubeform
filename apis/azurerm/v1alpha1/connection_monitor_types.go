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

type ConnectionMonitor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionMonitorSpec   `json:"spec,omitempty"`
	Status            ConnectionMonitorStatus `json:"status,omitempty"`
}

type ConnectionMonitorSpecDestination struct {
	// +optional
	Address string `json:"address,omitempty"`
	Port    int    `json:"port"`
	// +optional
	VirtualMachineId string `json:"virtual_machine_id,omitempty"`
}

type ConnectionMonitorSpecSource struct {
	// +optional
	Port             int    `json:"port,omitempty"`
	VirtualMachineId string `json:"virtual_machine_id"`
}

type ConnectionMonitorSpec struct {
	// +optional
	AutoStart bool `json:"auto_start,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Destination []ConnectionMonitorSpec `json:"destination"`
	// +optional
	IntervalInSeconds  int    `json:"interval_in_seconds,omitempty"`
	Location           string `json:"location"`
	Name               string `json:"name"`
	NetworkWatcherName string `json:"network_watcher_name"`
	ResourceGroupName  string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Source []ConnectionMonitorSpec `json:"source"`
}

type ConnectionMonitorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ConnectionMonitorList is a list of ConnectionMonitors
type ConnectionMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConnectionMonitor CRD objects
	Items []ConnectionMonitor `json:"items,omitempty"`
}
