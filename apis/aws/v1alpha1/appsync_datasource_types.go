package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	// +optional
	Region    string `json:"region,omitempty" tf:"region,omitempty"`
	TableName string `json:"tableName" tf:"table_name"`
	// +optional
	UseCallerCredentials bool `json:"useCallerCredentials,omitempty" tf:"use_caller_credentials,omitempty"`
}

type AppsyncDatasourceSpecElasticsearchConfig struct {
	Endpoint string `json:"endpoint" tf:"endpoint"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
}

type AppsyncDatasourceSpecHttpConfig struct {
	Endpoint string `json:"endpoint" tf:"endpoint"`
}

type AppsyncDatasourceSpecLambdaConfig struct {
	FunctionArn string `json:"functionArn" tf:"function_arn"`
}

type AppsyncDatasourceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ApiID string `json:"apiID" tf:"api_id"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DynamodbConfig []AppsyncDatasourceSpecDynamodbConfig `json:"dynamodbConfig,omitempty" tf:"dynamodb_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ElasticsearchConfig []AppsyncDatasourceSpecElasticsearchConfig `json:"elasticsearchConfig,omitempty" tf:"elasticsearch_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpConfig []AppsyncDatasourceSpecHttpConfig `json:"httpConfig,omitempty" tf:"http_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LambdaConfig []AppsyncDatasourceSpecLambdaConfig `json:"lambdaConfig,omitempty" tf:"lambda_config,omitempty"`
	Name         string                              `json:"name" tf:"name"`
	// +optional
	ServiceRoleArn string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`
	Type           string `json:"type" tf:"type"`
}

type AppsyncDatasourceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppsyncDatasourceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
