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

type AwsElasticsearchDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElasticsearchDomainSpec   `json:"spec,omitempty"`
	Status            AwsElasticsearchDomainStatus `json:"status,omitempty"`
}

type AwsElasticsearchDomainSpecEncryptAtRest struct {
	Enabled  bool   `json:"enabled"`
	KmsKeyId string `json:"kms_key_id"`
}

type AwsElasticsearchDomainSpecClusterConfig struct {
	ZoneAwarenessEnabled   bool   `json:"zone_awareness_enabled"`
	DedicatedMasterCount   int    `json:"dedicated_master_count"`
	DedicatedMasterEnabled bool   `json:"dedicated_master_enabled"`
	DedicatedMasterType    string `json:"dedicated_master_type"`
	InstanceCount          int    `json:"instance_count"`
	InstanceType           string `json:"instance_type"`
}

type AwsElasticsearchDomainSpecVpcOptions struct {
	AvailabilityZones []string `json:"availability_zones"`
	SecurityGroupIds  []string `json:"security_group_ids"`
	SubnetIds         []string `json:"subnet_ids"`
	VpcId             string   `json:"vpc_id"`
}

type AwsElasticsearchDomainSpecCognitoOptions struct {
	Enabled        bool   `json:"enabled"`
	UserPoolId     string `json:"user_pool_id"`
	IdentityPoolId string `json:"identity_pool_id"`
	RoleArn        string `json:"role_arn"`
}

type AwsElasticsearchDomainSpecSnapshotOptions struct {
	AutomatedSnapshotStartHour int `json:"automated_snapshot_start_hour"`
}

type AwsElasticsearchDomainSpecLogPublishingOptions struct {
	LogType               string `json:"log_type"`
	CloudwatchLogGroupArn string `json:"cloudwatch_log_group_arn"`
	Enabled               bool   `json:"enabled"`
}

type AwsElasticsearchDomainSpecNodeToNodeEncryption struct {
	Enabled bool `json:"enabled"`
}

type AwsElasticsearchDomainSpecEbsOptions struct {
	EbsEnabled bool   `json:"ebs_enabled"`
	Iops       int    `json:"iops"`
	VolumeSize int    `json:"volume_size"`
	VolumeType string `json:"volume_type"`
}

type AwsElasticsearchDomainSpec struct {
	Tags                 map[string]string            `json:"tags"`
	EncryptAtRest        []AwsElasticsearchDomainSpec `json:"encrypt_at_rest"`
	ClusterConfig        []AwsElasticsearchDomainSpec `json:"cluster_config"`
	VpcOptions           []AwsElasticsearchDomainSpec `json:"vpc_options"`
	KibanaEndpoint       string                       `json:"kibana_endpoint"`
	CognitoOptions       []AwsElasticsearchDomainSpec `json:"cognito_options"`
	AccessPolicies       string                       `json:"access_policies"`
	AdvancedOptions      map[string]string            `json:"advanced_options"`
	DomainId             string                       `json:"domain_id"`
	ElasticsearchVersion string                       `json:"elasticsearch_version"`
	Arn                  string                       `json:"arn"`
	SnapshotOptions      []AwsElasticsearchDomainSpec `json:"snapshot_options"`
	LogPublishingOptions []AwsElasticsearchDomainSpec `json:"log_publishing_options"`
	NodeToNodeEncryption []AwsElasticsearchDomainSpec `json:"node_to_node_encryption"`
	DomainName           string                       `json:"domain_name"`
	Endpoint             string                       `json:"endpoint"`
	EbsOptions           []AwsElasticsearchDomainSpec `json:"ebs_options"`
}

type AwsElasticsearchDomainStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsElasticsearchDomainList is a list of AwsElasticsearchDomains
type AwsElasticsearchDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticsearchDomain CRD objects
	Items []AwsElasticsearchDomain `json:"items,omitempty"`
}
