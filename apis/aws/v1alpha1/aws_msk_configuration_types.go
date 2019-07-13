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

type AwsMskConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsMskConfigurationSpec   `json:"spec,omitempty"`
	Status            AwsMskConfigurationStatus `json:"status,omitempty"`
}

type AwsMskConfigurationSpec struct {
	Arn              string   `json:"arn"`
	Description      string   `json:"description"`
	KafkaVersions    []string `json:"kafka_versions"`
	LatestRevision   int      `json:"latest_revision"`
	Name             string   `json:"name"`
	ServerProperties string   `json:"server_properties"`
}



type AwsMskConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsMskConfigurationList is a list of AwsMskConfigurations
type AwsMskConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsMskConfiguration CRD objects
	Items []AwsMskConfiguration `json:"items,omitempty"`
}