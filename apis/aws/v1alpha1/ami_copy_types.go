package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AmiCopy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AmiCopySpec   `json:"spec,omitempty"`
	Status            AmiCopyStatus `json:"status,omitempty"`
}

type AmiCopySpecEbsBlockDevice struct{}

type AmiCopySpecEphemeralBlockDevice struct{}

type AmiCopySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EbsBlockDevice []AmiCopySpecEbsBlockDevice `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`
	// +optional
	Encrypted bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EphemeralBlockDevice []AmiCopySpecEphemeralBlockDevice `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`
	// +optional
	KmsKeyID        string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	Name            string `json:"name" tf:"name"`
	SourceAmiID     string `json:"sourceAmiID" tf:"source_ami_id"`
	SourceAmiRegion string `json:"sourceAmiRegion" tf:"source_ami_region"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AmiCopyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AmiCopySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AmiCopyList is a list of AmiCopys
type AmiCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AmiCopy CRD objects
	Items []AmiCopy `json:"items,omitempty"`
}
