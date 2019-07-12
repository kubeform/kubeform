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

type AwsElasticBeanstalkApplicationVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElasticBeanstalkApplicationVersionSpec   `json:"spec,omitempty"`
	Status            AwsElasticBeanstalkApplicationVersionStatus `json:"status,omitempty"`
}

type AwsElasticBeanstalkApplicationVersionSpec struct {
	Arn         string            `json:"arn"`
	Description string            `json:"description"`
	Bucket      string            `json:"bucket"`
	Key         string            `json:"key"`
	Name        string            `json:"name"`
	ForceDelete bool              `json:"force_delete"`
	Tags        map[string]string `json:"tags"`
	Application string            `json:"application"`
}

type AwsElasticBeanstalkApplicationVersionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsElasticBeanstalkApplicationVersionList is a list of AwsElasticBeanstalkApplicationVersions
type AwsElasticBeanstalkApplicationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticBeanstalkApplicationVersion CRD objects
	Items []AwsElasticBeanstalkApplicationVersion `json:"items,omitempty"`
}
