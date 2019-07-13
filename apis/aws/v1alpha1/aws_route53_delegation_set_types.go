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

type AwsRoute53DelegationSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRoute53DelegationSetSpec   `json:"spec,omitempty"`
	Status            AwsRoute53DelegationSetStatus `json:"status,omitempty"`
}

type AwsRoute53DelegationSetSpec struct {
	ReferenceName string   `json:"reference_name"`
	NameServers   []string `json:"name_servers"`
}



type AwsRoute53DelegationSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRoute53DelegationSetList is a list of AwsRoute53DelegationSets
type AwsRoute53DelegationSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRoute53DelegationSet CRD objects
	Items []AwsRoute53DelegationSet `json:"items,omitempty"`
}