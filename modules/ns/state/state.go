package state

// Stater is for interfacing with a state management system
type Stater interface {
	Init(config []byte) error // Initialize the state system needs org_id,job_id and supporting connection details
	IsNew(zone string) bool   // Checks if the zone is new/has already been analyzed.
}