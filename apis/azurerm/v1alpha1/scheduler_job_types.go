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

type SchedulerJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchedulerJobSpec   `json:"spec,omitempty"`
	Status            SchedulerJobStatus `json:"status,omitempty"`
}

type SchedulerJobSpecActionStorageQueue struct {
	Message            string `json:"message"`
	SasToken           string `json:"sas_token"`
	StorageAccountName string `json:"storage_account_name"`
	StorageQueueName   string `json:"storage_queue_name"`
}

type SchedulerJobSpecActionWebAuthenticationActiveDirectory struct {
	ClientId string `json:"client_id"`
	Secret   string `json:"secret"`
	TenantId string `json:"tenant_id"`
}

type SchedulerJobSpecActionWebAuthenticationBasic struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type SchedulerJobSpecActionWebAuthenticationCertificate struct {
	Password string `json:"password"`
	Pfx      string `json:"pfx"`
}

type SchedulerJobSpecActionWeb struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticationActiveDirectory *[]SchedulerJobSpecActionWeb `json:"authentication_active_directory,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticationBasic *[]SchedulerJobSpecActionWeb `json:"authentication_basic,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticationCertificate *[]SchedulerJobSpecActionWeb `json:"authentication_certificate,omitempty"`
	// +optional
	Body string `json:"body,omitempty"`
	// +optional
	Headers map[string]string `json:"headers,omitempty"`
	Method  string            `json:"method"`
	Url     string            `json:"url"`
}

type SchedulerJobSpecErrorActionStorageQueue struct {
	Message            string `json:"message"`
	SasToken           string `json:"sas_token"`
	StorageAccountName string `json:"storage_account_name"`
	StorageQueueName   string `json:"storage_queue_name"`
}

type SchedulerJobSpecErrorActionWebAuthenticationActiveDirectory struct {
	ClientId string `json:"client_id"`
	Secret   string `json:"secret"`
	TenantId string `json:"tenant_id"`
}

type SchedulerJobSpecErrorActionWebAuthenticationBasic struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type SchedulerJobSpecErrorActionWebAuthenticationCertificate struct {
	Password string `json:"password"`
	Pfx      string `json:"pfx"`
}

type SchedulerJobSpecErrorActionWeb struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticationActiveDirectory *[]SchedulerJobSpecErrorActionWeb `json:"authentication_active_directory,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticationBasic *[]SchedulerJobSpecErrorActionWeb `json:"authentication_basic,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticationCertificate *[]SchedulerJobSpecErrorActionWeb `json:"authentication_certificate,omitempty"`
	// +optional
	Body string `json:"body,omitempty"`
	// +optional
	Headers map[string]string `json:"headers,omitempty"`
	Method  string            `json:"method"`
	Url     string            `json:"url"`
}

type SchedulerJobSpecRecurrenceMonthlyOccurrences struct {
	Day        string `json:"day"`
	Occurrence int    `json:"occurrence"`
}

type SchedulerJobSpecRecurrence struct {
	// +optional
	Count     int    `json:"count,omitempty"`
	Frequency string `json:"frequency"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Hours []int64 `json:"hours,omitempty"`
	// +optional
	Interval int `json:"interval,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Minutes []int64 `json:"minutes,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	MonthDays []int64 `json:"month_days,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	MonthlyOccurrences *[]SchedulerJobSpecRecurrence `json:"monthly_occurrences,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	WeekDays []string `json:"week_days,omitempty"`
}

type SchedulerJobSpecRetry struct {
	// +optional
	Count int `json:"count,omitempty"`
	// +optional
	Interval string `json:"interval,omitempty"`
}

type SchedulerJobSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ActionStorageQueue *[]SchedulerJobSpec `json:"action_storage_queue,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ActionWeb *[]SchedulerJobSpec `json:"action_web,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ErrorActionStorageQueue *[]SchedulerJobSpec `json:"error_action_storage_queue,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ErrorActionWeb    *[]SchedulerJobSpec `json:"error_action_web,omitempty"`
	JobCollectionName string              `json:"job_collection_name"`
	Name              string              `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Recurrence        *[]SchedulerJobSpec `json:"recurrence,omitempty"`
	ResourceGroupName string              `json:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Retry *[]SchedulerJobSpec `json:"retry,omitempty"`
}

type SchedulerJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SchedulerJobList is a list of SchedulerJobs
type SchedulerJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SchedulerJob CRD objects
	Items []SchedulerJob `json:"items,omitempty"`
}
