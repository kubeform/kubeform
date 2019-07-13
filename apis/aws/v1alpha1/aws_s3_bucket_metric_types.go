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

type AwsS3BucketMetric struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsS3BucketMetricSpec   `json:"spec,omitempty"`
	Status            AwsS3BucketMetricStatus `json:"status,omitempty"`
}

type AwsS3BucketMetricSpecFilter struct {
	Prefix string            `json:"prefix"`
	Tags   map[string]string `json:"tags"`
}

type AwsS3BucketMetricSpec struct {
	Name   string                  `json:"name"`
	Bucket string                  `json:"bucket"`
	Filter []AwsS3BucketMetricSpec `json:"filter"`
}



type AwsS3BucketMetricStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsS3BucketMetricList is a list of AwsS3BucketMetrics
type AwsS3BucketMetricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsS3BucketMetric CRD objects
	Items []AwsS3BucketMetric `json:"items,omitempty"`
}