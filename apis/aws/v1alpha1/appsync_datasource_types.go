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

type AppsyncDatasource struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppsyncDatasourceSpec   `json:"spec,omitempty"`
	Status            AppsyncDatasourceStatus `json:"status,omitempty"`
}

type AppsyncDatasourceSpecDynamodbConfig struct {
	TableName string `json:"table_name"`
	// +optional
	UseCallerCredentials bool `json:"use_caller_credentials,omitempty"`
}

type AppsyncDatasourceSpecElasticsearchConfig struct {
	Endpoint string `json:"endpoint"`
}

type AppsyncDatasourceSpecHttpConfig struct {
	Endpoint string `json:"endpoint"`
}

type AppsyncDatasourceSpecLambdaConfig struct {
	FunctionArn string `json:"function_arn"`
}

type AppsyncDatasourceSpec struct {
	ApiId string `json:"api_id"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DynamodbConfig *[]AppsyncDatasourceSpec `json:"dynamodb_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ElasticsearchConfig *[]AppsyncDatasourceSpec `json:"elasticsearch_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpConfig *[]AppsyncDatasourceSpec `json:"http_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LambdaConfig *[]AppsyncDatasourceSpec `json:"lambda_config,omitempty"`
	Name         string                   `json:"name"`
	// +optional
	ServiceRoleArn string `json:"service_role_arn,omitempty"`
	Type           string `json:"type"`
}

type AppsyncDatasourceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppsyncDatasourceList is a list of AppsyncDatasources
type AppsyncDatasourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppsyncDatasource CRD objects
	Items []AppsyncDatasource `json:"items,omitempty"`
}
