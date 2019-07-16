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

type SsmAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmAssociationSpec   `json:"spec,omitempty"`
	Status            SsmAssociationStatus `json:"status,omitempty"`
}

type SsmAssociationSpecOutputLocation struct {
	S3BucketName string `json:"s3_bucket_name"`
	// +optional
	S3KeyPrefix string `json:"s3_key_prefix,omitempty"`
}

type SsmAssociationSpec struct {
	// +optional
	AssociationName string `json:"association_name,omitempty"`
	// +optional
	ComplianceSeverity string `json:"compliance_severity,omitempty"`
	// +optional
	InstanceId string `json:"instance_id,omitempty"`
	// +optional
	MaxConcurrency string `json:"max_concurrency,omitempty"`
	// +optional
	MaxErrors string `json:"max_errors,omitempty"`
	Name      string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OutputLocation *[]SsmAssociationSpec `json:"output_location,omitempty"`
	// +optional
	ScheduleExpression string `json:"schedule_expression,omitempty"`
}

type SsmAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmAssociationList is a list of SsmAssociations
type SsmAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmAssociation CRD objects
	Items []SsmAssociation `json:"items,omitempty"`
}
