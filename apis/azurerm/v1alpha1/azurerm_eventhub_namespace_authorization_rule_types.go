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

type AzurermEventhubNamespaceAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermEventhubNamespaceAuthorizationRuleSpec   `json:"spec,omitempty"`
	Status            AzurermEventhubNamespaceAuthorizationRuleStatus `json:"status,omitempty"`
}

type AzurermEventhubNamespaceAuthorizationRuleSpec struct {
	Send                      bool   `json:"send"`
	ResourceGroupName         string `json:"resource_group_name"`
	Manage                    bool   `json:"manage"`
	PrimaryConnectionString   string `json:"primary_connection_string"`
	SecondaryKey              string `json:"secondary_key"`
	Listen                    bool   `json:"listen"`
	Name                      string `json:"name"`
	NamespaceName             string `json:"namespace_name"`
	Location                  string `json:"location"`
	PrimaryKey                string `json:"primary_key"`
	SecondaryConnectionString string `json:"secondary_connection_string"`
}



type AzurermEventhubNamespaceAuthorizationRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermEventhubNamespaceAuthorizationRuleList is a list of AzurermEventhubNamespaceAuthorizationRules
type AzurermEventhubNamespaceAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermEventhubNamespaceAuthorizationRule CRD objects
	Items []AzurermEventhubNamespaceAuthorizationRule `json:"items,omitempty"`
}