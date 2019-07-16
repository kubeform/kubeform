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

type LbListener struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbListenerSpec   `json:"spec,omitempty"`
	Status            LbListenerStatus `json:"status,omitempty"`
}

type LbListenerSpecDefaultActionAuthenticateCognito struct {
	// +optional
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params,omitempty"`
	UserPoolArn                      string            `json:"user_pool_arn"`
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
}

type LbListenerSpecDefaultActionAuthenticateOidc struct {
	// +optional
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params,omitempty"`
	AuthorizationEndpoint            string            `json:"authorization_endpoint"`
	ClientId                         string            `json:"client_id"`
	ClientSecret                     string            `json:"client_secret"`
	Issuer                           string            `json:"issuer"`
	TokenEndpoint                    string            `json:"token_endpoint"`
	UserInfoEndpoint                 string            `json:"user_info_endpoint"`
}

type LbListenerSpecDefaultActionFixedResponse struct {
	ContentType string `json:"content_type"`
	// +optional
	MessageBody string `json:"message_body,omitempty"`
}

type LbListenerSpecDefaultActionRedirect struct {
	// +optional
	Host string `json:"host,omitempty"`
	// +optional
	Path string `json:"path,omitempty"`
	// +optional
	Port string `json:"port,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty"`
	// +optional
	Query      string `json:"query,omitempty"`
	StatusCode string `json:"status_code"`
}

type LbListenerSpecDefaultAction struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticateCognito *[]LbListenerSpecDefaultAction `json:"authenticate_cognito,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticateOidc *[]LbListenerSpecDefaultAction `json:"authenticate_oidc,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FixedResponse *[]LbListenerSpecDefaultAction `json:"fixed_response,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Redirect *[]LbListenerSpecDefaultAction `json:"redirect,omitempty"`
	// +optional
	TargetGroupArn string `json:"target_group_arn,omitempty"`
	Type           string `json:"type"`
}

type LbListenerSpec struct {
	// +optional
	CertificateArn  string           `json:"certificate_arn,omitempty"`
	DefaultAction   []LbListenerSpec `json:"default_action"`
	LoadBalancerArn string           `json:"load_balancer_arn"`
	Port            int              `json:"port"`
	// +optional
	Protocol string `json:"protocol,omitempty"`
}

type LbListenerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbListenerList is a list of LbListeners
type LbListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbListener CRD objects
	Items []LbListener `json:"items,omitempty"`
}
