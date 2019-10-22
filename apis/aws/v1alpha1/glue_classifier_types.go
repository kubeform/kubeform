package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type GlueClassifier struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueClassifierSpec   `json:"spec,omitempty"`
	Status            GlueClassifierStatus `json:"status,omitempty"`
}

type GlueClassifierSpecGrokClassifier struct {
	Classification string `json:"classification" tf:"classification"`
	// +optional
	CustomPatterns string `json:"customPatterns,omitempty" tf:"custom_patterns,omitempty"`
	GrokPattern    string `json:"grokPattern" tf:"grok_pattern"`
}

type GlueClassifierSpecJsonClassifier struct {
	JsonPath string `json:"jsonPath" tf:"json_path"`
}

type GlueClassifierSpecXmlClassifier struct {
	Classification string `json:"classification" tf:"classification"`
	RowTag         string `json:"rowTag" tf:"row_tag"`
}

type GlueClassifierSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	GrokClassifier []GlueClassifierSpecGrokClassifier `json:"grokClassifier,omitempty" tf:"grok_classifier,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	JsonClassifier []GlueClassifierSpecJsonClassifier `json:"jsonClassifier,omitempty" tf:"json_classifier,omitempty"`
	Name           string                             `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	XmlClassifier []GlueClassifierSpecXmlClassifier `json:"xmlClassifier,omitempty" tf:"xml_classifier,omitempty"`
}

type GlueClassifierStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GlueClassifierSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
