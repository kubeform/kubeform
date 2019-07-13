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

type AzurermExpressRouteCircuitAuthorization struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermExpressRouteCircuitAuthorizationSpec   `json:"spec,omitempty"`
	Status            AzurermExpressRouteCircuitAuthorizationStatus `json:"status,omitempty"`
}

type AzurermExpressRouteCircuitAuthorizationSpec struct {
	ResourceGroupName       string `json:"resource_group_name"`
	ExpressRouteCircuitName string `json:"express_route_circuit_name"`
	AuthorizationKey        string `json:"authorization_key"`
	AuthorizationUseStatus  string `json:"authorization_use_status"`
	Name                    string `json:"name"`
}



type AzurermExpressRouteCircuitAuthorizationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermExpressRouteCircuitAuthorizationList is a list of AzurermExpressRouteCircuitAuthorizations
type AzurermExpressRouteCircuitAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermExpressRouteCircuitAuthorization CRD objects
	Items []AzurermExpressRouteCircuitAuthorization `json:"items,omitempty"`
}