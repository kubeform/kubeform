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

type GoogleServiceAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleServiceAccountSpec   `json:"spec,omitempty"`
	Status            GoogleServiceAccountStatus `json:"status,omitempty"`
}

type GoogleServiceAccountSpec struct {
	DisplayName string `json:"display_name"`
	Project     string `json:"project"`
	PolicyData  string `json:"policy_data"`
	Email       string `json:"email"`
	UniqueId    string `json:"unique_id"`
	Name        string `json:"name"`
	AccountId   string `json:"account_id"`
}



type GoogleServiceAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleServiceAccountList is a list of GoogleServiceAccounts
type GoogleServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleServiceAccount CRD objects
	Items []GoogleServiceAccount `json:"items,omitempty"`
}