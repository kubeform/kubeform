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

type SsmActivation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmActivationSpec   `json:"spec,omitempty"`
	Status            SsmActivationStatus `json:"status,omitempty"`
}

type SsmActivationSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	ExpirationDate string `json:"expiration_date,omitempty"`
	IamRole        string `json:"iam_role"`
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	RegistrationLimit int `json:"registration_limit,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type SsmActivationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmActivationList is a list of SsmActivations
type SsmActivationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmActivation CRD objects
	Items []SsmActivation `json:"items,omitempty"`
}
