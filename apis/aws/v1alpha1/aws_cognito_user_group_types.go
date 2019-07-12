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

type AwsCognitoUserGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCognitoUserGroupSpec   `json:"spec,omitempty"`
	Status            AwsCognitoUserGroupStatus `json:"status,omitempty"`
}

type AwsCognitoUserGroupSpec struct {
	Precedence  int    `json:"precedence"`
	RoleArn     string `json:"role_arn"`
	UserPoolId  string `json:"user_pool_id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type AwsCognitoUserGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCognitoUserGroupList is a list of AwsCognitoUserGroups
type AwsCognitoUserGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCognitoUserGroup CRD objects
	Items []AwsCognitoUserGroup `json:"items,omitempty"`
}
