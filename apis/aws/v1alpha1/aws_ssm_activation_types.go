package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmActivation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmActivationSpec   `json:"spec,omitempty"`
	Status            AwsSsmActivationStatus `json:"status,omitempty"`
}

type AwsSsmActivationSpec struct {
	Name              string            `json:"name"`
	ExpirationDate    string            `json:"expiration_date"`
	IamRole           string            `json:"iam_role"`
	Tags              map[string]string `json:"tags"`
	Description       string            `json:"description"`
	Expired           string            `json:"expired"`
	RegistrationLimit int               `json:"registration_limit"`
	RegistrationCount int               `json:"registration_count"`
	ActivationCode    string            `json:"activation_code"`
}



type AwsSsmActivationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmActivationList is a list of AwsSsmActivations
type AwsSsmActivationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmActivation CRD objects
	Items []AwsSsmActivation `json:"items,omitempty"`
}