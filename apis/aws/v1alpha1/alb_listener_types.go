package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AlbListener struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlbListenerSpec   `json:"spec,omitempty"`
	Status            AlbListenerStatus `json:"status,omitempty"`
}

type AlbListenerSpecDefaultActionAuthenticateCognito struct {
	// +optional
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`
	// +optional
	OnUnauthenticatedRequest string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`
	// +optional
	Scope string `json:"scope,omitempty" tf:"scope,omitempty"`
	// +optional
	SessionCookieName string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`
	// +optional
	SessionTimeout   int    `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`
	UserPoolArn      string `json:"userPoolArn" tf:"user_pool_arn"`
	UserPoolClientID string `json:"userPoolClientID" tf:"user_pool_client_id"`
	UserPoolDomain   string `json:"userPoolDomain" tf:"user_pool_domain"`
}

type AlbListenerSpecDefaultActionAuthenticateOidc struct {
	// +optional
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`
	AuthorizationEndpoint            string            `json:"authorizationEndpoint" tf:"authorization_endpoint"`
	ClientID                         string            `json:"clientID" tf:"client_id"`
	ClientSecret                     string            `json:"-" sensitive:"true" tf:"client_secret"`
	Issuer                           string            `json:"issuer" tf:"issuer"`
	// +optional
	OnUnauthenticatedRequest string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`
	// +optional
	Scope string `json:"scope,omitempty" tf:"scope,omitempty"`
	// +optional
	SessionCookieName string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`
	// +optional
	SessionTimeout   int    `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`
	TokenEndpoint    string `json:"tokenEndpoint" tf:"token_endpoint"`
	UserInfoEndpoint string `json:"userInfoEndpoint" tf:"user_info_endpoint"`
}

type AlbListenerSpecDefaultActionFixedResponse struct {
	ContentType string `json:"contentType" tf:"content_type"`
	// +optional
	MessageBody string `json:"messageBody,omitempty" tf:"message_body,omitempty"`
	// +optional
	StatusCode string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type AlbListenerSpecDefaultActionRedirect struct {
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	Port string `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// +optional
	Query      string `json:"query,omitempty" tf:"query,omitempty"`
	StatusCode string `json:"statusCode" tf:"status_code"`
}

type AlbListenerSpecDefaultAction struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticateCognito []AlbListenerSpecDefaultActionAuthenticateCognito `json:"authenticateCognito,omitempty" tf:"authenticate_cognito,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticateOidc []AlbListenerSpecDefaultActionAuthenticateOidc `json:"authenticateOidc,omitempty" tf:"authenticate_oidc,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FixedResponse []AlbListenerSpecDefaultActionFixedResponse `json:"fixedResponse,omitempty" tf:"fixed_response,omitempty"`
	// +optional
	Order int `json:"order,omitempty" tf:"order,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Redirect []AlbListenerSpecDefaultActionRedirect `json:"redirect,omitempty" tf:"redirect,omitempty"`
	// +optional
	TargetGroupArn string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`
	Type           string `json:"type" tf:"type"`
}

type AlbListenerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	CertificateArn  string                         `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`
	DefaultAction   []AlbListenerSpecDefaultAction `json:"defaultAction" tf:"default_action"`
	LoadBalancerArn string                         `json:"loadBalancerArn" tf:"load_balancer_arn"`
	Port            int                            `json:"port" tf:"port"`
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// +optional
	SslPolicy string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`
}

type AlbListenerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AlbListenerSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AlbListenerList is a list of AlbListeners
type AlbListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AlbListener CRD objects
	Items []AlbListener `json:"items,omitempty"`
}
