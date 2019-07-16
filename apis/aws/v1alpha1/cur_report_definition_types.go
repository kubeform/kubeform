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

type CurReportDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CurReportDefinitionSpec   `json:"spec,omitempty"`
	Status            CurReportDefinitionStatus `json:"status,omitempty"`
}

type CurReportDefinitionSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AdditionalArtifacts []string `json:"additional_artifacts,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	AdditionalSchemaElements []string `json:"additional_schema_elements"`
	Compression              string   `json:"compression"`
	Format                   string   `json:"format"`
	ReportName               string   `json:"report_name"`
	S3Bucket                 string   `json:"s3_bucket"`
	// +optional
	S3Prefix string `json:"s3_prefix,omitempty"`
	S3Region string `json:"s3_region"`
	TimeUnit string `json:"time_unit"`
}

type CurReportDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CurReportDefinitionList is a list of CurReportDefinitions
type CurReportDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CurReportDefinition CRD objects
	Items []CurReportDefinition `json:"items,omitempty"`
}
