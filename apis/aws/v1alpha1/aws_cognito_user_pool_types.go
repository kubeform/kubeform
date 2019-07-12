package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCognitoUserPoolSpec   `json:"spec,omitempty"`
	Status            AwsCognitoUserPoolStatus `json:"status,omitempty"`
}

type AwsCognitoUserPoolSpecEmailConfiguration struct {
	ReplyToEmailAddress string `json:"reply_to_email_address"`
	SourceArn           string `json:"source_arn"`
}

type AwsCognitoUserPoolSpecPasswordPolicy struct {
	RequireUppercase bool `json:"require_uppercase"`
	MinimumLength    int  `json:"minimum_length"`
	RequireLowercase bool `json:"require_lowercase"`
	RequireNumbers   bool `json:"require_numbers"`
	RequireSymbols   bool `json:"require_symbols"`
}

type AwsCognitoUserPoolSpecSchemaNumberAttributeConstraints struct {
	MinValue string `json:"min_value"`
	MaxValue string `json:"max_value"`
}

type AwsCognitoUserPoolSpecSchemaStringAttributeConstraints struct {
	MinLength string `json:"min_length"`
	MaxLength string `json:"max_length"`
}

type AwsCognitoUserPoolSpecSchema struct {
	NumberAttributeConstraints []AwsCognitoUserPoolSpecSchema `json:"number_attribute_constraints"`
	Required                   bool                           `json:"required"`
	StringAttributeConstraints []AwsCognitoUserPoolSpecSchema `json:"string_attribute_constraints"`
	AttributeDataType          string                         `json:"attribute_data_type"`
	DeveloperOnlyAttribute     bool                           `json:"developer_only_attribute"`
	Mutable                    bool                           `json:"mutable"`
	Name                       string                         `json:"name"`
}

type AwsCognitoUserPoolSpecSmsConfiguration struct {
	ExternalId   string `json:"external_id"`
	SnsCallerArn string `json:"sns_caller_arn"`
}

type AwsCognitoUserPoolSpecAdminCreateUserConfigInviteMessageTemplate struct {
	EmailMessage string `json:"email_message"`
	EmailSubject string `json:"email_subject"`
	SmsMessage   string `json:"sms_message"`
}

type AwsCognitoUserPoolSpecAdminCreateUserConfig struct {
	UnusedAccountValidityDays int                                           `json:"unused_account_validity_days"`
	AllowAdminCreateUserOnly  bool                                          `json:"allow_admin_create_user_only"`
	InviteMessageTemplate     []AwsCognitoUserPoolSpecAdminCreateUserConfig `json:"invite_message_template"`
}

type AwsCognitoUserPoolSpecUserPoolAddOns struct {
	AdvancedSecurityMode string `json:"advanced_security_mode"`
}

type AwsCognitoUserPoolSpecVerificationMessageTemplate struct {
	EmailMessage       string `json:"email_message"`
	EmailMessageByLink string `json:"email_message_by_link"`
	EmailSubject       string `json:"email_subject"`
	EmailSubjectByLink string `json:"email_subject_by_link"`
	SmsMessage         string `json:"sms_message"`
	DefaultEmailOption string `json:"default_email_option"`
}

type AwsCognitoUserPoolSpecDeviceConfiguration struct {
	DeviceOnlyRememberedOnUserPrompt bool `json:"device_only_remembered_on_user_prompt"`
	ChallengeRequiredOnNewDevice     bool `json:"challenge_required_on_new_device"`
}

type AwsCognitoUserPoolSpecLambdaConfig struct {
	PreSignUp                   string `json:"pre_sign_up"`
	CreateAuthChallenge         string `json:"create_auth_challenge"`
	PostConfirmation            string `json:"post_confirmation"`
	PreAuthentication           string `json:"pre_authentication"`
	PreTokenGeneration          string `json:"pre_token_generation"`
	UserMigration               string `json:"user_migration"`
	VerifyAuthChallengeResponse string `json:"verify_auth_challenge_response"`
	CustomMessage               string `json:"custom_message"`
	DefineAuthChallenge         string `json:"define_auth_challenge"`
	PostAuthentication          string `json:"post_authentication"`
}

type AwsCognitoUserPoolSpec struct {
	EmailConfiguration          []AwsCognitoUserPoolSpec `json:"email_configuration"`
	EmailVerificationSubject    string                   `json:"email_verification_subject"`
	MfaConfiguration            string                   `json:"mfa_configuration"`
	PasswordPolicy              []AwsCognitoUserPoolSpec `json:"password_policy"`
	Schema                      []AwsCognitoUserPoolSpec `json:"schema"`
	SmsConfiguration            []AwsCognitoUserPoolSpec `json:"sms_configuration"`
	AdminCreateUserConfig       []AwsCognitoUserPoolSpec `json:"admin_create_user_config"`
	Arn                         string                   `json:"arn"`
	LastModifiedDate            string                   `json:"last_modified_date"`
	UserPoolAddOns              []AwsCognitoUserPoolSpec `json:"user_pool_add_ons"`
	VerificationMessageTemplate []AwsCognitoUserPoolSpec `json:"verification_message_template"`
	AutoVerifiedAttributes      []string                 `json:"auto_verified_attributes"`
	CreationDate                string                   `json:"creation_date"`
	EmailVerificationMessage    string                   `json:"email_verification_message"`
	Endpoint                    string                   `json:"endpoint"`
	SmsAuthenticationMessage    string                   `json:"sms_authentication_message"`
	SmsVerificationMessage      string                   `json:"sms_verification_message"`
	Tags                        map[string]string        `json:"tags"`
	AliasAttributes             []string                 `json:"alias_attributes"`
	DeviceConfiguration         []AwsCognitoUserPoolSpec `json:"device_configuration"`
	LambdaConfig                []AwsCognitoUserPoolSpec `json:"lambda_config"`
	Name                        string                   `json:"name"`
	UsernameAttributes          []string                 `json:"username_attributes"`
}



type AwsCognitoUserPoolStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserPoolList is a list of AwsCognitoUserPools
type AwsCognitoUserPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCognitoUserPool CRD objects
	Items []AwsCognitoUserPool `json:"items,omitempty"`
}