package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type LicensemanagerLicenseConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LicensemanagerLicenseConfigurationSpec   `json:"spec,omitempty"`
	Status            LicensemanagerLicenseConfigurationStatus `json:"status,omitempty"`
}

type LicensemanagerLicenseConfigurationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	LicenseCount int `json:"licenseCount,omitempty" tf:"license_count,omitempty"`
	// +optional
	LicenseCountHardLimit bool   `json:"licenseCountHardLimit,omitempty" tf:"license_count_hard_limit,omitempty"`
	LicenseCountingType   string `json:"licenseCountingType" tf:"license_counting_type"`
	// +optional
	LicenseRules []string `json:"licenseRules,omitempty" tf:"license_rules,omitempty"`
	Name         string   `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LicensemanagerLicenseConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LicensemanagerLicenseConfigurationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
