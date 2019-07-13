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

type AwsOpsworksStack struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksStackSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksStackStatus `json:"status,omitempty"`
}

type AwsOpsworksStackSpecCustomCookbooksSource struct {
	Password string `json:"password"`
	Revision string `json:"revision"`
	SshKey   string `json:"ssh_key"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	Username string `json:"username"`
}

type AwsOpsworksStackSpec struct {
	ConfigurationManagerVersion string                 `json:"configuration_manager_version"`
	BerkshelfVersion            string                 `json:"berkshelf_version"`
	DefaultAvailabilityZone     string                 `json:"default_availability_zone"`
	Tags                        map[string]string      `json:"tags"`
	UseOpsworksSecurityGroups   bool                   `json:"use_opsworks_security_groups"`
	Region                      string                 `json:"region"`
	ServiceRoleArn              string                 `json:"service_role_arn"`
	Color                       string                 `json:"color"`
	VpcId                       string                 `json:"vpc_id"`
	CustomJson                  string                 `json:"custom_json"`
	DefaultOs                   string                 `json:"default_os"`
	DefaultSshKeyName           string                 `json:"default_ssh_key_name"`
	DefaultSubnetId             string                 `json:"default_subnet_id"`
	StackEndpoint               string                 `json:"stack_endpoint"`
	AgentVersion                string                 `json:"agent_version"`
	ConfigurationManagerName    string                 `json:"configuration_manager_name"`
	ManageBerkshelf             bool                   `json:"manage_berkshelf"`
	Name                        string                 `json:"name"`
	DefaultInstanceProfileArn   string                 `json:"default_instance_profile_arn"`
	HostnameTheme               string                 `json:"hostname_theme"`
	UseCustomCookbooks          bool                   `json:"use_custom_cookbooks"`
	Arn                         string                 `json:"arn"`
	CustomCookbooksSource       []AwsOpsworksStackSpec `json:"custom_cookbooks_source"`
	DefaultRootDeviceType       string                 `json:"default_root_device_type"`
}



type AwsOpsworksStackStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsOpsworksStackList is a list of AwsOpsworksStacks
type AwsOpsworksStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksStack CRD objects
	Items []AwsOpsworksStack `json:"items,omitempty"`
}