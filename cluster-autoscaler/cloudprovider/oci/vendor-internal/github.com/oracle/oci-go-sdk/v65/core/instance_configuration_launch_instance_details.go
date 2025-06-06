// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceConfigurationLaunchInstanceDetails Instance launch details for creating an instance from an instance configuration. Use the `sourceDetails`
// parameter to specify whether a boot volume or an image should be used to launch a new instance.
// See LaunchInstanceDetails for more information.
type InstanceConfigurationLaunchInstanceDetails struct {

	// The availability domain of the instance.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID of the compute capacity reservation this instance is launched under.
	CapacityReservationId *string `mandatory:"false" json:"capacityReservationId"`

	// The OCID of the compartment containing the instance.
	// Instances created from instance configurations are placed in the same compartment
	// as the instance that was used to create the instance configuration.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID of the cluster placement group of the instance.
	ClusterPlacementGroupId *string `mandatory:"false" json:"clusterPlacementGroupId"`

	CreateVnicDetails *InstanceConfigurationCreateVnicDetails `mandatory:"false" json:"createVnicDetails"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Security attributes (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels
	// for a resource that can be referenced in a Zero Trust Packet Routing (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm)
	// (ZPR) policy to control access to ZPR-supported resources.
	// Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Additional metadata key/value pairs that you provide. They serve the same purpose and
	// functionality as fields in the `metadata` object.
	// They are distinguished from `metadata` fields in that these can be nested JSON objects
	// (whereas `metadata` fields are string/string maps only).
	// The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of
	// 32,000 bytes.
	ExtendedMetadata map[string]interface{} `mandatory:"false" json:"extendedMetadata"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// This is an advanced option.
	// When a bare metal or virtual machine
	// instance boots, the iPXE firmware that runs on the instance is
	// configured to run an iPXE script to continue the boot process.
	// If you want more control over the boot process, you can provide
	// your own custom iPXE script that will run when the instance boots;
	// however, you should be aware that the same iPXE script will run
	// every time an instance boots; not only after the initial
	// LaunchInstance call.
	// The default iPXE script connects to the instance's local boot
	// volume over iSCSI and performs a network boot. If you use a custom iPXE
	// script and want to network-boot from the instance's local boot volume
	// over iSCSI the same way as the default iPXE script, you should use the
	// following iSCSI IP address: 169.254.0.2, and boot volume IQN:
	// iqn.2015-02.oracle.boot.
	// For more information about the Bring Your Own Image feature of
	// Oracle Cloud Infrastructure, see
	// Bring Your Own Image (https://docs.oracle.com/iaas/Content/Compute/References/bringyourownimage.htm).
	// For more information about iPXE, see http://ipxe.org.
	IpxeScript *string `mandatory:"false" json:"ipxeScript"`

	// Custom metadata key/value pairs that you provide, such as the SSH public key
	// required to connect to the instance.
	// A metadata service runs on every launched instance. The service is an HTTP
	// endpoint listening on 169.254.169.254. You can use the service to:
	// * Provide information to Cloud-Init (https://cloudinit.readthedocs.org/en/latest/)
	//   to be used for various system initialization tasks.
	// * Get information about the instance, including the custom metadata that you
	//   provide when you launch the instance.
	//  **Providing Cloud-Init Metadata**
	//  You can use the following metadata key names to provide information to
	//  Cloud-Init:
	//  **"ssh_authorized_keys"** - Provide one or more public SSH keys to be
	//  included in the `~/.ssh/authorized_keys` file for the default user on the
	//  instance. Use a newline character to separate multiple keys. The SSH
	//  keys must be in the format necessary for the `authorized_keys` file, as shown
	//  in the example below.
	//  **"user_data"** - Provide your own base64-encoded data to be used by
	//  Cloud-Init to run custom scripts or provide custom Cloud-Init configuration. For
	//  information about how to take advantage of user data, see the
	//  Cloud-Init Documentation (http://cloudinit.readthedocs.org/en/latest/topics/format.html).
	//  **Metadata Example**
	//       "metadata" : {
	//          "quake_bot_level" : "Severe",
	//          "ssh_authorized_keys" : "ssh-rsa <your_public_SSH_key>== rsa-key-20160227",
	//          "user_data" : "<your_public_SSH_key>=="
	//       }
	//  **Getting Metadata on the Instance**
	//  To get information about your instance, connect to the instance using SSH and issue any of the
	//  following GET requests:
	//      curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/
	//      curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/
	//      curl -H "Authorization: Bearer Oracle" http://169.254.169.254/opc/v2/instance/metadata/<any-key-name>
	//  You'll get back a response that includes all the instance information; only the metadata information; or
	//  the metadata information for the specified key name, respectively.
	//  The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of 32,000 bytes.
	Metadata map[string]string `mandatory:"false" json:"metadata"`

	// The shape of an instance. The shape determines the number of CPUs, amount of memory,
	// and other resources allocated to the instance.
	// You can enumerate all available shapes by calling ListShapes.
	Shape *string `mandatory:"false" json:"shape"`

	ShapeConfig *InstanceConfigurationLaunchInstanceShapeConfigDetails `mandatory:"false" json:"shapeConfig"`

	PlatformConfig InstanceConfigurationLaunchInstancePlatformConfig `mandatory:"false" json:"platformConfig"`

	SourceDetails InstanceConfigurationInstanceSourceDetails `mandatory:"false" json:"sourceDetails"`

	// A fault domain is a grouping of hardware and infrastructure within an availability domain.
	// Each availability domain contains three fault domains. Fault domains let you distribute your
	// instances so that they are not on the same physical hardware within a single availability domain.
	// A hardware failure or Compute hardware maintenance that affects one fault domain does not affect
	// instances in other fault domains.
	// If you do not specify the fault domain, the system selects one for you.
	//
	// To get a list of fault domains, use the
	// ListFaultDomains operation in the
	// Identity and Access Management Service API.
	// Example: `FAULT-DOMAIN-1`
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// The OCID of the dedicated virtual machine host to place the instance on.
	// Dedicated VM hosts can be used when launching individual instances from an instance configuration. They
	// cannot be used to launch instance pools.
	DedicatedVmHostId *string `mandatory:"false" json:"dedicatedVmHostId"`

	// Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
	// * `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for platform images.
	// * `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
	// * `PARAVIRTUALIZED` - VM instances launch with paravirtualized devices using VirtIO drivers.
	// * `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter.
	LaunchMode InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum `mandatory:"false" json:"launchMode,omitempty"`

	LaunchOptions *InstanceConfigurationLaunchOptions `mandatory:"false" json:"launchOptions"`

	AgentConfig *InstanceConfigurationLaunchInstanceAgentConfigDetails `mandatory:"false" json:"agentConfig"`

	// Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
	IsPvEncryptionInTransitEnabled *bool `mandatory:"false" json:"isPvEncryptionInTransitEnabled"`

	// The preferred maintenance action for an instance. The default is LIVE_MIGRATE, if live migration is supported.
	// * `LIVE_MIGRATE` - Run maintenance using a live migration.
	// * `REBOOT` - Run maintenance using a reboot.
	PreferredMaintenanceAction InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum `mandatory:"false" json:"preferredMaintenanceAction,omitempty"`

	InstanceOptions *InstanceConfigurationInstanceOptions `mandatory:"false" json:"instanceOptions"`

	AvailabilityConfig *InstanceConfigurationAvailabilityConfig `mandatory:"false" json:"availabilityConfig"`

	PreemptibleInstanceConfig *PreemptibleInstanceConfigDetails `mandatory:"false" json:"preemptibleInstanceConfig"`

	// List of licensing configurations associated with target launch values.
	LicensingConfigs []LaunchInstanceLicensingConfig `mandatory:"false" json:"licensingConfigs"`
}

func (m InstanceConfigurationLaunchInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceConfigurationLaunchInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInstanceConfigurationLaunchInstanceDetailsLaunchModeEnum(string(m.LaunchMode)); !ok && m.LaunchMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LaunchMode: %s. Supported values are: %s.", m.LaunchMode, strings.Join(GetInstanceConfigurationLaunchInstanceDetailsLaunchModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum(string(m.PreferredMaintenanceAction)); !ok && m.PreferredMaintenanceAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreferredMaintenanceAction: %s. Supported values are: %s.", m.PreferredMaintenanceAction, strings.Join(GetInstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *InstanceConfigurationLaunchInstanceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		AvailabilityDomain             *string                                                                  `json:"availabilityDomain"`
		CapacityReservationId          *string                                                                  `json:"capacityReservationId"`
		CompartmentId                  *string                                                                  `json:"compartmentId"`
		ClusterPlacementGroupId        *string                                                                  `json:"clusterPlacementGroupId"`
		CreateVnicDetails              *InstanceConfigurationCreateVnicDetails                                  `json:"createVnicDetails"`
		DefinedTags                    map[string]map[string]interface{}                                        `json:"definedTags"`
		SecurityAttributes             map[string]map[string]interface{}                                        `json:"securityAttributes"`
		DisplayName                    *string                                                                  `json:"displayName"`
		ExtendedMetadata               map[string]interface{}                                                   `json:"extendedMetadata"`
		FreeformTags                   map[string]string                                                        `json:"freeformTags"`
		IpxeScript                     *string                                                                  `json:"ipxeScript"`
		Metadata                       map[string]string                                                        `json:"metadata"`
		Shape                          *string                                                                  `json:"shape"`
		ShapeConfig                    *InstanceConfigurationLaunchInstanceShapeConfigDetails                   `json:"shapeConfig"`
		PlatformConfig                 instanceconfigurationlaunchinstanceplatformconfig                        `json:"platformConfig"`
		SourceDetails                  instanceconfigurationinstancesourcedetails                               `json:"sourceDetails"`
		FaultDomain                    *string                                                                  `json:"faultDomain"`
		DedicatedVmHostId              *string                                                                  `json:"dedicatedVmHostId"`
		LaunchMode                     InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum                 `json:"launchMode"`
		LaunchOptions                  *InstanceConfigurationLaunchOptions                                      `json:"launchOptions"`
		AgentConfig                    *InstanceConfigurationLaunchInstanceAgentConfigDetails                   `json:"agentConfig"`
		IsPvEncryptionInTransitEnabled *bool                                                                    `json:"isPvEncryptionInTransitEnabled"`
		PreferredMaintenanceAction     InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum `json:"preferredMaintenanceAction"`
		InstanceOptions                *InstanceConfigurationInstanceOptions                                    `json:"instanceOptions"`
		AvailabilityConfig             *InstanceConfigurationAvailabilityConfig                                 `json:"availabilityConfig"`
		PreemptibleInstanceConfig      *PreemptibleInstanceConfigDetails                                        `json:"preemptibleInstanceConfig"`
		LicensingConfigs               []launchinstancelicensingconfig                                          `json:"licensingConfigs"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.AvailabilityDomain = model.AvailabilityDomain

	m.CapacityReservationId = model.CapacityReservationId

	m.CompartmentId = model.CompartmentId

	m.ClusterPlacementGroupId = model.ClusterPlacementGroupId

	m.CreateVnicDetails = model.CreateVnicDetails

	m.DefinedTags = model.DefinedTags

	m.SecurityAttributes = model.SecurityAttributes

	m.DisplayName = model.DisplayName

	m.ExtendedMetadata = model.ExtendedMetadata

	m.FreeformTags = model.FreeformTags

	m.IpxeScript = model.IpxeScript

	m.Metadata = model.Metadata

	m.Shape = model.Shape

	m.ShapeConfig = model.ShapeConfig

	nn, e = model.PlatformConfig.UnmarshalPolymorphicJSON(model.PlatformConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PlatformConfig = nn.(InstanceConfigurationLaunchInstancePlatformConfig)
	} else {
		m.PlatformConfig = nil
	}

	nn, e = model.SourceDetails.UnmarshalPolymorphicJSON(model.SourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceDetails = nn.(InstanceConfigurationInstanceSourceDetails)
	} else {
		m.SourceDetails = nil
	}

	m.FaultDomain = model.FaultDomain

	m.DedicatedVmHostId = model.DedicatedVmHostId

	m.LaunchMode = model.LaunchMode

	m.LaunchOptions = model.LaunchOptions

	m.AgentConfig = model.AgentConfig

	m.IsPvEncryptionInTransitEnabled = model.IsPvEncryptionInTransitEnabled

	m.PreferredMaintenanceAction = model.PreferredMaintenanceAction

	m.InstanceOptions = model.InstanceOptions

	m.AvailabilityConfig = model.AvailabilityConfig

	m.PreemptibleInstanceConfig = model.PreemptibleInstanceConfig

	m.LicensingConfigs = make([]LaunchInstanceLicensingConfig, len(model.LicensingConfigs))
	for i, n := range model.LicensingConfigs {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.LicensingConfigs[i] = nn.(LaunchInstanceLicensingConfig)
		} else {
			m.LicensingConfigs[i] = nil
		}
	}
	return
}

// InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum Enum with underlying type: string
type InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum string

// Set of constants representing the allowable values for InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum
const (
	InstanceConfigurationLaunchInstanceDetailsLaunchModeNative          InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum = "NATIVE"
	InstanceConfigurationLaunchInstanceDetailsLaunchModeEmulated        InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum = "EMULATED"
	InstanceConfigurationLaunchInstanceDetailsLaunchModeParavirtualized InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum = "PARAVIRTUALIZED"
	InstanceConfigurationLaunchInstanceDetailsLaunchModeCustom          InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum = "CUSTOM"
)

var mappingInstanceConfigurationLaunchInstanceDetailsLaunchModeEnum = map[string]InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum{
	"NATIVE":          InstanceConfigurationLaunchInstanceDetailsLaunchModeNative,
	"EMULATED":        InstanceConfigurationLaunchInstanceDetailsLaunchModeEmulated,
	"PARAVIRTUALIZED": InstanceConfigurationLaunchInstanceDetailsLaunchModeParavirtualized,
	"CUSTOM":          InstanceConfigurationLaunchInstanceDetailsLaunchModeCustom,
}

var mappingInstanceConfigurationLaunchInstanceDetailsLaunchModeEnumLowerCase = map[string]InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum{
	"native":          InstanceConfigurationLaunchInstanceDetailsLaunchModeNative,
	"emulated":        InstanceConfigurationLaunchInstanceDetailsLaunchModeEmulated,
	"paravirtualized": InstanceConfigurationLaunchInstanceDetailsLaunchModeParavirtualized,
	"custom":          InstanceConfigurationLaunchInstanceDetailsLaunchModeCustom,
}

// GetInstanceConfigurationLaunchInstanceDetailsLaunchModeEnumValues Enumerates the set of values for InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum
func GetInstanceConfigurationLaunchInstanceDetailsLaunchModeEnumValues() []InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum {
	values := make([]InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum, 0)
	for _, v := range mappingInstanceConfigurationLaunchInstanceDetailsLaunchModeEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceConfigurationLaunchInstanceDetailsLaunchModeEnumStringValues Enumerates the set of values in String for InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum
func GetInstanceConfigurationLaunchInstanceDetailsLaunchModeEnumStringValues() []string {
	return []string{
		"NATIVE",
		"EMULATED",
		"PARAVIRTUALIZED",
		"CUSTOM",
	}
}

// GetMappingInstanceConfigurationLaunchInstanceDetailsLaunchModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceConfigurationLaunchInstanceDetailsLaunchModeEnum(val string) (InstanceConfigurationLaunchInstanceDetailsLaunchModeEnum, bool) {
	enum, ok := mappingInstanceConfigurationLaunchInstanceDetailsLaunchModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum Enum with underlying type: string
type InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum string

// Set of constants representing the allowable values for InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum
const (
	InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionLiveMigrate InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum = "LIVE_MIGRATE"
	InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionReboot      InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum = "REBOOT"
)

var mappingInstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum = map[string]InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum{
	"LIVE_MIGRATE": InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionLiveMigrate,
	"REBOOT":       InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionReboot,
}

var mappingInstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnumLowerCase = map[string]InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum{
	"live_migrate": InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionLiveMigrate,
	"reboot":       InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionReboot,
}

// GetInstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnumValues Enumerates the set of values for InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum
func GetInstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnumValues() []InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum {
	values := make([]InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum, 0)
	for _, v := range mappingInstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnumStringValues Enumerates the set of values in String for InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum
func GetInstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnumStringValues() []string {
	return []string{
		"LIVE_MIGRATE",
		"REBOOT",
	}
}

// GetMappingInstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum(val string) (InstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnum, bool) {
	enum, ok := mappingInstanceConfigurationLaunchInstanceDetailsPreferredMaintenanceActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
