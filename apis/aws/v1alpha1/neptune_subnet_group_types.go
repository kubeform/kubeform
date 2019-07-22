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

type NeptuneSubnetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneSubnetGroupSpec   `json:"spec,omitempty"`
	Status            NeptuneSubnetGroupStatus `json:"status,omitempty"`
}

type NeptuneSubnetGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS" tf:"subnet_ids"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NeptuneSubnetGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NeptuneSubnetGroupList is a list of NeptuneSubnetGroups
type NeptuneSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NeptuneSubnetGroup CRD objects
	Items []NeptuneSubnetGroup `json:"items,omitempty"`
}
