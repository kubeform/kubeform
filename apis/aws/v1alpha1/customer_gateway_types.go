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

type CustomerGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomerGatewaySpec   `json:"spec,omitempty"`
	Status            CustomerGatewayStatus `json:"status,omitempty"`
}

type CustomerGatewaySpec struct {
	BgpAsn    int    `json:"bgp_asn"`
	IpAddress string `json:"ip_address"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	Type string            `json:"type"`
}

type CustomerGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CustomerGatewayList is a list of CustomerGateways
type CustomerGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CustomerGateway CRD objects
	Items []CustomerGateway `json:"items,omitempty"`
}
