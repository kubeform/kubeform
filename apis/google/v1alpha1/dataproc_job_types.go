package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DataprocJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataprocJobSpec   `json:"spec,omitempty"`
	Status            DataprocJobStatus `json:"status,omitempty"`
}

type DataprocJobSpecHadoopConfigLoggingConfig struct {
	// +optional
	DriverLogLevels map[string]string `json:"driverLogLevels,omitempty" tf:"driver_log_levels,omitempty"`
}

type DataprocJobSpecHadoopConfig struct {
	// +optional
	ArchiveUris []string `json:"archiveUris,omitempty" tf:"archive_uris,omitempty"`
	// +optional
	Args []string `json:"args,omitempty" tf:"args,omitempty"`
	// +optional
	FileUris []string `json:"fileUris,omitempty" tf:"file_uris,omitempty"`
	// +optional
	JarFileUris []string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfig []DataprocJobSpecHadoopConfigLoggingConfig `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`
	// +optional
	MainClass string `json:"mainClass,omitempty" tf:"main_class,omitempty"`
	// +optional
	MainJarFileURI string `json:"mainJarFileURI,omitempty" tf:"main_jar_file_uri,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
}

type DataprocJobSpecHiveConfig struct {
	// +optional
	ContinueOnFailure bool `json:"continueOnFailure,omitempty" tf:"continue_on_failure,omitempty"`
	// +optional
	JarFileUris []string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
	// +optional
	QueryFileURI string `json:"queryFileURI,omitempty" tf:"query_file_uri,omitempty"`
	// +optional
	QueryList []string `json:"queryList,omitempty" tf:"query_list,omitempty"`
	// +optional
	ScriptVariables map[string]string `json:"scriptVariables,omitempty" tf:"script_variables,omitempty"`
}

type DataprocJobSpecPigConfigLoggingConfig struct {
	// +optional
	DriverLogLevels map[string]string `json:"driverLogLevels,omitempty" tf:"driver_log_levels,omitempty"`
}

type DataprocJobSpecPigConfig struct {
	// +optional
	ContinueOnFailure bool `json:"continueOnFailure,omitempty" tf:"continue_on_failure,omitempty"`
	// +optional
	JarFileUris []string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfig []DataprocJobSpecPigConfigLoggingConfig `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
	// +optional
	QueryFileURI string `json:"queryFileURI,omitempty" tf:"query_file_uri,omitempty"`
	// +optional
	QueryList []string `json:"queryList,omitempty" tf:"query_list,omitempty"`
	// +optional
	ScriptVariables map[string]string `json:"scriptVariables,omitempty" tf:"script_variables,omitempty"`
}

type DataprocJobSpecPlacement struct {
	ClusterName string `json:"clusterName" tf:"cluster_name"`
}

type DataprocJobSpecPysparkConfigLoggingConfig struct {
	// +optional
	DriverLogLevels map[string]string `json:"driverLogLevels,omitempty" tf:"driver_log_levels,omitempty"`
}

type DataprocJobSpecPysparkConfig struct {
	// +optional
	ArchiveUris []string `json:"archiveUris,omitempty" tf:"archive_uris,omitempty"`
	// +optional
	Args []string `json:"args,omitempty" tf:"args,omitempty"`
	// +optional
	FileUris []string `json:"fileUris,omitempty" tf:"file_uris,omitempty"`
	// +optional
	JarFileUris []string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfig     []DataprocJobSpecPysparkConfigLoggingConfig `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`
	MainPythonFileURI string                                      `json:"mainPythonFileURI" tf:"main_python_file_uri"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
	// +optional
	PythonFileUris []string `json:"pythonFileUris,omitempty" tf:"python_file_uris,omitempty"`
}

type DataprocJobSpecReference struct {
	// +optional
	JobID string `json:"jobID,omitempty" tf:"job_id,omitempty"`
}

type DataprocJobSpecScheduling struct {
	// +optional
	MaxFailuresPerHour int `json:"maxFailuresPerHour,omitempty" tf:"max_failures_per_hour,omitempty"`
}

type DataprocJobSpecSparkConfigLoggingConfig struct {
	// +optional
	DriverLogLevels map[string]string `json:"driverLogLevels,omitempty" tf:"driver_log_levels,omitempty"`
}

type DataprocJobSpecSparkConfig struct {
	// +optional
	ArchiveUris []string `json:"archiveUris,omitempty" tf:"archive_uris,omitempty"`
	// +optional
	Args []string `json:"args,omitempty" tf:"args,omitempty"`
	// +optional
	FileUris []string `json:"fileUris,omitempty" tf:"file_uris,omitempty"`
	// +optional
	JarFileUris []string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfig []DataprocJobSpecSparkConfigLoggingConfig `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`
	// +optional
	MainClass string `json:"mainClass,omitempty" tf:"main_class,omitempty"`
	// +optional
	MainJarFileURI string `json:"mainJarFileURI,omitempty" tf:"main_jar_file_uri,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
}

type DataprocJobSpecSparksqlConfigLoggingConfig struct {
	// +optional
	DriverLogLevels map[string]string `json:"driverLogLevels,omitempty" tf:"driver_log_levels,omitempty"`
}

type DataprocJobSpecSparksqlConfig struct {
	// +optional
	JarFileUris []string `json:"jarFileUris,omitempty" tf:"jar_file_uris,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfig []DataprocJobSpecSparksqlConfigLoggingConfig `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
	// +optional
	QueryFileURI string `json:"queryFileURI,omitempty" tf:"query_file_uri,omitempty"`
	// +optional
	QueryList []string `json:"queryList,omitempty" tf:"query_list,omitempty"`
	// +optional
	ScriptVariables map[string]string `json:"scriptVariables,omitempty" tf:"script_variables,omitempty"`
}

type DataprocJobSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	ForceDelete bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HadoopConfig []DataprocJobSpecHadoopConfig `json:"hadoopConfig,omitempty" tf:"hadoop_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HiveConfig []DataprocJobSpecHiveConfig `json:"hiveConfig,omitempty" tf:"hive_config,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PigConfig []DataprocJobSpecPigConfig `json:"pigConfig,omitempty" tf:"pig_config,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Placement []DataprocJobSpecPlacement `json:"placement" tf:"placement"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PysparkConfig []DataprocJobSpecPysparkConfig `json:"pysparkConfig,omitempty" tf:"pyspark_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Reference []DataprocJobSpecReference `json:"reference,omitempty" tf:"reference,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Scheduling []DataprocJobSpecScheduling `json:"scheduling,omitempty" tf:"scheduling,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SparkConfig []DataprocJobSpecSparkConfig `json:"sparkConfig,omitempty" tf:"spark_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SparksqlConfig []DataprocJobSpecSparksqlConfig `json:"sparksqlConfig,omitempty" tf:"sparksql_config,omitempty"`
}

type DataprocJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataprocJobList is a list of DataprocJobs
type DataprocJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataprocJob CRD objects
	Items []DataprocJob `json:"items,omitempty"`
}
