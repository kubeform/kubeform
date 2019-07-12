package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermEventhubNamespaceAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermEventhubNamespaceAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            AzurermEventhubNamespaceAuthorizationRuleStatus `json:"status,omitempty"`
}

type AzurermEventhubNamespaceAuthorizationRuleSpec struct {
	Name                      string `json:"name"`
	Send                      bool   `json:"send"`
	PrimaryKey                string `json:"primary_key"`
	SecondaryKey              string `json:"secondary_key"`
	SecondaryConnectionString string `json:"secondary_connection_string"`
	NamespaceName             string `json:"namespace_name"`
	ResourceGroupName         string `json:"resource_group_name"`
	Location                  string `json:"location"`
	Listen                    bool   `json:"listen"`
	Manage                    bool   `json:"manage"`
	PrimaryConnectionString   string `json:"primary_connection_string"`
}

type AzurermEventhubNamespaceAuthorizationRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermEventhubNamespaceAuthorizationRuleList is a list of AzurermEventhubNamespaceAuthorizationRules
type AzurermEventhubNamespaceAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermEventhubNamespaceAuthorizationRule CRD objects
	Items []AzurermEventhubNamespaceAuthorizationRule `json:"items,omitempty"`
}
