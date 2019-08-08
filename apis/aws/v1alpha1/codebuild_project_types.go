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

type CodebuildProject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodebuildProjectSpec   `json:"spec,omitempty"`
	Status            CodebuildProjectStatus `json:"status,omitempty"`
}

type CodebuildProjectSpecArtifacts struct {
	// +optional
	EncryptionDisabled bool `json:"encryptionDisabled,omitempty" tf:"encryption_disabled,omitempty"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamespaceType string `json:"namespaceType,omitempty" tf:"namespace_type,omitempty"`
	// +optional
	Packaging string `json:"packaging,omitempty" tf:"packaging,omitempty"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	Type string `json:"type" tf:"type"`
}

type CodebuildProjectSpecCache struct {
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	Modes []string `json:"modes,omitempty" tf:"modes,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type CodebuildProjectSpecEnvironmentEnvironmentVariable struct {
	Name string `json:"name" tf:"name"`
	// +optional
	Type  string `json:"type,omitempty" tf:"type,omitempty"`
	Value string `json:"value" tf:"value"`
}

type CodebuildProjectSpecEnvironmentRegistryCredential struct {
	Credential         string `json:"credential" tf:"credential"`
	CredentialProvider string `json:"credentialProvider" tf:"credential_provider"`
}

type CodebuildProjectSpecEnvironment struct {
	// +optional
	Certificate string `json:"certificate,omitempty" tf:"certificate,omitempty"`
	ComputeType string `json:"computeType" tf:"compute_type"`
	// +optional
	EnvironmentVariable []CodebuildProjectSpecEnvironmentEnvironmentVariable `json:"environmentVariable,omitempty" tf:"environment_variable,omitempty"`
	Image               string                                               `json:"image" tf:"image"`
	// +optional
	ImagePullCredentialsType string `json:"imagePullCredentialsType,omitempty" tf:"image_pull_credentials_type,omitempty"`
	// +optional
	PrivilegedMode bool `json:"privilegedMode,omitempty" tf:"privileged_mode,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RegistryCredential []CodebuildProjectSpecEnvironmentRegistryCredential `json:"registryCredential,omitempty" tf:"registry_credential,omitempty"`
	Type               string                                              `json:"type" tf:"type"`
}

type CodebuildProjectSpecLogsConfigCloudwatchLogs struct {
	// +optional
	GroupName string `json:"groupName,omitempty" tf:"group_name,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	StreamName string `json:"streamName,omitempty" tf:"stream_name,omitempty"`
}

type CodebuildProjectSpecLogsConfigS3Logs struct {
	// +optional
	EncryptionDisabled bool `json:"encryptionDisabled,omitempty" tf:"encryption_disabled,omitempty"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
}

type CodebuildProjectSpecLogsConfig struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CloudwatchLogs []CodebuildProjectSpecLogsConfigCloudwatchLogs `json:"cloudwatchLogs,omitempty" tf:"cloudwatch_logs,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3Logs []CodebuildProjectSpecLogsConfigS3Logs `json:"s3Logs,omitempty" tf:"s3_logs,omitempty"`
}

type CodebuildProjectSpecSecondaryArtifacts struct {
	ArtifactIdentifier string `json:"artifactIdentifier" tf:"artifact_identifier"`
	// +optional
	EncryptionDisabled bool `json:"encryptionDisabled,omitempty" tf:"encryption_disabled,omitempty"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamespaceType string `json:"namespaceType,omitempty" tf:"namespace_type,omitempty"`
	// +optional
	Packaging string `json:"packaging,omitempty" tf:"packaging,omitempty"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	Type string `json:"type" tf:"type"`
}

type CodebuildProjectSpecSecondarySourcesAuth struct {
	// +optional
	Resource string `json:"-" sensitive:"true" tf:"resource,omitempty"`
	Type     string `json:"type" tf:"type"`
}

type CodebuildProjectSpecSecondarySources struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Auth []CodebuildProjectSpecSecondarySourcesAuth `json:"auth,omitempty" tf:"auth,omitempty"`
	// +optional
	Buildspec string `json:"buildspec,omitempty" tf:"buildspec,omitempty"`
	// +optional
	GitCloneDepth int `json:"gitCloneDepth,omitempty" tf:"git_clone_depth,omitempty"`
	// +optional
	InsecureSsl bool `json:"insecureSsl,omitempty" tf:"insecure_ssl,omitempty"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	ReportBuildStatus bool   `json:"reportBuildStatus,omitempty" tf:"report_build_status,omitempty"`
	SourceIdentifier  string `json:"sourceIdentifier" tf:"source_identifier"`
	Type              string `json:"type" tf:"type"`
}

type CodebuildProjectSpecSourceAuth struct {
	// +optional
	Resource string `json:"-" sensitive:"true" tf:"resource,omitempty"`
	Type     string `json:"type" tf:"type"`
}

type CodebuildProjectSpecSource struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Auth []CodebuildProjectSpecSourceAuth `json:"auth,omitempty" tf:"auth,omitempty"`
	// +optional
	Buildspec string `json:"buildspec,omitempty" tf:"buildspec,omitempty"`
	// +optional
	GitCloneDepth int `json:"gitCloneDepth,omitempty" tf:"git_clone_depth,omitempty"`
	// +optional
	InsecureSsl bool `json:"insecureSsl,omitempty" tf:"insecure_ssl,omitempty"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	ReportBuildStatus bool   `json:"reportBuildStatus,omitempty" tf:"report_build_status,omitempty"`
	Type              string `json:"type" tf:"type"`
}

type CodebuildProjectSpecVpcConfig struct {
	// +kubebuilder:validation:MaxItems=5
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIDS []string `json:"securityGroupIDS" tf:"security_group_ids"`
	// +kubebuilder:validation:MaxItems=16
	// +kubebuilder:validation:UniqueItems=true
	Subnets []string `json:"subnets" tf:"subnets"`
	VpcID   string   `json:"vpcID" tf:"vpc_id"`
}

type CodebuildProjectSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	KubeFormSecret *core.LocalObjectReference `json:"kubeFormSecret,omitempty" tf:"-"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Artifacts []CodebuildProjectSpecArtifacts `json:"artifacts" tf:"artifacts"`
	// +optional
	BadgeEnabled bool `json:"badgeEnabled,omitempty" tf:"badge_enabled,omitempty"`
	// +optional
	BadgeURL string `json:"badgeURL,omitempty" tf:"badge_url,omitempty"`
	// +optional
	BuildTimeout int `json:"buildTimeout,omitempty" tf:"build_timeout,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Cache []CodebuildProjectSpecCache `json:"cache,omitempty" tf:"cache,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	EncryptionKey string `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Environment []CodebuildProjectSpecEnvironment `json:"environment" tf:"environment"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LogsConfig []CodebuildProjectSpecLogsConfig `json:"logsConfig,omitempty" tf:"logs_config,omitempty"`
	Name       string                           `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecondaryArtifacts []CodebuildProjectSpecSecondaryArtifacts `json:"secondaryArtifacts,omitempty" tf:"secondary_artifacts,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecondarySources []CodebuildProjectSpecSecondarySources `json:"secondarySources,omitempty" tf:"secondary_sources,omitempty"`
	ServiceRole      string                                 `json:"serviceRole" tf:"service_role"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Source []CodebuildProjectSpecSource `json:"source" tf:"source"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VpcConfig []CodebuildProjectSpecVpcConfig `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type CodebuildProjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CodebuildProjectSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CodebuildProjectList is a list of CodebuildProjects
type CodebuildProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CodebuildProject CRD objects
	Items []CodebuildProject `json:"items,omitempty"`
}
