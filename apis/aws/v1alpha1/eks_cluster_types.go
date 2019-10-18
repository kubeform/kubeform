package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type EksCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EksClusterSpec   `json:"spec,omitempty"`
	Status            EksClusterStatus `json:"status,omitempty"`
}

type EksClusterSpecCertificateAuthority struct {
	// +optional
	Data string `json:"data,omitempty" tf:"data,omitempty"`
}

type EksClusterSpecVpcConfig struct {
	// +optional
	EndpointPrivateAccess bool `json:"endpointPrivateAccess,omitempty" tf:"endpoint_private_access,omitempty"`
	// +optional
	EndpointPublicAccess bool `json:"endpointPublicAccess,omitempty" tf:"endpoint_public_access,omitempty"`
	// +optional
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids,omitempty"`
	// +kubebuilder:validation:MinItems=1
	SubnetIDS []string `json:"subnetIDS" tf:"subnet_ids"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
}

type EksClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CertificateAuthority []EksClusterSpecCertificateAuthority `json:"certificateAuthority,omitempty" tf:"certificate_authority,omitempty"`
	// +optional
	CreatedAt string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
	// +optional
	EnabledClusterLogTypes []string `json:"enabledClusterLogTypes,omitempty" tf:"enabled_cluster_log_types,omitempty"`
	// +optional
	Endpoint string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
	Name     string `json:"name" tf:"name"`
	// +optional
	PlatformVersion string `json:"platformVersion,omitempty" tf:"platform_version,omitempty"`
	RoleArn         string `json:"roleArn" tf:"role_arn"`
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
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
