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

type AzurermEventhubAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermEventhubAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            AzurermEventhubAuthorizationRuleStatus `json:"status,omitempty"`
}

type AzurermEventhubAuthorizationRuleSpec struct {
	EventhubName              string `json:"eventhub_name"`
	Name                      string `json:"name"`
	SecondaryKey              string `json:"secondary_key"`
	SecondaryConnectionString string `json:"secondary_connection_string"`
	PrimaryKey                string `json:"primary_key"`
	NamespaceName             string `json:"namespace_name"`
	Location                  string `json:"location"`
	Listen                    bool   `json:"listen"`
	Send                      bool   `json:"send"`
	Manage                    bool   `json:"manage"`
	PrimaryConnectionString   string `json:"primary_connection_string"`
	ResourceGroupName         string `json:"resource_group_name"`
}

type AzurermEventhubAuthorizationRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermEventhubAuthorizationRuleList is a list of AzurermEventhubAuthorizationRules
type AzurermEventhubAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermEventhubAuthorizationRule CRD objects
	Items []AzurermEventhubAuthorizationRule `json:"items,omitempty"`
}
