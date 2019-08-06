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

type CurReportDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CurReportDefinitionSpec   `json:"spec,omitempty"`
	Status            CurReportDefinitionStatus `json:"status,omitempty"`
}

type CurReportDefinitionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AdditionalArtifacts []string `json:"additionalArtifacts,omitempty" tf:"additional_artifacts,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	AdditionalSchemaElements []string `json:"additionalSchemaElements" tf:"additional_schema_elements"`
	Compression              string   `json:"compression" tf:"compression"`
	Format                   string   `json:"format" tf:"format"`
	ReportName               string   `json:"reportName" tf:"report_name"`
	S3Bucket                 string   `json:"s3Bucket" tf:"s3_bucket"`
	// +optional
	S3Prefix string `json:"s3Prefix,omitempty" tf:"s3_prefix,omitempty"`
	S3Region string `json:"s3Region" tf:"s3_region"`
	TimeUnit string `json:"timeUnit" tf:"time_unit"`
}

type CurReportDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CurReportDefinitionSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
