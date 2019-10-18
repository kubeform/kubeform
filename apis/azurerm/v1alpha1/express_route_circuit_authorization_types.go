package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ExpressRouteCircuitAuthorization struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitAuthorizationSpec   `json:"spec,omitempty"`
	Status            ExpressRouteCircuitAuthorizationStatus `json:"status,omitempty"`
}

type ExpressRouteCircuitAuthorizationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AuthorizationKey string `json:"-" sensitive:"true" tf:"authorization_key,omitempty"`
	// +optional
	AuthorizationUseStatus  string `json:"authorizationUseStatus,omitempty" tf:"authorization_use_status,omitempty"`
	ExpressRouteCircuitName string `json:"expressRouteCircuitName" tf:"express_route_circuit_name"`
	Name                    string `json:"name" tf:"name"`
	ResourceGroupName       string `json:"resourceGroupName" tf:"resource_group_name"`
}

type ExpressRouteCircuitAuthorizationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ExpressRouteCircuitAuthorizationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
