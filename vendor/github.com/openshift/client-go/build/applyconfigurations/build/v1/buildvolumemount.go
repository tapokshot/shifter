// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// BuildVolumeMountApplyConfiguration represents an declarative configuration of the BuildVolumeMount type for use
// with apply.
type BuildVolumeMountApplyConfiguration struct {
	DestinationPath *string `json:"destinationPath,omitempty"`
}

// BuildVolumeMountApplyConfiguration constructs an declarative configuration of the BuildVolumeMount type for use with
// apply.
func BuildVolumeMount() *BuildVolumeMountApplyConfiguration {
	return &BuildVolumeMountApplyConfiguration{}
}

// WithDestinationPath sets the DestinationPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DestinationPath field is set to the value of the last call.
func (b *BuildVolumeMountApplyConfiguration) WithDestinationPath(value string) *BuildVolumeMountApplyConfiguration {
	b.DestinationPath = &value
	return b
}
