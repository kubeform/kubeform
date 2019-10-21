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

type DynamodbGlobalTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DynamodbGlobalTableSpec   `json:"spec,omitempty"`
	Status            DynamodbGlobalTableStatus `json:"status,omitempty"`
}

type DynamodbGlobalTableSpecReplica struct {
	RegionName string `json:"regionName" tf:"region_name"`
}

type DynamodbGlobalTableSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn     string                           `json:"arn,omitempty" tf:"arn,omitempty"`
	Name    string                           `json:"name" tf:"name"`
	Replica []DynamodbGlobalTableSpecReplica `json:"replica" tf:"replica"`
}

type DynamodbGlobalTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DynamodbGlobalTableSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DynamodbGlobalTableList is a list of DynamodbGlobalTables
type DynamodbGlobalTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DynamodbGlobalTable CRD objects
	Items []DynamodbGlobalTable `json:"items,omitempty"`
}