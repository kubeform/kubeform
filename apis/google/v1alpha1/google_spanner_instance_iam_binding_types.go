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

type GoogleSpannerInstanceIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSpannerInstanceIamBindingSpec   `json:"spec,omitempty"`
	Status            GoogleSpannerInstanceIamBindingStatus `json:"status,omitempty"`
}

type GoogleSpannerInstanceIamBindingSpec struct {
	Members  []string `json:"members"`
	Etag     string   `json:"etag"`
	Role     string   `json:"role"`
	Project  string   `json:"project"`
	Instance string   `json:"instance"`
}

type GoogleSpannerInstanceIamBindingStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleSpannerInstanceIamBindingList is a list of GoogleSpannerInstanceIamBindings
type GoogleSpannerInstanceIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSpannerInstanceIamBinding CRD objects
	Items []GoogleSpannerInstanceIamBinding `json:"items,omitempty"`
}
