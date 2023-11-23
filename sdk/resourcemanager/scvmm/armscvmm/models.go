//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscvmm

import "time"

// AvailabilitySet - The AvailabilitySets resource definition.
type AvailabilitySet struct {
	// REQUIRED; The extended location.
	ExtendedLocation *ExtendedLocation

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Resource properties.
	Properties *AvailabilitySetProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// AvailabilitySetListItem - Availability Set model
type AvailabilitySetListItem struct {
	// Gets the ARM Id of the microsoft.scvmm/availabilitySets resource.
	ID *string

	// Gets or sets the name of the availability set.
	Name *string
}

// AvailabilitySetListResult - List of AvailabilitySets.
type AvailabilitySetListResult struct {
	// List of AvailabilitySets.
	Value []*AvailabilitySet

	// READ-ONLY; Url to follow for getting next page of resources.
	NextLink *string
}

// AvailabilitySetProperties - Defines the resource properties.
type AvailabilitySetProperties struct {
	// Name of the availability set.
	AvailabilitySetName *string

	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerID *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// Checkpoint - Defines the resource properties.
type Checkpoint struct {
	// Gets ID of the checkpoint.
	CheckpointID *string

	// Gets description of the checkpoint.
	Description *string

	// Gets name of the checkpoint.
	Name *string

	// Gets ID of parent of the checkpoint.
	ParentCheckpointID *string
}

// Cloud - The Clouds resource definition.
type Cloud struct {
	// REQUIRED; The extended location.
	ExtendedLocation *ExtendedLocation

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Resource properties.
	Properties *CloudProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// CloudCapacity - Cloud Capacity model
type CloudCapacity struct {
	// CPUCount specifies the maximum number of CPUs that can be allocated in the cloud.
	CPUCount *int64

	// MemoryMB specifies a memory usage limit in megabytes.
	MemoryMB *int64

	// VMCount gives the max number of VMs that can be deployed in the cloud.
	VMCount *int64
}

// CloudInventoryItem - The Cloud inventory item.
type CloudInventoryItem struct {
	// REQUIRED; They inventory type.
	InventoryType *InventoryType

	// READ-ONLY; Gets the Managed Object name in VMM for the inventory item.
	InventoryItemName *string

	// READ-ONLY; Gets the tracked resource id corresponding to the inventory resource.
	ManagedResourceID *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState

	// READ-ONLY; Gets the UUID (which is assigned by VMM) for the inventory item.
	UUID *string
}

// GetInventoryItemProperties implements the InventoryItemPropertiesClassification interface for type CloudInventoryItem.
func (c *CloudInventoryItem) GetInventoryItemProperties() *InventoryItemProperties {
	return &InventoryItemProperties{
		InventoryItemName: c.InventoryItemName,
		InventoryType:     c.InventoryType,
		ManagedResourceID: c.ManagedResourceID,
		ProvisioningState: c.ProvisioningState,
		UUID:              c.UUID,
	}
}

// CloudListResult - List of Clouds.
type CloudListResult struct {
	// List of Clouds.
	Value []*Cloud

	// READ-ONLY; Url to follow for getting next page of resources.
	NextLink *string
}

// CloudProperties - Defines the resource properties.
type CloudProperties struct {
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemID *string

	// Unique ID of the cloud.
	UUID *string

	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerID *string

	// READ-ONLY; Capacity of the cloud.
	CloudCapacity *CloudCapacity

	// READ-ONLY; Name of the cloud in VMMServer.
	CloudName *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState

	// READ-ONLY; List of QoS policies available for the cloud.
	StorageQoSPolicies []*StorageQoSPolicy
}

// ExtendedLocation - The extended location.
type ExtendedLocation struct {
	// The extended location name.
	Name *string

	// The extended location type.
	Type *string
}

// GuestAgent - Defines the GuestAgent.
type GuestAgent struct {
	// REQUIRED; Resource properties.
	Properties *GuestAgentProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// GuestAgentList - List of GuestAgent.
type GuestAgentList struct {
	// REQUIRED; Array of GuestAgent
	Value []*GuestAgent

	// READ-ONLY; Url to follow for getting next page of GuestAgent.
	NextLink *string
}

// GuestAgentProperties - Defines the resource properties.
type GuestAgentProperties struct {
	// Username / Password Credentials to provision guest agent.
	Credentials *GuestCredential

	// HTTP Proxy configuration for the VM.
	HTTPProxyConfig *HTTPProxyConfiguration

	// Gets or sets the guest agent provisioning action.
	ProvisioningAction *ProvisioningAction

	// READ-ONLY; Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState

	// READ-ONLY; Gets the guest agent status.
	Status *string

	// READ-ONLY; Gets a unique identifier for this resource.
	UUID *string
}

// GuestCredential - Username / Password Credentials to connect to guest.
type GuestCredential struct {
	// REQUIRED; Gets or sets the password to connect with the guest.
	Password *string

	// REQUIRED; Gets or sets username to connect with the guest.
	Username *string
}

// HTTPProxyConfiguration - HTTP Proxy configuration for the VM.
type HTTPProxyConfiguration struct {
	// Gets or sets httpsProxy url.
	HTTPSProxy *string
}

// HardwareProfile - Defines the resource properties.
type HardwareProfile struct {
	// Gets or sets the number of vCPUs for the vm.
	CPUCount *int32

	// Gets or sets a value indicating whether to enable dynamic memory or not.
	DynamicMemoryEnabled *DynamicMemoryEnabled

	// Gets or sets the max dynamic memory for the vm.
	DynamicMemoryMaxMB *int32

	// Gets or sets the min dynamic memory for the vm.
	DynamicMemoryMinMB *int32

	// Gets or sets a value indicating whether to enable processor compatibility mode for live migration of VMs.
	LimitCPUForMigration *LimitCPUForMigration

	// MemoryMB is the size of a virtual machine's memory, in MB.
	MemoryMB *int32

	// READ-ONLY; Gets highly available property.
	IsHighlyAvailable *IsHighlyAvailable
}

// HardwareProfileUpdate - Defines the resource properties.
type HardwareProfileUpdate struct {
	// Gets or sets the number of vCPUs for the vm.
	CPUCount *int32

	// Gets or sets a value indicating whether to enable dynamic memory or not.
	DynamicMemoryEnabled *DynamicMemoryEnabled

	// Gets or sets the max dynamic memory for the vm.
	DynamicMemoryMaxMB *int32

	// Gets or sets the min dynamic memory for the vm.
	DynamicMemoryMinMB *int32

	// Gets or sets a value indicating whether to enable processor compatibility mode for live migration of VMs.
	LimitCPUForMigration *LimitCPUForMigration

	// MemoryMB is the size of a virtual machine's memory, in MB.
	MemoryMB *int32
}

// InfrastructureProfile - Specifies the vmmServer infrastructure specific settings for the virtual machine instance.
type InfrastructureProfile struct {
	// Gets or sets the bios guid for the vm.
	BiosGUID *string

	// Type of checkpoint supported for the vm.
	CheckpointType *string

	// Checkpoints in the vm.
	Checkpoints []*Checkpoint

	// ARM Id of the cloud resource to use for deploying the vm.
	CloudID *string

	// Gets or sets the generation for the vm.
	Generation *int32

	// Gets or sets the inventory Item ID for the resource.
	InventoryItemID *string

	// ARM Id of the template resource to use for deploying the vm.
	TemplateID *string

	// Unique ID of the virtual machine.
	UUID *string

	// VMName is the name of VM on the SCVMM server.
	VMName *string

	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerID *string

	// READ-ONLY; Last restored checkpoint in the vm.
	LastRestoredVMCheckpoint *Checkpoint
}

// InfrastructureProfileUpdate - Specifies the vmmServer infrastructure specific settings for the virtual machine instance
// for update.
type InfrastructureProfileUpdate struct {
	// Type of checkpoint supported for the vm.
	CheckpointType *string
}

// InventoryItem - Defines the inventory item.
type InventoryItem struct {
	// REQUIRED; Resource properties.
	Properties InventoryItemPropertiesClassification

	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are
	// a kind of Microsoft.Web/sites type. If supported, the resource provider must
	// validate and persist this value.
	Kind *string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// InventoryItemDetails - Defines the resource properties.
type InventoryItemDetails struct {
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemID *string

	// Gets or sets the Managed Object name in VMM for the resource.
	InventoryItemName *string
}

// InventoryItemProperties - Defines the resource properties.
type InventoryItemProperties struct {
	// REQUIRED; They inventory type.
	InventoryType *InventoryType

	// READ-ONLY; Gets the Managed Object name in VMM for the inventory item.
	InventoryItemName *string

	// READ-ONLY; Gets the tracked resource id corresponding to the inventory resource.
	ManagedResourceID *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState

	// READ-ONLY; Gets the UUID (which is assigned by VMM) for the inventory item.
	UUID *string
}

// GetInventoryItemProperties implements the InventoryItemPropertiesClassification interface for type InventoryItemProperties.
func (i *InventoryItemProperties) GetInventoryItemProperties() *InventoryItemProperties { return i }

// InventoryItemsList - List of InventoryItems.
type InventoryItemsList struct {
	// REQUIRED; Array of InventoryItems
	Value []*InventoryItem

	// READ-ONLY; Url to follow for getting next page of InventoryItems.
	NextLink *string
}

// NetworkInterface - Network Interface model
type NetworkInterface struct {
	// Gets or sets the ipv4 address type.
	IPv4AddressType *AllocationMethod

	// Gets or sets the ipv6 address type.
	IPv6AddressType *AllocationMethod

	// Gets or sets the nic MAC address.
	MacAddress *string

	// Gets or sets the mac address type.
	MacAddressType *AllocationMethod

	// Gets or sets the name of the network interface.
	Name *string

	// Gets or sets the nic id.
	NicID *string

	// Gets or sets the ARM Id of the Microsoft.ScVmm/virtualNetwork resource to connect the nic.
	VirtualNetworkID *string

	// READ-ONLY; Gets the display name of the network interface as shown in the vmmServer. This is the fallback label for a NIC
	// when the name is not set.
	DisplayName *string

	// READ-ONLY; Gets the nic ipv4 addresses.
	IPv4Addresses []*string

	// READ-ONLY; Gets the nic ipv6 addresses.
	IPv6Addresses []*string

	// READ-ONLY; Gets the name of the virtual network in vmmServer that the nic is connected to.
	NetworkName *string
}

// NetworkInterfaceUpdate - Network Interface model
type NetworkInterfaceUpdate struct {
	// Gets or sets the ipv4 address type.
	IPv4AddressType *AllocationMethod

	// Gets or sets the ipv6 address type.
	IPv6AddressType *AllocationMethod

	// Gets or sets the nic MAC address.
	MacAddress *string

	// Gets or sets the mac address type.
	MacAddressType *AllocationMethod

	// Gets or sets the name of the network interface.
	Name *string

	// Gets or sets the nic id.
	NicID *string

	// Gets or sets the ARM Id of the Microsoft.ScVmm/virtualNetwork resource to connect the nic.
	VirtualNetworkID *string
}

// NetworkProfile - Defines the resource properties.
type NetworkProfile struct {
	// Gets or sets the list of network interfaces associated with the virtual machine.
	NetworkInterfaces []*NetworkInterface
}

// NetworkProfileUpdate - Defines the resource properties.
type NetworkProfileUpdate struct {
	// Gets or sets the list of network interfaces associated with the virtual machine.
	NetworkInterfaces []*NetworkInterfaceUpdate
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// OsProfileForVMInstance - Defines the resource properties.
type OsProfileForVMInstance struct {
	// Admin password of the virtual machine.
	AdminPassword *string

	// Gets or sets computer name.
	ComputerName *string

	// READ-ONLY; Gets os sku.
	OSSKU *string

	// READ-ONLY; Gets the type of the os.
	OSType *OsType

	// READ-ONLY; Gets os version.
	OSVersion *string
}

// ResourcePatch - Object containing tags updates for patch operations.
type ResourcePatch struct {
	// Resource tags.
	Tags map[string]*string
}

// StopVirtualMachineOptions - Defines the stop action properties.
type StopVirtualMachineOptions struct {
	// Gets or sets a value indicating whether to request non-graceful VM shutdown. True value for this flag indicates non-graceful
	// shutdown whereas false indicates otherwise. Defaults to false.
	SkipShutdown *SkipShutdown
}

// StorageProfile - Defines the resource properties.
type StorageProfile struct {
	// Gets or sets the list of virtual disks associated with the virtual machine.
	Disks []*VirtualDisk
}

// StorageProfileUpdate - Defines the resource properties.
type StorageProfileUpdate struct {
	// Gets or sets the list of virtual disks associated with the virtual machine.
	Disks []*VirtualDiskUpdate
}

// StorageQoSPolicy - The StorageQoSPolicy definition.
type StorageQoSPolicy struct {
	// The Bandwidth Limit for internet traffic.
	BandwidthLimit *int64

	// The ID of the QoS policy.
	ID *string

	// The maximum IO operations per second.
	IopsMaximum *int64

	// The minimum IO operations per second.
	IopsMinimum *int64

	// The name of the policy.
	Name *string

	// The underlying policy.
	PolicyID *string
}

// StorageQoSPolicyDetails - The StorageQoSPolicyDetails definition.
type StorageQoSPolicyDetails struct {
	// The ID of the QoS policy.
	ID *string

	// The name of the policy.
	Name *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// VMInstanceHybridIdentityMetadata - Defines the HybridIdentityMetadata.
type VMInstanceHybridIdentityMetadata struct {
	// REQUIRED; Resource properties.
	Properties *VMInstanceHybridIdentityMetadataProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// VMInstanceHybridIdentityMetadataList - List of HybridIdentityMetadata.
type VMInstanceHybridIdentityMetadataList struct {
	// REQUIRED; Array of HybridIdentityMetadata
	Value []*VMInstanceHybridIdentityMetadata

	// READ-ONLY; Url to follow for getting next page of HybridIdentityMetadata.
	NextLink *string
}

// VMInstanceHybridIdentityMetadataProperties - Describes the properties of Hybrid Identity Metadata for a Virtual Machine.
type VMInstanceHybridIdentityMetadataProperties struct {
	// Gets or sets the Public Key.
	PublicKey *string

	// The unique identifier for the resource.
	ResourceUID *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// VMMCredential - Credentials to connect to VMMServer.
type VMMCredential struct {
	// Password to use to connect to VMMServer.
	Password *string

	// Username to use to connect to VMMServer.
	Username *string
}

// VMMServer - The VmmServers resource definition.
type VMMServer struct {
	// REQUIRED; The extended location.
	ExtendedLocation *ExtendedLocation

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Resource properties.
	Properties *VMMServerProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// VMMServerListResult - List of VmmServers.
type VMMServerListResult struct {
	// List of VmmServers.
	Value []*VMMServer

	// READ-ONLY; Url to follow for getting next page of resources.
	NextLink *string
}

// VMMServerProperties - Defines the resource properties.
type VMMServerProperties struct {
	// REQUIRED; Fqdn is the hostname/ip of the vmmServer.
	Fqdn *string

	// Credentials to connect to VMMServer.
	Credentials *VMMCredential

	// Port is the port on which the vmmServer is listening.
	Port *int32

	// READ-ONLY; Gets the connection status to the vmmServer.
	ConnectionStatus *string

	// READ-ONLY; Gets any error message if connection to vmmServer is having any issue.
	ErrorMessage *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState

	// READ-ONLY; Unique ID of vmmServer.
	UUID *string

	// READ-ONLY; Version is the version of the vmmSever.
	Version *string
}

// VirtualDisk - Virtual disk model
type VirtualDisk struct {
	// Gets or sets the disk bus.
	Bus *int32

	// Gets or sets the disk bus type.
	BusType *string

	// Gets or sets a value indicating diff disk.
	CreateDiffDisk *CreateDiffDisk

	// Gets or sets the disk id.
	DiskID *string

	// Gets or sets the disk total size.
	DiskSizeGB *int32

	// Gets or sets the disk lun.
	Lun *int32

	// Gets or sets the name of the disk.
	Name *string

	// The QoS policy for the disk.
	StorageQoSPolicy *StorageQoSPolicyDetails

	// Gets or sets the disk id in the template.
	TemplateDiskID *string

	// Gets or sets the disk vhd type.
	VhdType *string

	// READ-ONLY; Gets the display name of the virtual disk as shown in the vmmServer. This is the fallback label for a disk when
	// the name is not set.
	DisplayName *string

	// READ-ONLY; Gets the max disk size.
	MaxDiskSizeGB *int32

	// READ-ONLY; Gets the disk vhd format type.
	VhdFormatType *string

	// READ-ONLY; Gets the disk volume type.
	VolumeType *string
}

// VirtualDiskUpdate - Virtual disk model
type VirtualDiskUpdate struct {
	// Gets or sets the disk bus.
	Bus *int32

	// Gets or sets the disk bus type.
	BusType *string

	// Gets or sets the disk id.
	DiskID *string

	// Gets or sets the disk total size.
	DiskSizeGB *int32

	// Gets or sets the disk lun.
	Lun *int32

	// Gets or sets the name of the disk.
	Name *string

	// The QoS policy for the disk.
	StorageQoSPolicy *StorageQoSPolicyDetails

	// Gets or sets the disk vhd type.
	VhdType *string
}

// VirtualMachineCreateCheckpoint - Defines the create checkpoint action properties.
type VirtualMachineCreateCheckpoint struct {
	// Description of the checkpoint.
	Description *string

	// Name of the checkpoint.
	Name *string
}

// VirtualMachineDeleteCheckpoint - Defines the delete checkpoint action properties.
type VirtualMachineDeleteCheckpoint struct {
	// ID of the checkpoint to be deleted.
	ID *string
}

// VirtualMachineInstance - Define the virtualMachineInstance.
type VirtualMachineInstance struct {
	// REQUIRED; Gets or sets the extended location.
	ExtendedLocation *ExtendedLocation

	// REQUIRED; Resource properties.
	Properties *VirtualMachineInstanceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// VirtualMachineInstanceListResult - List of VirtualMachineInstances.
type VirtualMachineInstanceListResult struct {
	// Array of VirtualMachineInstances.
	Value []*VirtualMachineInstance

	// READ-ONLY; Url to follow for getting next page of resources.
	NextLink *string
}

// VirtualMachineInstanceProperties - Defines the resource properties.
type VirtualMachineInstanceProperties struct {
	// Availability Sets in vm.
	AvailabilitySets []*AvailabilitySetListItem

	// Hardware properties.
	HardwareProfile *HardwareProfile

	// Gets the infrastructure profile.
	InfrastructureProfile *InfrastructureProfile

	// Network properties.
	NetworkProfile *NetworkProfile

	// OS properties.
	OSProfile *OsProfileForVMInstance

	// Storage properties.
	StorageProfile *StorageProfile

	// READ-ONLY; Gets the power state of the virtual machine.
	PowerState *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// VirtualMachineInstanceUpdate - Defines the virtualMachineInstanceUpdate.
type VirtualMachineInstanceUpdate struct {
	// Resource properties.
	Properties *VirtualMachineInstanceUpdateProperties
}

// VirtualMachineInstanceUpdateProperties - Defines the resource properties.
type VirtualMachineInstanceUpdateProperties struct {
	// Availability Sets in vm.
	AvailabilitySets []*AvailabilitySetListItem

	// Defines the resource properties.
	HardwareProfile *HardwareProfileUpdate

	// Gets the infrastructure profile.
	InfrastructureProfile *InfrastructureProfileUpdate

	// Defines the resource properties.
	NetworkProfile *NetworkProfileUpdate

	// Defines the resource properties.
	StorageProfile *StorageProfileUpdate
}

// VirtualMachineInventoryItem - The Virtual machine inventory item.
type VirtualMachineInventoryItem struct {
	// REQUIRED; They inventory type.
	InventoryType *InventoryType

	// Cloud inventory resource details where the VM is present.
	Cloud *InventoryItemDetails

	// Gets or sets the nic ip addresses.
	IPAddresses []*string

	// READ-ONLY; Gets the bios guid.
	BiosGUID *string

	// READ-ONLY; Gets the Managed Object name in VMM for the inventory item.
	InventoryItemName *string

	// READ-ONLY; Gets the tracked resource id corresponding to the inventory resource.
	ManagedMachineResourceID *string

	// READ-ONLY; Gets the tracked resource id corresponding to the inventory resource.
	ManagedResourceID *string

	// READ-ONLY; Gets os name.
	OSName *string

	// READ-ONLY; Gets the type of the os.
	OSType *OsType

	// READ-ONLY; Gets os version.
	OSVersion *string

	// READ-ONLY; Gets the power state of the virtual machine.
	PowerState *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState

	// READ-ONLY; Gets the UUID (which is assigned by VMM) for the inventory item.
	UUID *string
}

// GetInventoryItemProperties implements the InventoryItemPropertiesClassification interface for type VirtualMachineInventoryItem.
func (v *VirtualMachineInventoryItem) GetInventoryItemProperties() *InventoryItemProperties {
	return &InventoryItemProperties{
		InventoryItemName: v.InventoryItemName,
		InventoryType:     v.InventoryType,
		ManagedResourceID: v.ManagedResourceID,
		ProvisioningState: v.ProvisioningState,
		UUID:              v.UUID,
	}
}

// VirtualMachineRestoreCheckpoint - Defines the restore checkpoint action properties.
type VirtualMachineRestoreCheckpoint struct {
	// ID of the checkpoint to be restored to.
	ID *string
}

// VirtualMachineTemplate - The VirtualMachineTemplates resource definition.
type VirtualMachineTemplate struct {
	// REQUIRED; The extended location.
	ExtendedLocation *ExtendedLocation

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Resource properties.
	Properties *VirtualMachineTemplateProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// VirtualMachineTemplateInventoryItem - The Virtual machine template inventory item.
type VirtualMachineTemplateInventoryItem struct {
	// REQUIRED; They inventory type.
	InventoryType *InventoryType

	// READ-ONLY; Gets the desired number of vCPUs for the vm.
	CPUCount *int32

	// READ-ONLY; Gets the Managed Object name in VMM for the inventory item.
	InventoryItemName *string

	// READ-ONLY; Gets the tracked resource id corresponding to the inventory resource.
	ManagedResourceID *string

	// READ-ONLY; MemoryMB is the desired size of a virtual machine's memory, in MB.
	MemoryMB *int32

	// READ-ONLY; Gets os name.
	OSName *string

	// READ-ONLY; Gets the type of the os.
	OSType *OsType

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState

	// READ-ONLY; Gets the UUID (which is assigned by VMM) for the inventory item.
	UUID *string
}

// GetInventoryItemProperties implements the InventoryItemPropertiesClassification interface for type VirtualMachineTemplateInventoryItem.
func (v *VirtualMachineTemplateInventoryItem) GetInventoryItemProperties() *InventoryItemProperties {
	return &InventoryItemProperties{
		InventoryItemName: v.InventoryItemName,
		InventoryType:     v.InventoryType,
		ManagedResourceID: v.ManagedResourceID,
		ProvisioningState: v.ProvisioningState,
		UUID:              v.UUID,
	}
}

// VirtualMachineTemplateListResult - List of VirtualMachineTemplates.
type VirtualMachineTemplateListResult struct {
	// List of VirtualMachineTemplates.
	Value []*VirtualMachineTemplate

	// READ-ONLY; Url to follow for getting next page of resources.
	NextLink *string
}

// VirtualMachineTemplateProperties - Defines the resource properties.
type VirtualMachineTemplateProperties struct {
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemID *string

	// Unique ID of the virtual machine template.
	UUID *string

	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerID *string

	// READ-ONLY; Gets the desired number of vCPUs for the vm.
	CPUCount *int32

	// READ-ONLY; Gets computer name.
	ComputerName *string

	// READ-ONLY; Gets the disks of the template.
	Disks []*VirtualDisk

	// READ-ONLY; Gets a value indicating whether to enable dynamic memory or not.
	DynamicMemoryEnabled *DynamicMemoryEnabled

	// READ-ONLY; Gets the max dynamic memory for the vm.
	DynamicMemoryMaxMB *int32

	// READ-ONLY; Gets the min dynamic memory for the vm.
	DynamicMemoryMinMB *int32

	// READ-ONLY; Gets the generation for the vm.
	Generation *int32

	// READ-ONLY; Gets a value indicating whether the vm template is customizable or not.
	IsCustomizable *IsCustomizable

	// READ-ONLY; Gets highly available property.
	IsHighlyAvailable *IsHighlyAvailable

	// READ-ONLY; Gets a value indicating whether to enable processor compatibility mode for live migration of VMs.
	LimitCPUForMigration *LimitCPUForMigration

	// READ-ONLY; MemoryMB is the desired size of a virtual machine's memory, in MB.
	MemoryMB *int32

	// READ-ONLY; Gets the network interfaces of the template.
	NetworkInterfaces []*NetworkInterface

	// READ-ONLY; Gets os name.
	OSName *string

	// READ-ONLY; Gets the type of the os.
	OSType *OsType

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState
}

// VirtualNetwork - The VirtualNetworks resource definition.
type VirtualNetwork struct {
	// REQUIRED; The extended location.
	ExtendedLocation *ExtendedLocation

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Resource properties.
	Properties *VirtualNetworkProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// VirtualNetworkInventoryItem - The Virtual network inventory item.
type VirtualNetworkInventoryItem struct {
	// REQUIRED; They inventory type.
	InventoryType *InventoryType

	// READ-ONLY; Gets the Managed Object name in VMM for the inventory item.
	InventoryItemName *string

	// READ-ONLY; Gets the tracked resource id corresponding to the inventory resource.
	ManagedResourceID *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState

	// READ-ONLY; Gets the UUID (which is assigned by VMM) for the inventory item.
	UUID *string
}

// GetInventoryItemProperties implements the InventoryItemPropertiesClassification interface for type VirtualNetworkInventoryItem.
func (v *VirtualNetworkInventoryItem) GetInventoryItemProperties() *InventoryItemProperties {
	return &InventoryItemProperties{
		InventoryItemName: v.InventoryItemName,
		InventoryType:     v.InventoryType,
		ManagedResourceID: v.ManagedResourceID,
		ProvisioningState: v.ProvisioningState,
		UUID:              v.UUID,
	}
}

// VirtualNetworkListResult - List of VirtualNetworks.
type VirtualNetworkListResult struct {
	// List of VirtualNetworks.
	Value []*VirtualNetwork

	// READ-ONLY; Url to follow for getting next page of resources.
	NextLink *string
}

// VirtualNetworkProperties - Defines the resource properties.
type VirtualNetworkProperties struct {
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemID *string

	// Unique ID of the virtual network.
	UUID *string

	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerID *string

	// READ-ONLY; Name of the virtual network in vmmServer.
	NetworkName *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState
}
