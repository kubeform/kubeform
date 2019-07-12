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

type LinodeNodebalancerSpecTransfer struct {
	Out   float64 `json:"out"`
	Total float64 `json:"total"`
	In    float64 `json:"in"`
}

type LinodeNodebalancerSpec struct {
	Label              string            `json:"label"`
	ClientConnThrottle int               `json:"client_conn_throttle"`
	Hostname           string            `json:"hostname"`
	Updated            string            `json:"updated"`
	Region             string            `json:"region"`
	Ipv4               string            `json:"ipv4"`
	Ipv6               string            `json:"ipv6"`
	Created            string            `json:"created"`
	Transfer           map[string]string `json:"transfer"`
	Tags               []string          `json:"tags"`
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
