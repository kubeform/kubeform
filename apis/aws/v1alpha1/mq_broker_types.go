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

type MqBroker struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MqBrokerSpec   `json:"spec,omitempty"`
	Status            MqBrokerStatus `json:"status,omitempty"`
}

type MqBrokerSpecLogs struct {
	// +optional
	Audit bool `json:"audit,omitempty"`
	// +optional
	General bool `json:"general,omitempty"`
}

type MqBrokerSpecUser struct {
	// +optional
	ConsoleAccess bool `json:"console_access,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Groups   []string `json:"groups,omitempty"`
	Password string   `json:"password"`
	Username string   `json:"username"`
}

type MqBrokerSpec struct {
	// +optional
	ApplyImmediately bool `json:"apply_immediately,omitempty"`
	// +optional
	AutoMinorVersionUpgrade bool   `json:"auto_minor_version_upgrade,omitempty"`
	BrokerName              string `json:"broker_name"`
	// +optional
	DeploymentMode   string `json:"deployment_mode,omitempty"`
	EngineType       string `json:"engine_type"`
	EngineVersion    string `json:"engine_version"`
	HostInstanceType string `json:"host_instance_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Logs *[]MqBrokerSpec `json:"logs,omitempty"`
	// +optional
	PubliclyAccessible bool `json:"publicly_accessible,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"security_groups"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	User []MqBrokerSpec `json:"user"`
}

type MqBrokerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MqBrokerList is a list of MqBrokers
type MqBrokerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MqBroker CRD objects
	Items []MqBroker `json:"items,omitempty"`
}
