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

type CodebuildProject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodebuildProjectSpec   `json:"spec,omitempty"`
	Status            CodebuildProjectStatus `json:"status,omitempty"`
}

type CodebuildProjectSpecArtifacts struct {
	// +optional
	EncryptionDisabled bool `json:"encryption_disabled,omitempty"`
	// +optional
	Location string `json:"location,omitempty"`
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	NamespaceType string `json:"namespace_type,omitempty"`
	// +optional
	Packaging string `json:"packaging,omitempty"`
	// +optional
	Path string `json:"path,omitempty"`
	Type string `json:"type"`
}

type CodebuildProjectSpecCache struct {
	// +optional
	Location string `json:"location,omitempty"`
	// +optional
	Modes []string `json:"modes,omitempty"`
	// +optional
	Type string `json:"type,omitempty"`
}

type CodebuildProjectSpecEnvironmentRegistryCredential struct {
	Credential         string `json:"credential"`
	CredentialProvider string `json:"credential_provider"`
}

type CodebuildProjectSpecEnvironment struct {
	// +optional
	Certificate string `json:"certificate,omitempty"`
	ComputeType string `json:"compute_type"`
	Image       string `json:"image"`
	// +optional
	ImagePullCredentialsType string `json:"image_pull_credentials_type,omitempty"`
	// +optional
	PrivilegedMode bool `json:"privileged_mode,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RegistryCredential *[]CodebuildProjectSpecEnvironment `json:"registry_credential,omitempty"`
	Type               string                             `json:"type"`
}

type CodebuildProjectSpecLogsConfigCloudwatchLogs struct {
	// +optional
	GroupName string `json:"group_name,omitempty"`
	// +optional
	Status string `json:"status,omitempty"`
	// +optional
	StreamName string `json:"stream_name,omitempty"`
}

type CodebuildProjectSpecLogsConfigS3Logs struct {
	// +optional
	EncryptionDisabled bool `json:"encryption_disabled,omitempty"`
	// +optional
	Location string `json:"location,omitempty"`
	// +optional
	Status string `json:"status,omitempty"`
}

type CodebuildProjectSpecLogsConfig struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CloudwatchLogs *[]CodebuildProjectSpecLogsConfig `json:"cloudwatch_logs,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3Logs *[]CodebuildProjectSpecLogsConfig `json:"s3_logs,omitempty"`
}

type CodebuildProjectSpecSecondaryArtifacts struct {
	ArtifactIdentifier string `json:"artifact_identifier"`
	// +optional
	EncryptionDisabled bool `json:"encryption_disabled,omitempty"`
	// +optional
	Location string `json:"location,omitempty"`
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	NamespaceType string `json:"namespace_type,omitempty"`
	// +optional
	Packaging string `json:"packaging,omitempty"`
	// +optional
	Path string `json:"path,omitempty"`
	Type string `json:"type"`
}

type CodebuildProjectSpecSecondarySourcesAuth struct {
	// +optional
	Resource string `json:"resource,omitempty"`
	Type     string `json:"type"`
}

type CodebuildProjectSpecSecondarySources struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Auth *[]CodebuildProjectSpecSecondarySources `json:"auth,omitempty"`
	// +optional
	Buildspec string `json:"buildspec,omitempty"`
	// +optional
	GitCloneDepth int `json:"git_clone_depth,omitempty"`
	// +optional
	InsecureSsl bool `json:"insecure_ssl,omitempty"`
	// +optional
	Location string `json:"location,omitempty"`
	// +optional
	ReportBuildStatus bool   `json:"report_build_status,omitempty"`
	SourceIdentifier  string `json:"source_identifier"`
	Type              string `json:"type"`
}

type CodebuildProjectSpecSourceAuth struct {
	// +optional
	Resource string `json:"resource,omitempty"`
	Type     string `json:"type"`
}

type CodebuildProjectSpecSource struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Auth *[]CodebuildProjectSpecSource `json:"auth,omitempty"`
	// +optional
	Buildspec string `json:"buildspec,omitempty"`
	// +optional
	GitCloneDepth int `json:"git_clone_depth,omitempty"`
	// +optional
	InsecureSsl bool `json:"insecure_ssl,omitempty"`
	// +optional
	Location string `json:"location,omitempty"`
	// +optional
	ReportBuildStatus bool   `json:"report_build_status,omitempty"`
	Type              string `json:"type"`
}

type CodebuildProjectSpecVpcConfig struct {
	// +kubebuilder:validation:MaxItems=5
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIds []string `json:"security_group_ids"`
	// +kubebuilder:validation:MaxItems=16
	// +kubebuilder:validation:UniqueItems=true
	Subnets []string `json:"subnets"`
	VpcId   string   `json:"vpc_id"`
}

type CodebuildProjectSpec struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Artifacts []CodebuildProjectSpec `json:"artifacts"`
	// +optional
	BadgeEnabled bool `json:"badge_enabled,omitempty"`
	// +optional
	BuildTimeout int `json:"build_timeout,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Cache *[]CodebuildProjectSpec `json:"cache,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Environment []CodebuildProjectSpec `json:"environment"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LogsConfig *[]CodebuildProjectSpec `json:"logs_config,omitempty"`
	Name       string                  `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecondaryArtifacts *[]CodebuildProjectSpec `json:"secondary_artifacts,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecondarySources *[]CodebuildProjectSpec `json:"secondary_sources,omitempty"`
	ServiceRole      string                  `json:"service_role"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Source []CodebuildProjectSpec `json:"source"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VpcConfig *[]CodebuildProjectSpec `json:"vpc_config,omitempty"`
}

type CodebuildProjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
