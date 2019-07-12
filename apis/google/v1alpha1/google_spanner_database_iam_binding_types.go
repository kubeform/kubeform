package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleSpannerDatabaseIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSpannerDatabaseIamBindingSpec   `json:"spec,omitempty"`
	Status            GoogleSpannerDatabaseIamBindingStatus `json:"status,omitempty"`
}

type GoogleSpannerDatabaseIamBindingSpec struct {
	Project  string   `json:"project"`
	Role     string   `json:"role"`
	Members  []string `json:"members"`
	Etag     string   `json:"etag"`
	Instance string   `json:"instance"`
	Database string   `json:"database"`
}



type GoogleSpannerDatabaseIamBindingStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleSpannerDatabaseIamBindingList is a list of GoogleSpannerDatabaseIamBindings
type GoogleSpannerDatabaseIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSpannerDatabaseIamBinding CRD objects
	Items []GoogleSpannerDatabaseIamBinding `json:"items,omitempty"`
}