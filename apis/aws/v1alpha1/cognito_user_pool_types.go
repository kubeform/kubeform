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

type CognitoUserPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserPoolSpec   `json:"spec,omitempty"`
	Status            CognitoUserPoolStatus `json:"status,omitempty"`
}

type CognitoUserPoolSpecDeviceConfiguration struct {
	// +optional
	ChallengeRequiredOnNewDevice bool `json:"challenge_required_on_new_device,omitempty"`
	// +optional
	DeviceOnlyRememberedOnUserPrompt bool `json:"device_only_remembered_on_user_prompt,omitempty"`
}

type CognitoUserPoolSpecEmailConfiguration struct {
	// +optional
	EmailSendingAccount string `json:"email_sending_account,omitempty"`
	// +optional
	ReplyToEmailAddress string `json:"reply_to_email_address,omitempty"`
	// +optional
	SourceArn string `json:"source_arn,omitempty"`
}

type CognitoUserPoolSpecSchemaNumberAttributeConstraints struct {
	// +optional
	MaxValue string `json:"max_value,omitempty"`
	// +optional
	MinValue string `json:"min_value,omitempty"`
}

type CognitoUserPoolSpecSchemaStringAttributeConstraints struct {
	// +optional
	MaxLength string `json:"max_length,omitempty"`
	// +optional
	MinLength string `json:"min_length,omitempty"`
}

type CognitoUserPoolSpecSchema struct {
	AttributeDataType string `json:"attribute_data_type"`
	// +optional
	DeveloperOnlyAttribute bool `json:"developer_only_attribute,omitempty"`
	// +optional
	Mutable bool   `json:"mutable,omitempty"`
	Name    string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NumberAttributeConstraints *[]CognitoUserPoolSpecSchema `json:"number_attribute_constraints,omitempty"`
	// +optional
	Required bool `json:"required,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StringAttributeConstraints *[]CognitoUserPoolSpecSchema `json:"string_attribute_constraints,omitempty"`
}

type CognitoUserPoolSpecSmsConfiguration struct {
	ExternalId   string `json:"external_id"`
	SnsCallerArn string `json:"sns_caller_arn"`
}

type CognitoUserPoolSpecUserPoolAddOns struct {
	AdvancedSecurityMode string `json:"advanced_security_mode"`
}

type CognitoUserPoolSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AliasAttributes []string `json:"alias_attributes,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AutoVerifiedAttributes []string `json:"auto_verified_attributes,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DeviceConfiguration *[]CognitoUserPoolSpec `json:"device_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EmailConfiguration *[]CognitoUserPoolSpec `json:"email_configuration,omitempty"`
	// +optional
	MfaConfiguration string `json:"mfa_configuration,omitempty"`
	Name             string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Schema *[]CognitoUserPoolSpec `json:"schema,omitempty"`
	// +optional
	SmsAuthenticationMessage string `json:"sms_authentication_message,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SmsConfiguration *[]CognitoUserPoolSpec `json:"sms_configuration,omitempty"`
	// +optional
	SmsVerificationMessage string `json:"sms_verification_message,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	UserPoolAddOns *[]CognitoUserPoolSpec `json:"user_pool_add_ons,omitempty"`
	// +optional
	UsernameAttributes []string `json:"username_attributes,omitempty"`
}

type CognitoUserPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CognitoUserPoolList is a list of CognitoUserPools
type CognitoUserPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CognitoUserPool CRD objects
	Items []CognitoUserPool `json:"items,omitempty"`
}
