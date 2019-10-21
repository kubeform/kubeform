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
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type DnsMxRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsMxRecordSpec   `json:"spec,omitempty"`
	Status            DnsMxRecordStatus `json:"status,omitempty"`
}

type DnsMxRecordSpecRecord struct {
	Exchange   string `json:"exchange" tf:"exchange"`
	Preference string `json:"preference" tf:"preference"`
}

type DnsMxRecordSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Name              string                  `json:"name" tf:"name"`
	Record            []DnsMxRecordSpecRecord `json:"record" tf:"record"`
	ResourceGroupName string                  `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags     map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Ttl      int               `json:"ttl" tf:"ttl"`
	ZoneName string            `json:"zoneName" tf:"zone_name"`
}

type DnsMxRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DnsMxRecordSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsMxRecordList is a list of DnsMxRecords
type DnsMxRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsMxRecord CRD objects
	Items []DnsMxRecord `json:"items,omitempty"`
}