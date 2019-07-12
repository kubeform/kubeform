package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListenerRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLbListenerRuleSpec   `json:"spec,omitempty"`
	Status            AwsLbListenerRuleStatus `json:"status,omitempty"`
}

type AwsLbListenerRuleSpecActionAuthenticateOidc struct {
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	ClientId                         string            `json:"client_id"`
	ClientSecret                     string            `json:"client_secret"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	UserInfoEndpoint                 string            `json:"user_info_endpoint"`
	AuthorizationEndpoint            string            `json:"authorization_endpoint"`
	Issuer                           string            `json:"issuer"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   int               `json:"session_timeout"`
	TokenEndpoint                    string            `json:"token_endpoint"`
}

type AwsLbListenerRuleSpecActionRedirect struct {
	Host       string `json:"host"`
	Path       string `json:"path"`
	Port       string `json:"port"`
	Protocol   string `json:"protocol"`
	Query      string `json:"query"`
	StatusCode string `json:"status_code"`
}

type AwsLbListenerRuleSpecActionFixedResponse struct {
	ContentType string `json:"content_type"`
	MessageBody string `json:"message_body"`
	StatusCode  string `json:"status_code"`
}

type AwsLbListenerRuleSpecActionAuthenticateCognito struct {
	UserPoolArn                      string            `json:"user_pool_arn"`
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   int               `json:"session_timeout"`
}

type AwsLbListenerRuleSpecAction struct {
	AuthenticateOidc    []AwsLbListenerRuleSpecAction `json:"authenticate_oidc"`
	Type                string                        `json:"type"`
	Order               int                           `json:"order"`
	TargetGroupArn      string                        `json:"target_group_arn"`
	Redirect            []AwsLbListenerRuleSpecAction `json:"redirect"`
	FixedResponse       []AwsLbListenerRuleSpecAction `json:"fixed_response"`
	AuthenticateCognito []AwsLbListenerRuleSpecAction `json:"authenticate_cognito"`
}

type AwsLbListenerRuleSpecCondition struct {
	Field  string   `json:"field"`
	Values []string `json:"values"`
}

type AwsLbListenerRuleSpec struct {
	ListenerArn string                  `json:"listener_arn"`
	Priority    int                     `json:"priority"`
	Action      []AwsLbListenerRuleSpec `json:"action"`
	Condition   []AwsLbListenerRuleSpec `json:"condition"`
	Arn         string                  `json:"arn"`
}



type AwsLbListenerRuleStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbListenerRuleList is a list of AwsLbListenerRules
type AwsLbListenerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLbListenerRule CRD objects
	Items []AwsLbListenerRule `json:"items,omitempty"`
}