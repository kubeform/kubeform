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

type AthenaWorkgroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AthenaWorkgroupSpec   `json:"spec,omitempty"`
	Status            AthenaWorkgroupStatus `json:"status,omitempty"`
}

type AthenaWorkgroupSpecConfigurationResultConfigurationEncryptionConfiguration struct {
	// +optional
	EncryptionOption string `json:"encryption_option,omitempty"`
	// +optional
	KmsKeyArn string `json:"kms_key_arn,omitempty"`
}

type AthenaWorkgroupSpecConfigurationResultConfiguration struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionConfiguration *[]AthenaWorkgroupSpecConfigurationResultConfiguration `json:"encryption_configuration,omitempty"`
	// +optional
	OutputLocation string `json:"output_location,omitempty"`
}

type AthenaWorkgroupSpecConfiguration struct {
	// +optional
	BytesScannedCutoffPerQuery int `json:"bytes_scanned_cutoff_per_query,omitempty"`
	// +optional
	EnforceWorkgroupConfiguration bool `json:"enforce_workgroup_configuration,omitempty"`
	// +optional
	PublishCloudwatchMetricsEnabled bool `json:"publish_cloudwatch_metrics_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ResultConfiguration *[]AthenaWorkgroupSpecConfiguration `json:"result_configuration,omitempty"`
}

type AthenaWorkgroupSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Configuration *[]AthenaWorkgroupSpec `json:"configuration,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +optional
	State string `json:"state,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type AthenaWorkgroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AthenaWorkgroupList is a list of AthenaWorkgroups
type AthenaWorkgroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AthenaWorkgroup CRD objects
	Items []AthenaWorkgroup `json:"items,omitempty"`
}
