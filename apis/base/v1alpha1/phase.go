package v1alpha1

type Phase string

const (
	// used for Kubeform Resources that are currently running
	PhaseRunning Phase = "Running"
	// used for Kubeform Resources that are currently applying
	PhaseApplying Phase = "Applying"
	// used for Kubeform Resources that are currently initializing
	PhaseInitializing Phase = "Initializing"
	// used for Kubeform Resources that are Failed
	PhaseFailed Phase = "Failed"
	// used for Kubeform Resources that are Deleting
	PhaseDeleting Phase = "Deleting"
)
