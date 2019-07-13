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

type AwsMqConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsMqConfigurationSpec   `json:"spec,omitempty"`
	Status            AwsMqConfigurationStatus `json:"status,omitempty"`
}

type AwsMqConfigurationSpec struct {
	Name           string            `json:"name"`
	LatestRevision int               `json:"latest_revision"`
	Tags           map[string]string `json:"tags"`
	Arn            string            `json:"arn"`
	Data           string            `json:"data"`
	Description    string            `json:"description"`
	EngineType     string            `json:"engine_type"`
	EngineVersion  string            `json:"engine_version"`
}



type AwsMqConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsMqConfigurationList is a list of AwsMqConfigurations
type AwsMqConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsMqConfiguration CRD objects
	Items []AwsMqConfiguration `json:"items,omitempty"`
}