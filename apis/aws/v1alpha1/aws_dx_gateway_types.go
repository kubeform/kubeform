package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxGatewaySpec   `json:"spec,omitempty"`
	Status            AwsDxGatewayStatus `json:"status,omitempty"`
}

type AwsDxGatewaySpec struct {
	OwnerAccountId string `json:"owner_account_id"`
	Name           string `json:"name"`
	AmazonSideAsn  string `json:"amazon_side_asn"`
}



type AwsDxGatewayStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDxGatewayList is a list of AwsDxGateways
type AwsDxGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxGateway CRD objects
	Items []AwsDxGateway `json:"items,omitempty"`
}