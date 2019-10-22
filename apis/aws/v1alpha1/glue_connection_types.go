package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type GlueConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueConnectionSpec   `json:"spec,omitempty"`
	Status            GlueConnectionStatus `json:"status,omitempty"`
}

type GlueConnectionSpecPhysicalConnectionRequirements struct {
	// +optional
	AvailabilityZone string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	// +optional
	SecurityGroupIDList []string `json:"securityGroupIDList,omitempty" tf:"security_group_id_list,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
}

type GlueConnectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CatalogID            string            `json:"catalogID,omitempty" tf:"catalog_id,omitempty"`
	ConnectionProperties map[string]string `json:"connectionProperties" tf:"connection_properties"`
	// +optional
	ConnectionType string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	MatchCriteria []string `json:"matchCriteria,omitempty" tf:"match_criteria,omitempty"`
	Name          string   `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PhysicalConnectionRequirements []GlueConnectionSpecPhysicalConnectionRequirements `json:"physicalConnectionRequirements,omitempty" tf:"physical_connection_requirements,omitempty"`
}

type GlueConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GlueConnectionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlueConnectionList is a list of GlueConnections
type GlueConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlueConnection CRD objects
	Items []GlueConnection `json:"items,omitempty"`
}
