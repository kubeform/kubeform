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

type AwsEc2ClientVpnEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEc2ClientVpnEndpointSpec   `json:"spec,omitempty"`
	Status            AwsEc2ClientVpnEndpointStatus `json:"status,omitempty"`
}

type AwsEc2ClientVpnEndpointSpecAuthenticationOptions struct {
	Type                    string `json:"type"`
	ActiveDirectoryId       string `json:"active_directory_id"`
	RootCertificateChainArn string `json:"root_certificate_chain_arn"`
}

type AwsEc2ClientVpnEndpointSpecConnectionLogOptions struct {
	CloudwatchLogGroup  string `json:"cloudwatch_log_group"`
	CloudwatchLogStream string `json:"cloudwatch_log_stream"`
	Enabled             bool   `json:"enabled"`
}

type AwsEc2ClientVpnEndpointSpec struct {
	Status                string                        `json:"status"`
	DnsServers            []string                      `json:"dns_servers"`
	TransportProtocol     string                        `json:"transport_protocol"`
	AuthenticationOptions []AwsEc2ClientVpnEndpointSpec `json:"authentication_options"`
	ConnectionLogOptions  []AwsEc2ClientVpnEndpointSpec `json:"connection_log_options"`
	DnsName               string                        `json:"dns_name"`
	Tags                  map[string]string             `json:"tags"`
	Description           string                        `json:"description"`
	ClientCidrBlock       string                        `json:"client_cidr_block"`
	ServerCertificateArn  string                        `json:"server_certificate_arn"`
}



type AwsEc2ClientVpnEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsEc2ClientVpnEndpointList is a list of AwsEc2ClientVpnEndpoints
type AwsEc2ClientVpnEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEc2ClientVpnEndpoint CRD objects
	Items []AwsEc2ClientVpnEndpoint `json:"items,omitempty"`
}