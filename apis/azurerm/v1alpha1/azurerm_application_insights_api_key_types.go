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

type AzurermApplicationInsightsApiKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApplicationInsightsApiKeySpec   `json:"spec,omitempty"`
	Status            AzurermApplicationInsightsApiKeyStatus `json:"status,omitempty"`
}

type AzurermApplicationInsightsApiKeySpec struct {
	ApiKey                string   `json:"api_key"`
	Name                  string   `json:"name"`
	ApplicationInsightsId string   `json:"application_insights_id"`
	ReadPermissions       []string `json:"read_permissions"`
	WritePermissions      []string `json:"write_permissions"`
}



type AzurermApplicationInsightsApiKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApplicationInsightsApiKeyList is a list of AzurermApplicationInsightsApiKeys
type AzurermApplicationInsightsApiKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApplicationInsightsApiKey CRD objects
	Items []AzurermApplicationInsightsApiKey `json:"items,omitempty"`
}