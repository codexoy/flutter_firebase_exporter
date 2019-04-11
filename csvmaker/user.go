package csvmaker

import (
	"encoding/csv"
	"io"

	"github.com/minhajuddinkhan/flutter_pk_firebase_export/models"
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

func getVal(user models.User) ([]string, []string, error) {
	keys := []string{}
	values := []string{}
	return keys, values, GetFields(user, &keys, &values)
}

func (m *csvmaker) User(users []models.User) error {
	records := [][]string{}
	var keys []string
	for _, user := range users {
		k, values, err := getVal(user)
		keys = k
		if err != nil {
			return err
		}
		records = append(records, values[:])
	}
	records = append([][]string{keys}, records...)
	return m.Writer.WriteAll(records)
}
