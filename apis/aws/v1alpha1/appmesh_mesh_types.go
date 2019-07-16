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

type AppmeshMesh struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshMeshSpec   `json:"spec,omitempty"`
	Status            AppmeshMeshStatus `json:"status,omitempty"`
}

type AppmeshMeshSpecSpecEgressFilter struct {
	// +optional
	Type string `json:"type,omitempty"`
}

type AppmeshMeshSpecSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EgressFilter *[]AppmeshMeshSpecSpec `json:"egress_filter,omitempty"`
}

type AppmeshMeshSpec struct {
	Name string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Spec *[]AppmeshMeshSpec `json:"spec,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type AppmeshMeshStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppmeshMeshList is a list of AppmeshMeshs
type AppmeshMeshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppmeshMesh CRD objects
	Items []AppmeshMesh `json:"items,omitempty"`
}
