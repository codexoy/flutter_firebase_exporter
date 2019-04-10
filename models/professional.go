package models

type Professional struct {
	Designation      string `json:"designation"`
	OrganizationName string `json:"organizationName"`
	TechStack        string `json:"techStack"`
	YearsOfExps      string `json:"yearsOfExp"`
}
