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

type AwsAthenaWorkgroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAthenaWorkgroupSpec   `json:"spec,omitempty"`
	Status            AwsAthenaWorkgroupStatus `json:"status,omitempty"`
}

type AwsAthenaWorkgroupSpecConfigurationResultConfigurationEncryptionConfiguration struct {
	EncryptionOption string `json:"encryption_option"`
	KmsKeyArn        string `json:"kms_key_arn"`
}

type AwsAthenaWorkgroupSpecConfigurationResultConfiguration struct {
	EncryptionConfiguration []AwsAthenaWorkgroupSpecConfigurationResultConfiguration `json:"encryption_configuration"`
	OutputLocation          string                                                   `json:"output_location"`
}

type AwsAthenaWorkgroupSpecConfiguration struct {
	BytesScannedCutoffPerQuery      int                                   `json:"bytes_scanned_cutoff_per_query"`
	EnforceWorkgroupConfiguration   bool                                  `json:"enforce_workgroup_configuration"`
	PublishCloudwatchMetricsEnabled bool                                  `json:"publish_cloudwatch_metrics_enabled"`
	ResultConfiguration             []AwsAthenaWorkgroupSpecConfiguration `json:"result_configuration"`
}

type AwsAthenaWorkgroupSpec struct {
	Arn           string                   `json:"arn"`
	Configuration []AwsAthenaWorkgroupSpec `json:"configuration"`
	Description   string                   `json:"description"`
	Name          string                   `json:"name"`
	State         string                   `json:"state"`
	Tags          map[string]string        `json:"tags"`
}



type AwsAthenaWorkgroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAthenaWorkgroupList is a list of AwsAthenaWorkgroups
type AwsAthenaWorkgroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAthenaWorkgroup CRD objects
	Items []AwsAthenaWorkgroup `json:"items,omitempty"`
}