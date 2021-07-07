package compoas

type OAS struct {
	Openapi    string
	Info       Info
	Components Components
}

type Info struct {
	Title   string
	Version string
}

type Components struct {
	SecuritySchemes map[string]SecurityScheme
}

type SecurityScheme struct {
	Type         string
	Scheme       string
	BearerFormat string
}
