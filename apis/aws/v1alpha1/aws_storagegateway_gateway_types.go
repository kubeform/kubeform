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

type AwsStoragegatewayGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsStoragegatewayGatewaySpec   `json:"spec,omitempty"`
	Status            AwsStoragegatewayGatewayStatus `json:"status,omitempty"`
}

type AwsStoragegatewayGatewaySpecSmbActiveDirectorySettings struct {
	DomainName string `json:"domain_name"`
	Password   string `json:"password"`
	Username   string `json:"username"`
}

type AwsStoragegatewayGatewaySpec struct {
	Arn                        string                         `json:"arn"`
	ActivationKey              string                         `json:"activation_key"`
	GatewayId                  string                         `json:"gateway_id"`
	GatewayIpAddress           string                         `json:"gateway_ip_address"`
	GatewayTimezone            string                         `json:"gateway_timezone"`
	GatewayType                string                         `json:"gateway_type"`
	MediumChangerType          string                         `json:"medium_changer_type"`
	SmbActiveDirectorySettings []AwsStoragegatewayGatewaySpec `json:"smb_active_directory_settings"`
	GatewayName                string                         `json:"gateway_name"`
	SmbGuestPassword           string                         `json:"smb_guest_password"`
	TapeDriveType              string                         `json:"tape_drive_type"`
}



type AwsStoragegatewayGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsStoragegatewayGatewayList is a list of AwsStoragegatewayGateways
type AwsStoragegatewayGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsStoragegatewayGateway CRD objects
	Items []AwsStoragegatewayGateway `json:"items,omitempty"`
}