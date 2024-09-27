//nolint:gocritic
package codespec

import (
	"github.com/mongodb/terraform-provider-mongodbatlas/internal/common/conversion"
	"github.com/mongodb/terraform-provider-mongodbatlas/tools/codegen/config"
	"github.com/mongodb/terraform-provider-mongodbatlas/tools/codegen/openapi"
)

// using blank identifiers for now, will be removed in follow-up PRs once logic for conversion is added
func ToProviderSpecModel(atlasAdminAPISpecFilePath, configPath string, resourceName *string) *Model {
	_, err := openapi.ParseAtlasAdminAPI(atlasAdminAPISpecFilePath)
	if err != nil {
		panic(err)
	}

	genConfig, _ := config.ParseGenConfigYAML(configPath)

	// var resourceSpec config.Resource
	if resourceName != nil {
		_ = genConfig.Resources[*resourceName]
	}

	// TODO: remove after ToProviderSpecModel() implemented
	return TestExampleCodeSpecification()
}

func TestExampleCodeSpecification() *Model {
	testFieldDesc := "Test field description"
	return &Model{
		Resources: []Resource{{
			Schema: &Schema{
				Attributes: Attributes{
					Attribute{
						Name:                     "project_id",
						ComputedOptionalRequired: Required,
						String:                   &StringAttribute{},
						Description:              conversion.StringPtr("Overridden project_id description"),
					},
					Attribute{
						Name:                     "bucket_name",
						ComputedOptionalRequired: Required,
						String:                   &StringAttribute{},
						Description:              &testFieldDesc,
					},
					Attribute{
						Name:                     "iam_role_id",
						ComputedOptionalRequired: Required,
						String:                   &StringAttribute{},
						Description:              &testFieldDesc,
					},
					Attribute{
						Name:                     "state",
						ComputedOptionalRequired: Computed,
						String:                   &StringAttribute{},
						Description:              &testFieldDesc,
					},
					Attribute{
						Name:                     "prefix_path",
						String:                   &StringAttribute{},
						ComputedOptionalRequired: ComputedOptional,
						Description:              &testFieldDesc,
					},
				},
			},
			Name: "test_resource",
		},
		},
	}
}