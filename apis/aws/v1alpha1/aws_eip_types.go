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

type AwsEip struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEipSpec   `json:"spec,omitempty"`
	Status            AwsEipStatus `json:"status,omitempty"`
}

type AwsEipSpec struct {
	AllocationId           string            `json:"allocation_id"`
	AssociationId          string            `json:"association_id"`
	Domain                 string            `json:"domain"`
	PublicIp               string            `json:"public_ip"`
	PublicDns              string            `json:"public_dns"`
	Tags                   map[string]string `json:"tags"`
	Vpc                    bool              `json:"vpc"`
	NetworkInterface       string            `json:"network_interface"`
	PrivateIp              string            `json:"private_ip"`
	PrivateDns             string            `json:"private_dns"`
	AssociateWithPrivateIp string            `json:"associate_with_private_ip"`
	PublicIpv4Pool         string            `json:"public_ipv4_pool"`
	Instance               string            `json:"instance"`
}



type AwsEipStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEipList is a list of AwsEips
type AwsEipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEip CRD objects
	Items []AwsEip `json:"items,omitempty"`
}