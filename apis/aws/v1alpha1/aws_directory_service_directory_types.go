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
	Size            string                             `json:"size"`
	AccessUrl       string                             `json:"access_url"`
	Description     string                             `json:"description"`
	EnableSso       bool                               `json:"enable_sso"`
	Name            string                             `json:"name"`
	Password        string                             `json:"password"`
	Alias           string                             `json:"alias"`
	Type            string                             `json:"type"`
	Edition         string                             `json:"edition"`
	ShortName       string                             `json:"short_name"`
	ConnectSettings []AwsDirectoryServiceDirectorySpec `json:"connect_settings"`
	SecurityGroupId string                             `json:"security_group_id"`
	Tags            map[string]string                  `json:"tags"`
	VpcSettings     []AwsDirectoryServiceDirectorySpec `json:"vpc_settings"`
	DnsIpAddresses  []string                           `json:"dns_ip_addresses"`
}



type AwsDirectoryServiceDirectoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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