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

type EmrSecurityConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmrSecurityConfigurationSpec   `json:"spec,omitempty"`
	Status            EmrSecurityConfigurationStatus `json:"status,omitempty"`
}

type EmrSecurityConfigurationSpec struct {
	Configuration string `json:"configuration"`
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
}

type EmrSecurityConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EmrSecurityConfigurationList is a list of EmrSecurityConfigurations
type EmrSecurityConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EmrSecurityConfiguration CRD objects
	Items []EmrSecurityConfiguration `json:"items,omitempty"`
}
