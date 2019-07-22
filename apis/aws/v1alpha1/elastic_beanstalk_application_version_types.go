package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	Application string                    `json:"application" tf:"application"`
	Bucket      string                    `json:"bucket" tf:"bucket"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	ForceDelete bool   `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`
	Key         string `json:"key" tf:"key"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ElasticBeanstalkApplicationVersionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
