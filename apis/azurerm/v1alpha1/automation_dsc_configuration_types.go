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

type AutomationDscConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationDscConfigurationSpec   `json:"spec,omitempty"`
	Status            AutomationDscConfigurationStatus `json:"status,omitempty"`
}

type AutomationDscConfigurationSpec struct {
	AutomationAccountName string `json:"automation_account_name"`
	ContentEmbedded       string `json:"content_embedded"`
	// +optional
	Description string `json:"description,omitempty"`
	Location    string `json:"location"`
	// +optional
	LogVerbose        bool   `json:"log_verbose,omitempty"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
}

type AutomationDscConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutomationDscConfigurationList is a list of AutomationDscConfigurations
type AutomationDscConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutomationDscConfiguration CRD objects
	Items []AutomationDscConfiguration `json:"items,omitempty"`
}
