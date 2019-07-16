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

type ExpressRouteCircuitAuthorization struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitAuthorizationSpec   `json:"spec,omitempty"`
	Status            ExpressRouteCircuitAuthorizationStatus `json:"status,omitempty"`
}

type ExpressRouteCircuitAuthorizationSpec struct {
	ExpressRouteCircuitName string `json:"express_route_circuit_name"`
	Name                    string `json:"name"`
	ResourceGroupName       string `json:"resource_group_name"`
}

type ExpressRouteCircuitAuthorizationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ExpressRouteCircuitAuthorizationList is a list of ExpressRouteCircuitAuthorizations
type ExpressRouteCircuitAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ExpressRouteCircuitAuthorization CRD objects
	Items []ExpressRouteCircuitAuthorization `json:"items,omitempty"`
}
