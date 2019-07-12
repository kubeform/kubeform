package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LinodeRdns struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeRdnsSpec   `json:"spec,omitempty"`
	Status            LinodeRdnsStatus `json:"status,omitempty"`
}

type LinodeRdnsSpec struct {
	Rdns    string `json:"rdns"`
	Address string `json:"address"`
}



type LinodeRdnsStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinodeRdnsList is a list of LinodeRdnss
type LinodeRdnsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeRdns CRD objects
	Items []LinodeRdns `json:"items,omitempty"`
}