package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DbOptionGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbOptionGroupSpec   `json:"spec,omitempty"`
	Status            DbOptionGroupStatus `json:"status,omitempty"`
}

type DbOptionGroupSpecOptionOptionSettings struct {
	Name  string `json:"name" tf:"name"`
	Value string `json:"value" tf:"value"`
}

type DbOptionGroupSpecOption struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DbSecurityGroupMemberships []string `json:"dbSecurityGroupMemberships,omitempty" tf:"db_security_group_memberships,omitempty"`
	OptionName                 string   `json:"optionName" tf:"option_name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OptionSettings []DbOptionGroupSpecOptionOptionSettings `json:"optionSettings,omitempty" tf:"option_settings,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Version string `json:"version,omitempty" tf:"version,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VpcSecurityGroupMemberships []string `json:"vpcSecurityGroupMemberships,omitempty" tf:"vpc_security_group_memberships,omitempty"`
}

type DbOptionGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn                string `json:"arn,omitempty" tf:"arn,omitempty"`
	EngineName         string `json:"engineName" tf:"engine_name"`
	MajorEngineVersion string `json:"majorEngineVersion" tf:"major_engine_version"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Option []DbOptionGroupSpecOption `json:"option,omitempty" tf:"option,omitempty"`
	// +optional
	OptionGroupDescription string `json:"optionGroupDescription,omitempty" tf:"option_group_description,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DbOptionGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DbOptionGroupSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DbOptionGroupList is a list of DbOptionGroups
type DbOptionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DbOptionGroup CRD objects
	Items []DbOptionGroup `json:"items,omitempty"`
}
