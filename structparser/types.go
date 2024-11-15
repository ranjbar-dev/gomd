package structparser

// structInfo holds information about a struct and its fields
type structInfo struct {
	Name    string
	Comment string
	Fields  []fieldInfo
}

// fieldInfo holds information about a single struct field
type fieldInfo struct {
	Name    string
	Type    string
	JSONTag string
	Comment string
}
