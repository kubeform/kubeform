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

type AwsAlbListener struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAlbListenerSpec   `json:"spec,omitempty"`
	Status            AwsAlbListenerStatus `json:"status,omitempty"`
}

type AwsAlbListenerSpecDefaultActionAuthenticateCognito struct {
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   int               `json:"session_timeout"`
	UserPoolArn                      string            `json:"user_pool_arn"`
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
}

type AwsAlbListenerSpecDefaultActionAuthenticateOidc struct {
	SessionTimeout                   int               `json:"session_timeout"`
	UserInfoEndpoint                 string            `json:"user_info_endpoint"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	ClientSecret                     string            `json:"client_secret"`
	Issuer                           string            `json:"issuer"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	TokenEndpoint                    string            `json:"token_endpoint"`
	AuthorizationEndpoint            string            `json:"authorization_endpoint"`
	ClientId                         string            `json:"client_id"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
}

type AwsAlbListenerSpecDefaultActionRedirect struct {
	Host       string `json:"host"`
	Path       string `json:"path"`
	Port       string `json:"port"`
	Protocol   string `json:"protocol"`
	Query      string `json:"query"`
	StatusCode string `json:"status_code"`
}

type AwsAlbListenerSpecDefaultActionFixedResponse struct {
	ContentType string `json:"content_type"`
	MessageBody string `json:"message_body"`
	StatusCode  string `json:"status_code"`
}

type AwsAlbListenerSpecDefaultAction struct {
	AuthenticateCognito []AwsAlbListenerSpecDefaultAction `json:"authenticate_cognito"`
	AuthenticateOidc    []AwsAlbListenerSpecDefaultAction `json:"authenticate_oidc"`
	Type                string                            `json:"type"`
	Order               int                               `json:"order"`
	TargetGroupArn      string                            `json:"target_group_arn"`
	Redirect            []AwsAlbListenerSpecDefaultAction `json:"redirect"`
	FixedResponse       []AwsAlbListenerSpecDefaultAction `json:"fixed_response"`
}

type AwsAlbListenerSpec struct {
	LoadBalancerArn string               `json:"load_balancer_arn"`
	Port            int                  `json:"port"`
	Protocol        string               `json:"protocol"`
	SslPolicy       string               `json:"ssl_policy"`
	CertificateArn  string               `json:"certificate_arn"`
	DefaultAction   []AwsAlbListenerSpec `json:"default_action"`
	Arn             string               `json:"arn"`
}



type AwsAlbListenerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAlbListenerList is a list of AwsAlbListeners
type AwsAlbListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAlbListener CRD objects
	Items []AwsAlbListener `json:"items,omitempty"`
}