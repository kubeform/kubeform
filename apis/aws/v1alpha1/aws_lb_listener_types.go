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

type AwsLbListener struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLbListenerSpec   `json:"spec,omitempty"`
	Status            AwsLbListenerStatus `json:"status,omitempty"`
}

type AwsLbListenerSpecDefaultActionRedirect struct {
	Path       string `json:"path"`
	Port       string `json:"port"`
	Protocol   string `json:"protocol"`
	Query      string `json:"query"`
	StatusCode string `json:"status_code"`
	Host       string `json:"host"`
}

type AwsLbListenerSpecDefaultActionFixedResponse struct {
	StatusCode  string `json:"status_code"`
	ContentType string `json:"content_type"`
	MessageBody string `json:"message_body"`
}

type AwsLbListenerSpecDefaultActionAuthenticateCognito struct {
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   int               `json:"session_timeout"`
	UserPoolArn                      string            `json:"user_pool_arn"`
}

type AwsLbListenerSpecDefaultActionAuthenticateOidc struct {
	AuthorizationEndpoint            string            `json:"authorization_endpoint"`
	ClientId                         string            `json:"client_id"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	ClientSecret                     string            `json:"client_secret"`
	Issuer                           string            `json:"issuer"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	SessionTimeout                   int               `json:"session_timeout"`
	TokenEndpoint                    string            `json:"token_endpoint"`
	UserInfoEndpoint                 string            `json:"user_info_endpoint"`
}

type AwsLbListenerSpecDefaultAction struct {
	Order               int                              `json:"order"`
	TargetGroupArn      string                           `json:"target_group_arn"`
	Redirect            []AwsLbListenerSpecDefaultAction `json:"redirect"`
	FixedResponse       []AwsLbListenerSpecDefaultAction `json:"fixed_response"`
	AuthenticateCognito []AwsLbListenerSpecDefaultAction `json:"authenticate_cognito"`
	AuthenticateOidc    []AwsLbListenerSpecDefaultAction `json:"authenticate_oidc"`
	Type                string                           `json:"type"`
}

type AwsLbListenerSpec struct {
	CertificateArn  string              `json:"certificate_arn"`
	DefaultAction   []AwsLbListenerSpec `json:"default_action"`
	Arn             string              `json:"arn"`
	LoadBalancerArn string              `json:"load_balancer_arn"`
	Port            int                 `json:"port"`
	Protocol        string              `json:"protocol"`
	SslPolicy       string              `json:"ssl_policy"`
}

type AwsLbListenerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsLbListenerList is a list of AwsLbListeners
type AwsLbListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLbListener CRD objects
	Items []AwsLbListener `json:"items,omitempty"`
}
