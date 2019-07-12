package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppsyncDatasource struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppsyncDatasourceSpec   `json:"spec,omitempty"`
	Status            AwsAppsyncDatasourceStatus `json:"status,omitempty"`
}

type AwsAppsyncDatasourceSpecElasticsearchConfig struct {
	Region   string `json:"region"`
	Endpoint string `json:"endpoint"`
}

type AwsAppsyncDatasourceSpecDynamodbConfig struct {
	TableName            string `json:"table_name"`
	UseCallerCredentials bool   `json:"use_caller_credentials"`
	Region               string `json:"region"`
}

type AwsAppsyncDatasourceSpecHttpConfig struct {
	Endpoint string `json:"endpoint"`
}

type AwsAppsyncDatasourceSpecLambdaConfig struct {
	FunctionArn string `json:"function_arn"`
}

type AwsAppsyncDatasourceSpec struct {
	Name                string                     `json:"name"`
	Type                string                     `json:"type"`
	ElasticsearchConfig []AwsAppsyncDatasourceSpec `json:"elasticsearch_config"`
	ServiceRoleArn      string                     `json:"service_role_arn"`
	ApiId               string                     `json:"api_id"`
	Description         string                     `json:"description"`
	DynamodbConfig      []AwsAppsyncDatasourceSpec `json:"dynamodb_config"`
	HttpConfig          []AwsAppsyncDatasourceSpec `json:"http_config"`
	LambdaConfig        []AwsAppsyncDatasourceSpec `json:"lambda_config"`
	Arn                 string                     `json:"arn"`
}

type AwsAppsyncDatasourceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppsyncDatasourceList is a list of AwsAppsyncDatasources
type AwsAppsyncDatasourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppsyncDatasource CRD objects
	Items []AwsAppsyncDatasource `json:"items,omitempty"`
}
