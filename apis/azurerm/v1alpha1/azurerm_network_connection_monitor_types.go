package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermNetworkConnectionMonitor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNetworkConnectionMonitorSpec   `json:"spec,omitempty"`
	Status            AzurermNetworkConnectionMonitorStatus `json:"status,omitempty"`
}

type AzurermNetworkConnectionMonitorSpecSource struct {
	VirtualMachineId string `json:"virtual_machine_id"`
	Port             int    `json:"port"`
}

type AzurermNetworkConnectionMonitorSpecDestination struct {
	Address          string `json:"address"`
	Port             int    `json:"port"`
	VirtualMachineId string `json:"virtual_machine_id"`
}

type AzurermNetworkConnectionMonitorSpec struct {
	Name               string                                `json:"name"`
	AutoStart          bool                                  `json:"auto_start"`
	IntervalInSeconds  int                                   `json:"interval_in_seconds"`
	Source             []AzurermNetworkConnectionMonitorSpec `json:"source"`
	Destination        []AzurermNetworkConnectionMonitorSpec `json:"destination"`
	ResourceGroupName  string                                `json:"resource_group_name"`
	NetworkWatcherName string                                `json:"network_watcher_name"`
	Location           string                                `json:"location"`
	Tags               map[string]string                     `json:"tags"`
}



type AzurermNetworkConnectionMonitorStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermNetworkConnectionMonitorList is a list of AzurermNetworkConnectionMonitors
type AzurermNetworkConnectionMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkConnectionMonitor CRD objects
	Items []AzurermNetworkConnectionMonitor `json:"items,omitempty"`
}