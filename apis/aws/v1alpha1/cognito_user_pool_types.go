package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CognitoUserPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserPoolSpec   `json:"spec,omitempty"`
	Status            CognitoUserPoolStatus `json:"status,omitempty"`
}

type CognitoUserPoolSpecAdminCreateUserConfigInviteMessageTemplate struct {
	// +optional
	EmailMessage string `json:"emailMessage,omitempty" tf:"email_message,omitempty"`
	// +optional
	EmailSubject string `json:"emailSubject,omitempty" tf:"email_subject,omitempty"`
	// +optional
	SmsMessage string `json:"smsMessage,omitempty" tf:"sms_message,omitempty"`
}

type CognitoUserPoolSpecAdminCreateUserConfig struct {
	// +optional
	AllowAdminCreateUserOnly bool `json:"allowAdminCreateUserOnly,omitempty" tf:"allow_admin_create_user_only,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	InviteMessageTemplate []CognitoUserPoolSpecAdminCreateUserConfigInviteMessageTemplate `json:"inviteMessageTemplate,omitempty" tf:"invite_message_template,omitempty"`
	// +optional
	UnusedAccountValidityDays int `json:"unusedAccountValidityDays,omitempty" tf:"unused_account_validity_days,omitempty"`
}

type CognitoUserPoolSpecDeviceConfiguration struct {
	// +optional
	ChallengeRequiredOnNewDevice bool `json:"challengeRequiredOnNewDevice,omitempty" tf:"challenge_required_on_new_device,omitempty"`
	// +optional
	DeviceOnlyRememberedOnUserPrompt bool `json:"deviceOnlyRememberedOnUserPrompt,omitempty" tf:"device_only_remembered_on_user_prompt,omitempty"`
}

type CognitoUserPoolSpecEmailConfiguration struct {
	// +optional
	EmailSendingAccount string `json:"emailSendingAccount,omitempty" tf:"email_sending_account,omitempty"`
	// +optional
	ReplyToEmailAddress string `json:"replyToEmailAddress,omitempty" tf:"reply_to_email_address,omitempty"`
	// +optional
	SourceArn string `json:"sourceArn,omitempty" tf:"source_arn,omitempty"`
}

type CognitoUserPoolSpecLambdaConfig struct {
	// +optional
	CreateAuthChallenge string `json:"createAuthChallenge,omitempty" tf:"create_auth_challenge,omitempty"`
	// +optional
	CustomMessage string `json:"customMessage,omitempty" tf:"custom_message,omitempty"`
	// +optional
	DefineAuthChallenge string `json:"defineAuthChallenge,omitempty" tf:"define_auth_challenge,omitempty"`
	// +optional
	PostAuthentication string `json:"postAuthentication,omitempty" tf:"post_authentication,omitempty"`
	// +optional
	PostConfirmation string `json:"postConfirmation,omitempty" tf:"post_confirmation,omitempty"`
	// +optional
	PreAuthentication string `json:"preAuthentication,omitempty" tf:"pre_authentication,omitempty"`
	// +optional
	PreSignUp string `json:"preSignUp,omitempty" tf:"pre_sign_up,omitempty"`
	// +optional
	PreTokenGeneration string `json:"preTokenGeneration,omitempty" tf:"pre_token_generation,omitempty"`
	// +optional
	UserMigration string `json:"userMigration,omitempty" tf:"user_migration,omitempty"`
	// +optional
	VerifyAuthChallengeResponse string `json:"verifyAuthChallengeResponse,omitempty" tf:"verify_auth_challenge_response,omitempty"`
}

type CognitoUserPoolSpecPasswordPolicy struct {
	// +optional
	MinimumLength int `json:"minimumLength,omitempty" tf:"minimum_length,omitempty"`
	// +optional
	RequireLowercase bool `json:"requireLowercase,omitempty" tf:"require_lowercase,omitempty"`
	// +optional
	RequireNumbers bool `json:"requireNumbers,omitempty" tf:"require_numbers,omitempty"`
	// +optional
	RequireSymbols bool `json:"requireSymbols,omitempty" tf:"require_symbols,omitempty"`
	// +optional
	RequireUppercase bool `json:"requireUppercase,omitempty" tf:"require_uppercase,omitempty"`
}

type CognitoUserPoolSpecSchemaNumberAttributeConstraints struct {
	// +optional
	MaxValue string `json:"maxValue,omitempty" tf:"max_value,omitempty"`
	// +optional
	MinValue string `json:"minValue,omitempty" tf:"min_value,omitempty"`
}

type CognitoUserPoolSpecSchemaStringAttributeConstraints struct {
	// +optional
	MaxLength string `json:"maxLength,omitempty" tf:"max_length,omitempty"`
	// +optional
	MinLength string `json:"minLength,omitempty" tf:"min_length,omitempty"`
}

type CognitoUserPoolSpecSchema struct {
	AttributeDataType string `json:"attributeDataType" tf:"attribute_data_type"`
	// +optional
	DeveloperOnlyAttribute bool `json:"developerOnlyAttribute,omitempty" tf:"developer_only_attribute,omitempty"`
	// +optional
	Mutable bool   `json:"mutable,omitempty" tf:"mutable,omitempty"`
	Name    string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NumberAttributeConstraints []CognitoUserPoolSpecSchemaNumberAttributeConstraints `json:"numberAttributeConstraints,omitempty" tf:"number_attribute_constraints,omitempty"`
	// +optional
	Required bool `json:"required,omitempty" tf:"required,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StringAttributeConstraints []CognitoUserPoolSpecSchemaStringAttributeConstraints `json:"stringAttributeConstraints,omitempty" tf:"string_attribute_constraints,omitempty"`
}

type CognitoUserPoolSpecSmsConfiguration struct {
	ExternalID   string `json:"externalID" tf:"external_id"`
	SnsCallerArn string `json:"snsCallerArn" tf:"sns_caller_arn"`
}

type CognitoUserPoolSpecUserPoolAddOns struct {
	AdvancedSecurityMode string `json:"advancedSecurityMode" tf:"advanced_security_mode"`
}

type CognitoUserPoolSpecVerificationMessageTemplate struct {
	// +optional
	DefaultEmailOption string `json:"defaultEmailOption,omitempty" tf:"default_email_option,omitempty"`
	// +optional
	EmailMessage string `json:"emailMessage,omitempty" tf:"email_message,omitempty"`
	// +optional
	EmailMessageByLink string `json:"emailMessageByLink,omitempty" tf:"email_message_by_link,omitempty"`
	// +optional
	EmailSubject string `json:"emailSubject,omitempty" tf:"email_subject,omitempty"`
	// +optional
	EmailSubjectByLink string `json:"emailSubjectByLink,omitempty" tf:"email_subject_by_link,omitempty"`
	// +optional
	SmsMessage string `json:"smsMessage,omitempty" tf:"sms_message,omitempty"`
}

type CognitoUserPoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AdminCreateUserConfig []CognitoUserPoolSpecAdminCreateUserConfig `json:"adminCreateUserConfig,omitempty" tf:"admin_create_user_config,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AliasAttributes []string `json:"aliasAttributes,omitempty" tf:"alias_attributes,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AutoVerifiedAttributes []string `json:"autoVerifiedAttributes,omitempty" tf:"auto_verified_attributes,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DeviceConfiguration []CognitoUserPoolSpecDeviceConfiguration `json:"deviceConfiguration,omitempty" tf:"device_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EmailConfiguration []CognitoUserPoolSpecEmailConfiguration `json:"emailConfiguration,omitempty" tf:"email_configuration,omitempty"`
	// +optional
	EmailVerificationMessage string `json:"emailVerificationMessage,omitempty" tf:"email_verification_message,omitempty"`
	// +optional
	EmailVerificationSubject string `json:"emailVerificationSubject,omitempty" tf:"email_verification_subject,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LambdaConfig []CognitoUserPoolSpecLambdaConfig `json:"lambdaConfig,omitempty" tf:"lambda_config,omitempty"`
	// +optional
	MfaConfiguration string `json:"mfaConfiguration,omitempty" tf:"mfa_configuration,omitempty"`
	Name             string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PasswordPolicy []CognitoUserPoolSpecPasswordPolicy `json:"passwordPolicy,omitempty" tf:"password_policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Schema []CognitoUserPoolSpecSchema `json:"schema,omitempty" tf:"schema,omitempty"`
	// +optional
	SmsAuthenticationMessage string `json:"smsAuthenticationMessage,omitempty" tf:"sms_authentication_message,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SmsConfiguration []CognitoUserPoolSpecSmsConfiguration `json:"smsConfiguration,omitempty" tf:"sms_configuration,omitempty"`
	// +optional
	SmsVerificationMessage string `json:"smsVerificationMessage,omitempty" tf:"sms_verification_message,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	UserPoolAddOns []CognitoUserPoolSpecUserPoolAddOns `json:"userPoolAddOns,omitempty" tf:"user_pool_add_ons,omitempty"`
	// +optional
	UsernameAttributes []string `json:"usernameAttributes,omitempty" tf:"username_attributes,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VerificationMessageTemplate []CognitoUserPoolSpecVerificationMessageTemplate `json:"verificationMessageTemplate,omitempty" tf:"verification_message_template,omitempty"`
}

type CognitoUserPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
