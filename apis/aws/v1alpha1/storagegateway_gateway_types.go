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

type StoragegatewayGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StoragegatewayGatewaySpec   `json:"spec,omitempty"`
	Status            StoragegatewayGatewayStatus `json:"status,omitempty"`
}

type StoragegatewayGatewaySpecSmbActiveDirectorySettings struct {
	DomainName string `json:"domainName" tf:"domain_name"`
	Password   string `json:"-" sensitive:"true" tf:"password"`
	Username   string `json:"username" tf:"username"`
}

type StoragegatewayGatewaySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	KubeFormSecret *core.LocalObjectReference `json:"kubeFormSecret,omitempty" tf:"-"`

	// +optional
	ActivationKey string `json:"activationKey,omitempty" tf:"activation_key,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	GatewayID string `json:"gatewayID,omitempty" tf:"gateway_id,omitempty"`
	// +optional
	GatewayIPAddress string `json:"gatewayIPAddress,omitempty" tf:"gateway_ip_address,omitempty"`
	GatewayName      string `json:"gatewayName" tf:"gateway_name"`
	GatewayTimezone  string `json:"gatewayTimezone" tf:"gateway_timezone"`
	// +optional
	GatewayType string `json:"gatewayType,omitempty" tf:"gateway_type,omitempty"`
	// +optional
	MediumChangerType string `json:"mediumChangerType,omitempty" tf:"medium_changer_type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SmbActiveDirectorySettings []StoragegatewayGatewaySpecSmbActiveDirectorySettings `json:"smbActiveDirectorySettings,omitempty" tf:"smb_active_directory_settings,omitempty"`
	// +optional
	SmbGuestPassword string `json:"-" sensitive:"true" tf:"smb_guest_password,omitempty"`
	// +optional
	TapeDriveType string `json:"tapeDriveType,omitempty" tf:"tape_drive_type,omitempty"`
}

type StoragegatewayGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StoragegatewayGatewaySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
