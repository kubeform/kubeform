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

type AwsLightsailInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLightsailInstanceSpec   `json:"spec,omitempty"`
	Status            AwsLightsailInstanceStatus `json:"status,omitempty"`
}

type AwsLightsailInstanceSpec struct {
	BundleId         string            `json:"bundle_id"`
	RamSize          int               `json:"ram_size"`
	IsStaticIp       bool              `json:"is_static_ip"`
	Tags             map[string]string `json:"tags"`
	Name             string            `json:"name"`
	BlueprintId      string            `json:"blueprint_id"`
	PrivateIpAddress string            `json:"private_ip_address"`
	KeyPairName      string            `json:"key_pair_name"`
	UserData         string            `json:"user_data"`
	Arn              string            `json:"arn"`
	Ipv6Address      string            `json:"ipv6_address"`
	Username         string            `json:"username"`
	AvailabilityZone string            `json:"availability_zone"`
	CreatedAt        string            `json:"created_at"`
	CpuCount         int               `json:"cpu_count"`
	PublicIpAddress  string            `json:"public_ip_address"`
}

type AwsLightsailInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLightsailInstanceList is a list of AwsLightsailInstances
type AwsLightsailInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLightsailInstance CRD objects
	Items []AwsLightsailInstance `json:"items,omitempty"`
}
