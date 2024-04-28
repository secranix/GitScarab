package env

// Options represents a function type that modifies the behavior of the env struct.
type Options func(e *env)

// WithOptionalFlag flags the environment variable as optional.
func WithOptionalFlag() Options {
	return func(e *env) {
		// Setting required field to false indicates the variable is optional.
		e.required = false
	}
}
