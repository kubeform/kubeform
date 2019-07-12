package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApplicationSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApplicationSecurityGroupSpec   `json:"spec,omitempty"`
	Status            AzurermApplicationSecurityGroupStatus `json:"status,omitempty"`
}

type AzurermApplicationSecurityGroupSpec struct {
	ResourceGroupName string            `json:"resource_group_name"`
	Tags              map[string]string `json:"tags"`
	Name              string            `json:"name"`
	Location          string            `json:"location"`
}



type AzurermApplicationSecurityGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApplicationSecurityGroupList is a list of AzurermApplicationSecurityGroups
type AzurermApplicationSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApplicationSecurityGroup CRD objects
	Items []AzurermApplicationSecurityGroup `json:"items,omitempty"`
}