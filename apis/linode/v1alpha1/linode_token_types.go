package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LinodeToken struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeTokenSpec   `json:"spec,omitempty"`
	Status            LinodeTokenStatus `json:"status,omitempty"`
}

type LinodeTokenSpec struct {
	Expiry  string `json:"expiry"`
	Created string `json:"created"`
	Token   string `json:"token"`
	Label   string `json:"label"`
	Scopes  string `json:"scopes"`
}



type LinodeTokenStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinodeTokenList is a list of LinodeTokens
type LinodeTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeToken CRD objects
	Items []LinodeToken `json:"items,omitempty"`
}