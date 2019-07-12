package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCodebuildProject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCodebuildProjectSpec   `json:"spec,omitempty"`
	Status            AwsCodebuildProjectStatus `json:"status,omitempty"`
}

type AwsCodebuildProjectSpecEnvironmentEnvironmentVariable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type AwsCodebuildProjectSpecEnvironment struct {
	ImagePullCredentialsType string                               `json:"image_pull_credentials_type"`
	PrivilegedMode           bool                                 `json:"privileged_mode"`
	Certificate              string                               `json:"certificate"`
	ComputeType              string                               `json:"compute_type"`
	EnvironmentVariable      []AwsCodebuildProjectSpecEnvironment `json:"environment_variable"`
	Image                    string                               `json:"image"`
	Type                     string                               `json:"type"`
}

type AwsCodebuildProjectSpecSourceAuth struct {
	Resource string `json:"resource"`
	Type     string `json:"type"`
}

type AwsCodebuildProjectSpecSource struct {
	Location          string                          `json:"location"`
	Type              string                          `json:"type"`
	GitCloneDepth     int                             `json:"git_clone_depth"`
	InsecureSsl       bool                            `json:"insecure_ssl"`
	ReportBuildStatus bool                            `json:"report_build_status"`
	Auth              []AwsCodebuildProjectSpecSource `json:"auth"`
	Buildspec         string                          `json:"buildspec"`
}

type AwsCodebuildProjectSpecVpcConfig struct {
	VpcId            string   `json:"vpc_id"`
	Subnets          []string `json:"subnets"`
	SecurityGroupIds []string `json:"security_group_ids"`
}

type AwsCodebuildProjectSpecArtifacts struct {
	EncryptionDisabled bool   `json:"encryption_disabled"`
	Location           string `json:"location"`
	NamespaceType      string `json:"namespace_type"`
	Packaging          string `json:"packaging"`
	Path               string `json:"path"`
	Type               string `json:"type"`
	Name               string `json:"name"`
}

type AwsCodebuildProjectSpecCache struct {
	Type     string `json:"type"`
	Location string `json:"location"`
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

type AwsCodebuildProjectSpecSecondarySourcesAuth struct {
	Resource string `json:"resource"`
	Type     string `json:"type"`
}

type AwsCodebuildProjectSpecSecondarySources struct {
	Auth              []AwsCodebuildProjectSpecSecondarySources `json:"auth"`
	Buildspec         string                                    `json:"buildspec"`
	Location          string                                    `json:"location"`
	Type              string                                    `json:"type"`
	GitCloneDepth     int                                       `json:"git_clone_depth"`
	InsecureSsl       bool                                      `json:"insecure_ssl"`
	ReportBuildStatus bool                                      `json:"report_build_status"`
	SourceIdentifier  string                                    `json:"source_identifier"`
}

type AwsCodebuildProjectSpec struct {
	Environment        []AwsCodebuildProjectSpec `json:"environment"`
	Name               string                    `json:"name"`
	Source             []AwsCodebuildProjectSpec `json:"source"`
	VpcConfig          []AwsCodebuildProjectSpec `json:"vpc_config"`
	Artifacts          []AwsCodebuildProjectSpec `json:"artifacts"`
	Cache              []AwsCodebuildProjectSpec `json:"cache"`
	SecondaryArtifacts []AwsCodebuildProjectSpec `json:"secondary_artifacts"`
	SecondarySources   []AwsCodebuildProjectSpec `json:"secondary_sources"`
	BadgeEnabled       bool                      `json:"badge_enabled"`
	BadgeUrl           string                    `json:"badge_url"`
	Description        string                    `json:"description"`
	EncryptionKey      string                    `json:"encryption_key"`
	Tags               map[string]string         `json:"tags"`
	Arn                string                    `json:"arn"`
	ServiceRole        string                    `json:"service_role"`
	BuildTimeout       int                       `json:"build_timeout"`
}

type AwsCodebuildProjectStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodebuildProjectList is a list of AwsCodebuildProjects
type AwsCodebuildProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCodebuildProject CRD objects
	Items []AwsCodebuildProject `json:"items,omitempty"`
}
