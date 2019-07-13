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

type AwsDmsEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDmsEndpointSpec   `json:"spec,omitempty"`
	Status            AwsDmsEndpointStatus `json:"status,omitempty"`
}

type AwsDmsEndpointSpecMongodbSettings struct {
	AuthMechanism     string `json:"auth_mechanism"`
	NestingLevel      string `json:"nesting_level"`
	ExtractDocId      string `json:"extract_doc_id"`
	DocsToInvestigate string `json:"docs_to_investigate"`
	AuthSource        string `json:"auth_source"`
	AuthType          string `json:"auth_type"`
}

type AwsDmsEndpointSpecS3Settings struct {
	ServiceAccessRoleArn    string `json:"service_access_role_arn"`
	ExternalTableDefinition string `json:"external_table_definition"`
	CsvRowDelimiter         string `json:"csv_row_delimiter"`
	CsvDelimiter            string `json:"csv_delimiter"`
	BucketFolder            string `json:"bucket_folder"`
	BucketName              string `json:"bucket_name"`
	CompressionType         string `json:"compression_type"`
}

type AwsDmsEndpointSpec struct {
	Port                      int                  `json:"port"`
	ServerName                string               `json:"server_name"`
	CertificateArn            string               `json:"certificate_arn"`
	SslMode                   string               `json:"ssl_mode"`
	MongodbSettings           []AwsDmsEndpointSpec `json:"mongodb_settings"`
	DatabaseName              string               `json:"database_name"`
	EngineName                string               `json:"engine_name"`
	Password                  string               `json:"password"`
	EndpointId                string               `json:"endpoint_id"`
	ServiceAccessRole         string               `json:"service_access_role"`
	EndpointType              string               `json:"endpoint_type"`
	ExtraConnectionAttributes string               `json:"extra_connection_attributes"`
	KmsKeyArn                 string               `json:"kms_key_arn"`
	Tags                      map[string]string    `json:"tags"`
	Username                  string               `json:"username"`
	S3Settings                []AwsDmsEndpointSpec `json:"s3_settings"`
	EndpointArn               string               `json:"endpoint_arn"`
}



type AwsDmsEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDmsEndpointList is a list of AwsDmsEndpoints
type AwsDmsEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDmsEndpoint CRD objects
	Items []AwsDmsEndpoint `json:"items,omitempty"`
}