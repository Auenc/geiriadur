package util

func AppendMap(toAppend map[string][]string, m map[string]string) map[string][]string {
	for key, value := range m {
		if _, ok := toAppend[key]; !ok {
			toAppend[key] = []string{}
		}
		toAppend[key] = append(toAppend[key], value)
	}
	return toAppend
}
