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

type AwsCurReportDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCurReportDefinitionSpec   `json:"spec,omitempty"`
	Status            AwsCurReportDefinitionStatus `json:"status,omitempty"`
}

type AwsCurReportDefinitionSpec struct {
	Format                   string   `json:"format"`
	Compression              string   `json:"compression"`
	AdditionalSchemaElements []string `json:"additional_schema_elements"`
	AdditionalArtifacts      []string `json:"additional_artifacts"`
	S3Region                 string   `json:"s3_region"`
	ReportName               string   `json:"report_name"`
	TimeUnit                 string   `json:"time_unit"`
	S3Bucket                 string   `json:"s3_bucket"`
	S3Prefix                 string   `json:"s3_prefix"`
}



type AwsCurReportDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCurReportDefinitionList is a list of AwsCurReportDefinitions
type AwsCurReportDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCurReportDefinition CRD objects
	Items []AwsCurReportDefinition `json:"items,omitempty"`
}