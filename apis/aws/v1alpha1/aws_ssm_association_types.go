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

type AwsSsmAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmAssociationSpec   `json:"spec,omitempty"`
	Status            AwsSsmAssociationStatus `json:"status,omitempty"`
}

type AwsSsmAssociationSpecTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type AwsSsmAssociationSpecOutputLocation struct {
	S3BucketName string `json:"s3_bucket_name"`
	S3KeyPrefix  string `json:"s3_key_prefix"`
}

type AwsSsmAssociationSpec struct {
	AssociationName    string                  `json:"association_name"`
	DocumentVersion    string                  `json:"document_version"`
	MaxConcurrency     string                  `json:"max_concurrency"`
	Name               string                  `json:"name"`
	ScheduleExpression string                  `json:"schedule_expression"`
	Targets            []AwsSsmAssociationSpec `json:"targets"`
	AssociationId      string                  `json:"association_id"`
	InstanceId         string                  `json:"instance_id"`
	MaxErrors          string                  `json:"max_errors"`
	Parameters         map[string]string       `json:"parameters"`
	OutputLocation     []AwsSsmAssociationSpec `json:"output_location"`
	ComplianceSeverity string                  `json:"compliance_severity"`
}



type AwsSsmAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSsmAssociationList is a list of AwsSsmAssociations
type AwsSsmAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmAssociation CRD objects
	Items []AwsSsmAssociation `json:"items,omitempty"`
}