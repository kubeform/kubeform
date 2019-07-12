package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmParameter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmParameterSpec   `json:"spec,omitempty"`
	Status            AwsSsmParameterStatus `json:"status,omitempty"`
}

type AwsSsmParameterSpec struct {
	Description    string            `json:"description"`
	Tier           string            `json:"tier"`
	Value          string            `json:"value"`
	Arn            string            `json:"arn"`
	KeyId          string            `json:"key_id"`
	Tags           map[string]string `json:"tags"`
	Name           string            `json:"name"`
	Type           string            `json:"type"`
	Overwrite      bool              `json:"overwrite"`
	AllowedPattern string            `json:"allowed_pattern"`
}

type AwsSsmParameterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmParameterList is a list of AwsSsmParameters
type AwsSsmParameterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmParameter CRD objects
	Items []AwsSsmParameter `json:"items,omitempty"`
}
