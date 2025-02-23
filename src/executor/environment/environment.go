package parser

type Environment struct {
	store map[string]any
}

func NewEnvironment() Environment {
	return Environment{
		store: make(map[string]any),
	}
}

func (env Environment) Set(key string, value any) {
	env.store[key] = value
}

// --- for better error logging, do so in the caller to get
func (env Environment) Get(key string) (any, bool) {
	val, ok := env.store[key]
	return val, ok
}
