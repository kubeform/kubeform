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

type ElasticBeanstalkApplicationVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticBeanstalkApplicationVersionSpec   `json:"spec,omitempty"`
	Status            ElasticBeanstalkApplicationVersionStatus `json:"status,omitempty"`
}

type ElasticBeanstalkApplicationVersionSpec struct {
	Application string `json:"application"`
	Bucket      string `json:"bucket"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	ForceDelete bool   `json:"force_delete,omitempty"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type ElasticBeanstalkApplicationVersionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticBeanstalkApplicationVersionList is a list of ElasticBeanstalkApplicationVersions
type ElasticBeanstalkApplicationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticBeanstalkApplicationVersion CRD objects
	Items []ElasticBeanstalkApplicationVersion `json:"items,omitempty"`
}
