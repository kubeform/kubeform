package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleSpannerDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSpannerDatabaseSpec   `json:"spec,omitempty"`
	Status            GoogleSpannerDatabaseStatus `json:"status,omitempty"`
}

type GoogleSpannerDatabaseSpec struct {
	Ddl      []string `json:"ddl"`
	State    string   `json:"state"`
	Instance string   `json:"instance"`
	Name     string   `json:"name"`
	Project  string   `json:"project"`
}



type GoogleSpannerDatabaseStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleSpannerDatabaseList is a list of GoogleSpannerDatabases
type GoogleSpannerDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSpannerDatabase CRD objects
	Items []GoogleSpannerDatabase `json:"items,omitempty"`
}