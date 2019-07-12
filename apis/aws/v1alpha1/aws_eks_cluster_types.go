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

type AwsEksCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEksClusterSpec   `json:"spec,omitempty"`
	Status            AwsEksClusterStatus `json:"status,omitempty"`
}

type AwsEksClusterSpecVpcConfig struct {
	VpcId                 string   `json:"vpc_id"`
	EndpointPrivateAccess bool     `json:"endpoint_private_access"`
	EndpointPublicAccess  bool     `json:"endpoint_public_access"`
	SecurityGroupIds      []string `json:"security_group_ids"`
	SubnetIds             []string `json:"subnet_ids"`
}

type AwsEksClusterSpecCertificateAuthority struct {
	Data string `json:"data"`
}

type AwsEksClusterSpec struct {
	CreatedAt              string              `json:"created_at"`
	Endpoint               string              `json:"endpoint"`
	Name                   string              `json:"name"`
	PlatformVersion        string              `json:"platform_version"`
	VpcConfig              []AwsEksClusterSpec `json:"vpc_config"`
	Arn                    string              `json:"arn"`
	CertificateAuthority   []AwsEksClusterSpec `json:"certificate_authority"`
	RoleArn                string              `json:"role_arn"`
	Version                string              `json:"version"`
	EnabledClusterLogTypes []string            `json:"enabled_cluster_log_types"`
}

type AwsEksClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEksClusterList is a list of AwsEksClusters
type AwsEksClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEksCluster CRD objects
	Items []AwsEksCluster `json:"items,omitempty"`
}
