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

type AzurermAvailabilitySet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAvailabilitySetSpec   `json:"spec,omitempty"`
	Status            AzurermAvailabilitySetStatus `json:"status,omitempty"`
}

type AzurermAvailabilitySetSpec struct {
	Name                      string            `json:"name"`
	ResourceGroupName         string            `json:"resource_group_name"`
	Location                  string            `json:"location"`
	PlatformUpdateDomainCount int               `json:"platform_update_domain_count"`
	PlatformFaultDomainCount  int               `json:"platform_fault_domain_count"`
	Managed                   bool              `json:"managed"`
	Tags                      map[string]string `json:"tags"`
}

type AzurermAvailabilitySetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAvailabilitySetList is a list of AzurermAvailabilitySets
type AzurermAvailabilitySetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAvailabilitySet CRD objects
	Items []AzurermAvailabilitySet `json:"items,omitempty"`
}
