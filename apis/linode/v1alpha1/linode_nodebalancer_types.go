package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LinodeNodebalancer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeNodebalancerSpec   `json:"spec,omitempty"`
	Status            LinodeNodebalancerStatus `json:"status,omitempty"`
}

type TransferSpec struct {
	Total float64 `json:"total"`
	In    float64 `json:"in"`
	Out   float64 `json:"out"`
}

type LinodeNodebalancerSpec struct {
	Hostname           string            `json:"hostname"`
	Ipv6               string            `json:"ipv6"`
	Created            string            `json:"created"`
	Updated            string            `json:"updated"`
	Tags               []string          `json:"tags"`
	Label              string            `json:"label"`
	Region             string            `json:"region"`
	ClientConnThrottle int               `json:"client_conn_throttle"`
	Ipv4               string            `json:"ipv4"`
	Transfer           map[string]string `json:"transfer"`
}



type LinodeNodebalancerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinodeNodebalancerList is a list of LinodeNodebalancers
type LinodeNodebalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeNodebalancer CRD objects
	Items []LinodeNodebalancer `json:"items,omitempty"`
}