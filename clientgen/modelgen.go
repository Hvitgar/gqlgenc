package clientgen

import (
	"fmt"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/vektah/gqlparser/v2/ast"
	"go/constant"
	"go/types"
)

func (r *SourceGenerator) GenFromDefinition(typeName string, def *ast.Definition) types.Type {
	switch def.Kind {
	case ast.Object, ast.InputObject:
		vars := make([]*types.Var, 0, len(def.Fields))
		tags := make([]string, 0, len(def.Fields))

		for _, field := range def.Fields {
			fieldDef := r.cfg.Schema.Types[field.Type.Name()]

			typ := r.namedType("", fieldDef.Name, func() types.Type {
				return r.GenFromDefinition(fieldDef.Name, fieldDef)
			})

			typ = r.binder.CopyModifiersFromAst(field.Type, typ)

			if isStruct(typ) && (fieldDef.Kind == ast.Object || fieldDef.Kind == ast.InputObject) {
				typ = types.NewPointer(typ)
			}

			name := field.Name
			if nameOveride := r.cfg.Models[def.Name].Fields[field.Name].FieldName; nameOveride != "" {
				name = nameOveride
			}

			vars = append(vars, types.NewVar(0, nil, templates.ToGo(name), typ))
			tags = append(tags, `json:"`+name+`"`)
		}

		return types.NewStruct(vars, tags)
	case ast.Enum:
		genType := r.GetGenType(typeName)

		consts := make([]*types.Const, 0)
		for _, v := range def.EnumValues {
			consts = append(consts, types.NewConst(
				0,
				r.client.Pkg(),
				fmt.Sprintf("%v%v", def.Name, v.Name),
				genType.RefType,
				constant.MakeString(v.Name),
			))
		}

		genType.Consts = consts

		return types.Typ[types.String]
	case ast.Scalar:
		panic("scalars must be predeclared: " + def.Name)
	}

	panic("cannot generate type for def: " + def.Name)
}

func isStruct(t types.Type) bool {
	_, is := t.Underlying().(*types.Struct)
	return is
}
