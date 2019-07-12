package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMqConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsMqConfigurationSpec   `json:"spec,omitempty"`
	Status            AwsMqConfigurationStatus `json:"status,omitempty"`
}

type AwsMqConfigurationSpec struct {
	LatestRevision int               `json:"latest_revision"`
	Tags           map[string]string `json:"tags"`
	Arn            string            `json:"arn"`
	Data           string            `json:"data"`
	Description    string            `json:"description"`
	EngineType     string            `json:"engine_type"`
	EngineVersion  string            `json:"engine_version"`
	Name           string            `json:"name"`
}



type AwsMqConfigurationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMqConfigurationList is a list of AwsMqConfigurations
type AwsMqConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsMqConfiguration CRD objects
	Items []AwsMqConfiguration `json:"items,omitempty"`
}