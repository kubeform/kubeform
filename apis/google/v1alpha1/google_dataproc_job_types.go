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

type GoogleDataprocJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleDataprocJobSpec   `json:"spec,omitempty"`
	Status            GoogleDataprocJobStatus `json:"status,omitempty"`
}

type GoogleDataprocJobSpecStatus struct {
	State          string `json:"state"`
	Details        string `json:"details"`
	StateStartTime string `json:"state_start_time"`
	Substate       string `json:"substate"`
}

type GoogleDataprocJobSpecScheduling struct {
	MaxFailuresPerHour int `json:"max_failures_per_hour"`
}

type GoogleDataprocJobSpecHadoopConfigLoggingConfig struct {
	DriverLogLevels map[string]string `json:"driver_log_levels"`
}

type GoogleDataprocJobSpecHadoopConfig struct {
	JarFileUris    []string                            `json:"jar_file_uris"`
	FileUris       []string                            `json:"file_uris"`
	ArchiveUris    []string                            `json:"archive_uris"`
	Properties     map[string]string                   `json:"properties"`
	LoggingConfig  []GoogleDataprocJobSpecHadoopConfig `json:"logging_config"`
	MainClass      string                              `json:"main_class"`
	MainJarFileUri string                              `json:"main_jar_file_uri"`
	Args           []string                            `json:"args"`
}

type GoogleDataprocJobSpecHiveConfig struct {
	QueryList         []string          `json:"query_list"`
	QueryFileUri      string            `json:"query_file_uri"`
	ContinueOnFailure bool              `json:"continue_on_failure"`
	ScriptVariables   map[string]string `json:"script_variables"`
	Properties        map[string]string `json:"properties"`
	JarFileUris       []string          `json:"jar_file_uris"`
}

type GoogleDataprocJobSpecSparksqlConfigLoggingConfig struct {
	DriverLogLevels map[string]string `json:"driver_log_levels"`
}

type GoogleDataprocJobSpecSparksqlConfig struct {
	QueryList       []string                              `json:"query_list"`
	QueryFileUri    string                                `json:"query_file_uri"`
	ScriptVariables map[string]string                     `json:"script_variables"`
	Properties      map[string]string                     `json:"properties"`
	JarFileUris     []string                              `json:"jar_file_uris"`
	LoggingConfig   []GoogleDataprocJobSpecSparksqlConfig `json:"logging_config"`
}

type GoogleDataprocJobSpecPysparkConfigLoggingConfig struct {
	DriverLogLevels map[string]string `json:"driver_log_levels"`
}

type GoogleDataprocJobSpecPysparkConfig struct {
	Properties        map[string]string                    `json:"properties"`
	LoggingConfig     []GoogleDataprocJobSpecPysparkConfig `json:"logging_config"`
	MainPythonFileUri string                               `json:"main_python_file_uri"`
	Args              []string                             `json:"args"`
	PythonFileUris    []string                             `json:"python_file_uris"`
	JarFileUris       []string                             `json:"jar_file_uris"`
	FileUris          []string                             `json:"file_uris"`
	ArchiveUris       []string                             `json:"archive_uris"`
}

type GoogleDataprocJobSpecReference struct {
	JobId string `json:"job_id"`
}

type GoogleDataprocJobSpecPlacement struct {
	ClusterName string `json:"cluster_name"`
	ClusterUuid string `json:"cluster_uuid"`
}

type GoogleDataprocJobSpecSparkConfigLoggingConfig struct {
	DriverLogLevels map[string]string `json:"driver_log_levels"`
}

type GoogleDataprocJobSpecSparkConfig struct {
	Args           []string                           `json:"args"`
	JarFileUris    []string                           `json:"jar_file_uris"`
	FileUris       []string                           `json:"file_uris"`
	ArchiveUris    []string                           `json:"archive_uris"`
	Properties     map[string]string                  `json:"properties"`
	LoggingConfig  []GoogleDataprocJobSpecSparkConfig `json:"logging_config"`
	MainClass      string                             `json:"main_class"`
	MainJarFileUri string                             `json:"main_jar_file_uri"`
}

type GoogleDataprocJobSpecPigConfigLoggingConfig struct {
	DriverLogLevels map[string]string `json:"driver_log_levels"`
}

type GoogleDataprocJobSpecPigConfig struct {
	QueryList         []string                         `json:"query_list"`
	QueryFileUri      string                           `json:"query_file_uri"`
	ContinueOnFailure bool                             `json:"continue_on_failure"`
	ScriptVariables   map[string]string                `json:"script_variables"`
	Properties        map[string]string                `json:"properties"`
	JarFileUris       []string                         `json:"jar_file_uris"`
	LoggingConfig     []GoogleDataprocJobSpecPigConfig `json:"logging_config"`
}

type GoogleDataprocJobSpec struct {
	Status                  []GoogleDataprocJobSpec `json:"status"`
	Scheduling              []GoogleDataprocJobSpec `json:"scheduling"`
	HadoopConfig            []GoogleDataprocJobSpec `json:"hadoop_config"`
	HiveConfig              []GoogleDataprocJobSpec `json:"hive_config"`
	Region                  string                  `json:"region"`
	SparksqlConfig          []GoogleDataprocJobSpec `json:"sparksql_config"`
	PysparkConfig           []GoogleDataprocJobSpec `json:"pyspark_config"`
	DriverOutputResourceUri string                  `json:"driver_output_resource_uri"`
	Labels                  map[string]string       `json:"labels"`
	ForceDelete             bool                    `json:"force_delete"`
	Reference               []GoogleDataprocJobSpec `json:"reference"`
	Placement               []GoogleDataprocJobSpec `json:"placement"`
	DriverControlsFilesUri  string                  `json:"driver_controls_files_uri"`
	SparkConfig             []GoogleDataprocJobSpec `json:"spark_config"`
	PigConfig               []GoogleDataprocJobSpec `json:"pig_config"`
	Project                 string                  `json:"project"`
}

type GoogleDataprocJobStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleDataprocJobList is a list of GoogleDataprocJobs
type GoogleDataprocJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleDataprocJob CRD objects
	Items []GoogleDataprocJob `json:"items,omitempty"`
}
