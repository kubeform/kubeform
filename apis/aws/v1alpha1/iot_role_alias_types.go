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

type IotRoleAlias struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotRoleAliasSpec   `json:"spec,omitempty"`
	Status            IotRoleAliasStatus `json:"status,omitempty"`
}

type IotRoleAliasSpec struct {
	Alias string `json:"alias"`
	// +optional
	CredentialDuration int    `json:"credential_duration,omitempty"`
	RoleArn            string `json:"role_arn"`
}

type IotRoleAliasStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IotRoleAliasList is a list of IotRoleAliass
type IotRoleAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IotRoleAlias CRD objects
	Items []IotRoleAlias `json:"items,omitempty"`
}
