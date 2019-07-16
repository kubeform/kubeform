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

type AmiFromInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AmiFromInstanceSpec   `json:"spec,omitempty"`
	Status            AmiFromInstanceStatus `json:"status,omitempty"`
}

type AmiFromInstanceSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +optional
	SnapshotWithoutReboot bool   `json:"snapshot_without_reboot,omitempty"`
	SourceInstanceId      string `json:"source_instance_id"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type AmiFromInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AmiFromInstanceList is a list of AmiFromInstances
type AmiFromInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AmiFromInstance CRD objects
	Items []AmiFromInstance `json:"items,omitempty"`
}
