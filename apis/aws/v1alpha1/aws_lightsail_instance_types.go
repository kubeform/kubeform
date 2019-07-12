package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLightsailInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLightsailInstanceSpec   `json:"spec,omitempty"`
	Status            AwsLightsailInstanceStatus `json:"status,omitempty"`
}

type AwsLightsailInstanceSpec struct {
	KeyPairName      string `json:"key_pair_name"`
	Arn              string `json:"arn"`
	CpuCount         int    `json:"cpu_count"`
	Name             string `json:"name"`
	PublicIpAddress  string `json:"public_ip_address"`
	BlueprintId      string `json:"blueprint_id"`
	UserData         string `json:"user_data"`
	CreatedAt        string `json:"created_at"`
	RamSize          int    `json:"ram_size"`
	IsStaticIp       bool   `json:"is_static_ip"`
	PrivateIpAddress string `json:"private_ip_address"`
	AvailabilityZone string `json:"availability_zone"`
	BundleId         string `json:"bundle_id"`
	Ipv6Address      string `json:"ipv6_address"`
	Username         string `json:"username"`
}

type AwsLightsailInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLightsailInstanceList is a list of AwsLightsailInstances
type AwsLightsailInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLightsailInstance CRD objects
	Items []AwsLightsailInstance `json:"items,omitempty"`
}
