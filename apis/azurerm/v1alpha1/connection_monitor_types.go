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

type ConnectionMonitor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionMonitorSpec   `json:"spec,omitempty"`
	Status            ConnectionMonitorStatus `json:"status,omitempty"`
}

type ConnectionMonitorSpecDestination struct {
	// +optional
	Address string `json:"address,omitempty" tf:"address,omitempty"`
	Port    int    `json:"port" tf:"port"`
	// +optional
	VirtualMachineID string `json:"virtualMachineID,omitempty" tf:"virtual_machine_id,omitempty"`
}

type ConnectionMonitorSpecSource struct {
	// +optional
	Port             int    `json:"port,omitempty" tf:"port,omitempty"`
	VirtualMachineID string `json:"virtualMachineID" tf:"virtual_machine_id"`
}

type ConnectionMonitorSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoStart bool `json:"autoStart,omitempty" tf:"auto_start,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Destination []ConnectionMonitorSpecDestination `json:"destination" tf:"destination"`
	// +optional
	IntervalInSeconds  int    `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`
	Location           string `json:"location" tf:"location"`
	Name               string `json:"name" tf:"name"`
	NetworkWatcherName string `json:"networkWatcherName" tf:"network_watcher_name"`
	ResourceGroupName  string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Source []ConnectionMonitorSpecSource `json:"source" tf:"source"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConnectionMonitorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ConnectionMonitorSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
