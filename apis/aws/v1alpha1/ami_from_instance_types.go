package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AmiFromInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AmiFromInstanceSpec   `json:"spec,omitempty"`
	Status            AmiFromInstanceStatus `json:"status,omitempty"`
}

type AmiFromInstanceSpecEbsBlockDevice struct{}

type AmiFromInstanceSpecEphemeralBlockDevice struct{}

type AmiFromInstanceSpec struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EbsBlockDevice []AmiFromInstanceSpecEbsBlockDevice `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EphemeralBlockDevice []AmiFromInstanceSpecEphemeralBlockDevice `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`
	Name                 string                                    `json:"name" tf:"name"`
	// +optional
	SnapshotWithoutReboot bool   `json:"snapshotWithoutReboot,omitempty" tf:"snapshot_without_reboot,omitempty"`
	SourceInstanceID      string `json:"sourceInstanceID" tf:"source_instance_id"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AmiFromInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
