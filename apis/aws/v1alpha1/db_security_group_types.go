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

type DbSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbSecurityGroupSpec   `json:"spec,omitempty"`
	Status            DbSecurityGroupStatus `json:"status,omitempty"`
}

type DbSecurityGroupSpecIngress struct {
	// +optional
	Cidr string `json:"cidr,omitempty"`
}

type DbSecurityGroupSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Ingress []DbSecurityGroupSpec `json:"ingress"`
	Name    string                `json:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DbSecurityGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DbSecurityGroupList is a list of DbSecurityGroups
type DbSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DbSecurityGroup CRD objects
	Items []DbSecurityGroup `json:"items,omitempty"`
}
