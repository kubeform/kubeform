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

type SpannerInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpannerInstanceSpec   `json:"spec,omitempty"`
	Status            SpannerInstanceStatus `json:"status,omitempty"`
}

type SpannerInstanceSpec struct {
	Config      string `json:"config"`
	DisplayName string `json:"display_name"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	// +optional
	NumNodes int `json:"num_nodes,omitempty"`
}

type SpannerInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpannerInstanceList is a list of SpannerInstances
type SpannerInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpannerInstance CRD objects
	Items []SpannerInstance `json:"items,omitempty"`
}
