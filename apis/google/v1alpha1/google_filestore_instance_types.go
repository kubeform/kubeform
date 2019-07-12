package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleFilestoreInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleFilestoreInstanceSpec   `json:"spec,omitempty"`
	Status            GoogleFilestoreInstanceStatus `json:"status,omitempty"`
}

type GoogleFilestoreInstanceSpecNetworks struct {
	Modes           []string `json:"modes"`
	Network         string   `json:"network"`
	ReservedIpRange string   `json:"reserved_ip_range"`
	IpAddresses     []string `json:"ip_addresses"`
}

type GoogleFilestoreInstanceSpecFileShares struct {
	CapacityGb int    `json:"capacity_gb"`
	Name       string `json:"name"`
}

type GoogleFilestoreInstanceSpec struct {
	Tier        string                        `json:"tier"`
	Networks    []GoogleFilestoreInstanceSpec `json:"networks"`
	Zone        string                        `json:"zone"`
	Description string                        `json:"description"`
	Labels      map[string]string             `json:"labels"`
	CreateTime  string                        `json:"create_time"`
	Etag        string                        `json:"etag"`
	FileShares  []GoogleFilestoreInstanceSpec `json:"file_shares"`
	Name        string                        `json:"name"`
	Project     string                        `json:"project"`
}

type GoogleFilestoreInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleFilestoreInstanceList is a list of GoogleFilestoreInstances
type GoogleFilestoreInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleFilestoreInstance CRD objects
	Items []GoogleFilestoreInstance `json:"items,omitempty"`
}
