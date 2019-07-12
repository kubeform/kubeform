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

type AwsAlbListenerRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAlbListenerRuleSpec   `json:"spec,omitempty"`
	Status            AwsAlbListenerRuleStatus `json:"status,omitempty"`
}

type AwsAlbListenerRuleSpecCondition struct {
	Field  string   `json:"field"`
	Values []string `json:"values"`
}

type AwsAlbListenerRuleSpecActionRedirect struct {
	Host       string `json:"host"`
	Path       string `json:"path"`
	Port       string `json:"port"`
	Protocol   string `json:"protocol"`
	Query      string `json:"query"`
	StatusCode string `json:"status_code"`
}

type AwsAlbListenerRuleSpecActionFixedResponse struct {
	ContentType string `json:"content_type"`
	MessageBody string `json:"message_body"`
	StatusCode  string `json:"status_code"`
}

type AwsAlbListenerRuleSpecActionAuthenticateCognito struct {
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   int               `json:"session_timeout"`
	UserPoolArn                      string            `json:"user_pool_arn"`
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
}

type AwsAlbListenerRuleSpecActionAuthenticateOidc struct {
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	TokenEndpoint                    string            `json:"token_endpoint"`
	UserInfoEndpoint                 string            `json:"user_info_endpoint"`
	AuthorizationEndpoint            string            `json:"authorization_endpoint"`
	ClientId                         string            `json:"client_id"`
	ClientSecret                     string            `json:"client_secret"`
	Issuer                           string            `json:"issuer"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	SessionTimeout                   int               `json:"session_timeout"`
}

type AwsAlbListenerRuleSpecAction struct {
	Type                string                         `json:"type"`
	Order               int                            `json:"order"`
	TargetGroupArn      string                         `json:"target_group_arn"`
	Redirect            []AwsAlbListenerRuleSpecAction `json:"redirect"`
	FixedResponse       []AwsAlbListenerRuleSpecAction `json:"fixed_response"`
	AuthenticateCognito []AwsAlbListenerRuleSpecAction `json:"authenticate_cognito"`
	AuthenticateOidc    []AwsAlbListenerRuleSpecAction `json:"authenticate_oidc"`
}

type AwsAlbListenerRuleSpec struct {
	Condition   []AwsAlbListenerRuleSpec `json:"condition"`
	Arn         string                   `json:"arn"`
	ListenerArn string                   `json:"listener_arn"`
	Priority    int                      `json:"priority"`
	Action      []AwsAlbListenerRuleSpec `json:"action"`
}

type AwsAlbListenerRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAlbListenerRuleList is a list of AwsAlbListenerRules
type AwsAlbListenerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAlbListenerRule CRD objects
	Items []AwsAlbListenerRule `json:"items,omitempty"`
}
