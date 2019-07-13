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

type AwsDbOptionGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDbOptionGroupSpec   `json:"spec,omitempty"`
	Status            AwsDbOptionGroupStatus `json:"status,omitempty"`
}

type AwsDbOptionGroupSpecOptionOptionSettings struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsDbOptionGroupSpecOption struct {
	OptionName                  string                       `json:"option_name"`
	OptionSettings              []AwsDbOptionGroupSpecOption `json:"option_settings"`
	Port                        int                          `json:"port"`
	DbSecurityGroupMemberships  []string                     `json:"db_security_group_memberships"`
	VpcSecurityGroupMemberships []string                     `json:"vpc_security_group_memberships"`
	Version                     string                       `json:"version"`
}

type AwsDbOptionGroupSpec struct {
	Option                 []AwsDbOptionGroupSpec `json:"option"`
	Tags                   map[string]string      `json:"tags"`
	Arn                    string                 `json:"arn"`
	Name                   string                 `json:"name"`
	NamePrefix             string                 `json:"name_prefix"`
	EngineName             string                 `json:"engine_name"`
	MajorEngineVersion     string                 `json:"major_engine_version"`
	OptionGroupDescription string                 `json:"option_group_description"`
}



type AwsDbOptionGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDbOptionGroupList is a list of AwsDbOptionGroups
type AwsDbOptionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDbOptionGroup CRD objects
	Items []AwsDbOptionGroup `json:"items,omitempty"`
}