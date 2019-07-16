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

type FunctionApp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionAppSpec   `json:"spec,omitempty"`
	Status            FunctionAppStatus `json:"status,omitempty"`
}

type FunctionAppSpec struct {
	AppServicePlanId string `json:"app_service_plan_id"`
	// +optional
	AppSettings map[string]string `json:"app_settings,omitempty"`
	// +optional
	EnableBuiltinLogging bool `json:"enable_builtin_logging,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	HttpsOnly               bool   `json:"https_only,omitempty"`
	Location                string `json:"location"`
	Name                    string `json:"name"`
	ResourceGroupName       string `json:"resource_group_name"`
	StorageConnectionString string `json:"storage_connection_string"`
	// +optional
	Version string `json:"version,omitempty"`
}

type FunctionAppStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FunctionAppList is a list of FunctionApps
type FunctionAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FunctionApp CRD objects
	Items []FunctionApp `json:"items,omitempty"`
}
