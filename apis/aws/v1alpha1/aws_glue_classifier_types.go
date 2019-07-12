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

type AwsGlueClassifier struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlueClassifierSpec   `json:"spec,omitempty"`
	Status            AwsGlueClassifierStatus `json:"status,omitempty"`
}

type AwsGlueClassifierSpecXmlClassifier struct {
	Classification string `json:"classification"`
	RowTag         string `json:"row_tag"`
}

type AwsGlueClassifierSpecGrokClassifier struct {
	GrokPattern    string `json:"grok_pattern"`
	Classification string `json:"classification"`
	CustomPatterns string `json:"custom_patterns"`
}

type AwsGlueClassifierSpecJsonClassifier struct {
	JsonPath string `json:"json_path"`
}

type AwsGlueClassifierSpec struct {
	XmlClassifier  []AwsGlueClassifierSpec `json:"xml_classifier"`
	GrokClassifier []AwsGlueClassifierSpec `json:"grok_classifier"`
	JsonClassifier []AwsGlueClassifierSpec `json:"json_classifier"`
	Name           string                  `json:"name"`
}

type AwsGlueClassifierStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGlueClassifierList is a list of AwsGlueClassifiers
type AwsGlueClassifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlueClassifier CRD objects
	Items []AwsGlueClassifier `json:"items,omitempty"`
}
