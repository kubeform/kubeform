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

type LinodeDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeDomainSpec   `json:"spec,omitempty"`
	Status            LinodeDomainStatus `json:"status,omitempty"`
}

type LinodeDomainSpec struct {
	Tags        []string `json:"tags"`
	Type        string   `json:"type"`
	Status      string   `json:"status"`
	AxfrIps     []string `json:"axfr_ips"`
	ExpireSec   int      `json:"expire_sec"`
	RefreshSec  int      `json:"refresh_sec"`
	SoaEmail    string   `json:"soa_email"`
	Domain      string   `json:"domain"`
	Group       string   `json:"group"`
	Description string   `json:"description"`
	MasterIps   []string `json:"master_ips"`
	TtlSec      int      `json:"ttl_sec"`
	RetrySec    int      `json:"retry_sec"`
}



type LinodeDomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LinodeDomainList is a list of LinodeDomains
type LinodeDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeDomain CRD objects
	Items []LinodeDomain `json:"items,omitempty"`
}