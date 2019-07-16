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

type GlueConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueConnectionSpec   `json:"spec,omitempty"`
	Status            GlueConnectionStatus `json:"status,omitempty"`
}

type GlueConnectionSpecPhysicalConnectionRequirements struct {
	// +optional
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// +optional
	SecurityGroupIdList []string `json:"security_group_id_list,omitempty"`
	// +optional
	SubnetId string `json:"subnet_id,omitempty"`
}

type GlueConnectionSpec struct {
	ConnectionProperties map[string]string `json:"connection_properties"`
	// +optional
	ConnectionType string `json:"connection_type,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	MatchCriteria []string `json:"match_criteria,omitempty"`
	Name          string   `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PhysicalConnectionRequirements *[]GlueConnectionSpec `json:"physical_connection_requirements,omitempty"`
}

type GlueConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
