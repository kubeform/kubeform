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

type OrganizationsAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationsAccountSpec   `json:"spec,omitempty"`
	Status            OrganizationsAccountStatus `json:"status,omitempty"`
}

type OrganizationsAccountSpec struct {
	Email string `json:"email"`
	// +optional
	IamUserAccessToBilling string `json:"iam_user_access_to_billing,omitempty"`
	Name                   string `json:"name"`
	// +optional
	RoleName string `json:"role_name,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type OrganizationsAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OrganizationsAccountList is a list of OrganizationsAccounts
type OrganizationsAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OrganizationsAccount CRD objects
	Items []OrganizationsAccount `json:"items,omitempty"`
}
