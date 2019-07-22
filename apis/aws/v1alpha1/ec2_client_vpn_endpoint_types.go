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

type Ec2ClientVPNEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2ClientVPNEndpointSpec   `json:"spec,omitempty"`
	Status            Ec2ClientVPNEndpointStatus `json:"status,omitempty"`
}

type Ec2ClientVPNEndpointSpecAuthenticationOptions struct {
	// +optional
	ActiveDirectoryID string `json:"activeDirectoryID,omitempty" tf:"active_directory_id,omitempty"`
	// +optional
	RootCertificateChainArn string `json:"rootCertificateChainArn,omitempty" tf:"root_certificate_chain_arn,omitempty"`
	Type                    string `json:"type" tf:"type"`
}

type Ec2ClientVPNEndpointSpecConnectionLogOptions struct {
	// +optional
	CloudwatchLogGroup string `json:"cloudwatchLogGroup,omitempty" tf:"cloudwatch_log_group,omitempty"`
	// +optional
	CloudwatchLogStream string `json:"cloudwatchLogStream,omitempty" tf:"cloudwatch_log_stream,omitempty"`
	Enabled             bool   `json:"enabled" tf:"enabled"`
}

type Ec2ClientVPNEndpointSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +kubebuilder:validation:MaxItems=1
	AuthenticationOptions []Ec2ClientVPNEndpointSpecAuthenticationOptions `json:"authenticationOptions" tf:"authentication_options"`
	ClientCIDRBlock       string                                          `json:"clientCIDRBlock" tf:"client_cidr_block"`
	// +kubebuilder:validation:MaxItems=1
	ConnectionLogOptions []Ec2ClientVPNEndpointSpecConnectionLogOptions `json:"connectionLogOptions" tf:"connection_log_options"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DnsServers           []string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`
	ServerCertificateArn string   `json:"serverCertificateArn" tf:"server_certificate_arn"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TransportProtocol string `json:"transportProtocol,omitempty" tf:"transport_protocol,omitempty"`
}

type Ec2ClientVPNEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Ec2ClientVPNEndpointList is a list of Ec2ClientVPNEndpoints
type Ec2ClientVPNEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ec2ClientVPNEndpoint CRD objects
	Items []Ec2ClientVPNEndpoint `json:"items,omitempty"`
}
