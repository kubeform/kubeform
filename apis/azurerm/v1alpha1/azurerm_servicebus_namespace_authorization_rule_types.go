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

type AzurermServicebusNamespaceAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermServicebusNamespaceAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            AzurermServicebusNamespaceAuthorizationRuleStatus `json:"status,omitempty"`
}

type AzurermServicebusNamespaceAuthorizationRuleSpec struct {
	NamespaceName             string `json:"namespace_name"`
	SecondaryConnectionString string `json:"secondary_connection_string"`
	Send                      bool   `json:"send"`
	Manage                    bool   `json:"manage"`
	SecondaryKey              string `json:"secondary_key"`
	Name                      string `json:"name"`
	ResourceGroupName         string `json:"resource_group_name"`
	Listen                    bool   `json:"listen"`
	PrimaryKey                string `json:"primary_key"`
	PrimaryConnectionString   string `json:"primary_connection_string"`
}

type AzurermServicebusNamespaceAuthorizationRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermServicebusNamespaceAuthorizationRuleList is a list of AzurermServicebusNamespaceAuthorizationRules
type AzurermServicebusNamespaceAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermServicebusNamespaceAuthorizationRule CRD objects
	Items []AzurermServicebusNamespaceAuthorizationRule `json:"items,omitempty"`
}
