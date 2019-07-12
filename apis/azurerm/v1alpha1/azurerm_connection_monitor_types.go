package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	NetworkWatcherName string                         `json:"network_watcher_name"`
	Location           string                         `json:"location"`
	Source             []AzurermConnectionMonitorSpec `json:"source"`
	Name               string                         `json:"name"`
	AutoStart          bool                           `json:"auto_start"`
	IntervalInSeconds  int                            `json:"interval_in_seconds"`
	Destination        []AzurermConnectionMonitorSpec `json:"destination"`
	Tags               map[string]string              `json:"tags"`
	ResourceGroupName  string                         `json:"resource_group_name"`
}

type AzurermConnectionMonitorStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermConnectionMonitorList is a list of AzurermConnectionMonitors
type AzurermConnectionMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermConnectionMonitor CRD objects
	Items []AzurermConnectionMonitor `json:"items,omitempty"`
}
