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

type MskCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MskClusterSpec   `json:"spec,omitempty"`
	Status            MskClusterStatus `json:"status,omitempty"`
}

type MskClusterSpecBrokerNodeGroupInfo struct {
	// +optional
	AzDistribution string   `json:"az_distribution,omitempty"`
	ClientSubnets  []string `json:"client_subnets"`
	EbsVolumeSize  int      `json:"ebs_volume_size"`
	InstanceType   string   `json:"instance_type"`
	SecurityGroups []string `json:"security_groups"`
}

type MskClusterSpecClientAuthenticationTls struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CertificateAuthorityArns []string `json:"certificate_authority_arns,omitempty"`
}

type MskClusterSpecClientAuthentication struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Tls *[]MskClusterSpecClientAuthentication `json:"tls,omitempty"`
}

type MskClusterSpecConfigurationInfo struct {
	Arn      string `json:"arn"`
	Revision int    `json:"revision"`
}

type MskClusterSpecEncryptionInfoEncryptionInTransit struct {
	// +optional
	ClientBroker string `json:"client_broker,omitempty"`
	// +optional
	InCluster bool `json:"in_cluster,omitempty"`
}

type MskClusterSpecEncryptionInfo struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionInTransit *[]MskClusterSpecEncryptionInfo `json:"encryption_in_transit,omitempty"`
}

type MskClusterSpec struct {
	// +kubebuilder:validation:MaxItems=1
	BrokerNodeGroupInfo []MskClusterSpec `json:"broker_node_group_info"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ClientAuthentication *[]MskClusterSpec `json:"client_authentication,omitempty"`
	ClusterName          string            `json:"cluster_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ConfigurationInfo *[]MskClusterSpec `json:"configuration_info,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionInfo *[]MskClusterSpec `json:"encryption_info,omitempty"`
	// +optional
	EnhancedMonitoring  string `json:"enhanced_monitoring,omitempty"`
	KafkaVersion        string `json:"kafka_version"`
	NumberOfBrokerNodes int    `json:"number_of_broker_nodes"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type MskClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MskClusterList is a list of MskClusters
type MskClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MskCluster CRD objects
	Items []MskCluster `json:"items,omitempty"`
}
