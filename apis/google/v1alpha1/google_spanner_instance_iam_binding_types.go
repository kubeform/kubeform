package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleSpannerInstanceIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSpannerInstanceIamBindingSpec   `json:"spec,omitempty"`
	Status            GoogleSpannerInstanceIamBindingStatus `json:"status,omitempty"`
}

type GoogleSpannerInstanceIamBindingSpec struct {
	Project  string   `json:"project"`
	Instance string   `json:"instance"`
	Members  []string `json:"members"`
	Etag     string   `json:"etag"`
	Role     string   `json:"role"`
}



type GoogleSpannerInstanceIamBindingStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleSpannerInstanceIamBindingList is a list of GoogleSpannerInstanceIamBindings
type GoogleSpannerInstanceIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSpannerInstanceIamBinding CRD objects
	Items []GoogleSpannerInstanceIamBinding `json:"items,omitempty"`
}