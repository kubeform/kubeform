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

type AwsElasticsearchDomainSpecNodeToNodeEncryption struct {
	Enabled bool `json:"enabled"`
}

type AwsElasticsearchDomainSpecLogPublishingOptions struct {
	CloudwatchLogGroupArn string `json:"cloudwatch_log_group_arn"`
	Enabled               bool   `json:"enabled"`
	LogType               string `json:"log_type"`
}

type AwsElasticsearchDomainSpecCognitoOptions struct {
	Enabled        bool   `json:"enabled"`
	UserPoolId     string `json:"user_pool_id"`
	IdentityPoolId string `json:"identity_pool_id"`
	RoleArn        string `json:"role_arn"`
}

type AwsElasticsearchDomainSpecEbsOptions struct {
	EbsEnabled bool   `json:"ebs_enabled"`
	Iops       int    `json:"iops"`
	VolumeSize int    `json:"volume_size"`
	VolumeType string `json:"volume_type"`
}

type AwsElasticsearchDomainSpecSnapshotOptions struct {
	AutomatedSnapshotStartHour int `json:"automated_snapshot_start_hour"`
}

type AwsElasticsearchDomainSpecEncryptAtRest struct {
	Enabled  bool   `json:"enabled"`
	KmsKeyId string `json:"kms_key_id"`
}

type AwsElasticsearchDomainSpecClusterConfig struct {
	DedicatedMasterEnabled bool   `json:"dedicated_master_enabled"`
	DedicatedMasterType    string `json:"dedicated_master_type"`
	InstanceCount          int    `json:"instance_count"`
	InstanceType           string `json:"instance_type"`
	ZoneAwarenessEnabled   bool   `json:"zone_awareness_enabled"`
	DedicatedMasterCount   int    `json:"dedicated_master_count"`
}

type AwsElasticsearchDomainSpecVpcOptions struct {
	AvailabilityZones []string `json:"availability_zones"`
	SecurityGroupIds  []string `json:"security_group_ids"`
	SubnetIds         []string `json:"subnet_ids"`
	VpcId             string   `json:"vpc_id"`
}

type AwsElasticsearchDomainSpec struct {
	Endpoint             string                       `json:"endpoint"`
	NodeToNodeEncryption []AwsElasticsearchDomainSpec `json:"node_to_node_encryption"`
	LogPublishingOptions []AwsElasticsearchDomainSpec `json:"log_publishing_options"`
	CognitoOptions       []AwsElasticsearchDomainSpec `json:"cognito_options"`
	DomainName           string                       `json:"domain_name"`
	Arn                  string                       `json:"arn"`
	DomainId             string                       `json:"domain_id"`
	ElasticsearchVersion string                       `json:"elasticsearch_version"`
	Tags                 map[string]string            `json:"tags"`
	AdvancedOptions      map[string]string            `json:"advanced_options"`
	EbsOptions           []AwsElasticsearchDomainSpec `json:"ebs_options"`
	SnapshotOptions      []AwsElasticsearchDomainSpec `json:"snapshot_options"`
	AccessPolicies       string                       `json:"access_policies"`
	KibanaEndpoint       string                       `json:"kibana_endpoint"`
	EncryptAtRest        []AwsElasticsearchDomainSpec `json:"encrypt_at_rest"`
	ClusterConfig        []AwsElasticsearchDomainSpec `json:"cluster_config"`
	VpcOptions           []AwsElasticsearchDomainSpec `json:"vpc_options"`
}



type AwsElasticsearchDomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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