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

type GlueClassifier struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueClassifierSpec   `json:"spec,omitempty"`
	Status            GlueClassifierStatus `json:"status,omitempty"`
}

type GlueClassifierSpecGrokClassifier struct {
	Classification string `json:"classification"`
	// +optional
	CustomPatterns string `json:"custom_patterns,omitempty"`
	GrokPattern    string `json:"grok_pattern"`
}

type GlueClassifierSpecJsonClassifier struct {
	JsonPath string `json:"json_path"`
}

type GlueClassifierSpecXmlClassifier struct {
	Classification string `json:"classification"`
	RowTag         string `json:"row_tag"`
}

type GlueClassifierSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	GrokClassifier *[]GlueClassifierSpec `json:"grok_classifier,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	JsonClassifier *[]GlueClassifierSpec `json:"json_classifier,omitempty"`
	Name           string                `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	XmlClassifier *[]GlueClassifierSpec `json:"xml_classifier,omitempty"`
}

type GlueClassifierStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlueClassifierList is a list of GlueClassifiers
type GlueClassifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlueClassifier CRD objects
	Items []GlueClassifier `json:"items,omitempty"`
}
