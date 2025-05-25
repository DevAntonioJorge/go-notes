package path

func NormalizePath(path string) string {
	if path == "" {
		return "/"
	}
	if path[0] != '/' {
		path = "/" + path
	}
	if path[len(path)-1] != '/' {
		path += "/"
	}
	return path
}
func IsValidPath(path string) bool {
	if path == "" {
		return false
	}
	if path[0] != '/' {
		return false
	}
	if path[len(path)-1] != '/' {
		return false
	}
	return true
}
