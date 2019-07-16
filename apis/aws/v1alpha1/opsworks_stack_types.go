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

type OpsworksStack struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksStackSpec   `json:"spec,omitempty"`
	Status            OpsworksStackStatus `json:"status,omitempty"`
}

type OpsworksStackSpec struct {
	// +optional
	BerkshelfVersion string `json:"berkshelf_version,omitempty"`
	// +optional
	Color string `json:"color,omitempty"`
	// +optional
	ConfigurationManagerName string `json:"configuration_manager_name,omitempty"`
	// +optional
	ConfigurationManagerVersion string `json:"configuration_manager_version,omitempty"`
	// +optional
	CustomJson                string `json:"custom_json,omitempty"`
	DefaultInstanceProfileArn string `json:"default_instance_profile_arn"`
	// +optional
	DefaultOs string `json:"default_os,omitempty"`
	// +optional
	DefaultRootDeviceType string `json:"default_root_device_type,omitempty"`
	// +optional
	DefaultSshKeyName string `json:"default_ssh_key_name,omitempty"`
	// +optional
	HostnameTheme string `json:"hostname_theme,omitempty"`
	// +optional
	ManageBerkshelf bool   `json:"manage_berkshelf,omitempty"`
	Name            string `json:"name"`
	Region          string `json:"region"`
	ServiceRoleArn  string `json:"service_role_arn"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	UseCustomCookbooks bool `json:"use_custom_cookbooks,omitempty"`
	// +optional
	UseOpsworksSecurityGroups bool `json:"use_opsworks_security_groups,omitempty"`
}

type OpsworksStackStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksStackList is a list of OpsworksStacks
type OpsworksStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksStack CRD objects
	Items []OpsworksStack `json:"items,omitempty"`
}
