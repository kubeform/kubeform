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

type EmrSecurityConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmrSecurityConfigurationSpec   `json:"spec,omitempty"`
	Status            EmrSecurityConfigurationStatus `json:"status,omitempty"`
}

type EmrSecurityConfigurationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Configuration string `json:"configuration" tf:"configuration"`
	// +optional
	CreationDate string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
}

type EmrSecurityConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EmrSecurityConfigurationSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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