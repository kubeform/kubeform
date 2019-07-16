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

type AthenaDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AthenaDatabaseSpec   `json:"spec,omitempty"`
	Status            AthenaDatabaseStatus `json:"status,omitempty"`
}

type AthenaDatabaseSpecEncryptionConfiguration struct {
	EncryptionOption string `json:"encryption_option"`
	// +optional
	KmsKey string `json:"kms_key,omitempty"`
}

type AthenaDatabaseSpec struct {
	Bucket string `json:"bucket"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionConfiguration *[]AthenaDatabaseSpec `json:"encryption_configuration,omitempty"`
	// +optional
	ForceDestroy bool   `json:"force_destroy,omitempty"`
	Name         string `json:"name"`
}

type AthenaDatabaseStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AthenaDatabaseList is a list of AthenaDatabases
type AthenaDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AthenaDatabase CRD objects
	Items []AthenaDatabase `json:"items,omitempty"`
}
