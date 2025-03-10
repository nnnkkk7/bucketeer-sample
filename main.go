package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer"
	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer/user"
)

func main() {
	// Set up context with timeout
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 1. Set up required information
	apiKey := "YOUR_API_KEY"           // API key created in the admin console
	apiEndpoint := "YOUR_API_ENDPOINT" // API URL provided by the administrator
	featureTag := "YOUR_FEATURE_TAG"   // Tag set when creating the flag
	endUserID := "END_USER_ID"         // ID of the end user

	// 2. Initialize the Bucketeer client
	client, err := bucketeer.NewSDK(
		ctx,
		bucketeer.WithAPIKey(apiKey),
		bucketeer.WithHost(apiEndpoint),
		bucketeer.WithTag(featureTag),
	)
	if err != nil {
		log.Fatalf("Failed to initialize the new client: %v", err)
	}
	defer client.Close(ctx)

	// 3. Create a user object
	// Use an attribute map instead of nil if you need to add attributes
	userObj := user.NewUser(
		endUserID,
		nil, // User attributes are optional
	)

	// 4. Evaluate a boolean flag
	flagID := "YOUR_FEATURE_FLAG_ID" // Flag ID created in the admin console
	defaultValue := false            // Default value if flag is not found or error occurs

	// Execute the flag evaluation
	showNewFeature := client.BoolVariation(ctx, userObj, flagID, defaultValue)

	// Branch processing based on flag value
	if showNewFeature {
		fmt.Println("New feature is enabled - executing new feature code")
		// Implement new feature code here
	} else {
		fmt.Println("New feature is disabled - executing legacy code")
		// Implement legacy code here
	}

	// Examples of other flag types

	// String flag example
	// stringFlagID := "YOUR_STRING_FLAG_ID"
	// stringDefaultValue := "default-value"
	// stringValue := client.StringVariation(ctx, userObj, stringFlagID, stringDefaultValue)
	// fmt.Printf("String flag value: %s\n", stringValue)

	// 5. Wait a bit before application exit to ensure events are sent
	time.Sleep(time.Second)
}
