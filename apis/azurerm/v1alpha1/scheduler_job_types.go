package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type SchedulerJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SchedulerJobSpec   `json:"spec,omitempty"`
	Status            SchedulerJobStatus `json:"status,omitempty"`
}

type SchedulerJobSpecActionStorageQueue struct {
	Message            string `json:"message" tf:"message"`
	SasToken           string `json:"sasToken" tf:"sas_token"`
	StorageAccountName string `json:"storageAccountName" tf:"storage_account_name"`
	StorageQueueName   string `json:"storageQueueName" tf:"storage_queue_name"`
}

type SchedulerJobSpecActionWebAuthenticationActiveDirectory struct {
	// +optional
	Audience string `json:"audience,omitempty" tf:"audience,omitempty"`
	ClientID string `json:"clientID" tf:"client_id"`
	Secret   string `json:"-" sensitive:"true" tf:"secret"`
	TenantID string `json:"tenantID" tf:"tenant_id"`
}

type SchedulerJobSpecActionWebAuthenticationBasic struct {
	Password string `json:"-" sensitive:"true" tf:"password"`
	Username string `json:"username" tf:"username"`
}

type SchedulerJobSpecActionWebAuthenticationCertificate struct {
	// +optional
	Expiration string `json:"expiration,omitempty" tf:"expiration,omitempty"`
	Password   string `json:"-" sensitive:"true" tf:"password"`
	Pfx        string `json:"-" sensitive:"true" tf:"pfx"`
	// +optional
	SubjectName string `json:"subjectName,omitempty" tf:"subject_name,omitempty"`
	// +optional
	Thumbprint string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type SchedulerJobSpecActionWeb struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticationActiveDirectory []SchedulerJobSpecActionWebAuthenticationActiveDirectory `json:"authenticationActiveDirectory,omitempty" tf:"authentication_active_directory,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticationBasic []SchedulerJobSpecActionWebAuthenticationBasic `json:"authenticationBasic,omitempty" tf:"authentication_basic,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticationCertificate []SchedulerJobSpecActionWebAuthenticationCertificate `json:"authenticationCertificate,omitempty" tf:"authentication_certificate,omitempty"`
	// +optional
	Body string `json:"body,omitempty" tf:"body,omitempty"`
	// +optional
	Headers map[string]string `json:"headers,omitempty" tf:"headers,omitempty"`
	Method  string            `json:"method" tf:"method"`
	Url     string            `json:"url" tf:"url"`
}

type SchedulerJobSpecErrorActionStorageQueue struct {
	Message            string `json:"message" tf:"message"`
	SasToken           string `json:"sasToken" tf:"sas_token"`
	StorageAccountName string `json:"storageAccountName" tf:"storage_account_name"`
	StorageQueueName   string `json:"storageQueueName" tf:"storage_queue_name"`
}

type SchedulerJobSpecErrorActionWebAuthenticationActiveDirectory struct {
	// +optional
	Audience string `json:"audience,omitempty" tf:"audience,omitempty"`
	ClientID string `json:"clientID" tf:"client_id"`
	Secret   string `json:"-" sensitive:"true" tf:"secret"`
	TenantID string `json:"tenantID" tf:"tenant_id"`
}

type SchedulerJobSpecErrorActionWebAuthenticationBasic struct {
	Password string `json:"-" sensitive:"true" tf:"password"`
	Username string `json:"username" tf:"username"`
}

type SchedulerJobSpecErrorActionWebAuthenticationCertificate struct {
	// +optional
	Expiration string `json:"expiration,omitempty" tf:"expiration,omitempty"`
	Password   string `json:"-" sensitive:"true" tf:"password"`
	Pfx        string `json:"-" sensitive:"true" tf:"pfx"`
	// +optional
	SubjectName string `json:"subjectName,omitempty" tf:"subject_name,omitempty"`
	// +optional
	Thumbprint string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type SchedulerJobSpecErrorActionWeb struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticationActiveDirectory []SchedulerJobSpecErrorActionWebAuthenticationActiveDirectory `json:"authenticationActiveDirectory,omitempty" tf:"authentication_active_directory,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticationBasic []SchedulerJobSpecErrorActionWebAuthenticationBasic `json:"authenticationBasic,omitempty" tf:"authentication_basic,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticationCertificate []SchedulerJobSpecErrorActionWebAuthenticationCertificate `json:"authenticationCertificate,omitempty" tf:"authentication_certificate,omitempty"`
	// +optional
	Body string `json:"body,omitempty" tf:"body,omitempty"`
	// +optional
	Headers map[string]string `json:"headers,omitempty" tf:"headers,omitempty"`
	Method  string            `json:"method" tf:"method"`
	Url     string            `json:"url" tf:"url"`
}

type SchedulerJobSpecRecurrenceMonthlyOccurrences struct {
	Day        string `json:"day" tf:"day"`
	Occurrence int    `json:"occurrence" tf:"occurrence"`
}

type SchedulerJobSpecRecurrence struct {
	// +optional
	Count int `json:"count,omitempty" tf:"count,omitempty"`
	// +optional
	EndTime   string `json:"endTime,omitempty" tf:"end_time,omitempty"`
	Frequency string `json:"frequency" tf:"frequency"`
	// +optional
	Hours []int64 `json:"hours,omitempty" tf:"hours,omitempty"`
	// +optional
	Interval int `json:"interval,omitempty" tf:"interval,omitempty"`
	// +optional
	Minutes []int64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	MonthDays []int64 `json:"monthDays,omitempty" tf:"month_days,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	MonthlyOccurrences []SchedulerJobSpecRecurrenceMonthlyOccurrences `json:"monthlyOccurrences,omitempty" tf:"monthly_occurrences,omitempty"`
	// +optional
	WeekDays []string `json:"weekDays,omitempty" tf:"week_days,omitempty"`
}

type SchedulerJobSpecRetry struct {
	// +optional
	Count int `json:"count,omitempty" tf:"count,omitempty"`
	// +optional
	Interval string `json:"interval,omitempty" tf:"interval,omitempty"`
}

type SchedulerJobSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	ActionStorageQueue []SchedulerJobSpecActionStorageQueue `json:"actionStorageQueue,omitempty" tf:"action_storage_queue,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ActionWeb []SchedulerJobSpecActionWeb `json:"actionWeb,omitempty" tf:"action_web,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ErrorActionStorageQueue []SchedulerJobSpecErrorActionStorageQueue `json:"errorActionStorageQueue,omitempty" tf:"error_action_storage_queue,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ErrorActionWeb    []SchedulerJobSpecErrorActionWeb `json:"errorActionWeb,omitempty" tf:"error_action_web,omitempty"`
	JobCollectionName string                           `json:"jobCollectionName" tf:"job_collection_name"`
	Name              string                           `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Recurrence        []SchedulerJobSpecRecurrence `json:"recurrence,omitempty" tf:"recurrence,omitempty"`
	ResourceGroupName string                       `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Retry []SchedulerJobSpecRetry `json:"retry,omitempty" tf:"retry,omitempty"`
	// +optional
	StartTime string `json:"startTime,omitempty" tf:"start_time,omitempty"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
}

type SchedulerJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SchedulerJobSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
