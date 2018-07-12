package am

// JobConfig contains all necessary information for starting a job.
type JobConfig struct {
	OrgID            int
	ScanGroupID      int
	ConfiguredUserID int       // UserID that configured the job
	LaunchUserID     int       // UserId that started the job
	InputID          int64     // The ID of the original input list
	Modules          []*Module // Modules used in this job config
	MaxHostsAllowed  int64     // Max number of hosts allowed to identify (10k etc)
}
