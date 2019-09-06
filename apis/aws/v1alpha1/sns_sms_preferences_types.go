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

type SnsSmsPreferences struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnsSmsPreferencesSpec   `json:"spec,omitempty"`
	Status            SnsSmsPreferencesStatus `json:"status,omitempty"`
}

type SnsSmsPreferencesSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DefaultSenderID string `json:"defaultSenderID,omitempty" tf:"default_sender_id,omitempty"`
	// +optional
	DefaultSmsType string `json:"defaultSmsType,omitempty" tf:"default_sms_type,omitempty"`
	// +optional
	DeliveryStatusIamRoleArn string `json:"deliveryStatusIamRoleArn,omitempty" tf:"delivery_status_iam_role_arn,omitempty"`
	// +optional
	DeliveryStatusSuccessSamplingRate string `json:"deliveryStatusSuccessSamplingRate,omitempty" tf:"delivery_status_success_sampling_rate,omitempty"`
	// +optional
	MonthlySpendLimit string `json:"monthlySpendLimit,omitempty" tf:"monthly_spend_limit,omitempty"`
	// +optional
	UsageReportS3Bucket string `json:"usageReportS3Bucket,omitempty" tf:"usage_report_s3_bucket,omitempty"`
}

type SnsSmsPreferencesStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SnsSmsPreferencesSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
