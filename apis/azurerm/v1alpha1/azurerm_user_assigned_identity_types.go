package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermUserAssignedIdentity struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermUserAssignedIdentitySpec   `json:"spec,omitempty"`
	Status            AzurermUserAssignedIdentityStatus `json:"status,omitempty"`
}

type AzurermUserAssignedIdentitySpec struct {
	PrincipalId       string            `json:"principal_id"`
	ClientId          string            `json:"client_id"`
	Name              string            `json:"name"`
	ResourceGroupName string            `json:"resource_group_name"`
	Location          string            `json:"location"`
	Tags              map[string]string `json:"tags"`
}

type AzurermUserAssignedIdentityStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermUserAssignedIdentityList is a list of AzurermUserAssignedIdentitys
type AzurermUserAssignedIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermUserAssignedIdentity CRD objects
	Items []AzurermUserAssignedIdentity `json:"items,omitempty"`
}
