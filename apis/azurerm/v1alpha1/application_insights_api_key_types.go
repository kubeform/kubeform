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

type ApplicationInsightsApiKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationInsightsApiKeySpec   `json:"spec,omitempty"`
	Status            ApplicationInsightsApiKeyStatus `json:"status,omitempty"`
}

type ApplicationInsightsApiKeySpec struct {
	ApplicationInsightsId string `json:"application_insights_id"`
	Name                  string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ReadPermissions []string `json:"read_permissions,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	WritePermissions []string `json:"write_permissions,omitempty"`
}

type ApplicationInsightsApiKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApplicationInsightsApiKeyList is a list of ApplicationInsightsApiKeys
type ApplicationInsightsApiKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApplicationInsightsApiKey CRD objects
	Items []ApplicationInsightsApiKey `json:"items,omitempty"`
}
