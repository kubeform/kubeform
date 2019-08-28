package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MqBroker struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MqBrokerSpec   `json:"spec,omitempty"`
	Status            MqBrokerStatus `json:"status,omitempty"`
}

type MqBrokerSpecConfiguration struct {
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	Revision int `json:"revision,omitempty" tf:"revision,omitempty"`
}

type MqBrokerSpecInstances struct {
	// +optional
	ConsoleURL string `json:"consoleURL,omitempty" tf:"console_url,omitempty"`
	// +optional
	Endpoints []string `json:"endpoints,omitempty" tf:"endpoints,omitempty"`
	// +optional
	IpAddress string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
}

type MqBrokerSpecLogs struct {
	// +optional
	Audit bool `json:"audit,omitempty" tf:"audit,omitempty"`
	// +optional
	General bool `json:"general,omitempty" tf:"general,omitempty"`
}

type MqBrokerSpecMaintenanceWindowStartTime struct {
	DayOfWeek string `json:"dayOfWeek" tf:"day_of_week"`
	TimeOfDay string `json:"timeOfDay" tf:"time_of_day"`
	TimeZone  string `json:"timeZone" tf:"time_zone"`
}

type MqBrokerSpecUser struct {
	// +optional
	ConsoleAccess bool `json:"consoleAccess,omitempty" tf:"console_access,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Groups   []string `json:"groups,omitempty" tf:"groups,omitempty"`
	Password string   `json:"-" sensitive:"true" tf:"password"`
	Username string   `json:"username" tf:"username"`
}

type MqBrokerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	ApplyImmediately bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	AutoMinorVersionUpgrade bool   `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`
	BrokerName              string `json:"brokerName" tf:"broker_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Configuration []MqBrokerSpecConfiguration `json:"configuration,omitempty" tf:"configuration,omitempty"`
	// +optional
	DeploymentMode   string `json:"deploymentMode,omitempty" tf:"deployment_mode,omitempty"`
	EngineType       string `json:"engineType" tf:"engine_type"`
	EngineVersion    string `json:"engineVersion" tf:"engine_version"`
	HostInstanceType string `json:"hostInstanceType" tf:"host_instance_type"`
	// +optional
	Instances []MqBrokerSpecInstances `json:"instances,omitempty" tf:"instances,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Logs []MqBrokerSpecLogs `json:"logs,omitempty" tf:"logs,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	MaintenanceWindowStartTime []MqBrokerSpecMaintenanceWindowStartTime `json:"maintenanceWindowStartTime,omitempty" tf:"maintenance_window_start_time,omitempty"`
	// +optional
	PubliclyAccessible bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"securityGroups" tf:"security_groups"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS,omitempty" tf:"subnet_ids,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	User []MqBrokerSpecUser `json:"user" tf:"user"`
}



type MqBrokerStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *MqBrokerSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
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