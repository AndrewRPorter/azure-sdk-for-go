//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/NameSpaces/PrivateEndPointConnectionList.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewPrivateEndpointConnectionsClient("subID", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("SDK-ServiceBus-4794", "sdk-Namespace-5828", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PrivateEndpointConnectionListResult = armservicebus.PrivateEndpointConnectionListResult{
		// 	Value: []*armservicebus.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("5dc668b3-70e4-437f-b61c-a3c1e594be7a"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces/PrivateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-ServiceBus-7182/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-5705-new/privateEndpointConnections/5dc668b3-70e4-437f-b61c-a3c1e594be7a"),
		// 			Properties: &armservicebus.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armservicebus.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-ServiceBus-7182/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-5705-new"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armservicebus.ConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					Status: to.Ptr(armservicebus.PrivateLinkConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armservicebus.EndPointProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/NameSpaces/PrivateEndPointConnectionCreate.json
func ExamplePrivateEndpointConnectionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewPrivateEndpointConnectionsClient("subID", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "ArunMonocle", "sdk-Namespace-2924", "privateEndpointConnectionName", armservicebus.PrivateEndpointConnection{
		Properties: &armservicebus.PrivateEndpointConnectionProperties{
			PrivateEndpoint: &armservicebus.PrivateEndpoint{
				ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-ServiceBus-8396/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-2847"),
			},
			PrivateLinkServiceConnectionState: &armservicebus.ConnectionState{
				Description: to.Ptr("testing"),
				Status:      to.Ptr(armservicebus.PrivateLinkConnectionStatusRejected),
			},
			ProvisioningState: to.Ptr(armservicebus.EndPointProvisioningStateSucceeded),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armservicebus.PrivateEndpointConnection{
	// 	Name: to.Ptr("928c44d5-b7c6-423b-b6fa-811e0c27b3e0"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/PrivateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-ServiceBus-4794/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-5828/privateEndpointConnections/928c44d5-b7c6-423b-b6fa-811e0c27b3e0"),
	// 	Properties: &armservicebus.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armservicebus.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-ServiceBus-4794/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-5828"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armservicebus.ConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			Status: to.Ptr(armservicebus.PrivateLinkConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armservicebus.EndPointProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/NameSpaces/PrivateEndPointConnectionDelete.json
func ExamplePrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewPrivateEndpointConnectionsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "ArunMonocle", "sdk-Namespace-3285", "928c44d5-b7c6-423b-b6fa-811e0c27b3e0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/NameSpaces/PrivateEndPointConnectionGet.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewPrivateEndpointConnectionsClient("subID", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "SDK-ServiceBus-4794", "sdk-Namespace-5828", "privateEndpointConnectionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armservicebus.PrivateEndpointConnection{
	// 	Name: to.Ptr("privateEndpointConnectionName"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/PrivateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-ServiceBus-4794/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-5828/privateEndpointConnections/privateEndpointConnectionName"),
	// 	Properties: &armservicebus.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armservicebus.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/dbedb4e0-40e6-4145-81f3-f1314c150774/resourceGroups/SDK-ServiceBus-4794/providers/Microsoft.Network/privateEndpoints/sdk-Namespace-5828"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armservicebus.ConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			Status: to.Ptr(armservicebus.PrivateLinkConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armservicebus.EndPointProvisioningStateSucceeded),
	// 	},
	// }
}
