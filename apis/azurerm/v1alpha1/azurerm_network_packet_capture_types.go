package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermNetworkPacketCapture struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNetworkPacketCaptureSpec   `json:"spec,omitempty"`
	Status            AzurermNetworkPacketCaptureStatus `json:"status,omitempty"`
}

type AzurermNetworkPacketCaptureSpecFilter struct {
	Protocol        string `json:"protocol"`
	RemoteIpAddress string `json:"remote_ip_address"`
	RemotePort      string `json:"remote_port"`
	LocalIpAddress  string `json:"local_ip_address"`
	LocalPort       string `json:"local_port"`
}

type AzurermNetworkPacketCaptureSpecStorageLocation struct {
	FilePath         string `json:"file_path"`
	StorageAccountId string `json:"storage_account_id"`
	StoragePath      string `json:"storage_path"`
}

type AzurermNetworkPacketCaptureSpec struct {
	ResourceGroupName      string                            `json:"resource_group_name"`
	Filter                 []AzurermNetworkPacketCaptureSpec `json:"filter"`
	Name                   string                            `json:"name"`
	NetworkWatcherName     string                            `json:"network_watcher_name"`
	TargetResourceId       string                            `json:"target_resource_id"`
	MaximumBytesPerPacket  int                               `json:"maximum_bytes_per_packet"`
	MaximumBytesPerSession int                               `json:"maximum_bytes_per_session"`
	MaximumCaptureDuration int                               `json:"maximum_capture_duration"`
	StorageLocation        []AzurermNetworkPacketCaptureSpec `json:"storage_location"`
}

type AzurermNetworkPacketCaptureStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermNetworkPacketCaptureList is a list of AzurermNetworkPacketCaptures
type AzurermNetworkPacketCaptureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkPacketCapture CRD objects
	Items []AzurermNetworkPacketCapture `json:"items,omitempty"`
}
