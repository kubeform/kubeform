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

type AwsCustomerGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCustomerGatewaySpec   `json:"spec,omitempty"`
	Status            AwsCustomerGatewayStatus `json:"status,omitempty"`
}

type AwsCustomerGatewaySpec struct {
	IpAddress string            `json:"ip_address"`
	Type      string            `json:"type"`
	Tags      map[string]string `json:"tags"`
	BgpAsn    int               `json:"bgp_asn"`
}



type AwsCustomerGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCustomerGatewayList is a list of AwsCustomerGateways
type AwsCustomerGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCustomerGateway CRD objects
	Items []AwsCustomerGateway `json:"items,omitempty"`
}