package gc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

// GetInstanceStatus gets the instance state given the instance id
func GetInstanceStatus(writer http.ResponseWriter, request *http.Request) {
	InitHeaders(writer)
	log.Print("Instance status request invoked!")

	params := mux.Vars(request)
	if len(params) != 0 {
		instanceID := getInstance(params["project"], params["zone"], params["instance"])
		json.NewEncoder(writer).Encode(fmt.Sprintf("Instance status is %#v", instanceID))
	} else {
		log.Fatal("Failed to retrieve query arguments for status")
	}
}

func getInstance(project string, zone string, instance string) uint64 {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := computeService.Instances.Get(project, zone, instance).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	return resp.Id
}

// CreateInstance provisisons new instance in GCP
func CreateInstance(writer http.ResponseWriter, request *http.Request) {
	InitHeaders(writer)
	log.Print("Instance create/provsion requested!")

	params := mux.Vars(request)
	log.Print(fmt.Printf("POST params are %#v", params))
	if len(params) != 0 {
		instanceID := provisionInstance(params["project"], params["zone"], params["username"], params["userpass"])
		json.NewEncoder(writer).Encode(fmt.Sprintf("Instance provisioned and instance ID is %#v", instanceID))
	} else {
		log.Fatal("Failed to provision new instance!")
	}

}

func provisionInstance(project string, zone string, username string, userpass string) uint64 {
	// customize the attributes based on the instance to be provisioned
	var diskSize int64
	region := "us-west1"
	instance := "demo-stack"
	instanceType := "g1-small"
	diskType := "pd-standard"
	image := "projects/ubuntu-os-cloud/global/images/ubuntu-1604-xenial-v20200611"
	diskSize = 10
	accEmail := "gcp service account"
	enAutoRestart := true

}