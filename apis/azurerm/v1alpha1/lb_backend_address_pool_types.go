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

type LbBackendAddressPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbBackendAddressPoolSpec   `json:"spec,omitempty"`
	Status            LbBackendAddressPoolStatus `json:"status,omitempty"`
}

type LbBackendAddressPoolSpec struct {
	LoadbalancerId string `json:"loadbalancer_id"`
	// +optional
	// Deprecated
	Location          string `json:"location,omitempty"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
}

type LbBackendAddressPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbBackendAddressPoolList is a list of LbBackendAddressPools
type LbBackendAddressPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbBackendAddressPool CRD objects
	Items []LbBackendAddressPool `json:"items,omitempty"`
}
