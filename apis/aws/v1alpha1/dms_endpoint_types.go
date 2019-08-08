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

type DmsEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsEndpointSpec   `json:"spec,omitempty"`
	Status            DmsEndpointStatus `json:"status,omitempty"`
}

type DmsEndpointSpecMongodbSettings struct {
	// +optional
	AuthMechanism string `json:"authMechanism,omitempty" tf:"auth_mechanism,omitempty"`
	// +optional
	AuthSource string `json:"authSource,omitempty" tf:"auth_source,omitempty"`
	// +optional
	AuthType string `json:"authType,omitempty" tf:"auth_type,omitempty"`
	// +optional
	DocsToInvestigate string `json:"docsToInvestigate,omitempty" tf:"docs_to_investigate,omitempty"`
	// +optional
	ExtractDocID string `json:"extractDocID,omitempty" tf:"extract_doc_id,omitempty"`
	// +optional
	NestingLevel string `json:"nestingLevel,omitempty" tf:"nesting_level,omitempty"`
}

type DmsEndpointSpecS3Settings struct {
	// +optional
	BucketFolder string `json:"bucketFolder,omitempty" tf:"bucket_folder,omitempty"`
	// +optional
	BucketName string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`
	// +optional
	CompressionType string `json:"compressionType,omitempty" tf:"compression_type,omitempty"`
	// +optional
	CsvDelimiter string `json:"csvDelimiter,omitempty" tf:"csv_delimiter,omitempty"`
	// +optional
	CsvRowDelimiter string `json:"csvRowDelimiter,omitempty" tf:"csv_row_delimiter,omitempty"`
	// +optional
	ExternalTableDefinition string `json:"externalTableDefinition,omitempty" tf:"external_table_definition,omitempty"`
	// +optional
	ServiceAccessRoleArn string `json:"serviceAccessRoleArn,omitempty" tf:"service_access_role_arn,omitempty"`
}

type DmsEndpointSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	KubeFormSecret *core.LocalObjectReference `json:"kubeFormSecret,omitempty" tf:"-"`

	// +optional
	CertificateArn string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`
	// +optional
	DatabaseName string `json:"databaseName,omitempty" tf:"database_name,omitempty"`
	// +optional
	EndpointArn  string `json:"endpointArn,omitempty" tf:"endpoint_arn,omitempty"`
	EndpointID   string `json:"endpointID" tf:"endpoint_id"`
	EndpointType string `json:"endpointType" tf:"endpoint_type"`
	EngineName   string `json:"engineName" tf:"engine_name"`
	// +optional
	ExtraConnectionAttributes string `json:"extraConnectionAttributes,omitempty" tf:"extra_connection_attributes,omitempty"`
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MongodbSettings []DmsEndpointSpecMongodbSettings `json:"mongodbSettings,omitempty" tf:"mongodb_settings,omitempty"`
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3Settings []DmsEndpointSpecS3Settings `json:"s3Settings,omitempty" tf:"s3_settings,omitempty"`
	// +optional
	ServerName string `json:"serverName,omitempty" tf:"server_name,omitempty"`
	// +optional
	ServiceAccessRole string `json:"serviceAccessRole,omitempty" tf:"service_access_role,omitempty"`
	// +optional
	SslMode string `json:"sslMode,omitempty" tf:"ssl_mode,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
}

type DmsEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DmsEndpointSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
