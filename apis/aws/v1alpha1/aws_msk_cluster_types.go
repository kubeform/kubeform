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

type AwsMskCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsMskClusterSpec   `json:"spec,omitempty"`
	Status            AwsMskClusterStatus `json:"status,omitempty"`
}

type AwsMskClusterSpecClientAuthenticationTls struct {
	CertificateAuthorityArns []string `json:"certificate_authority_arns"`
}

type AwsMskClusterSpecClientAuthentication struct {
	Tls []AwsMskClusterSpecClientAuthentication `json:"tls"`
}

type AwsMskClusterSpecConfigurationInfo struct {
	Arn      string `json:"arn"`
	Revision int    `json:"revision"`
}

type AwsMskClusterSpecEncryptionInfoEncryptionInTransit struct {
	InCluster    bool   `json:"in_cluster"`
	ClientBroker string `json:"client_broker"`
}

type AwsMskClusterSpecEncryptionInfo struct {
	EncryptionAtRestKmsKeyArn string                            `json:"encryption_at_rest_kms_key_arn"`
	EncryptionInTransit       []AwsMskClusterSpecEncryptionInfo `json:"encryption_in_transit"`
}

type AwsMskClusterSpecBrokerNodeGroupInfo struct {
	EbsVolumeSize  int      `json:"ebs_volume_size"`
	AzDistribution string   `json:"az_distribution"`
	ClientSubnets  []string `json:"client_subnets"`
	InstanceType   string   `json:"instance_type"`
	SecurityGroups []string `json:"security_groups"`
}

type AwsMskClusterSpec struct {
	CurrentVersion         string              `json:"current_version"`
	Tags                   map[string]string   `json:"tags"`
	ClientAuthentication   []AwsMskClusterSpec `json:"client_authentication"`
	ClusterName            string              `json:"cluster_name"`
	ConfigurationInfo      []AwsMskClusterSpec `json:"configuration_info"`
	KafkaVersion           string              `json:"kafka_version"`
	ZookeeperConnectString string              `json:"zookeeper_connect_string"`
	Arn                    string              `json:"arn"`
	EncryptionInfo         []AwsMskClusterSpec `json:"encryption_info"`
	NumberOfBrokerNodes    int                 `json:"number_of_broker_nodes"`
	BootstrapBrokers       string              `json:"bootstrap_brokers"`
	BootstrapBrokersTls    string              `json:"bootstrap_brokers_tls"`
	BrokerNodeGroupInfo    []AwsMskClusterSpec `json:"broker_node_group_info"`
	EnhancedMonitoring     string              `json:"enhanced_monitoring"`
}

type AwsMskClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsMskClusterList is a list of AwsMskClusters
type AwsMskClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsMskCluster CRD objects
	Items []AwsMskCluster `json:"items,omitempty"`
}
