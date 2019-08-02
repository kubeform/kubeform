package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	SamlMetadata string `json:"samlMetadata" tf:"saml_metadata"`
	Type         string `json:"type" tf:"type"`
}

type WorklinkFleetSpecNetwork struct {
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIDS []string `json:"securityGroupIDS" tf:"security_group_ids"`
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS" tf:"subnet_ids"`
	VpcID     string   `json:"vpcID" tf:"vpc_id"`
}

type WorklinkFleetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AuditStreamArn string `json:"auditStreamArn,omitempty" tf:"audit_stream_arn,omitempty"`
	// +optional
	DeviceCaCertificate string `json:"deviceCaCertificate,omitempty" tf:"device_ca_certificate,omitempty"`
	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	IdentityProvider []WorklinkFleetSpecIdentityProvider `json:"identityProvider,omitempty" tf:"identity_provider,omitempty"`
	Name             string                              `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Network []WorklinkFleetSpecNetwork `json:"network,omitempty" tf:"network,omitempty"`
	// +optional
	OptimizeForEndUserLocation bool `json:"optimizeForEndUserLocation,omitempty" tf:"optimize_for_end_user_location,omitempty"`
}

type WorklinkFleetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *WorklinkFleetSpec `json:"state,omitempty"`
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
