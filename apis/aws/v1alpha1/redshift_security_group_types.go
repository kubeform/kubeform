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

type RedshiftSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedshiftSecurityGroupSpec   `json:"spec,omitempty"`
	Status            RedshiftSecurityGroupStatus `json:"status,omitempty"`
}

type RedshiftSecurityGroupSpecIngress struct {
	// +optional
	Cidr string `json:"cidr,omitempty"`
}

type RedshiftSecurityGroupSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Ingress []RedshiftSecurityGroupSpec `json:"ingress"`
	Name    string                      `json:"name"`
}

type RedshiftSecurityGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedshiftSecurityGroupList is a list of RedshiftSecurityGroups
type RedshiftSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RedshiftSecurityGroup CRD objects
	Items []RedshiftSecurityGroup `json:"items,omitempty"`
}
