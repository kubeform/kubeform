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

type ComputeRouter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeRouterSpec   `json:"spec,omitempty"`
	Status            ComputeRouterStatus `json:"status,omitempty"`
}

type ComputeRouterSpecBgpAdvertisedIpRanges struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Range string `json:"range,omitempty"`
}

type ComputeRouterSpecBgp struct {
	// +optional
	AdvertiseMode string `json:"advertise_mode,omitempty"`
	// +optional
	AdvertisedGroups []string `json:"advertised_groups,omitempty"`
	// +optional
	AdvertisedIpRanges *[]ComputeRouterSpecBgp `json:"advertised_ip_ranges,omitempty"`
	Asn                int                     `json:"asn"`
}

type ComputeRouterSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Bgp *[]ComputeRouterSpec `json:"bgp,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Network     string `json:"network"`
}

type ComputeRouterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeRouterList is a list of ComputeRouters
type ComputeRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeRouter CRD objects
	Items []ComputeRouter `json:"items,omitempty"`
}
