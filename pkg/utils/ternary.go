package utils

func If_key_exist_else[T any](dict map[string]any, key string, def T) T {
	if val, ok := dict[key]; ok {
		return val.(T)
	}
	return def
}
