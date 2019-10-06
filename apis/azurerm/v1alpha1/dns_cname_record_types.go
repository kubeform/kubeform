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

type DnsCnameRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsCnameRecordSpec   `json:"spec,omitempty"`
	Status            DnsCnameRecordStatus `json:"status,omitempty"`
}

type DnsCnameRecordSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Name              string `json:"name" tf:"name"`
	Record            string `json:"record" tf:"record"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags     map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Ttl      int               `json:"ttl" tf:"ttl"`
	ZoneName string            `json:"zoneName" tf:"zone_name"`
}

type DnsCnameRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DnsCnameRecordSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsCnameRecordList is a list of DnsCnameRecords
type DnsCnameRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsCnameRecord CRD objects
	Items []DnsCnameRecord `json:"items,omitempty"`
}
