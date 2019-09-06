package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DnsRecordSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsRecordSetSpec   `json:"spec,omitempty"`
	Status            DnsRecordSetStatus `json:"status,omitempty"`
}

type DnsRecordSetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ManagedZone string `json:"managedZone" tf:"managed_zone"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Project string   `json:"project,omitempty" tf:"project,omitempty"`
	Rrdatas []string `json:"rrdatas" tf:"rrdatas"`
	Ttl     int      `json:"ttl" tf:"ttl"`
	Type    string   `json:"type" tf:"type"`
}

type DnsRecordSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DnsRecordSetSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsRecordSetList is a list of DnsRecordSets
type DnsRecordSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsRecordSet CRD objects
	Items []DnsRecordSet `json:"items,omitempty"`
}
