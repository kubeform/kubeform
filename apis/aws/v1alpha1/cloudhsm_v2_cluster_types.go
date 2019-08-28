package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CloudhsmV2Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudhsmV2ClusterSpec   `json:"spec,omitempty"`
	Status            CloudhsmV2ClusterStatus `json:"status,omitempty"`
}

type CloudhsmV2ClusterSpecClusterCertificates struct {
	// +optional
	AwsHardwareCertificate string `json:"awsHardwareCertificate,omitempty" tf:"aws_hardware_certificate,omitempty"`
	// +optional
	ClusterCertificate string `json:"clusterCertificate,omitempty" tf:"cluster_certificate,omitempty"`
	// +optional
	ClusterCsr string `json:"clusterCsr,omitempty" tf:"cluster_csr,omitempty"`
	// +optional
	HsmCertificate string `json:"hsmCertificate,omitempty" tf:"hsm_certificate,omitempty"`
	// +optional
	ManufacturerHardwareCertificate string `json:"manufacturerHardwareCertificate,omitempty" tf:"manufacturer_hardware_certificate,omitempty"`
}

type CloudhsmV2ClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	ClusterCertificates []CloudhsmV2ClusterSpecClusterCertificates `json:"clusterCertificates,omitempty" tf:"cluster_certificates,omitempty"`
	// +optional
	ClusterID string `json:"clusterID,omitempty" tf:"cluster_id,omitempty"`
	// +optional
	ClusterState string `json:"clusterState,omitempty" tf:"cluster_state,omitempty"`
	HsmType      string `json:"hsmType" tf:"hsm_type"`
	// +optional
	SecurityGroupID string `json:"securityGroupID,omitempty" tf:"security_group_id,omitempty"`
	// +optional
	SourceBackupIdentifier string `json:"sourceBackupIdentifier,omitempty" tf:"source_backup_identifier,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS" tf:"subnet_ids"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
}



type CloudhsmV2ClusterStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *CloudhsmV2ClusterSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudhsmV2ClusterList is a list of CloudhsmV2Clusters
type CloudhsmV2ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudhsmV2Cluster CRD objects
	Items []CloudhsmV2Cluster `json:"items,omitempty"`
}