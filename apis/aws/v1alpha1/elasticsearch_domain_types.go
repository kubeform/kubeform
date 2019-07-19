package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ElasticsearchDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticsearchDomainSpec   `json:"spec,omitempty"`
	Status            ElasticsearchDomainStatus `json:"status,omitempty"`
}

type ElasticsearchDomainSpecClusterConfig struct {
	// +optional
	DedicatedMasterCount int `json:"dedicatedMasterCount,omitempty" tf:"dedicated_master_count,omitempty"`
	// +optional
	DedicatedMasterEnabled bool `json:"dedicatedMasterEnabled,omitempty" tf:"dedicated_master_enabled,omitempty"`
	// +optional
	DedicatedMasterType string `json:"dedicatedMasterType,omitempty" tf:"dedicated_master_type,omitempty"`
	// +optional
	InstanceCount int `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`
	// +optional
	InstanceType string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`
	// +optional
	ZoneAwarenessEnabled bool `json:"zoneAwarenessEnabled,omitempty" tf:"zone_awareness_enabled,omitempty"`
}

type ElasticsearchDomainSpecCognitoOptions struct {
	// +optional
	Enabled        bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	IdentityPoolID string `json:"identityPoolID" tf:"identity_pool_id"`
	RoleArn        string `json:"roleArn" tf:"role_arn"`
	UserPoolID     string `json:"userPoolID" tf:"user_pool_id"`
}

type ElasticsearchDomainSpecEbsOptions struct {
	EbsEnabled bool `json:"ebsEnabled" tf:"ebs_enabled"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	VolumeSize int `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
	// +optional
	VolumeType string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type ElasticsearchDomainSpecEncryptAtRest struct {
	Enabled bool `json:"enabled" tf:"enabled"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
}

type ElasticsearchDomainSpecLogPublishingOptions struct {
	CloudwatchLogGroupArn string `json:"cloudwatchLogGroupArn" tf:"cloudwatch_log_group_arn"`
	// +optional
	Enabled bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	LogType string `json:"logType" tf:"log_type"`
}

type ElasticsearchDomainSpecNodeToNodeEncryption struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type ElasticsearchDomainSpecSnapshotOptions struct {
	AutomatedSnapshotStartHour int `json:"automatedSnapshotStartHour" tf:"automated_snapshot_start_hour"`
}

type ElasticsearchDomainSpecVpcOptions struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS,omitempty" tf:"subnet_ids,omitempty"`
}

type ElasticsearchDomainSpec struct {
	// +optional
	AccessPolicies string `json:"accessPolicies,omitempty" tf:"access_policies,omitempty"`
	// +optional
	AdvancedOptions map[string]string `json:"advancedOptions,omitempty" tf:"advanced_options,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ClusterConfig []ElasticsearchDomainSpecClusterConfig `json:"clusterConfig,omitempty" tf:"cluster_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CognitoOptions []ElasticsearchDomainSpecCognitoOptions `json:"cognitoOptions,omitempty" tf:"cognito_options,omitempty"`
	DomainName     string                                  `json:"domainName" tf:"domain_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EbsOptions []ElasticsearchDomainSpecEbsOptions `json:"ebsOptions,omitempty" tf:"ebs_options,omitempty"`
	// +optional
	ElasticsearchVersion string `json:"elasticsearchVersion,omitempty" tf:"elasticsearch_version,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptAtRest []ElasticsearchDomainSpecEncryptAtRest `json:"encryptAtRest,omitempty" tf:"encrypt_at_rest,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	LogPublishingOptions []ElasticsearchDomainSpecLogPublishingOptions `json:"logPublishingOptions,omitempty" tf:"log_publishing_options,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NodeToNodeEncryption []ElasticsearchDomainSpecNodeToNodeEncryption `json:"nodeToNodeEncryption,omitempty" tf:"node_to_node_encryption,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SnapshotOptions []ElasticsearchDomainSpecSnapshotOptions `json:"snapshotOptions,omitempty" tf:"snapshot_options,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VpcOptions  []ElasticsearchDomainSpecVpcOptions `json:"vpcOptions,omitempty" tf:"vpc_options,omitempty"`
	ProviderRef core.LocalObjectReference           `json:"providerRef" tf:"-"`
}

type ElasticsearchDomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticsearchDomainList is a list of ElasticsearchDomains
type ElasticsearchDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticsearchDomain CRD objects
	Items []ElasticsearchDomain `json:"items,omitempty"`
}
