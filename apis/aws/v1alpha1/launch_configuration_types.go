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

type LaunchConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LaunchConfigurationSpec   `json:"spec,omitempty"`
	Status            LaunchConfigurationStatus `json:"status,omitempty"`
}

type LaunchConfigurationSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type LaunchConfigurationSpec struct {
	// +optional
	AssociatePublicIpAddress bool `json:"associate_public_ip_address,omitempty"`
	// +optional
	EnableMonitoring bool `json:"enable_monitoring,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EphemeralBlockDevice *[]LaunchConfigurationSpec `json:"ephemeral_block_device,omitempty"`
	// +optional
	IamInstanceProfile string `json:"iam_instance_profile,omitempty"`
	ImageId            string `json:"image_id"`
	InstanceType       string `json:"instance_type"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	PlacementTenancy string `json:"placement_tenancy,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"security_groups,omitempty"`
	// +optional
	SpotPrice string `json:"spot_price,omitempty"`
	// +optional
	UserData string `json:"user_data,omitempty"`
	// +optional
	UserDataBase64 string `json:"user_data_base64,omitempty"`
	// +optional
	VpcClassicLinkId string `json:"vpc_classic_link_id,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VpcClassicLinkSecurityGroups []string `json:"vpc_classic_link_security_groups,omitempty"`
}

type LaunchConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LaunchConfigurationList is a list of LaunchConfigurations
type LaunchConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LaunchConfiguration CRD objects
	Items []LaunchConfiguration `json:"items,omitempty"`
}
