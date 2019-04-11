package csvmaker

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/minhajuddinkhan/flutter_pk_firebase_export/models"
	"github.com/stretchr/testify/assert"
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
	expected := fmt.Sprintf("%s,%s,%s,%t,%t,%t,%t,%s,%s,%s,%s,%t,%s,%s,%s,%s",
		user.Name,
		user.Email,
		user.MobileNumber,
		user.IsContributor,
		user.IsRegistered,
		user.IsPresent,
		user.IsRegistrationConfirmed,
		user.PhotoURL,
		user.Registration.Competition,
		user.Registration.Occupation,
		user.Registration.ReasonToAttend,
		user.Registration.IsBringingLaptop,
		user.ProfessionalDetails.Designation,
		user.ProfessionalDetails.OrganizationName,
		user.ProfessionalDetails.TechStack,
		user.ProfessionalDetails.YearsOfExps,
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

func TestGetFields(t *testing.T) {

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
	keys := []string{}
	values := []string{}
	err := GetFields(user, &keys, &values)
	assert.Nil(t, err)
}

func TestGetFieldsWithIntStructs(t *testing.T) {

	keys := []string{}
	values := []string{}
	x := struct {
		Int   int
		Int8  int8
		Int16 int16
		Int32 int32
		Int64 int64
	}{
		Int:   1,
		Int8:  2,
		Int16: 3,
		Int32: 4,
		Int64: 5,
	}
	err := GetFields(x, &keys, &values)
	assert.Nil(t, err)
	v := reflect.ValueOf(x)
	for i := 0; i < v.NumField(); i++ {
		assert.Equal(t, fmt.Sprintf("%d", v.Field(i).Interface()), values[i])
	}
}

func TestGetFieldsWithSlices(t *testing.T) {
	x := struct {
		Slice []string
	}{
		Slice: []string{"1", "2", "3"},
	}
	keys := []string{}
	values := []string{}
	err := GetFields(x, &keys, &values)
	assert.Nil(t, err)
	v := reflect.ValueOf(x)
	for i := 0; i < v.NumField(); i++ {
		for j := 0; j < (v.Field(i).Len()); j++ {
			assert.Equal(t, v.Field(i).Index(j).Interface().(string), x.Slice[j])
		}
	}

	tp := reflect.TypeOf(x)
	for i := 0; i < tp.NumField(); i++ {
		assert.Equal(t, keys[i], fmt.Sprintf("%s[%d]", tp.Field(i).Name, i))
	}
}
