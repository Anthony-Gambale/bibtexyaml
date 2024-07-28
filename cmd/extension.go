package cmd

func isYamlFile(name string) bool {
	if len(name) < 5 {
		return false
	}
	return name[len(name)-5:] == ".yaml"
}
