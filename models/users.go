package models

// User a user model
type User struct {
	Name                    string        `json:"name"`
	Email                   string        `json:"email"`
	MobileNumber            string        `json:"mobileNumber"`
	IsContributor           bool          `json:"isContributor"`
	IsRegistered            bool          `json:"isRegistered"`
	IsPresent               bool          `json:"isPresent"`
	IsRegistrationConfirmed bool          `json:"isRegistrationConfirmed"`
	PhotoURL                string        `json:"photoUrl"`
	Registration            *Registration `json:"registration"`
	ProfessionalDetails     *Professional `json:"professionalDetails"`
	StudentDetails          *Student      `json:"studentDetails"`
}
