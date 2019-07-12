package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCurReportDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCurReportDefinitionSpec   `json:"spec,omitempty"`
	Status            AwsCurReportDefinitionStatus `json:"status,omitempty"`
}

type AwsCurReportDefinitionSpec struct {
	TimeUnit                 string   `json:"time_unit"`
	Format                   string   `json:"format"`
	AdditionalSchemaElements []string `json:"additional_schema_elements"`
	S3Region                 string   `json:"s3_region"`
	AdditionalArtifacts      []string `json:"additional_artifacts"`
	ReportName               string   `json:"report_name"`
	Compression              string   `json:"compression"`
	S3Bucket                 string   `json:"s3_bucket"`
	S3Prefix                 string   `json:"s3_prefix"`
}

type AwsCurReportDefinitionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCurReportDefinitionList is a list of AwsCurReportDefinitions
type AwsCurReportDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCurReportDefinition CRD objects
	Items []AwsCurReportDefinition `json:"items,omitempty"`
}
