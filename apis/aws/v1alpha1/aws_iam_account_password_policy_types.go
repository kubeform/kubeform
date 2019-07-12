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

type AwsIamAccountPasswordPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamAccountPasswordPolicySpec   `json:"spec,omitempty"`
	Status            AwsIamAccountPasswordPolicyStatus `json:"status,omitempty"`
}

type AwsIamAccountPasswordPolicySpec struct {
	MinimumPasswordLength      int  `json:"minimum_password_length"`
	PasswordReusePrevention    int  `json:"password_reuse_prevention"`
	RequireLowercaseCharacters bool `json:"require_lowercase_characters"`
	MaxPasswordAge             int  `json:"max_password_age"`
	RequireNumbers             bool `json:"require_numbers"`
	RequireSymbols             bool `json:"require_symbols"`
	RequireUppercaseCharacters bool `json:"require_uppercase_characters"`
	AllowUsersToChangePassword bool `json:"allow_users_to_change_password"`
	ExpirePasswords            bool `json:"expire_passwords"`
	HardExpiry                 bool `json:"hard_expiry"`
}

type AwsIamAccountPasswordPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIamAccountPasswordPolicyList is a list of AwsIamAccountPasswordPolicys
type AwsIamAccountPasswordPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamAccountPasswordPolicy CRD objects
	Items []AwsIamAccountPasswordPolicy `json:"items,omitempty"`
}
