package models

import (
	"context"
)

// SchedulerUnit is implemented by supported schedulers
type SchedulerUnit interface {
	GetName() string

	// GetTemplatePath returns the path for template files used during spec
	// compilation. Each scheduler needs to have a template file stored under
	// resources/pack/template/scheduler dir. This path needs to be relative
	// from pack folder.
	GetTemplatePath() string

	// GetJobsDir returns the directory to use while storing compiled
	// scheduler specific input files
	GetJobsDir() string

	// GetJobsExtension provides extension to use for input files of scheduler
	GetJobsExtension() string

	// Bootstrap will be executed per project when the application boots up
	// this can be used to do adhoc commands for initialization of scheduler
	Bootstrap(context.Context, ProjectSpec) error
}

// Scheduler is a single unit initialized at the start of application
// based on config. This will be used to perform adhoc operations
// to support target scheduling engine
var Scheduler SchedulerUnit
