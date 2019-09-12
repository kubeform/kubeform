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

type DnsNsRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsNsRecordSpec   `json:"spec,omitempty"`
	Status            DnsNsRecordStatus `json:"status,omitempty"`
}

type DnsNsRecordSpecRecord struct {
	Nsdname string `json:"nsdname" tf:"nsdname"`
}

type DnsNsRecordSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Name string `json:"name" tf:"name"`
	// +optional
	// Deprecated
	Record []DnsNsRecordSpecRecord `json:"record,omitempty" tf:"record,omitempty"`
	// +optional
	Records           []string `json:"records,omitempty" tf:"records,omitempty"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags     map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Ttl      int               `json:"ttl" tf:"ttl"`
	ZoneName string            `json:"zoneName" tf:"zone_name"`
}

type DnsNsRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DnsNsRecordSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsNsRecordList is a list of DnsNsRecords
type DnsNsRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsNsRecord CRD objects
	Items []DnsNsRecord `json:"items,omitempty"`
}
