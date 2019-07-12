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

type AwsDmsEndpointSpecS3Settings struct {
	BucketName              string `json:"bucket_name"`
	CompressionType         string `json:"compression_type"`
	ServiceAccessRoleArn    string `json:"service_access_role_arn"`
	ExternalTableDefinition string `json:"external_table_definition"`
	CsvRowDelimiter         string `json:"csv_row_delimiter"`
	CsvDelimiter            string `json:"csv_delimiter"`
	BucketFolder            string `json:"bucket_folder"`
}

type AwsDmsEndpointSpecMongodbSettings struct {
	AuthType          string `json:"auth_type"`
	AuthMechanism     string `json:"auth_mechanism"`
	NestingLevel      string `json:"nesting_level"`
	ExtractDocId      string `json:"extract_doc_id"`
	DocsToInvestigate string `json:"docs_to_investigate"`
	AuthSource        string `json:"auth_source"`
}

type AwsDmsEndpointSpec struct {
	SslMode                   string               `json:"ssl_mode"`
	CertificateArn            string               `json:"certificate_arn"`
	DatabaseName              string               `json:"database_name"`
	ServiceAccessRole         string               `json:"service_access_role"`
	ExtraConnectionAttributes string               `json:"extra_connection_attributes"`
	Port                      int                  `json:"port"`
	S3Settings                []AwsDmsEndpointSpec `json:"s3_settings"`
	EndpointType              string               `json:"endpoint_type"`
	EngineName                string               `json:"engine_name"`
	KmsKeyArn                 string               `json:"kms_key_arn"`
	Password                  string               `json:"password"`
	ServerName                string               `json:"server_name"`
	Tags                      map[string]string    `json:"tags"`
	EndpointArn               string               `json:"endpoint_arn"`
	EndpointId                string               `json:"endpoint_id"`
	Username                  string               `json:"username"`
	MongodbSettings           []AwsDmsEndpointSpec `json:"mongodb_settings"`
}

type AwsDmsEndpointStatus struct {
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
