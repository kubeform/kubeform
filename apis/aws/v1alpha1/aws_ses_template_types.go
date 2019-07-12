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

type AwsSesTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesTemplateSpec   `json:"spec,omitempty"`
	Status            AwsSesTemplateStatus `json:"status,omitempty"`
}

type AwsSesTemplateSpec struct {
	Subject string `json:"subject"`
	Text    string `json:"text"`
	Name    string `json:"name"`
	Html    string `json:"html"`
}

type AwsSesTemplateStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSesTemplateList is a list of AwsSesTemplates
type AwsSesTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesTemplate CRD objects
	Items []AwsSesTemplate `json:"items,omitempty"`
}
