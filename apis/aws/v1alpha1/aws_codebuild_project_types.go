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

type AwsCodebuildProject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCodebuildProjectSpec   `json:"spec,omitempty"`
	Status            AwsCodebuildProjectStatus `json:"status,omitempty"`
}

type AwsCodebuildProjectSpecVpcConfig struct {
	VpcId            string   `json:"vpc_id"`
	Subnets          []string `json:"subnets"`
	SecurityGroupIds []string `json:"security_group_ids"`
}

type AwsCodebuildProjectSpecArtifacts struct {
	Path               string `json:"path"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	EncryptionDisabled bool   `json:"encryption_disabled"`
	Location           string `json:"location"`
	NamespaceType      string `json:"namespace_type"`
	Packaging          string `json:"packaging"`
}

type AwsCodebuildProjectSpecEnvironmentRegistryCredential struct {
	Credential         string `json:"credential"`
	CredentialProvider string `json:"credential_provider"`
}

type AwsCodebuildProjectSpecEnvironmentEnvironmentVariable struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsCodebuildProjectSpecEnvironment struct {
	Image                    string                               `json:"image"`
	Type                     string                               `json:"type"`
	ImagePullCredentialsType string                               `json:"image_pull_credentials_type"`
	PrivilegedMode           bool                                 `json:"privileged_mode"`
	Certificate              string                               `json:"certificate"`
	RegistryCredential       []AwsCodebuildProjectSpecEnvironment `json:"registry_credential"`
	ComputeType              string                               `json:"compute_type"`
	EnvironmentVariable      []AwsCodebuildProjectSpecEnvironment `json:"environment_variable"`
}

type AwsCodebuildProjectSpecLogsConfigCloudwatchLogs struct {
	Status     string `json:"status"`
	GroupName  string `json:"group_name"`
	StreamName string `json:"stream_name"`
}

type AwsCodebuildProjectSpecLogsConfigS3Logs struct {
	Status             string `json:"status"`
	Location           string `json:"location"`
	EncryptionDisabled bool   `json:"encryption_disabled"`
}

type AwsCodebuildProjectSpecLogsConfig struct {
	CloudwatchLogs []AwsCodebuildProjectSpecLogsConfig `json:"cloudwatch_logs"`
	S3Logs         []AwsCodebuildProjectSpecLogsConfig `json:"s3_logs"`
}

type AwsCodebuildProjectSpecSecondaryArtifacts struct {
	Type               string `json:"type"`
	Name               string `json:"name"`
	EncryptionDisabled bool   `json:"encryption_disabled"`
	Location           string `json:"location"`
	NamespaceType      string `json:"namespace_type"`
	Packaging          string `json:"packaging"`
	Path               string `json:"path"`
	ArtifactIdentifier string `json:"artifact_identifier"`
}

type AwsCodebuildProjectSpecCache struct {
	Type     string   `json:"type"`
	Location string   `json:"location"`
	Modes    []string `json:"modes"`
}

type AwsCodebuildProjectSpecSecondarySourcesAuth struct {
	Resource string `json:"resource"`
	Type     string `json:"type"`
}

type AwsCodebuildProjectSpecSecondarySources struct {
	SourceIdentifier  string                                    `json:"source_identifier"`
	Auth              []AwsCodebuildProjectSpecSecondarySources `json:"auth"`
	Buildspec         string                                    `json:"buildspec"`
	Location          string                                    `json:"location"`
	Type              string                                    `json:"type"`
	GitCloneDepth     int                                       `json:"git_clone_depth"`
	InsecureSsl       bool                                      `json:"insecure_ssl"`
	ReportBuildStatus bool                                      `json:"report_build_status"`
}

type AwsCodebuildProjectSpecSourceAuth struct {
	Resource string `json:"resource"`
	Type     string `json:"type"`
}

type AwsCodebuildProjectSpecSource struct {
	GitCloneDepth     int                             `json:"git_clone_depth"`
	InsecureSsl       bool                            `json:"insecure_ssl"`
	ReportBuildStatus bool                            `json:"report_build_status"`
	Auth              []AwsCodebuildProjectSpecSource `json:"auth"`
	Buildspec         string                          `json:"buildspec"`
	Location          string                          `json:"location"`
	Type              string                          `json:"type"`
}

type AwsCodebuildProjectSpec struct {
	VpcConfig          []AwsCodebuildProjectSpec `json:"vpc_config"`
	Artifacts          []AwsCodebuildProjectSpec `json:"artifacts"`
	Description        string                    `json:"description"`
	Environment        []AwsCodebuildProjectSpec `json:"environment"`
	LogsConfig         []AwsCodebuildProjectSpec `json:"logs_config"`
	Name               string                    `json:"name"`
	SecondaryArtifacts []AwsCodebuildProjectSpec `json:"secondary_artifacts"`
	Cache              []AwsCodebuildProjectSpec `json:"cache"`
	Tags               map[string]string         `json:"tags"`
	Arn                string                    `json:"arn"`
	EncryptionKey      string                    `json:"encryption_key"`
	SecondarySources   []AwsCodebuildProjectSpec `json:"secondary_sources"`
	Source             []AwsCodebuildProjectSpec `json:"source"`
	ServiceRole        string                    `json:"service_role"`
	BuildTimeout       int                       `json:"build_timeout"`
	BadgeEnabled       bool                      `json:"badge_enabled"`
	BadgeUrl           string                    `json:"badge_url"`
}



type AwsCodebuildProjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCodebuildProjectList is a list of AwsCodebuildProjects
type AwsCodebuildProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCodebuildProject CRD objects
	Items []AwsCodebuildProject `json:"items,omitempty"`
}