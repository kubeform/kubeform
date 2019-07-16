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

type NetworkPacketCapture struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkPacketCaptureSpec   `json:"spec,omitempty"`
	Status            NetworkPacketCaptureStatus `json:"status,omitempty"`
}

type NetworkPacketCaptureSpecFilter struct {
	// +optional
	LocalIpAddress string `json:"local_ip_address,omitempty"`
	// +optional
	LocalPort string `json:"local_port,omitempty"`
	Protocol  string `json:"protocol"`
	// +optional
	RemoteIpAddress string `json:"remote_ip_address,omitempty"`
	// +optional
	RemotePort string `json:"remote_port,omitempty"`
}

type NetworkPacketCaptureSpecStorageLocation struct {
	// +optional
	FilePath string `json:"file_path,omitempty"`
	// +optional
	StorageAccountId string `json:"storage_account_id,omitempty"`
}

type NetworkPacketCaptureSpec struct {
	// +optional
	Filter *[]NetworkPacketCaptureSpec `json:"filter,omitempty"`
	// +optional
	MaximumBytesPerPacket int `json:"maximum_bytes_per_packet,omitempty"`
	// +optional
	MaximumBytesPerSession int `json:"maximum_bytes_per_session,omitempty"`
	// +optional
	MaximumCaptureDuration int    `json:"maximum_capture_duration,omitempty"`
	Name                   string `json:"name"`
	NetworkWatcherName     string `json:"network_watcher_name"`
	ResourceGroupName      string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	StorageLocation  []NetworkPacketCaptureSpec `json:"storage_location"`
	TargetResourceId string                     `json:"target_resource_id"`
}

type NetworkPacketCaptureStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
