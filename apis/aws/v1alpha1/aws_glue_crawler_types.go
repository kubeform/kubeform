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

type AwsGlueCrawlerSpecS3Target struct {
	Path       string   `json:"path"`
	Exclusions []string `json:"exclusions"`
}

type AwsGlueCrawlerSpecJdbcTarget struct {
	Exclusions     []string `json:"exclusions"`
	ConnectionName string   `json:"connection_name"`
	Path           string   `json:"path"`
}

type AwsGlueCrawlerSpecSchemaChangePolicy struct {
	DeleteBehavior string `json:"delete_behavior"`
	UpdateBehavior string `json:"update_behavior"`
}

type AwsGlueCrawlerSpecDynamodbTarget struct {
	Path string `json:"path"`
}

type AwsGlueCrawlerSpec struct {
	Classifiers           []string             `json:"classifiers"`
	TablePrefix           string               `json:"table_prefix"`
	S3Target              []AwsGlueCrawlerSpec `json:"s3_target"`
	Name                  string               `json:"name"`
	Description           string               `json:"description"`
	DatabaseName          string               `json:"database_name"`
	SecurityConfiguration string               `json:"security_configuration"`
	Role                  string               `json:"role"`
	JdbcTarget            []AwsGlueCrawlerSpec `json:"jdbc_target"`
	SchemaChangePolicy    []AwsGlueCrawlerSpec `json:"schema_change_policy"`
	DynamodbTarget        []AwsGlueCrawlerSpec `json:"dynamodb_target"`
	Configuration         string               `json:"configuration"`
	Arn                   string               `json:"arn"`
	Schedule              string               `json:"schedule"`
}



type AwsGlueCrawlerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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