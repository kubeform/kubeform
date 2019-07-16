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

type GuarddutyThreatintelset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GuarddutyThreatintelsetSpec   `json:"spec,omitempty"`
	Status            GuarddutyThreatintelsetStatus `json:"status,omitempty"`
}

type GuarddutyThreatintelsetSpec struct {
	Activate   bool   `json:"activate"`
	DetectorId string `json:"detector_id"`
	Format     string `json:"format"`
	Location   string `json:"location"`
	Name       string `json:"name"`
}

type GuarddutyThreatintelsetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GuarddutyThreatintelsetList is a list of GuarddutyThreatintelsets
type GuarddutyThreatintelsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GuarddutyThreatintelset CRD objects
	Items []GuarddutyThreatintelset `json:"items,omitempty"`
}
