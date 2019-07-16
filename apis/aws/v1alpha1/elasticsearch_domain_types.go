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

type ElasticsearchDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticsearchDomainSpec   `json:"spec,omitempty"`
	Status            ElasticsearchDomainStatus `json:"status,omitempty"`
}

type ElasticsearchDomainSpecCognitoOptions struct {
	// +optional
	Enabled        bool   `json:"enabled,omitempty"`
	IdentityPoolId string `json:"identity_pool_id"`
	RoleArn        string `json:"role_arn"`
	UserPoolId     string `json:"user_pool_id"`
}

type ElasticsearchDomainSpecLogPublishingOptions struct {
	CloudwatchLogGroupArn string `json:"cloudwatch_log_group_arn"`
	// +optional
	Enabled bool   `json:"enabled,omitempty"`
	LogType string `json:"log_type"`
}

type ElasticsearchDomainSpecSnapshotOptions struct {
	AutomatedSnapshotStartHour int `json:"automated_snapshot_start_hour"`
}

type ElasticsearchDomainSpecVpcOptions struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubnetIds []string `json:"subnet_ids,omitempty"`
}

type ElasticsearchDomainSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CognitoOptions *[]ElasticsearchDomainSpec `json:"cognito_options,omitempty"`
	DomainName     string                     `json:"domain_name"`
	// +optional
	ElasticsearchVersion string `json:"elasticsearch_version,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	LogPublishingOptions *[]ElasticsearchDomainSpec `json:"log_publishing_options,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SnapshotOptions *[]ElasticsearchDomainSpec `json:"snapshot_options,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VpcOptions *[]ElasticsearchDomainSpec `json:"vpc_options,omitempty"`
}

type ElasticsearchDomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
