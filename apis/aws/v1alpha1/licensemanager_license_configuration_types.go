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

type LicensemanagerLicenseConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LicensemanagerLicenseConfigurationSpec   `json:"spec,omitempty"`
	Status            LicensemanagerLicenseConfigurationStatus `json:"status,omitempty"`
}

type LicensemanagerLicenseConfigurationSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	LicenseCount int `json:"license_count,omitempty"`
	// +optional
	LicenseCountHardLimit bool   `json:"license_count_hard_limit,omitempty"`
	LicenseCountingType   string `json:"license_counting_type"`
	// +optional
	LicenseRules []string `json:"license_rules,omitempty"`
	Name         string   `json:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type LicensemanagerLicenseConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LicensemanagerLicenseConfigurationList is a list of LicensemanagerLicenseConfigurations
type LicensemanagerLicenseConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LicensemanagerLicenseConfiguration CRD objects
	Items []LicensemanagerLicenseConfiguration `json:"items,omitempty"`
}
