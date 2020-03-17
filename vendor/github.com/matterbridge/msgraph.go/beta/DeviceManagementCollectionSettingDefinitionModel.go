// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementCollectionSettingDefinition Entity representing the defintion for a collection setting
type DeviceManagementCollectionSettingDefinition struct {
	// DeviceManagementSettingDefinition is the base model of DeviceManagementCollectionSettingDefinition
	DeviceManagementSettingDefinition
	// ElementDefinitionID The Setting Definition ID that describes what each element of the collection looks like
	ElementDefinitionID *string `json:"elementDefinitionId,omitempty"`
}