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

type AzurermNetworkPacketCapture struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNetworkPacketCaptureSpec   `json:"spec,omitempty"`
	Status            AzurermNetworkPacketCaptureStatus `json:"status,omitempty"`
}

type AzurermNetworkPacketCaptureSpecStorageLocation struct {
	FilePath         string `json:"file_path"`
	StorageAccountId string `json:"storage_account_id"`
	StoragePath      string `json:"storage_path"`
}

type AzurermNetworkPacketCaptureSpecFilter struct {
	RemotePort      string `json:"remote_port"`
	LocalIpAddress  string `json:"local_ip_address"`
	LocalPort       string `json:"local_port"`
	Protocol        string `json:"protocol"`
	RemoteIpAddress string `json:"remote_ip_address"`
}

type AzurermNetworkPacketCaptureSpec struct {
	StorageLocation        []AzurermNetworkPacketCaptureSpec `json:"storage_location"`
	NetworkWatcherName     string                            `json:"network_watcher_name"`
	MaximumBytesPerPacket  int                               `json:"maximum_bytes_per_packet"`
	MaximumBytesPerSession int                               `json:"maximum_bytes_per_session"`
	MaximumCaptureDuration int                               `json:"maximum_capture_duration"`
	Name                   string                            `json:"name"`
	ResourceGroupName      string                            `json:"resource_group_name"`
	TargetResourceId       string                            `json:"target_resource_id"`
	Filter                 []AzurermNetworkPacketCaptureSpec `json:"filter"`
}

type AzurermNetworkPacketCaptureStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermNetworkPacketCaptureList is a list of AzurermNetworkPacketCaptures
type AzurermNetworkPacketCaptureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkPacketCapture CRD objects
	Items []AzurermNetworkPacketCapture `json:"items,omitempty"`
}
