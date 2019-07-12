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

type AwsCodebuildProjectSpecSourceAuth struct {
	Resource string `json:"resource"`
	Type     string `json:"type"`
}

type AwsCodebuildProjectSpecSource struct {
	Type              string                          `json:"type"`
	GitCloneDepth     int                             `json:"git_clone_depth"`
	InsecureSsl       bool                            `json:"insecure_ssl"`
	ReportBuildStatus bool                            `json:"report_build_status"`
	Auth              []AwsCodebuildProjectSpecSource `json:"auth"`
	Buildspec         string                          `json:"buildspec"`
	Location          string                          `json:"location"`
}

type AwsCodebuildProjectSpecEnvironmentRegistryCredential struct {
	Credential         string `json:"credential"`
	CredentialProvider string `json:"credential_provider"`
}

type AwsCodebuildProjectSpecEnvironmentEnvironmentVariable struct {
	Value string `json:"value"`
	Type  string `json:"type"`
	Name  string `json:"name"`
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

type AwsCodebuildProjectSpecSecondaryArtifacts struct {
	Name               string `json:"name"`
	EncryptionDisabled bool   `json:"encryption_disabled"`
	Location           string `json:"location"`
	NamespaceType      string `json:"namespace_type"`
	Packaging          string `json:"packaging"`
	Path               string `json:"path"`
	ArtifactIdentifier string `json:"artifact_identifier"`
	Type               string `json:"type"`
}

type AwsCodebuildProjectSpecArtifacts struct {
	Packaging          string `json:"packaging"`
	Path               string `json:"path"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	EncryptionDisabled bool   `json:"encryption_disabled"`
	Location           string `json:"location"`
	NamespaceType      string `json:"namespace_type"`
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

type AwsCodebuildProjectSpecVpcConfig struct {
	VpcId            string   `json:"vpc_id"`
	Subnets          []string `json:"subnets"`
	SecurityGroupIds []string `json:"security_group_ids"`
}

type AwsCodebuildProjectSpecCache struct {
	Modes    []string `json:"modes"`
	Type     string   `json:"type"`
	Location string   `json:"location"`
}

type AwsCodebuildProjectSpecSecondarySourcesAuth struct {
	Resource string `json:"resource"`
	Type     string `json:"type"`
}

type AwsCodebuildProjectSpecSecondarySources struct {
	Buildspec         string                                    `json:"buildspec"`
	Location          string                                    `json:"location"`
	Type              string                                    `json:"type"`
	GitCloneDepth     int                                       `json:"git_clone_depth"`
	InsecureSsl       bool                                      `json:"insecure_ssl"`
	ReportBuildStatus bool                                      `json:"report_build_status"`
	SourceIdentifier  string                                    `json:"source_identifier"`
	Auth              []AwsCodebuildProjectSpecSecondarySources `json:"auth"`
}

type AwsCodebuildProjectSpec struct {
	Source             []AwsCodebuildProjectSpec `json:"source"`
	BadgeEnabled       bool                      `json:"badge_enabled"`
	Arn                string                    `json:"arn"`
	EncryptionKey      string                    `json:"encryption_key"`
	Environment        []AwsCodebuildProjectSpec `json:"environment"`
	Name               string                    `json:"name"`
	SecondaryArtifacts []AwsCodebuildProjectSpec `json:"secondary_artifacts"`
	ServiceRole        string                    `json:"service_role"`
	BadgeUrl           string                    `json:"badge_url"`
	Artifacts          []AwsCodebuildProjectSpec `json:"artifacts"`
	Description        string                    `json:"description"`
	Tags               map[string]string         `json:"tags"`
	LogsConfig         []AwsCodebuildProjectSpec `json:"logs_config"`
	VpcConfig          []AwsCodebuildProjectSpec `json:"vpc_config"`
	Cache              []AwsCodebuildProjectSpec `json:"cache"`
	SecondarySources   []AwsCodebuildProjectSpec `json:"secondary_sources"`
	BuildTimeout       int                       `json:"build_timeout"`
}

type AwsCodebuildProjectStatus struct {
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
