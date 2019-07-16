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

type GuarddutyIpset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GuarddutyIpsetSpec   `json:"spec,omitempty"`
	Status            GuarddutyIpsetStatus `json:"status,omitempty"`
}

type GuarddutyIpsetSpec struct {
	Activate   bool   `json:"activate"`
	DetectorId string `json:"detector_id"`
	Format     string `json:"format"`
	Location   string `json:"location"`
	Name       string `json:"name"`
}

type GuarddutyIpsetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GuarddutyIpsetList is a list of GuarddutyIpsets
type GuarddutyIpsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GuarddutyIpset CRD objects
	Items []GuarddutyIpset `json:"items,omitempty"`
}
