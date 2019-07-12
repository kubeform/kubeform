package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudhsmV2Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudhsmV2ClusterSpec   `json:"spec,omitempty"`
	Status            AwsCloudhsmV2ClusterStatus `json:"status,omitempty"`
}

type AwsCloudhsmV2ClusterSpecClusterCertificates struct {
	ManufacturerHardwareCertificate string `json:"manufacturer_hardware_certificate"`
	ClusterCertificate              string `json:"cluster_certificate"`
	ClusterCsr                      string `json:"cluster_csr"`
	AwsHardwareCertificate          string `json:"aws_hardware_certificate"`
	HsmCertificate                  string `json:"hsm_certificate"`
}

type AwsCloudhsmV2ClusterSpec struct {
	SubnetIds              []string                   `json:"subnet_ids"`
	VpcId                  string                     `json:"vpc_id"`
	SecurityGroupId        string                     `json:"security_group_id"`
	ClusterState           string                     `json:"cluster_state"`
	HsmType                string                     `json:"hsm_type"`
	ClusterId              string                     `json:"cluster_id"`
	ClusterCertificates    []AwsCloudhsmV2ClusterSpec `json:"cluster_certificates"`
	Tags                   map[string]string          `json:"tags"`
	SourceBackupIdentifier string                     `json:"source_backup_identifier"`
}



type AwsCloudhsmV2ClusterStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudhsmV2ClusterList is a list of AwsCloudhsmV2Clusters
type AwsCloudhsmV2ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudhsmV2Cluster CRD objects
	Items []AwsCloudhsmV2Cluster `json:"items,omitempty"`
}