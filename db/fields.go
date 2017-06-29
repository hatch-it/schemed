package db

type Fields map[string]interface{}

func (fields *Fields) Keys(separator string) string {
	keys := make([]string, 0, len(fields))
	for key := range fields {
		keys = append(keys, key)
	}
	return strings.Join(keys, separator)
}

func (fields *Fields) Values(separator string) string {
	values := make([]string, 0, len(fields))
	for _, value := range fields {
		values = append(values, value)
	}
	return strings.Join(values, separator)
}
