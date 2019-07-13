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

type GoogleComputeSubnetworkIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeSubnetworkIamBindingSpec   `json:"spec,omitempty"`
	Status            GoogleComputeSubnetworkIamBindingStatus `json:"status,omitempty"`
}

type GoogleComputeSubnetworkIamBindingSpec struct {
	Project    string   `json:"project"`
	Region     string   `json:"region"`
	Etag       string   `json:"etag"`
	Role       string   `json:"role"`
	Members    []string `json:"members"`
	Subnetwork string   `json:"subnetwork"`
}



type GoogleComputeSubnetworkIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeSubnetworkIamBindingList is a list of GoogleComputeSubnetworkIamBindings
type GoogleComputeSubnetworkIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeSubnetworkIamBinding CRD objects
	Items []GoogleComputeSubnetworkIamBinding `json:"items,omitempty"`
}