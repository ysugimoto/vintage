package transformer

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/vintage"
)

const lineFeed = "\n"

func (tf *CoreTransformer) transformAcl(acl *ast.AclDeclaration) ([]byte, error) {
	var buf bytes.Buffer

	name := acl.Name.String()
	tf.acls[name] = newExpressionValue(vintage.ACL, "acl__"+name)

	buf.WriteString(
		fmt.Sprintf(`var acl__%s = vintage.NewAcl("%s",`+lineFeed, name, name),
	)

	for _, cidr := range acl.CIDRs {
		mask := "32"
		if cidr.Mask != nil {
			mask = cidr.Mask.String()
		}
		var inverse bool
		if cidr.Inverse != nil {
			inverse = cidr.Inverse.Value
		}
		buf.WriteString(
			fmt.Sprintf(`vintage.AclEntry("%s/%s", %t),`+lineFeed, cidr.IP.String(), mask, inverse),
		)
	}

	buf.WriteString(")" + lineFeed)
	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformBackend(backend *ast.BackendDeclaration) ([]byte, error) {
	var buf bytes.Buffer

	name := backend.Name.String()
	// We will use first found backend as default
	isDefault := len(tf.backends) == 0
	tf.backends[name] = newExpressionValue(vintage.BACKEND, "backend__"+name)

	buf.WriteString(
		fmt.Sprintf(`var backend__%s = vintage.NewBackend("%s", %t)`+lineFeed, name, name, isDefault),
	)

	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformDirector(director *ast.DirectorDeclaration) ([]byte, error) {
	var buf bytes.Buffer

	name := director.Name.String()
	tf.backends[name] = newExpressionValue(vintage.BACKEND, "director__"+name)

	buf.WriteString(
		fmt.Sprintf(
			`var director__%s = vintage.NewDirector("%s", "%s",`+lineFeed,
			name, name, director.DirectorType.String(),
		),
	)
	for _, prop := range director.Properties {
		switch p := prop.(type) {
		case *ast.DirectorProperty:
			buf.WriteString(
				fmt.Sprintf(`vintage.DirectorProperty("%s", %s),`+lineFeed, p.Key.Value, toString(p.Value)),
			)
		case *ast.DirectorBackendObject:
			buf.WriteString(`vintage.DirectorBackend(` + lineFeed)
			for _, v := range p.Values {
				buf.WriteString(
					fmt.Sprintf(`vintage.DirectorProperty("%s", %s),`+lineFeed, v.Key.Value, toString(v.Value)),
				)
			}
			buf.WriteString(")," + lineFeed)
		}
	}

	buf.WriteString(")" + lineFeed)

	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformTable(table *ast.TableDeclaration) ([]byte, error) {
	var buf bytes.Buffer

	name := table.Name.String()
	tf.tables[name] = newExpressionValue(vintage.IDENT, "table__"+name)

	tableType := "STRING"
	if table.ValueType != nil {
		tableType = table.ValueType.String()
	}

	buf.WriteString(
		fmt.Sprintf(`var table__%s = vintage.NewTable("%s", "%s",`+lineFeed, name, name, tableType),
	)

	for _, prop := range table.Properties {
		buf.WriteString(
			fmt.Sprintf(`vintage.TableItem("%s", %s),`+lineFeed, prop.Key.Value, prop.Value),
		)
	}

	buf.WriteString(")" + lineFeed)

	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformSubroutine(sub *ast.SubroutineDeclaration) ([]byte, error) {
	var buf bytes.Buffer

	name := sub.Name.String()
	if sub.ReturnType != nil {
		tf.functionSubroutines[name] = newExpressionValue(vintage.IDENT, name)
		buf.WriteString(fmt.Sprintf(
			"func %s(ctx *fastly.Runtime) (%s, error) {"+lineFeed,
			name,
			vintage.GoTypeString(vintage.VCLType(sub.ReturnType.Value)),
		))
		inside, _, err := tf.transformBlockStatement(sub.Block.Statements)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.Write(inside)
		buf.WriteString("}\n")
		return buf.Bytes(), nil
	}

	tf.subroutines[name] = newExpressionValue(vintage.IDENT, name)

	buf.WriteString(
		fmt.Sprintf("func %s(ctx *fastly.Runtime) (vintage.State, error) {"+lineFeed, name),
	)
	inside, rs, err := tf.transformBlockStatement(sub.Block.Statements)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf.Write(inside)
	if !rs {
		buf.WriteString(`return vintage.NONE, nil` + lineFeed)
	}
	buf.WriteString("}\n")

	return buf.Bytes(), nil

}
