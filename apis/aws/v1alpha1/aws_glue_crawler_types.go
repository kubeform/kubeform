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

type AwsGlueCrawler struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlueCrawlerSpec   `json:"spec,omitempty"`
	Status            AwsGlueCrawlerStatus `json:"status,omitempty"`
}

type AwsGlueCrawlerSpecDynamodbTarget struct {
	Path string `json:"path"`
}

type AwsGlueCrawlerSpecS3Target struct {
	Path       string   `json:"path"`
	Exclusions []string `json:"exclusions"`
}

type AwsGlueCrawlerSpecJdbcTarget struct {
	ConnectionName string   `json:"connection_name"`
	Path           string   `json:"path"`
	Exclusions     []string `json:"exclusions"`
}

type AwsGlueCrawlerSpecSchemaChangePolicy struct {
	DeleteBehavior string `json:"delete_behavior"`
	UpdateBehavior string `json:"update_behavior"`
}

type AwsGlueCrawlerSpec struct {
	Role                  string               `json:"role"`
	Description           string               `json:"description"`
	TablePrefix           string               `json:"table_prefix"`
	DynamodbTarget        []AwsGlueCrawlerSpec `json:"dynamodb_target"`
	SecurityConfiguration string               `json:"security_configuration"`
	Classifiers           []string             `json:"classifiers"`
	S3Target              []AwsGlueCrawlerSpec `json:"s3_target"`
	JdbcTarget            []AwsGlueCrawlerSpec `json:"jdbc_target"`
	Name                  string               `json:"name"`
	Arn                   string               `json:"arn"`
	DatabaseName          string               `json:"database_name"`
	Schedule              string               `json:"schedule"`
	SchemaChangePolicy    []AwsGlueCrawlerSpec `json:"schema_change_policy"`
	Configuration         string               `json:"configuration"`
}

type AwsGlueCrawlerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGlueCrawlerList is a list of AwsGlueCrawlers
type AwsGlueCrawlerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlueCrawler CRD objects
	Items []AwsGlueCrawler `json:"items,omitempty"`
}
