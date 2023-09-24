package gcp

import (
	"context"
	"fmt"
	// "io"
	"strings"

	iam "google.golang.org/api/iam/v1"
)

func CheckServiceAccountExists(projectID string, serviceAccountName string) (bool, error) {
	ctx := context.Background()
	var check bool
	var err error

	service, err := iam.NewService(ctx)
	if err != nil {
		fmt.Println(err)
		return true, err
	}

	fullServiceAccountName := fmt.Sprintf("projects/" + projectID + "/serviceAccounts/" + serviceAccountName + "@" + projectID + ".iam.gserviceaccount.com")

	_, err = service.Projects.ServiceAccounts.Get(fullServiceAccountName).Do()
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			check = false
			fmt.Println("projects/" + projectID + "/serviceAccounts/" + serviceAccountName + " not found")
			return check, nil
		} else {
			fmt.Println(err)
			return true, err
		}
	} else {
		check = true
	}

	return check, err
}

func CreateServiceAccount(projectID string, serviceAccountName string) (*iam.ServiceAccount, error) {
	ctx := context.Background()
	service, err := iam.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("iam.NewService: %w", err)
	}

	request := &iam.CreateServiceAccountRequest{
		AccountId: serviceAccountName,
		ServiceAccount: &iam.ServiceAccount{
			DisplayName: "Created by golang for " + serviceAccountName,
		},
	}

	account, err := service.Projects.ServiceAccounts.Create("projects/"+projectID, request).Do()
	if err != nil {
		return nil, fmt.Errorf("creation failed with: %w", err)
	}
	return account, nil

}

func DeleteServiceAccount(projectID string, serviceAccountName string) error {
	ctx := context.Background()
	service, err := iam.NewService(ctx)
	if err != nil {
		return fmt.Errorf("iam.NewService: %w", err)
	}

	fullServiceAccountName := fmt.Sprintf("projects/" + projectID + "/serviceAccounts/" + serviceAccountName + "@" + projectID + ".iam.gserviceaccount.com")

	_, err = service.Projects.ServiceAccounts.Delete(fullServiceAccountName).Do()
	if err != nil {
		return fmt.Errorf("Projects.ServiceAccounts.Delete: %w", err)
	}
	fmt.Printf("Deleted service account: %v\n", serviceAccountName)
	return nil

}
