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

type RedshiftSubnetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedshiftSubnetGroupSpec   `json:"spec,omitempty"`
	Status            RedshiftSubnetGroupStatus `json:"status,omitempty"`
}

type RedshiftSubnetGroupSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +kubebuilder:validation:UniqueItems=true
	SubnetIds []string `json:"subnet_ids"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type RedshiftSubnetGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedshiftSubnetGroupList is a list of RedshiftSubnetGroups
type RedshiftSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RedshiftSubnetGroup CRD objects
	Items []RedshiftSubnetGroup `json:"items,omitempty"`
}
