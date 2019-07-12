package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLicensemanagerLicenseConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLicensemanagerLicenseConfigurationSpec   `json:"spec,omitempty"`
	Status            AwsLicensemanagerLicenseConfigurationStatus `json:"status,omitempty"`
}

type AwsLicensemanagerLicenseConfigurationSpec struct {
	LicenseCount          int               `json:"license_count"`
	LicenseCountHardLimit bool              `json:"license_count_hard_limit"`
	LicenseCountingType   string            `json:"license_counting_type"`
	LicenseRules          []string          `json:"license_rules"`
	Name                  string            `json:"name"`
	Tags                  map[string]string `json:"tags"`
	Description           string            `json:"description"`
}

type AwsLicensemanagerLicenseConfigurationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLicensemanagerLicenseConfigurationList is a list of AwsLicensemanagerLicenseConfigurations
type AwsLicensemanagerLicenseConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLicensemanagerLicenseConfiguration CRD objects
	Items []AwsLicensemanagerLicenseConfiguration `json:"items,omitempty"`
}
