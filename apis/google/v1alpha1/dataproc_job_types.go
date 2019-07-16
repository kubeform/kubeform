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

type DataprocJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataprocJobSpec   `json:"spec,omitempty"`
	Status            DataprocJobStatus `json:"status,omitempty"`
}

type DataprocJobSpecHadoopConfig struct {
	// +optional
	ArchiveUris []string `json:"archive_uris,omitempty"`
	// +optional
	Args []string `json:"args,omitempty"`
	// +optional
	FileUris []string `json:"file_uris,omitempty"`
	// +optional
	JarFileUris []string `json:"jar_file_uris,omitempty"`
	// +optional
	MainClass string `json:"main_class,omitempty"`
	// +optional
	MainJarFileUri string `json:"main_jar_file_uri,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty"`
}

type DataprocJobSpecHiveConfig struct {
	// +optional
	ContinueOnFailure bool `json:"continue_on_failure,omitempty"`
	// +optional
	JarFileUris []string `json:"jar_file_uris,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty"`
	// +optional
	QueryFileUri string `json:"query_file_uri,omitempty"`
	// +optional
	QueryList []string `json:"query_list,omitempty"`
	// +optional
	ScriptVariables map[string]string `json:"script_variables,omitempty"`
}

type DataprocJobSpecPigConfig struct {
	// +optional
	ContinueOnFailure bool `json:"continue_on_failure,omitempty"`
	// +optional
	JarFileUris []string `json:"jar_file_uris,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty"`
	// +optional
	QueryFileUri string `json:"query_file_uri,omitempty"`
	// +optional
	QueryList []string `json:"query_list,omitempty"`
	// +optional
	ScriptVariables map[string]string `json:"script_variables,omitempty"`
}

type DataprocJobSpecPlacement struct {
	ClusterName string `json:"cluster_name"`
}

type DataprocJobSpecPysparkConfig struct {
	// +optional
	ArchiveUris []string `json:"archive_uris,omitempty"`
	// +optional
	Args []string `json:"args,omitempty"`
	// +optional
	FileUris []string `json:"file_uris,omitempty"`
	// +optional
	JarFileUris       []string `json:"jar_file_uris,omitempty"`
	MainPythonFileUri string   `json:"main_python_file_uri"`
	// +optional
	Properties map[string]string `json:"properties,omitempty"`
	// +optional
	PythonFileUris []string `json:"python_file_uris,omitempty"`
}

type DataprocJobSpecScheduling struct {
	// +optional
	MaxFailuresPerHour int `json:"max_failures_per_hour,omitempty"`
}

type DataprocJobSpecSparkConfig struct {
	// +optional
	ArchiveUris []string `json:"archive_uris,omitempty"`
	// +optional
	Args []string `json:"args,omitempty"`
	// +optional
	FileUris []string `json:"file_uris,omitempty"`
	// +optional
	JarFileUris []string `json:"jar_file_uris,omitempty"`
	// +optional
	MainClass string `json:"main_class,omitempty"`
	// +optional
	MainJarFileUri string `json:"main_jar_file_uri,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty"`
}

type DataprocJobSpecSparksqlConfig struct {
	// +optional
	JarFileUris []string `json:"jar_file_uris,omitempty"`
	// +optional
	Properties map[string]string `json:"properties,omitempty"`
	// +optional
	QueryFileUri string `json:"query_file_uri,omitempty"`
	// +optional
	QueryList []string `json:"query_list,omitempty"`
	// +optional
	ScriptVariables map[string]string `json:"script_variables,omitempty"`
}

type DataprocJobSpec struct {
	// +optional
	ForceDelete bool `json:"force_delete,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HadoopConfig *[]DataprocJobSpec `json:"hadoop_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HiveConfig *[]DataprocJobSpec `json:"hive_config,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PigConfig *[]DataprocJobSpec `json:"pig_config,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Placement []DataprocJobSpec `json:"placement"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PysparkConfig *[]DataprocJobSpec `json:"pyspark_config,omitempty"`
	// +optional
	Region string `json:"region,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Scheduling *[]DataprocJobSpec `json:"scheduling,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SparkConfig *[]DataprocJobSpec `json:"spark_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SparksqlConfig *[]DataprocJobSpec `json:"sparksql_config,omitempty"`
}

type DataprocJobStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
