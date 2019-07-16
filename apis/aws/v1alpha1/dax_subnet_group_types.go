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

type DaxSubnetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DaxSubnetGroupSpec   `json:"spec,omitempty"`
	Status            DaxSubnetGroupStatus `json:"status,omitempty"`
}

type DaxSubnetGroupSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +kubebuilder:validation:UniqueItems=true
	SubnetIds []string `json:"subnet_ids"`
}

type DaxSubnetGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DaxSubnetGroupList is a list of DaxSubnetGroups
type DaxSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DaxSubnetGroup CRD objects
	Items []DaxSubnetGroup `json:"items,omitempty"`
}
