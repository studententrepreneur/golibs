package utils

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

var Secrets = make(map[string]interface{})

var Ctx context.Context
var ProjectID string

func AccessSecretVersion(name string) ([]byte, error) {

	client, err := secretmanager.NewClient(Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create secretmanager client: %v", err)
	}
	defer client.Close()
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}
	result, err := client.AccessSecretVersion(Ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to access secret version: %v", err)
	}
	Secrets[name] = result.Payload.Data

	return result.Payload.Data, nil
}

func RegisterContext(projectID string, ctx context.Context) {
	ProjectID = projectID
	Ctx = ctx
}
