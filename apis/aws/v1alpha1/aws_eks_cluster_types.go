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

type AwsEksClusterSpecCertificateAuthority struct {
	Data string `json:"data"`
}

type AwsEksClusterSpecVpcConfig struct {
	EndpointPrivateAccess bool     `json:"endpoint_private_access"`
	EndpointPublicAccess  bool     `json:"endpoint_public_access"`
	SecurityGroupIds      []string `json:"security_group_ids"`
	SubnetIds             []string `json:"subnet_ids"`
	VpcId                 string   `json:"vpc_id"`
}

type AwsEksClusterSpec struct {
	Arn                    string              `json:"arn"`
	Endpoint               string              `json:"endpoint"`
	PlatformVersion        string              `json:"platform_version"`
	RoleArn                string              `json:"role_arn"`
	Version                string              `json:"version"`
	CertificateAuthority   []AwsEksClusterSpec `json:"certificate_authority"`
	CreatedAt              string              `json:"created_at"`
	Name                   string              `json:"name"`
	VpcConfig              []AwsEksClusterSpec `json:"vpc_config"`
	EnabledClusterLogTypes []string            `json:"enabled_cluster_log_types"`
}



type AwsEksClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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