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

type EksCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EksClusterSpec   `json:"spec,omitempty"`
	Status            EksClusterStatus `json:"status,omitempty"`
}

type EksClusterSpecCertificateAuthority struct {
	Data string `json:"data" tf:"data"`
}

type EksClusterSpecVpcConfig struct {
	// +optional
	EndpointPrivateAccess bool `json:"endpointPrivateAccess,omitempty" tf:"endpoint_private_access,omitempty"`
	// +optional
	EndpointPublicAccess bool `json:"endpointPublicAccess,omitempty" tf:"endpoint_public_access,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids,omitempty"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS" tf:"subnet_ids"`
	VpcID     string   `json:"vpcID" tf:"vpc_id"`
}

type EksClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Arn string `json:"arn" tf:"arn"`
	// +kubebuilder:validation:MaxItems=1
	CertificateAuthority []EksClusterSpecCertificateAuthority `json:"certificateAuthority" tf:"certificate_authority"`
	CreatedAt            string                               `json:"createdAt" tf:"created_at"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EnabledClusterLogTypes []string `json:"enabledClusterLogTypes,omitempty" tf:"enabled_cluster_log_types,omitempty"`
	Endpoint               string   `json:"endpoint" tf:"endpoint"`
	Name                   string   `json:"name" tf:"name"`
	PlatformVersion        string   `json:"platformVersion" tf:"platform_version"`
	RoleArn                string   `json:"roleArn" tf:"role_arn"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	VpcConfig []EksClusterSpecVpcConfig `json:"vpcConfig" tf:"vpc_config"`
}

type EksClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EksClusterSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EksClusterList is a list of EksClusters
type EksClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EksCluster CRD objects
	Items []EksCluster `json:"items,omitempty"`
}
