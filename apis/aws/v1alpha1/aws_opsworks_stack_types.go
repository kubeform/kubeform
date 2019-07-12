package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsOpsworksStack struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsOpsworksStackSpec   `json:"spec,omitempty"`
	Status            AwsOpsworksStackStatus `json:"status,omitempty"`
}

type AwsOpsworksStackSpecCustomCookbooksSource struct {
	Revision string `json:"revision"`
	SshKey   string `json:"ssh_key"`
	Type     string `json:"type"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AwsOpsworksStackSpec struct {
	CustomJson                  string                 `json:"custom_json"`
	StackEndpoint               string                 `json:"stack_endpoint"`
	DefaultAvailabilityZone     string                 `json:"default_availability_zone"`
	Tags                        map[string]string      `json:"tags"`
	UseOpsworksSecurityGroups   bool                   `json:"use_opsworks_security_groups"`
	AgentVersion                string                 `json:"agent_version"`
	Arn                         string                 `json:"arn"`
	Name                        string                 `json:"name"`
	DefaultInstanceProfileArn   string                 `json:"default_instance_profile_arn"`
	ConfigurationManagerVersion string                 `json:"configuration_manager_version"`
	DefaultRootDeviceType       string                 `json:"default_root_device_type"`
	DefaultSubnetId             string                 `json:"default_subnet_id"`
	UseCustomCookbooks          bool                   `json:"use_custom_cookbooks"`
	Region                      string                 `json:"region"`
	ServiceRoleArn              string                 `json:"service_role_arn"`
	Color                       string                 `json:"color"`
	BerkshelfVersion            string                 `json:"berkshelf_version"`
	DefaultOs                   string                 `json:"default_os"`
	VpcId                       string                 `json:"vpc_id"`
	ConfigurationManagerName    string                 `json:"configuration_manager_name"`
	ManageBerkshelf             bool                   `json:"manage_berkshelf"`
	CustomCookbooksSource       []AwsOpsworksStackSpec `json:"custom_cookbooks_source"`
	DefaultSshKeyName           string                 `json:"default_ssh_key_name"`
	HostnameTheme               string                 `json:"hostname_theme"`
}



type AwsOpsworksStackStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksStackList is a list of AwsOpsworksStacks
type AwsOpsworksStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsOpsworksStack CRD objects
	Items []AwsOpsworksStack `json:"items,omitempty"`
}