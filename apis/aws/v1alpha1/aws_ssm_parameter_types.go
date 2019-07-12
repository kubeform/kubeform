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

type AwsSsmParameter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmParameterSpec   `json:"spec,omitempty"`
	Status            AwsSsmParameterStatus `json:"status,omitempty"`
}

type AwsSsmParameterSpec struct {
	Arn            string            `json:"arn"`
	Overwrite      bool              `json:"overwrite"`
	AllowedPattern string            `json:"allowed_pattern"`
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	Type           string            `json:"type"`
	Version        int               `json:"version"`
	Tags           map[string]string `json:"tags"`
	Tier           string            `json:"tier"`
	Value          string            `json:"value"`
	KeyId          string            `json:"key_id"`
}

type AwsSsmParameterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSsmParameterList is a list of AwsSsmParameters
type AwsSsmParameterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmParameter CRD objects
	Items []AwsSsmParameter `json:"items,omitempty"`
}
