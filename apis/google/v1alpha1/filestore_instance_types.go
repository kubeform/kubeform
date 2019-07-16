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

type FilestoreInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FilestoreInstanceSpec   `json:"spec,omitempty"`
	Status            FilestoreInstanceStatus `json:"status,omitempty"`
}

type FilestoreInstanceSpecFileShares struct {
	CapacityGb int    `json:"capacity_gb"`
	Name       string `json:"name"`
}

type FilestoreInstanceSpecNetworks struct {
	Modes   []string `json:"modes"`
	Network string   `json:"network"`
}

type FilestoreInstanceSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	FileShares []FilestoreInstanceSpec `json:"file_shares"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	Name   string            `json:"name"`
	// +kubebuilder:validation:MinItems=1
	Networks []FilestoreInstanceSpec `json:"networks"`
	Tier     string                  `json:"tier"`
	Zone     string                  `json:"zone"`
}

type FilestoreInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FilestoreInstanceList is a list of FilestoreInstances
type FilestoreInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FilestoreInstance CRD objects
	Items []FilestoreInstance `json:"items,omitempty"`
}
