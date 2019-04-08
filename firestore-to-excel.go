package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.TODO()

	privateKeyPath := "wtq-key.json"
	if len(os.Args) > 1 {
		privateKeyPath := os.Args[1]
		validatePrivateKeyPath(privateKeyPath)
	}

	client := createFirestoreClient(ctx, privateKeyPath)
	defer client.Close()

	writer, file := createCsvWriter("result.csv")
	defer file.Close()
	defer writer.Flush()

	writeDataToFile(writer, client.Collection("users").Documents(ctx))
}

func validatePrivateKeyPath(path string) {
	if path == "" {
		path = "wtq-key.json"
	}

	_, err := os.Stat(path)

	if err != nil {
		log.Fatalln(err)
	}
}

func createFirestoreClient(ctx context.Context, privateKeyPath string) *firestore.Client {
	sa := option.WithCredentialsFile(privateKeyPath)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

func createCsvWriter(path string) (*csv.Writer, *os.File) {
	file, err := os.Create(path)
	checkError("Cannot create file", err)

	writer := csv.NewWriter(file)

	return writer, file
}

func writeDataToFile(writer *csv.Writer, iter *firestore.DocumentIterator) {
	err := writer.Write([]string{"name", "email", "phone", "occupation", "institution", "designation", "experience", "tech", "competition", "bringingLaptop"})
	checkError("Cannot write to file", err)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		data := doc.Data()
		fmt.Print("*** new record **** ")
		fmt.Println(data)

		if !boolValue(data["isRegistered"]) {
			continue
		}

		err = writer.Write([]string{
			stringValue(data["name"]),
			stringValue(data["email"]),
			stringValue(data["mobileNumber"]),
			stringValue(subValue(data["registration"])["occupation"]),
			getInstitution(data),
			getDesignation(data),
			getExperience(data),
			stringValue(subValue(data["professionalDetails"])["techStack"]),
			stringValue(subValue(data["registration"])["competition"]),
			stringValue(data["laptop"]),
		})
		checkError("Cannot write to file", err)
	}
}

func getInstitution(data map[string]interface{}) string {
	return getFieldValue(data, "uniName", "organizationName")
}

func getExperience(data map[string]interface{}) string {
	return getFieldValue(data, "currentYear", "yearsOfExp")
}

func getDesignation(data map[string]interface{}) string {
	return getFieldValue(data, "program", "designation")
}

func getFieldValue(data map[string]interface{}, studentField, professionalField string) string {
	if detail, ok := data["studentDetails"]; ok && detail != nil {
		return stringValue(subValue(detail)[studentField])
	}

	if detail, ok := data["professionalDetails"]; ok && detail != nil {
		return stringValue(subValue(detail)[professionalField])
	}

	return ""
}

func stringValue(value interface{}) string {
	if value != nil {
		return value.(string)
	}
	return ""
}

func boolValue(value interface{}) bool {
	if value != nil {
		return value.(bool)
	}
	return false
}

func subValue(value interface{}) map[string]interface{} {
	if value != nil {
		return value.(map[string]interface{})
	}
	return nil
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
