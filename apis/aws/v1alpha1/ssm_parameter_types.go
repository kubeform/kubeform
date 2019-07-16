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

type SsmParameter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmParameterSpec   `json:"spec,omitempty"`
	Status            SsmParameterStatus `json:"status,omitempty"`
}

type SsmParameterSpec struct {
	// +optional
	AllowedPattern string `json:"allowed_pattern,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +optional
	Overwrite bool `json:"overwrite,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	Tier  string `json:"tier,omitempty"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type SsmParameterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmParameterList is a list of SsmParameters
type SsmParameterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmParameter CRD objects
	Items []SsmParameter `json:"items,omitempty"`
}
