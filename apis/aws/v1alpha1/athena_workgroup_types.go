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

type AthenaWorkgroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AthenaWorkgroupSpec   `json:"spec,omitempty"`
	Status            AthenaWorkgroupStatus `json:"status,omitempty"`
}

type AthenaWorkgroupSpecConfigurationResultConfigurationEncryptionConfiguration struct {
	// +optional
	EncryptionOption string `json:"encryptionOption,omitempty" tf:"encryption_option,omitempty"`
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
}

type AthenaWorkgroupSpecConfigurationResultConfiguration struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionConfiguration []AthenaWorkgroupSpecConfigurationResultConfigurationEncryptionConfiguration `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`
	// +optional
	OutputLocation string `json:"outputLocation,omitempty" tf:"output_location,omitempty"`
}

type AthenaWorkgroupSpecConfiguration struct {
	// +optional
	BytesScannedCutoffPerQuery int `json:"bytesScannedCutoffPerQuery,omitempty" tf:"bytes_scanned_cutoff_per_query,omitempty"`
	// +optional
	EnforceWorkgroupConfiguration bool `json:"enforceWorkgroupConfiguration,omitempty" tf:"enforce_workgroup_configuration,omitempty"`
	// +optional
	PublishCloudwatchMetricsEnabled bool `json:"publishCloudwatchMetricsEnabled,omitempty" tf:"publish_cloudwatch_metrics_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ResultConfiguration []AthenaWorkgroupSpecConfigurationResultConfiguration `json:"resultConfiguration,omitempty" tf:"result_configuration,omitempty"`
}

type AthenaWorkgroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Arn string `json:"arn" tf:"arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Configuration []AthenaWorkgroupSpecConfiguration `json:"configuration,omitempty" tf:"configuration,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AthenaWorkgroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AthenaWorkgroupSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
