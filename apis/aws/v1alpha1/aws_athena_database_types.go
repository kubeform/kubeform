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

type AwsAthenaDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAthenaDatabaseSpec   `json:"spec,omitempty"`
	Status            AwsAthenaDatabaseStatus `json:"status,omitempty"`
}

type AwsAthenaDatabaseSpecEncryptionConfiguration struct {
	KmsKey           string `json:"kms_key"`
	EncryptionOption string `json:"encryption_option"`
}

type AwsAthenaDatabaseSpec struct {
	Name                    string                  `json:"name"`
	Bucket                  string                  `json:"bucket"`
	ForceDestroy            bool                    `json:"force_destroy"`
	EncryptionConfiguration []AwsAthenaDatabaseSpec `json:"encryption_configuration"`
}

type AwsAthenaDatabaseStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAthenaDatabaseList is a list of AwsAthenaDatabases
type AwsAthenaDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAthenaDatabase CRD objects
	Items []AwsAthenaDatabase `json:"items,omitempty"`
}
