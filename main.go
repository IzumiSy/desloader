package main

import (
	"context"
	"desloader/schema"
	"fmt"
	"io/ioutil"

	flags "github.com/jessevdk/go-flags"
	"golang.org/x/oauth2/google"
)

var opts struct {
	ProjectID  string `long:"projectId" description:"Target GCP project ID"`
	SchemaPath string `long:"schema" description:"A schema definition in YAML format" required:"true"`
	SourcePath string `long:"source" description:"A source file path in CSV format" required:"true"`
	Kind       string `long:"kind" description:"Target kind on Datastore" required:"true"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		return
	}

	ctx := context.Background()
	projectID := opts.ProjectID

	if projectID == "" {
		credentials, err := google.FindDefaultCredentials(ctx)
		if err != nil {
			panic(err)
		}
		if credentials == nil {
			fmt.Println("Error: invalid credential")
			return
		}
		projectID = credentials.ProjectID

		if projectID == "" {
			fmt.Println("Error: Failed retrieving project ID from gcloud credentials.")
			fmt.Println("Could you manually provide your project ID with --projectID option?")
			return
		}
	}

	schemaBytes, err := ioutil.ReadFile(opts.SchemaPath)
	if err != nil {
		fmt.Printf("Failed reading schema: %s", err.Error())
		return
	}

	s, err := schema.New(schemaBytes)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	fmt.Printf("%v", s)

	return
}
