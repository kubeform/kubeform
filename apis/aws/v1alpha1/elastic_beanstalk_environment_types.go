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

type ElasticBeanstalkEnvironment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticBeanstalkEnvironmentSpec   `json:"spec,omitempty"`
	Status            ElasticBeanstalkEnvironmentStatus `json:"status,omitempty"`
}

type ElasticBeanstalkEnvironmentSpecSetting struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	// +optional
	Resource string `json:"resource,omitempty"`
	Value    string `json:"value"`
}

type ElasticBeanstalkEnvironmentSpec struct {
	Application string `json:"application"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +optional
	PollInterval string `json:"poll_interval,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Setting *[]ElasticBeanstalkEnvironmentSpec `json:"setting,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	TemplateName string `json:"template_name,omitempty"`
	// +optional
	Tier string `json:"tier,omitempty"`
	// +optional
	WaitForReadyTimeout string `json:"wait_for_ready_timeout,omitempty"`
}

type ElasticBeanstalkEnvironmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticBeanstalkEnvironmentList is a list of ElasticBeanstalkEnvironments
type ElasticBeanstalkEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticBeanstalkEnvironment CRD objects
	Items []ElasticBeanstalkEnvironment `json:"items,omitempty"`
}
