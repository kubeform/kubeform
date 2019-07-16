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

type OpsworksApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksApplicationSpec   `json:"spec,omitempty"`
	Status            OpsworksApplicationStatus `json:"status,omitempty"`
}

type OpsworksApplicationSpecEnvironment struct {
	Key string `json:"key"`
	// +optional
	Secure bool   `json:"secure,omitempty"`
	Value  string `json:"value"`
}

type OpsworksApplicationSpecSslConfiguration struct {
	Certificate string `json:"certificate"`
	// +optional
	Chain      string `json:"chain,omitempty"`
	PrivateKey string `json:"private_key"`
}

type OpsworksApplicationSpec struct {
	// +optional
	AutoBundleOnDeploy string `json:"auto_bundle_on_deploy,omitempty"`
	// +optional
	AwsFlowRubySettings string `json:"aws_flow_ruby_settings,omitempty"`
	// +optional
	DataSourceArn string `json:"data_source_arn,omitempty"`
	// +optional
	DataSourceDatabaseName string `json:"data_source_database_name,omitempty"`
	// +optional
	DataSourceType string `json:"data_source_type,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	DocumentRoot string `json:"document_root,omitempty"`
	// +optional
	Domains []string `json:"domains,omitempty"`
	// +optional
	EnableSsl bool `json:"enable_ssl,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Environment *[]OpsworksApplicationSpec `json:"environment,omitempty"`
	Name        string                     `json:"name"`
	// +optional
	RailsEnv string `json:"rails_env,omitempty"`
	// +optional
	SslConfiguration *[]OpsworksApplicationSpec `json:"ssl_configuration,omitempty"`
	StackId          string                     `json:"stack_id"`
	Type             string                     `json:"type"`
}

type OpsworksApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksApplicationList is a list of OpsworksApplications
type OpsworksApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksApplication CRD objects
	Items []OpsworksApplication `json:"items,omitempty"`
}
