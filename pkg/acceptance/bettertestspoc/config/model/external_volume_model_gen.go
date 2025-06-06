// Code generated by config model builder generator; DO NOT EDIT.

package model

import (
	"encoding/json"

	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/config"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/provider/resources"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

type ExternalVolumeModel struct {
	Name               tfconfig.Variable `json:"name,omitempty"`
	AllowWrites        tfconfig.Variable `json:"allow_writes,omitempty"`
	Comment            tfconfig.Variable `json:"comment,omitempty"`
	FullyQualifiedName tfconfig.Variable `json:"fully_qualified_name,omitempty"`
	StorageLocation    tfconfig.Variable `json:"storage_location,omitempty"`

	DynamicBlock *config.DynamicBlock `json:"dynamic,omitempty"`

	*config.ResourceModelMeta
}

/////////////////////////////////////////////////
// Basic builders (resource name and required) //
/////////////////////////////////////////////////

func ExternalVolume(
	resourceName string,
	name string,
	storageLocation []sdk.ExternalVolumeStorageLocation,
) *ExternalVolumeModel {
	e := &ExternalVolumeModel{ResourceModelMeta: config.Meta(resourceName, resources.ExternalVolume)}
	e.WithName(name)
	e.WithStorageLocation(storageLocation)
	return e
}

func ExternalVolumeWithDefaultMeta(
	name string,
	storageLocation []sdk.ExternalVolumeStorageLocation,
) *ExternalVolumeModel {
	e := &ExternalVolumeModel{ResourceModelMeta: config.DefaultMeta(resources.ExternalVolume)}
	e.WithName(name)
	e.WithStorageLocation(storageLocation)
	return e
}

///////////////////////////////////////////////////////////////////////
// set proper json marshalling, handle depends on and dynamic blocks //
///////////////////////////////////////////////////////////////////////

func (e *ExternalVolumeModel) MarshalJSON() ([]byte, error) {
	type Alias ExternalVolumeModel
	return json.Marshal(&struct {
		*Alias
		DependsOn []string `json:"depends_on,omitempty"`
	}{
		Alias:     (*Alias)(e),
		DependsOn: e.DependsOn(),
	})
}

func (e *ExternalVolumeModel) WithDependsOn(values ...string) *ExternalVolumeModel {
	e.SetDependsOn(values...)
	return e
}

func (e *ExternalVolumeModel) WithDynamicBlock(dynamicBlock *config.DynamicBlock) *ExternalVolumeModel {
	e.DynamicBlock = dynamicBlock
	return e
}

/////////////////////////////////
// below all the proper values //
/////////////////////////////////

func (e *ExternalVolumeModel) WithName(name string) *ExternalVolumeModel {
	e.Name = tfconfig.StringVariable(name)
	return e
}

func (e *ExternalVolumeModel) WithAllowWrites(allowWrites string) *ExternalVolumeModel {
	e.AllowWrites = tfconfig.StringVariable(allowWrites)
	return e
}

func (e *ExternalVolumeModel) WithComment(comment string) *ExternalVolumeModel {
	e.Comment = tfconfig.StringVariable(comment)
	return e
}

func (e *ExternalVolumeModel) WithFullyQualifiedName(fullyQualifiedName string) *ExternalVolumeModel {
	e.FullyQualifiedName = tfconfig.StringVariable(fullyQualifiedName)
	return e
}

// storage_location attribute type is not yet supported, so WithStorageLocation can't be generated

//////////////////////////////////////////
// below it's possible to set any value //
//////////////////////////////////////////

func (e *ExternalVolumeModel) WithNameValue(value tfconfig.Variable) *ExternalVolumeModel {
	e.Name = value
	return e
}

func (e *ExternalVolumeModel) WithAllowWritesValue(value tfconfig.Variable) *ExternalVolumeModel {
	e.AllowWrites = value
	return e
}

func (e *ExternalVolumeModel) WithCommentValue(value tfconfig.Variable) *ExternalVolumeModel {
	e.Comment = value
	return e
}

func (e *ExternalVolumeModel) WithFullyQualifiedNameValue(value tfconfig.Variable) *ExternalVolumeModel {
	e.FullyQualifiedName = value
	return e
}

func (e *ExternalVolumeModel) WithStorageLocationValue(value tfconfig.Variable) *ExternalVolumeModel {
	e.StorageLocation = value
	return e
}
