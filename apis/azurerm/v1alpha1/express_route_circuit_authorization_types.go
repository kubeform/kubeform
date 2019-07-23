package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ExpressRouteCircuitName string `json:"expressRouteCircuitName" tf:"express_route_circuit_name"`
	Name                    string `json:"name" tf:"name"`
	ResourceGroupName       string `json:"resourceGroupName" tf:"resource_group_name"`
}

type ExpressRouteCircuitAuthorizationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
