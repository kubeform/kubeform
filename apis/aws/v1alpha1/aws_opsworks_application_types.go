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

type AwsOpsworksApplicationSpecAppSource struct {
	Type     string `json:"type"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Revision string `json:"revision"`
	SshKey   string `json:"ssh_key"`
}

type AwsOpsworksApplicationSpecEnvironment struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Secure bool   `json:"secure"`
}

type AwsOpsworksApplicationSpecSslConfiguration struct {
	Chain       string `json:"chain"`
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

type AwsOpsworksApplicationSpec struct {
	Description            string                       `json:"description"`
	EnableSsl              bool                         `json:"enable_ssl"`
	AppSource              []AwsOpsworksApplicationSpec `json:"app_source"`
	StackId                string                       `json:"stack_id"`
	RailsEnv               string                       `json:"rails_env"`
	AwsFlowRubySettings    string                       `json:"aws_flow_ruby_settings"`
	DataSourceType         string                       `json:"data_source_type"`
	DataSourceDatabaseName string                       `json:"data_source_database_name"`
	Domains                []string                     `json:"domains"`
	Environment            []AwsOpsworksApplicationSpec `json:"environment"`
	Type                   string                       `json:"type"`
	SslConfiguration       []AwsOpsworksApplicationSpec `json:"ssl_configuration"`
	ShortName              string                       `json:"short_name"`
	DocumentRoot           string                       `json:"document_root"`
	AutoBundleOnDeploy     string                       `json:"auto_bundle_on_deploy"`
	DataSourceArn          string                       `json:"data_source_arn"`
	Name                   string                       `json:"name"`
}



type AwsOpsworksApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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