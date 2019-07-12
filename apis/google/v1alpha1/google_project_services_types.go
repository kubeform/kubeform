package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleProjectServices struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleProjectServicesSpec   `json:"spec,omitempty"`
	Status            GoogleProjectServicesStatus `json:"status,omitempty"`
}

type GoogleProjectServicesSpec struct {
	Project          string   `json:"project"`
	Services         []string `json:"services"`
	DisableOnDestroy bool     `json:"disable_on_destroy"`
}

type GoogleProjectServicesStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleProjectServicesList is a list of GoogleProjectServicess
type GoogleProjectServicesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleProjectServices CRD objects
	Items []GoogleProjectServices `json:"items,omitempty"`
}
