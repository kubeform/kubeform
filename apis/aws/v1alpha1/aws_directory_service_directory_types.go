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

type AwsDirectoryServiceDirectory struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDirectoryServiceDirectorySpec   `json:"spec,omitempty"`
	Status            AwsDirectoryServiceDirectoryStatus `json:"status,omitempty"`
}

type AwsDirectoryServiceDirectorySpecConnectSettings struct {
	CustomerUsername string   `json:"customer_username"`
	CustomerDnsIps   []string `json:"customer_dns_ips"`
	SubnetIds        []string `json:"subnet_ids"`
	VpcId            string   `json:"vpc_id"`
}

type AwsDirectoryServiceDirectorySpecVpcSettings struct {
	SubnetIds []string `json:"subnet_ids"`
	VpcId     string   `json:"vpc_id"`
}

type AwsDirectoryServiceDirectorySpec struct {
	Tags            map[string]string                  `json:"tags"`
	DnsIpAddresses  []string                           `json:"dns_ip_addresses"`
	SecurityGroupId string                             `json:"security_group_id"`
	Type            string                             `json:"type"`
	Alias           string                             `json:"alias"`
	Description     string                             `json:"description"`
	AccessUrl       string                             `json:"access_url"`
	ShortName       string                             `json:"short_name"`
	EnableSso       bool                               `json:"enable_sso"`
	ConnectSettings []AwsDirectoryServiceDirectorySpec `json:"connect_settings"`
	Password        string                             `json:"password"`
	VpcSettings     []AwsDirectoryServiceDirectorySpec `json:"vpc_settings"`
	Edition         string                             `json:"edition"`
	Name            string                             `json:"name"`
	Size            string                             `json:"size"`
}

type AwsDirectoryServiceDirectoryStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDirectoryServiceDirectoryList is a list of AwsDirectoryServiceDirectorys
type AwsDirectoryServiceDirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDirectoryServiceDirectory CRD objects
	Items []AwsDirectoryServiceDirectory `json:"items,omitempty"`
}
