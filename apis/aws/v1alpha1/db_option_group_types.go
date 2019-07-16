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

type DbOptionGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbOptionGroupSpec   `json:"spec,omitempty"`
	Status            DbOptionGroupStatus `json:"status,omitempty"`
}

type DbOptionGroupSpecOptionOptionSettings struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DbOptionGroupSpecOption struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	DbSecurityGroupMemberships []string `json:"db_security_group_memberships,omitempty"`
	OptionName                 string   `json:"option_name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OptionSettings *[]DbOptionGroupSpecOption `json:"option_settings,omitempty"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	Version string `json:"version,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VpcSecurityGroupMemberships []string `json:"vpc_security_group_memberships,omitempty"`
}

type DbOptionGroupSpec struct {
	EngineName         string `json:"engine_name"`
	MajorEngineVersion string `json:"major_engine_version"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Option *[]DbOptionGroupSpec `json:"option,omitempty"`
	// +optional
	OptionGroupDescription string `json:"option_group_description,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DbOptionGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
