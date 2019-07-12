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

type AwsIotThing struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIotThingSpec   `json:"spec,omitempty"`
	Status            AwsIotThingStatus `json:"status,omitempty"`
}

type AwsIotThingSpec struct {
	Attributes      map[string]string `json:"attributes"`
	ThingTypeName   string            `json:"thing_type_name"`
	DefaultClientId string            `json:"default_client_id"`
	Version         int               `json:"version"`
	Arn             string            `json:"arn"`
	Name            string            `json:"name"`
}

type AwsIotThingStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIotThingList is a list of AwsIotThings
type AwsIotThingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIotThing CRD objects
	Items []AwsIotThing `json:"items,omitempty"`
}
