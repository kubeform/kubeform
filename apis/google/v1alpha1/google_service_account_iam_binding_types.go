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

type GoogleServiceAccountIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleServiceAccountIamBindingSpec   `json:"spec,omitempty"`
	Status            GoogleServiceAccountIamBindingStatus `json:"status,omitempty"`
}

type GoogleServiceAccountIamBindingSpec struct {
	Members          []string `json:"members"`
	Etag             string   `json:"etag"`
	Role             string   `json:"role"`
	ServiceAccountId string   `json:"service_account_id"`
}



type GoogleServiceAccountIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleServiceAccountIamBindingList is a list of GoogleServiceAccountIamBindings
type GoogleServiceAccountIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleServiceAccountIamBinding CRD objects
	Items []GoogleServiceAccountIamBinding `json:"items,omitempty"`
}