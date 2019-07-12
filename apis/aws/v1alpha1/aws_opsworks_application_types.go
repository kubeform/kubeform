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

type AwsOpsworksApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksApplicationSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksApplicationStatus `json:"status,omitempty"`
}

type AwsOpsworksApplicationSpecSslConfiguration struct {
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
	Chain       string `json:"chain"`
}

type AwsOpsworksApplicationSpecAppSource struct {
	Revision string `json:"revision"`
	SshKey   string `json:"ssh_key"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AwsOpsworksApplicationSpecEnvironment struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Secure bool   `json:"secure"`
}

type AwsOpsworksApplicationSpec struct {
	Type                   string                       `json:"type"`
	EnableSsl              bool                         `json:"enable_ssl"`
	SslConfiguration       []AwsOpsworksApplicationSpec `json:"ssl_configuration"`
	DocumentRoot           string                       `json:"document_root"`
	DataSourceType         string                       `json:"data_source_type"`
	DataSourceArn          string                       `json:"data_source_arn"`
	ShortName              string                       `json:"short_name"`
	StackId                string                       `json:"stack_id"`
	AutoBundleOnDeploy     string                       `json:"auto_bundle_on_deploy"`
	DataSourceDatabaseName string                       `json:"data_source_database_name"`
	Description            string                       `json:"description"`
	Domains                []string                     `json:"domains"`
	Name                   string                       `json:"name"`
	RailsEnv               string                       `json:"rails_env"`
	AwsFlowRubySettings    string                       `json:"aws_flow_ruby_settings"`
	AppSource              []AwsOpsworksApplicationSpec `json:"app_source"`
	Environment            []AwsOpsworksApplicationSpec `json:"environment"`
}

type AwsOpsworksApplicationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksApplicationList is a list of AwsOpsworksApplications
type AwsOpsworksApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksApplication CRD objects
	Items []AwsOpsworksApplication `json:"items,omitempty"`
}
