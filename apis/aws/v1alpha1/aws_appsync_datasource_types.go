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

type AwsAppsyncDatasource struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppsyncDatasourceSpec   `json:"spec,omitempty"`
	Status            AwsAppsyncDatasourceStatus `json:"status,omitempty"`
}

type AwsAppsyncDatasourceSpecDynamodbConfig struct {
	UseCallerCredentials bool   `json:"use_caller_credentials"`
	Region               string `json:"region"`
	TableName            string `json:"table_name"`
}

type AwsAppsyncDatasourceSpecElasticsearchConfig struct {
	Region   string `json:"region"`
	Endpoint string `json:"endpoint"`
}

type AwsAppsyncDatasourceSpecLambdaConfig struct {
	FunctionArn string `json:"function_arn"`
}

type AwsAppsyncDatasourceSpecHttpConfig struct {
	Endpoint string `json:"endpoint"`
}

type AwsAppsyncDatasourceSpec struct {
	ServiceRoleArn      string                     `json:"service_role_arn"`
	ApiId               string                     `json:"api_id"`
	Type                string                     `json:"type"`
	Description         string                     `json:"description"`
	DynamodbConfig      []AwsAppsyncDatasourceSpec `json:"dynamodb_config"`
	ElasticsearchConfig []AwsAppsyncDatasourceSpec `json:"elasticsearch_config"`
	LambdaConfig        []AwsAppsyncDatasourceSpec `json:"lambda_config"`
	Name                string                     `json:"name"`
	HttpConfig          []AwsAppsyncDatasourceSpec `json:"http_config"`
	Arn                 string                     `json:"arn"`
}

type AwsAppsyncDatasourceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAppsyncDatasourceList is a list of AwsAppsyncDatasources
type AwsAppsyncDatasourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppsyncDatasource CRD objects
	Items []AwsAppsyncDatasource `json:"items,omitempty"`
}
