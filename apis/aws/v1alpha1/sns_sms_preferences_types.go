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

type SnsSmsPreferences struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnsSmsPreferencesSpec   `json:"spec,omitempty"`
	Status            SnsSmsPreferencesStatus `json:"status,omitempty"`
}

type SnsSmsPreferencesSpec struct {
	// +optional
	DefaultSenderId string `json:"default_sender_id,omitempty"`
	// +optional
	DefaultSmsType string `json:"default_sms_type,omitempty"`
	// +optional
	DeliveryStatusIamRoleArn string `json:"delivery_status_iam_role_arn,omitempty"`
	// +optional
	DeliveryStatusSuccessSamplingRate string `json:"delivery_status_success_sampling_rate,omitempty"`
	// +optional
	MonthlySpendLimit string `json:"monthly_spend_limit,omitempty"`
	// +optional
	UsageReportS3Bucket string `json:"usage_report_s3_bucket,omitempty"`
}

type SnsSmsPreferencesStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SnsSmsPreferencesList is a list of SnsSmsPreferencess
type SnsSmsPreferencesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SnsSmsPreferences CRD objects
	Items []SnsSmsPreferences `json:"items,omitempty"`
}
