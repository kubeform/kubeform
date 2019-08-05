package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type NetworkPacketCapture struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkPacketCaptureSpec   `json:"spec,omitempty"`
	Status            NetworkPacketCaptureStatus `json:"status,omitempty"`
}

type NetworkPacketCaptureSpecFilter struct {
	// +optional
	LocalIPAddress string `json:"localIPAddress,omitempty" tf:"local_ip_address,omitempty"`
	// +optional
	LocalPort string `json:"localPort,omitempty" tf:"local_port,omitempty"`
	Protocol  string `json:"protocol" tf:"protocol"`
	// +optional
	RemoteIPAddress string `json:"remoteIPAddress,omitempty" tf:"remote_ip_address,omitempty"`
	// +optional
	RemotePort string `json:"remotePort,omitempty" tf:"remote_port,omitempty"`
}

type NetworkPacketCaptureSpecStorageLocation struct {
	// +optional
	FilePath string `json:"filePath,omitempty" tf:"file_path,omitempty"`
	// +optional
	StorageAccountID string `json:"storageAccountID,omitempty" tf:"storage_account_id,omitempty"`
}

type NetworkPacketCaptureSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Filter []NetworkPacketCaptureSpecFilter `json:"filter,omitempty" tf:"filter,omitempty"`
	// +optional
	MaximumBytesPerPacket int `json:"maximumBytesPerPacket,omitempty" tf:"maximum_bytes_per_packet,omitempty"`
	// +optional
	MaximumBytesPerSession int `json:"maximumBytesPerSession,omitempty" tf:"maximum_bytes_per_session,omitempty"`
	// +optional
	MaximumCaptureDuration int    `json:"maximumCaptureDuration,omitempty" tf:"maximum_capture_duration,omitempty"`
	Name                   string `json:"name" tf:"name"`
	NetworkWatcherName     string `json:"networkWatcherName" tf:"network_watcher_name"`
	ResourceGroupName      string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	StorageLocation  []NetworkPacketCaptureSpecStorageLocation `json:"storageLocation" tf:"storage_location"`
	TargetResourceID string                                    `json:"targetResourceID" tf:"target_resource_id"`
}

type NetworkPacketCaptureStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NetworkPacketCaptureSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkPacketCaptureList is a list of NetworkPacketCaptures
type NetworkPacketCaptureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkPacketCapture CRD objects
	Items []NetworkPacketCapture `json:"items,omitempty"`
}
