package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIotThingType struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIotThingTypeSpec   `json:"spec,omitempty"`
	Status            AwsIotThingTypeStatus `json:"status,omitempty"`
}

type AwsIotThingTypeSpecProperties struct {
	Description          string   `json:"description"`
	SearchableAttributes []string `json:"searchable_attributes"`
}

type AwsIotThingTypeSpec struct {
	Name       string                `json:"name"`
	Properties []AwsIotThingTypeSpec `json:"properties"`
	Deprecated bool                  `json:"deprecated"`
	Arn        string                `json:"arn"`
}

type AwsIotThingTypeStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIotThingTypeList is a list of AwsIotThingTypes
type AwsIotThingTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIotThingType CRD objects
	Items []AwsIotThingType `json:"items,omitempty"`
}
