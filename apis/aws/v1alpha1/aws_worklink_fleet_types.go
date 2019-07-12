package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWorklinkFleet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWorklinkFleetSpec   `json:"spec,omitempty"`
	Status            AwsWorklinkFleetStatus `json:"status,omitempty"`
}

type AwsWorklinkFleetSpecNetwork struct {
	VpcId            string   `json:"vpc_id"`
	SecurityGroupIds []string `json:"security_group_ids"`
	SubnetIds        []string `json:"subnet_ids"`
}

type AwsWorklinkFleetSpecIdentityProvider struct {
	Type         string `json:"type"`
	SamlMetadata string `json:"saml_metadata"`
}

type AwsWorklinkFleetSpec struct {
	Name                       string                 `json:"name"`
	DisplayName                string                 `json:"display_name"`
	DeviceCaCertificate        string                 `json:"device_ca_certificate"`
	LastUpdatedTime            string                 `json:"last_updated_time"`
	OptimizeForEndUserLocation bool                   `json:"optimize_for_end_user_location"`
	Arn                        string                 `json:"arn"`
	AuditStreamArn             string                 `json:"audit_stream_arn"`
	Network                    []AwsWorklinkFleetSpec `json:"network"`
	IdentityProvider           []AwsWorklinkFleetSpec `json:"identity_provider"`
	CompanyCode                string                 `json:"company_code"`
	CreatedTime                string                 `json:"created_time"`
}



type AwsWorklinkFleetStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWorklinkFleetList is a list of AwsWorklinkFleets
type AwsWorklinkFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWorklinkFleet CRD objects
	Items []AwsWorklinkFleet `json:"items,omitempty"`
}