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
	Port       string `json:"port"`
	Protocol   string `json:"protocol"`
	Query      string `json:"query"`
	StatusCode string `json:"status_code"`
	Host       string `json:"host"`
	Path       string `json:"path"`
}

type AwsLbListenerSpecDefaultActionFixedResponse struct {
	ContentType string `json:"content_type"`
	MessageBody string `json:"message_body"`
	StatusCode  string `json:"status_code"`
}

type AwsLbListenerSpecDefaultActionAuthenticateCognito struct {
	SessionTimeout                   int               `json:"session_timeout"`
	UserPoolArn                      string            `json:"user_pool_arn"`
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
}

type AwsLbListenerSpecDefaultActionAuthenticateOidc struct {
	ClientId                         string            `json:"client_id"`
	ClientSecret                     string            `json:"client_secret"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   int               `json:"session_timeout"`
	AuthorizationEndpoint            string            `json:"authorization_endpoint"`
	Issuer                           string            `json:"issuer"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	TokenEndpoint                    string            `json:"token_endpoint"`
	UserInfoEndpoint                 string            `json:"user_info_endpoint"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
}

type AwsLbListenerSpecDefaultAction struct {
	Type                string                           `json:"type"`
	Order               int                              `json:"order"`
	TargetGroupArn      string                           `json:"target_group_arn"`
	Redirect            []AwsLbListenerSpecDefaultAction `json:"redirect"`
	FixedResponse       []AwsLbListenerSpecDefaultAction `json:"fixed_response"`
	AuthenticateCognito []AwsLbListenerSpecDefaultAction `json:"authenticate_cognito"`
	AuthenticateOidc    []AwsLbListenerSpecDefaultAction `json:"authenticate_oidc"`
}

type AwsLbListenerSpec struct {
	Port            int                 `json:"port"`
	Protocol        string              `json:"protocol"`
	SslPolicy       string              `json:"ssl_policy"`
	CertificateArn  string              `json:"certificate_arn"`
	DefaultAction   []AwsLbListenerSpec `json:"default_action"`
	Arn             string              `json:"arn"`
	LoadBalancerArn string              `json:"load_balancer_arn"`
}



type AwsLbListenerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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