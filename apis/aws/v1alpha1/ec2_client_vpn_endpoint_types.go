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

type Ec2ClientVpnEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2ClientVpnEndpointSpec   `json:"spec,omitempty"`
	Status            Ec2ClientVpnEndpointStatus `json:"status,omitempty"`
}

type Ec2ClientVpnEndpointSpecAuthenticationOptions struct {
	// +optional
	ActiveDirectoryId string `json:"active_directory_id,omitempty"`
	// +optional
	RootCertificateChainArn string `json:"root_certificate_chain_arn,omitempty"`
	Type                    string `json:"type"`
}

type Ec2ClientVpnEndpointSpecConnectionLogOptions struct {
	// +optional
	CloudwatchLogGroup string `json:"cloudwatch_log_group,omitempty"`
	// +optional
	CloudwatchLogStream string `json:"cloudwatch_log_stream,omitempty"`
	Enabled             bool   `json:"enabled"`
}

type Ec2ClientVpnEndpointSpec struct {
	// +kubebuilder:validation:MaxItems=1
	AuthenticationOptions []Ec2ClientVpnEndpointSpec `json:"authentication_options"`
	ClientCidrBlock       string                     `json:"client_cidr_block"`
	// +kubebuilder:validation:MaxItems=1
	ConnectionLogOptions []Ec2ClientVpnEndpointSpec `json:"connection_log_options"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DnsServers           []string `json:"dns_servers,omitempty"`
	ServerCertificateArn string   `json:"server_certificate_arn"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	TransportProtocol string `json:"transport_protocol,omitempty"`
}

type Ec2ClientVpnEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Ec2ClientVpnEndpointList is a list of Ec2ClientVpnEndpoints
type Ec2ClientVpnEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ec2ClientVpnEndpoint CRD objects
	Items []Ec2ClientVpnEndpoint `json:"items,omitempty"`
}
