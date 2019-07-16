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

type ElasticBeanstalkApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticBeanstalkApplicationSpec   `json:"spec,omitempty"`
	Status            ElasticBeanstalkApplicationStatus `json:"status,omitempty"`
}

type ElasticBeanstalkApplicationSpecAppversionLifecycle struct {
	// +optional
	DeleteSourceFromS3 bool `json:"delete_source_from_s3,omitempty"`
	// +optional
	MaxAgeInDays int `json:"max_age_in_days,omitempty"`
	// +optional
	MaxCount    int    `json:"max_count,omitempty"`
	ServiceRole string `json:"service_role"`
}

type ElasticBeanstalkApplicationSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AppversionLifecycle *[]ElasticBeanstalkApplicationSpec `json:"appversion_lifecycle,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type ElasticBeanstalkApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticBeanstalkApplicationList is a list of ElasticBeanstalkApplications
type ElasticBeanstalkApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticBeanstalkApplication CRD objects
	Items []ElasticBeanstalkApplication `json:"items,omitempty"`
}
