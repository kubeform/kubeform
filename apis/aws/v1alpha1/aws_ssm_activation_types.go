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

type AwsSsmActivation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmActivationSpec   `json:"spec,omitempty"`
	Status            AwsSsmActivationStatus `json:"status,omitempty"`
}

type AwsSsmActivationSpec struct {
	Expired           string            `json:"expired"`
	RegistrationLimit int               `json:"registration_limit"`
	RegistrationCount int               `json:"registration_count"`
	Description       string            `json:"description"`
	ExpirationDate    string            `json:"expiration_date"`
	IamRole           string            `json:"iam_role"`
	ActivationCode    string            `json:"activation_code"`
	Tags              map[string]string `json:"tags"`
	Name              string            `json:"name"`
}



type AwsSsmActivationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSsmActivationList is a list of AwsSsmActivations
type AwsSsmActivationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmActivation CRD objects
	Items []AwsSsmActivation `json:"items,omitempty"`
}