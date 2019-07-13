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

type GoogleSpannerDatabaseIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSpannerDatabaseIamBindingSpec   `json:"spec,omitempty"`
	Status            GoogleSpannerDatabaseIamBindingStatus `json:"status,omitempty"`
}

type GoogleSpannerDatabaseIamBindingSpec struct {
	Members  []string `json:"members"`
	Etag     string   `json:"etag"`
	Instance string   `json:"instance"`
	Database string   `json:"database"`
	Project  string   `json:"project"`
	Role     string   `json:"role"`
}



type GoogleSpannerDatabaseIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleSpannerDatabaseIamBindingList is a list of GoogleSpannerDatabaseIamBindings
type GoogleSpannerDatabaseIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSpannerDatabaseIamBinding CRD objects
	Items []GoogleSpannerDatabaseIamBinding `json:"items,omitempty"`
}