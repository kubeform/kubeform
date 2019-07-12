package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeRouter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeRouterSpec   `json:"spec,omitempty"`
	Status            GoogleComputeRouterStatus `json:"status,omitempty"`
}

type GoogleComputeRouterSpecBgpAdvertisedIpRanges struct {
	Description string `json:"description"`
	Range       string `json:"range"`
}

type GoogleComputeRouterSpecBgp struct {
	Asn                int                          `json:"asn"`
	AdvertiseMode      string                       `json:"advertise_mode"`
	AdvertisedGroups   []string                     `json:"advertised_groups"`
	AdvertisedIpRanges []GoogleComputeRouterSpecBgp `json:"advertised_ip_ranges"`
}

type GoogleComputeRouterSpec struct {
	Description       string                    `json:"description"`
	Region            string                    `json:"region"`
	CreationTimestamp string                    `json:"creation_timestamp"`
	Project           string                    `json:"project"`
	SelfLink          string                    `json:"self_link"`
	Name              string                    `json:"name"`
	Network           string                    `json:"network"`
	Bgp               []GoogleComputeRouterSpec `json:"bgp"`
}



type GoogleComputeRouterStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeRouterList is a list of GoogleComputeRouters
type GoogleComputeRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeRouter CRD objects
	Items []GoogleComputeRouter `json:"items,omitempty"`
}