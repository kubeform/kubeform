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

type IamAccountPasswordPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamAccountPasswordPolicySpec   `json:"spec,omitempty"`
	Status            IamAccountPasswordPolicyStatus `json:"status,omitempty"`
}

type IamAccountPasswordPolicySpec struct {
	// +optional
	AllowUsersToChangePassword bool `json:"allow_users_to_change_password,omitempty"`
	// +optional
	MinimumPasswordLength int `json:"minimum_password_length,omitempty"`
}

type IamAccountPasswordPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamAccountPasswordPolicyList is a list of IamAccountPasswordPolicys
type IamAccountPasswordPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamAccountPasswordPolicy CRD objects
	Items []IamAccountPasswordPolicy `json:"items,omitempty"`
}
