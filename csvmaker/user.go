package csvmaker

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/minhajuddinkhan/flutter_pk_firebase_export/models"
)

var (
	userCSVHeaders = []string{"name", "email", "phone", "occupation", "institution", "designation", "experience", "tech", "competition", "bringingLaptop"}
)

// CSVMaker CSVMaker
type CSVMaker interface {
	User(users []models.User) error
}
type csvmaker struct {
	Writer *csv.Writer
}

// New creates a new csv maker
func New(w io.Writer) CSVMaker {
	return &csvmaker{Writer: csv.NewWriter(w)}
}

// User creates user csv
func (m *csvmaker) User(users []models.User) error {
	records := [][]string{
		userCSVHeaders,
	}
	for _, user := range users {
		records = append(records, mapUser(user))
	}

	return m.Writer.WriteAll(records)
}

func mapUser(user models.User) []string {
	return []string{
		user.Name,
		user.Email,
		user.MobileNumber,

		//Occupation
		func(user models.User) string {
			if user.Registration != nil {
				return user.Registration.Occupation
			}
			return ""
		}(user),

		//Institution
		func(user models.User) string {
			if user.StudentDetails != nil {
				return user.StudentDetails.UniName
			}
			if user.ProfessionalDetails != nil {
				return user.ProfessionalDetails.OrganizationName
			}
			return ""
		}(user),

		//Designation
		func(user models.User) string {
			if user.ProfessionalDetails != nil {
				return user.ProfessionalDetails.Designation
			}
			return ""
		}(user),

		//Exp
		func(user models.User) string {
			if user.ProfessionalDetails != nil {
				return user.ProfessionalDetails.YearsOfExps
			}
			return ""
		}(user),

		//Tech
		func(user models.User) string {
			if user.ProfessionalDetails != nil {
				return user.ProfessionalDetails.TechStack
			}
			return ""
		}(user),

		//Competition
		func(user models.User) string {
			if user.Registration != nil {
				return user.Registration.Competition
			}
			return ""
		}(user),
		//Bringing Laptop
		func(user models.User) string {
			if user.Registration != nil {
				return fmt.Sprintf("%t", user.Registration.IsBringingLaptop)
			}
			return "false"
		}(user),
	}
}
