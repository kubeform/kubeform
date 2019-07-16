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

type CloudhsmV2Hsm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudhsmV2HsmSpec   `json:"spec,omitempty"`
	Status            CloudhsmV2HsmStatus `json:"status,omitempty"`
}

type CloudhsmV2HsmSpec struct {
	ClusterId string `json:"cluster_id"`
}

type CloudhsmV2HsmStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudhsmV2HsmList is a list of CloudhsmV2Hsms
type CloudhsmV2HsmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudhsmV2Hsm CRD objects
	Items []CloudhsmV2Hsm `json:"items,omitempty"`
}
