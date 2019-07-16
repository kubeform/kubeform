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

type GlueCrawler struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueCrawlerSpec   `json:"spec,omitempty"`
	Status            GlueCrawlerStatus `json:"status,omitempty"`
}

type GlueCrawlerSpecDynamodbTarget struct {
	Path string `json:"path"`
}

type GlueCrawlerSpecJdbcTarget struct {
	ConnectionName string `json:"connection_name"`
	// +optional
	Exclusions []string `json:"exclusions,omitempty"`
	Path       string   `json:"path"`
}

type GlueCrawlerSpecS3Target struct {
	// +optional
	Exclusions []string `json:"exclusions,omitempty"`
	Path       string   `json:"path"`
}

type GlueCrawlerSpecSchemaChangePolicy struct {
	// +optional
	DeleteBehavior string `json:"delete_behavior,omitempty"`
	// +optional
	UpdateBehavior string `json:"update_behavior,omitempty"`
}

type GlueCrawlerSpec struct {
	// +optional
	Classifiers []string `json:"classifiers,omitempty"`
	// +optional
	Configuration string `json:"configuration,omitempty"`
	DatabaseName  string `json:"database_name"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	DynamodbTarget *[]GlueCrawlerSpec `json:"dynamodb_target,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	JdbcTarget *[]GlueCrawlerSpec `json:"jdbc_target,omitempty"`
	Name       string             `json:"name"`
	Role       string             `json:"role"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	S3Target *[]GlueCrawlerSpec `json:"s3_target,omitempty"`
	// +optional
	Schedule string `json:"schedule,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SchemaChangePolicy *[]GlueCrawlerSpec `json:"schema_change_policy,omitempty"`
	// +optional
	SecurityConfiguration string `json:"security_configuration,omitempty"`
	// +optional
	TablePrefix string `json:"table_prefix,omitempty"`
}

type GlueCrawlerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlueCrawlerList is a list of GlueCrawlers
type GlueCrawlerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlueCrawler CRD objects
	Items []GlueCrawler `json:"items,omitempty"`
}
