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

type GoogleDnsRecordSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleDnsRecordSetSpec   `json:"spec,omitempty"`
	Status            GoogleDnsRecordSetStatus `json:"status,omitempty"`
}

type GoogleDnsRecordSetSpec struct {
	Ttl         int      `json:"ttl"`
	Type        string   `json:"type"`
	Project     string   `json:"project"`
	ManagedZone string   `json:"managed_zone"`
	Name        string   `json:"name"`
	Rrdatas     []string `json:"rrdatas"`
}



type GoogleDnsRecordSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleDnsRecordSetList is a list of GoogleDnsRecordSets
type GoogleDnsRecordSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleDnsRecordSet CRD objects
	Items []GoogleDnsRecordSet `json:"items,omitempty"`
}