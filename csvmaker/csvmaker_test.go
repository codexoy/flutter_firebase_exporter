package csvmaker

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/minhajuddinkhan/flutter_pk_firebase_export/models"
)

func TestCSVMaker(t *testing.T) {

	user := models.User{
		Name:         "Ali",
		Email:        "hello@gmail.com",
		MobileNumber: "123342",
		Registration: &models.Registration{
			Occupation:       "occup",
			Competition:      "code",
			IsBringingLaptop: true,
		},
		ProfessionalDetails: &models.Professional{
			OrganizationName: "orgName",
			Designation:      "dev",
			YearsOfExps:      "5",
			TechStack:        "go",
		},
	}
	//"Ali,hello@gmail.com,123342,occup,orgName,dev,5,go,code,true
	expected := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%t",
		user.Name,
		user.Email,
		user.MobileNumber,
		user.Registration.Occupation,
		user.ProfessionalDetails.OrganizationName,
		user.ProfessionalDetails.Designation,
		user.ProfessionalDetails.YearsOfExps,
		user.ProfessionalDetails.TechStack,
		user.Registration.Competition,
		user.Registration.IsBringingLaptop,
	)
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	maker := New(writer)
	err := maker.User([]models.User{user})
	assert.Nil(t, err)

	var outputRows []string
	for _, s := range strings.Split(string(b.Bytes()), "\n") {
		if s == "" {
			continue
		}
		outputRows = append(outputRows, s)
	}
	assert.Greater(t, len(outputRows), 1)
	for i := 1; i < len(outputRows); i++ {
		assert.Equal(t, expected, outputRows[i])
	}
}
