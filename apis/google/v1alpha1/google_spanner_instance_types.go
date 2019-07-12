package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleSpannerInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSpannerInstanceSpec   `json:"spec,omitempty"`
	Status            GoogleSpannerInstanceStatus `json:"status,omitempty"`
}

type GoogleSpannerInstanceSpec struct {
	Config      string            `json:"config"`
	Name        string            `json:"name"`
	DisplayName string            `json:"display_name"`
	NumNodes    int               `json:"num_nodes"`
	Labels      map[string]string `json:"labels"`
	Project     string            `json:"project"`
	State       string            `json:"state"`
}

type GoogleSpannerInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleSpannerInstanceList is a list of GoogleSpannerInstances
type GoogleSpannerInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSpannerInstance CRD objects
	Items []GoogleSpannerInstance `json:"items,omitempty"`
}
