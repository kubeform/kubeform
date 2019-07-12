package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDirectoryServiceDirectory struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDirectoryServiceDirectorySpec   `json:"spec,omitempty"`
	Status            AwsDirectoryServiceDirectoryStatus `json:"status,omitempty"`
}

type AwsDirectoryServiceDirectorySpecConnectSettings struct {
	CustomerDnsIps   []string `json:"customer_dns_ips"`
	SubnetIds        []string `json:"subnet_ids"`
	VpcId            string   `json:"vpc_id"`
	CustomerUsername string   `json:"customer_username"`
}

type AwsDirectoryServiceDirectorySpecVpcSettings struct {
	SubnetIds []string `json:"subnet_ids"`
	VpcId     string   `json:"vpc_id"`
}

type AwsDirectoryServiceDirectorySpec struct {
	ConnectSettings []AwsDirectoryServiceDirectorySpec `json:"connect_settings"`
	Password        string                             `json:"password"`
	Alias           string                             `json:"alias"`
	DnsIpAddresses  []string                           `json:"dns_ip_addresses"`
	Type            string                             `json:"type"`
	Edition         string                             `json:"edition"`
	VpcSettings     []AwsDirectoryServiceDirectorySpec `json:"vpc_settings"`
	AccessUrl       string                             `json:"access_url"`
	EnableSso       bool                               `json:"enable_sso"`
	Size            string                             `json:"size"`
	Tags            map[string]string                  `json:"tags"`
	ShortName       string                             `json:"short_name"`
	SecurityGroupId string                             `json:"security_group_id"`
	Name            string                             `json:"name"`
	Description     string                             `json:"description"`
}



type AwsDirectoryServiceDirectoryStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDirectoryServiceDirectoryList is a list of AwsDirectoryServiceDirectorys
type AwsDirectoryServiceDirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDirectoryServiceDirectory CRD objects
	Items []AwsDirectoryServiceDirectory `json:"items,omitempty"`
}