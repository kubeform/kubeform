package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DnsCaaRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsCaaRecordSpec   `json:"spec,omitempty"`
	Status            DnsCaaRecordStatus `json:"status,omitempty"`
}

type DnsCaaRecordSpecRecord struct {
	Flags int    `json:"flags" tf:"flags"`
	Tag   string `json:"tag" tf:"tag"`
	Value string `json:"value" tf:"value"`
}

type DnsCaaRecordSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	Name        string                    `json:"name" tf:"name"`
	// +kubebuilder:validation:UniqueItems=true
	Record            []DnsCaaRecordSpecRecord `json:"record" tf:"record"`
	ResourceGroupName string                   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags     map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Ttl      int               `json:"ttl" tf:"ttl"`
	ZoneName string            `json:"zoneName" tf:"zone_name"`
}

type DnsCaaRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsCaaRecordList is a list of DnsCaaRecords
type DnsCaaRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsCaaRecord CRD objects
	Items []DnsCaaRecord `json:"items,omitempty"`
}
