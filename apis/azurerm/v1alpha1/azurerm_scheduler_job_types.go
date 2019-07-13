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

type AzurermSchedulerJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSchedulerJobSpec   `json:"spec,omitempty"`
	Status            AzurermSchedulerJobStatus `json:"status,omitempty"`
}

type AzurermSchedulerJobSpecActionStorageQueue struct {
	StorageAccountName string `json:"storage_account_name"`
	StorageQueueName   string `json:"storage_queue_name"`
	SasToken           string `json:"sas_token"`
	Message            string `json:"message"`
}

type AzurermSchedulerJobSpecErrorActionWebAuthenticationBasic struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermSchedulerJobSpecErrorActionWebAuthenticationCertificate struct {
	SubjectName string `json:"subject_name"`
	Pfx         string `json:"pfx"`
	Password    string `json:"password"`
	Thumbprint  string `json:"thumbprint"`
	Expiration  string `json:"expiration"`
}

type AzurermSchedulerJobSpecErrorActionWebAuthenticationActiveDirectory struct {
	TenantId string `json:"tenant_id"`
	ClientId string `json:"client_id"`
	Secret   string `json:"secret"`
	Audience string `json:"audience"`
}

type AzurermSchedulerJobSpecErrorActionWeb struct {
	Body                          string                                  `json:"body"`
	Headers                       map[string]string                       `json:"headers"`
	AuthenticationBasic           []AzurermSchedulerJobSpecErrorActionWeb `json:"authentication_basic"`
	AuthenticationCertificate     []AzurermSchedulerJobSpecErrorActionWeb `json:"authentication_certificate"`
	AuthenticationActiveDirectory []AzurermSchedulerJobSpecErrorActionWeb `json:"authentication_active_directory"`
	Url                           string                                  `json:"url"`
	Method                        string                                  `json:"method"`
}

type AzurermSchedulerJobSpecErrorActionStorageQueue struct {
	StorageAccountName string `json:"storage_account_name"`
	StorageQueueName   string `json:"storage_queue_name"`
	SasToken           string `json:"sas_token"`
	Message            string `json:"message"`
}

type AzurermSchedulerJobSpecActionWebAuthenticationBasic struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermSchedulerJobSpecActionWebAuthenticationCertificate struct {
	Pfx         string `json:"pfx"`
	Password    string `json:"password"`
	Thumbprint  string `json:"thumbprint"`
	Expiration  string `json:"expiration"`
	SubjectName string `json:"subject_name"`
}

type AzurermSchedulerJobSpecActionWebAuthenticationActiveDirectory struct {
	Audience string `json:"audience"`
	TenantId string `json:"tenant_id"`
	ClientId string `json:"client_id"`
	Secret   string `json:"secret"`
}

type AzurermSchedulerJobSpecActionWeb struct {
	Headers                       map[string]string                  `json:"headers"`
	AuthenticationBasic           []AzurermSchedulerJobSpecActionWeb `json:"authentication_basic"`
	AuthenticationCertificate     []AzurermSchedulerJobSpecActionWeb `json:"authentication_certificate"`
	AuthenticationActiveDirectory []AzurermSchedulerJobSpecActionWeb `json:"authentication_active_directory"`
	Url                           string                             `json:"url"`
	Method                        string                             `json:"method"`
	Body                          string                             `json:"body"`
}

type AzurermSchedulerJobSpecRetry struct {
	Interval string `json:"interval"`
	Count    int    `json:"count"`
}

type AzurermSchedulerJobSpecRecurrenceMonthlyOccurrences struct {
	Day        string `json:"day"`
	Occurrence int    `json:"occurrence"`
}

type AzurermSchedulerJobSpecRecurrence struct {
	Count              int                                 `json:"count"`
	EndTime            string                              `json:"end_time"`
	Minutes            []int64                             `json:"minutes"`
	Hours              []int64                             `json:"hours"`
	MonthDays          []int64                             `json:"month_days"`
	Frequency          string                              `json:"frequency"`
	Interval           int                                 `json:"interval"`
	WeekDays           []string                            `json:"week_days"`
	MonthlyOccurrences []AzurermSchedulerJobSpecRecurrence `json:"monthly_occurrences"`
}

type AzurermSchedulerJobSpec struct {
	Name                    string                    `json:"name"`
	ResourceGroupName       string                    `json:"resource_group_name"`
	JobCollectionName       string                    `json:"job_collection_name"`
	ActionStorageQueue      []AzurermSchedulerJobSpec `json:"action_storage_queue"`
	ErrorActionWeb          []AzurermSchedulerJobSpec `json:"error_action_web"`
	ErrorActionStorageQueue []AzurermSchedulerJobSpec `json:"error_action_storage_queue"`
	ActionWeb               []AzurermSchedulerJobSpec `json:"action_web"`
	Retry                   []AzurermSchedulerJobSpec `json:"retry"`
	Recurrence              []AzurermSchedulerJobSpec `json:"recurrence"`
	StartTime               string                    `json:"start_time"`
	State                   string                    `json:"state"`
}



type AzurermSchedulerJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermSchedulerJobList is a list of AzurermSchedulerJobs
type AzurermSchedulerJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSchedulerJob CRD objects
	Items []AzurermSchedulerJob `json:"items,omitempty"`
}