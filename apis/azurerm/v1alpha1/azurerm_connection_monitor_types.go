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

type AzurermConnectionMonitor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermConnectionMonitorSpec   `json:"spec,omitempty"`
	Status            AzurermConnectionMonitorStatus `json:"status,omitempty"`
}

type AzurermConnectionMonitorSpecSource struct {
	VirtualMachineId string `json:"virtual_machine_id"`
	Port             int    `json:"port"`
}

type AzurermConnectionMonitorSpecDestination struct {
	VirtualMachineId string `json:"virtual_machine_id"`
	Address          string `json:"address"`
	Port             int    `json:"port"`
}

type AzurermConnectionMonitorSpec struct {
	AutoStart          bool                           `json:"auto_start"`
	IntervalInSeconds  int                            `json:"interval_in_seconds"`
	Source             []AzurermConnectionMonitorSpec `json:"source"`
	Destination        []AzurermConnectionMonitorSpec `json:"destination"`
	Tags               map[string]string              `json:"tags"`
	ResourceGroupName  string                         `json:"resource_group_name"`
	Location           string                         `json:"location"`
	Name               string                         `json:"name"`
	NetworkWatcherName string                         `json:"network_watcher_name"`
}

type AzurermConnectionMonitorStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermConnectionMonitorList is a list of AzurermConnectionMonitors
type AzurermConnectionMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermConnectionMonitor CRD objects
	Items []AzurermConnectionMonitor `json:"items,omitempty"`
}
