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

type AwsSnsSmsPreferences struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSnsSmsPreferencesSpec   `json:"spec,omitempty"`
	Status            AwsSnsSmsPreferencesStatus `json:"status,omitempty"`
}

type AwsSnsSmsPreferencesSpec struct {
	MonthlySpendLimit                 string `json:"monthly_spend_limit"`
	DeliveryStatusIamRoleArn          string `json:"delivery_status_iam_role_arn"`
	DeliveryStatusSuccessSamplingRate string `json:"delivery_status_success_sampling_rate"`
	DefaultSenderId                   string `json:"default_sender_id"`
	DefaultSmsType                    string `json:"default_sms_type"`
	UsageReportS3Bucket               string `json:"usage_report_s3_bucket"`
}

type AwsSnsSmsPreferencesStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSnsSmsPreferencesList is a list of AwsSnsSmsPreferencess
type AwsSnsSmsPreferencesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSnsSmsPreferences CRD objects
	Items []AwsSnsSmsPreferences `json:"items,omitempty"`
}
