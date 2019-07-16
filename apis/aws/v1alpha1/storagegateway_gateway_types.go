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

type StoragegatewayGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewayGatewaySpec   `json:"spec,omitempty"`
	Status            StoragegatewayGatewayStatus `json:"status,omitempty"`
}

type StoragegatewayGatewaySpecSmbActiveDirectorySettings struct {
	DomainName string `json:"domain_name"`
	Password   string `json:"password"`
	Username   string `json:"username"`
}

type StoragegatewayGatewaySpec struct {
	GatewayName     string `json:"gateway_name"`
	GatewayTimezone string `json:"gateway_timezone"`
	// +optional
	GatewayType string `json:"gateway_type,omitempty"`
	// +optional
	MediumChangerType string `json:"medium_changer_type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SmbActiveDirectorySettings *[]StoragegatewayGatewaySpec `json:"smb_active_directory_settings,omitempty"`
	// +optional
	SmbGuestPassword string `json:"smb_guest_password,omitempty"`
	// +optional
	TapeDriveType string `json:"tape_drive_type,omitempty"`
}

type StoragegatewayGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StoragegatewayGatewayList is a list of StoragegatewayGateways
type StoragegatewayGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StoragegatewayGateway CRD objects
	Items []StoragegatewayGateway `json:"items,omitempty"`
}
