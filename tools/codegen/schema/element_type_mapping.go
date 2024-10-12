package schema

import (
	"fmt"

	"github.com/mongodb/terraform-provider-mongodbatlas/tools/codegen/codespec"
)

var elementTypeToString = map[codespec.ElemType]string{
	codespec.Bool:    "types.BoolType",
	codespec.Float64: "types.Float64",
	codespec.Int64:   "types.Int64",
	codespec.Number:  "types.NumberType",
	codespec.String:  "types.StringType",
}

func ElementTypeProperty(elementType codespec.ElemType) CodeStatement {
	result := elementTypeToString[elementType]
	return CodeStatement{
		Result:  fmt.Sprintf("ElementType: %s", result),
		Imports: []string{"github.com/hashicorp/terraform-plugin-framework/types"},
	}
}
