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

type AmiCopy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AmiCopySpec   `json:"spec,omitempty"`
	Status            AmiCopyStatus `json:"status,omitempty"`
}

type AmiCopySpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Encrypted       bool   `json:"encrypted,omitempty"`
	Name            string `json:"name"`
	SourceAmiId     string `json:"source_ami_id"`
	SourceAmiRegion string `json:"source_ami_region"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type AmiCopyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
