package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleProjectService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleProjectServiceSpec   `json:"spec,omitempty"`
	Status            GoogleProjectServiceStatus `json:"status,omitempty"`
}

type GoogleProjectServiceSpec struct {
	Project          string `json:"project"`
	DisableOnDestroy bool   `json:"disable_on_destroy"`
	Service          string `json:"service"`
}



type GoogleProjectServiceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleProjectServiceList is a list of GoogleProjectServices
type GoogleProjectServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleProjectService CRD objects
	Items []GoogleProjectService `json:"items,omitempty"`
}