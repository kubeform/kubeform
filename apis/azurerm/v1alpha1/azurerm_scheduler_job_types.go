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

type AzurermSchedulerJobSpecRetry struct {
	Interval string `json:"interval"`
	Count    int    `json:"count"`
}

type AzurermSchedulerJobSpecErrorActionWebAuthenticationActiveDirectory struct {
	TenantId string `json:"tenant_id"`
	ClientId string `json:"client_id"`
	Secret   string `json:"secret"`
	Audience string `json:"audience"`
}

type AzurermSchedulerJobSpecErrorActionWebAuthenticationBasic struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AzurermSchedulerJobSpecErrorActionWebAuthenticationCertificate struct {
	Pfx         string `json:"pfx"`
	Password    string `json:"password"`
	Thumbprint  string `json:"thumbprint"`
	Expiration  string `json:"expiration"`
	SubjectName string `json:"subject_name"`
}

type AzurermSchedulerJobSpecErrorActionWeb struct {
	AuthenticationActiveDirectory []AzurermSchedulerJobSpecErrorActionWeb `json:"authentication_active_directory"`
	Url                           string                                  `json:"url"`
	Method                        string                                  `json:"method"`
	Body                          string                                  `json:"body"`
	Headers                       map[string]string                       `json:"headers"`
	AuthenticationBasic           []AzurermSchedulerJobSpecErrorActionWeb `json:"authentication_basic"`
	AuthenticationCertificate     []AzurermSchedulerJobSpecErrorActionWeb `json:"authentication_certificate"`
}

type AzurermSchedulerJobSpecErrorActionStorageQueue struct {
	StorageAccountName string `json:"storage_account_name"`
	StorageQueueName   string `json:"storage_queue_name"`
	SasToken           string `json:"sas_token"`
	Message            string `json:"message"`
}

type AzurermSchedulerJobSpecActionStorageQueue struct {
	StorageAccountName string `json:"storage_account_name"`
	StorageQueueName   string `json:"storage_queue_name"`
	SasToken           string `json:"sas_token"`
	Message            string `json:"message"`
}

type AzurermSchedulerJobSpecRecurrenceMonthlyOccurrences struct {
	Day        string `json:"day"`
	Occurrence int    `json:"occurrence"`
}

type AzurermSchedulerJobSpecRecurrence struct {
	Interval           int                                 `json:"interval"`
	Count              int                                 `json:"count"`
	WeekDays           []string                            `json:"week_days"`
	Frequency          string                              `json:"frequency"`
	EndTime            string                              `json:"end_time"`
	Minutes            []int64                             `json:"minutes"`
	Hours              []int64                             `json:"hours"`
	MonthDays          []int64                             `json:"month_days"`
	MonthlyOccurrences []AzurermSchedulerJobSpecRecurrence `json:"monthly_occurrences"`
}

type AzurermSchedulerJobSpecActionWebAuthenticationBasic struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type AzurermSchedulerJobSpecActionWebAuthenticationCertificate struct {
	Pfx         string `json:"pfx"`
	Password    string `json:"password"`
	Thumbprint  string `json:"thumbprint"`
	Expiration  string `json:"expiration"`
	SubjectName string `json:"subject_name"`
}

type AzurermSchedulerJobSpecActionWebAuthenticationActiveDirectory struct {
	TenantId string `json:"tenant_id"`
	ClientId string `json:"client_id"`
	Secret   string `json:"secret"`
	Audience string `json:"audience"`
}

type AzurermSchedulerJobSpecActionWeb struct {
	Body                          string                             `json:"body"`
	Headers                       map[string]string                  `json:"headers"`
	AuthenticationBasic           []AzurermSchedulerJobSpecActionWeb `json:"authentication_basic"`
	AuthenticationCertificate     []AzurermSchedulerJobSpecActionWeb `json:"authentication_certificate"`
	AuthenticationActiveDirectory []AzurermSchedulerJobSpecActionWeb `json:"authentication_active_directory"`
	Url                           string                             `json:"url"`
	Method                        string                             `json:"method"`
}

type AzurermSchedulerJobSpec struct {
	Retry                   []AzurermSchedulerJobSpec `json:"retry"`
	ResourceGroupName       string                    `json:"resource_group_name"`
	ErrorActionWeb          []AzurermSchedulerJobSpec `json:"error_action_web"`
	ErrorActionStorageQueue []AzurermSchedulerJobSpec `json:"error_action_storage_queue"`
	ActionStorageQueue      []AzurermSchedulerJobSpec `json:"action_storage_queue"`
	Recurrence              []AzurermSchedulerJobSpec `json:"recurrence"`
	StartTime               string                    `json:"start_time"`
	State                   string                    `json:"state"`
	Name                    string                    `json:"name"`
	JobCollectionName       string                    `json:"job_collection_name"`
	ActionWeb               []AzurermSchedulerJobSpec `json:"action_web"`
}

type AzurermSchedulerJobStatus struct {
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
