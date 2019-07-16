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

type DmsEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsEndpointSpec   `json:"spec,omitempty"`
	Status            DmsEndpointStatus `json:"status,omitempty"`
}

type DmsEndpointSpecMongodbSettings struct {
	// +optional
	AuthMechanism string `json:"auth_mechanism,omitempty"`
	// +optional
	AuthSource string `json:"auth_source,omitempty"`
	// +optional
	AuthType string `json:"auth_type,omitempty"`
	// +optional
	DocsToInvestigate string `json:"docs_to_investigate,omitempty"`
	// +optional
	ExtractDocId string `json:"extract_doc_id,omitempty"`
	// +optional
	NestingLevel string `json:"nesting_level,omitempty"`
}

type DmsEndpointSpecS3Settings struct {
	// +optional
	BucketFolder string `json:"bucket_folder,omitempty"`
	// +optional
	BucketName string `json:"bucket_name,omitempty"`
	// +optional
	CompressionType string `json:"compression_type,omitempty"`
	// +optional
	CsvDelimiter string `json:"csv_delimiter,omitempty"`
	// +optional
	CsvRowDelimiter string `json:"csv_row_delimiter,omitempty"`
	// +optional
	ExternalTableDefinition string `json:"external_table_definition,omitempty"`
	// +optional
	ServiceAccessRoleArn string `json:"service_access_role_arn,omitempty"`
}

type DmsEndpointSpec struct {
	// +optional
	DatabaseName string `json:"database_name,omitempty"`
	EndpointId   string `json:"endpoint_id"`
	EndpointType string `json:"endpoint_type"`
	EngineName   string `json:"engine_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MongodbSettings *[]DmsEndpointSpec `json:"mongodb_settings,omitempty"`
	// +optional
	Password string `json:"password,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3Settings *[]DmsEndpointSpec `json:"s3_settings,omitempty"`
	// +optional
	ServerName string `json:"server_name,omitempty"`
	// +optional
	ServiceAccessRole string `json:"service_access_role,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	Username string `json:"username,omitempty"`
}

type DmsEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DmsEndpointList is a list of DmsEndpoints
type DmsEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DmsEndpoint CRD objects
	Items []DmsEndpoint `json:"items,omitempty"`
}
