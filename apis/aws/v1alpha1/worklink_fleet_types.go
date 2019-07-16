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

type WorklinkFleet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorklinkFleetSpec   `json:"spec,omitempty"`
	Status            WorklinkFleetStatus `json:"status,omitempty"`
}

type WorklinkFleetSpecIdentityProvider struct {
	SamlMetadata string `json:"saml_metadata"`
	Type         string `json:"type"`
}

type WorklinkFleetSpecNetwork struct {
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIds []string `json:"security_group_ids"`
	// +kubebuilder:validation:UniqueItems=true
	SubnetIds []string `json:"subnet_ids"`
	VpcId     string   `json:"vpc_id"`
}

type WorklinkFleetSpec struct {
	// +optional
	AuditStreamArn string `json:"audit_stream_arn,omitempty"`
	// +optional
	DeviceCaCertificate string `json:"device_ca_certificate,omitempty"`
	// +optional
	DisplayName string `json:"display_name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	IdentityProvider *[]WorklinkFleetSpec `json:"identity_provider,omitempty"`
	Name             string               `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Network *[]WorklinkFleetSpec `json:"network,omitempty"`
	// +optional
	OptimizeForEndUserLocation bool `json:"optimize_for_end_user_location,omitempty"`
}

type WorklinkFleetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WorklinkFleetList is a list of WorklinkFleets
type WorklinkFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WorklinkFleet CRD objects
	Items []WorklinkFleet `json:"items,omitempty"`
}
