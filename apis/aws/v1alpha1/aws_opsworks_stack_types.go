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
	Type     string `json:"type"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Revision string `json:"revision"`
	SshKey   string `json:"ssh_key"`
}

type AwsOpsworksStackSpec struct {
	Color                       string                 `json:"color"`
	ConfigurationManagerVersion string                 `json:"configuration_manager_version"`
	BerkshelfVersion            string                 `json:"berkshelf_version"`
	DefaultOs                   string                 `json:"default_os"`
	DefaultRootDeviceType       string                 `json:"default_root_device_type"`
	Arn                         string                 `json:"arn"`
	Name                        string                 `json:"name"`
	Region                      string                 `json:"region"`
	VpcId                       string                 `json:"vpc_id"`
	HostnameTheme               string                 `json:"hostname_theme"`
	UseCustomCookbooks          bool                   `json:"use_custom_cookbooks"`
	UseOpsworksSecurityGroups   bool                   `json:"use_opsworks_security_groups"`
	Tags                        map[string]string      `json:"tags"`
	StackEndpoint               string                 `json:"stack_endpoint"`
	ServiceRoleArn              string                 `json:"service_role_arn"`
	CustomJson                  string                 `json:"custom_json"`
	DefaultSubnetId             string                 `json:"default_subnet_id"`
	ManageBerkshelf             bool                   `json:"manage_berkshelf"`
	DefaultSshKeyName           string                 `json:"default_ssh_key_name"`
	AgentVersion                string                 `json:"agent_version"`
	DefaultInstanceProfileArn   string                 `json:"default_instance_profile_arn"`
	ConfigurationManagerName    string                 `json:"configuration_manager_name"`
	CustomCookbooksSource       []AwsOpsworksStackSpec `json:"custom_cookbooks_source"`
	DefaultAvailabilityZone     string                 `json:"default_availability_zone"`
}

type AwsOpsworksStackStatus struct {
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
