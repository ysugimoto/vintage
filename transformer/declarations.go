package transformer

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
)

const LF = "\n"

func (t *transformer) transformAcl(acl *ast.AclDeclaration) ([]byte, error) {
	var buf bytes.Buffer

	name := acl.Name.String()
	t.acls[name] = "acl__" + name

	buf.WriteString(
		fmt.Sprintf(`var acl__%s = vintage.NewAcl(`+LF, name),
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
			fmt.Sprintf(`vintage.AclEntry("%s/%s", %t),`+LF, cidr.IP.String(), mask, inverse),
		)
	}

	buf.WriteString(")" + LF)
	return buf.Bytes(), nil
}

func (t *transformer) transformBackend(backend *ast.BackendDeclaration) ([]byte, error) {
	var buf bytes.Buffer

	name := backend.Name.String()
	t.backends[name] = "backend__" + name

	buf.WriteString(
		fmt.Sprintf(`var backend__%s = vintage.NewBackend("%s")`+LF, name, name),
	)

	return buf.Bytes(), nil
}

func (t *transformer) transformDirector(director *ast.DirectorDeclaration) ([]byte, error) {
	var buf bytes.Buffer

	name := director.Name.String()
	t.backends[name] = "director__" + name

	buf.WriteString(
		fmt.Sprintf(
			`var director__%s = vintage.NewDirector("%s", "%s",`+LF,
			name, name, director.DirectorType.String(),
		),
	)
	for _, prop := range director.Properties {
		switch p := prop.(type) {
		case *ast.DirectorProperty:
			buf.WriteString(
				fmt.Sprintf(`vintage.DirectorProperty("%s", %s),`+LF, p.Key.Value, toString(p.Value)),
			)
		case *ast.DirectorBackendObject:
			buf.WriteString(`vintage.DirectorBackend(` + LF)
			for _, v := range p.Values {
				buf.WriteString(
					fmt.Sprintf(`vintage.DirectorProperty("%s", %s),`+LF, v.Key.Value, toString(v.Value)),
				)
			}
			buf.WriteString(")," + LF)
		}
	}

	buf.WriteString(")" + LF)

	return buf.Bytes(), nil
}

func (t *transformer) transformTable(table *ast.TableDeclaration) ([]byte, error) {
	var buf bytes.Buffer

	name := table.Name.String()
	t.tables[name] = "table__" + name

	tableType := "STRING"
	if table.ValueType != nil {
		tableType = table.ValueType.String()
	}

	buf.WriteString(
		fmt.Sprintf(`var table__%s = vintage.NewTable("%s", "%s",`+LF, name, name, tableType),
	)

	for _, prop := range table.Properties {
		buf.WriteString(
			fmt.Sprintf(`vintage.TableItem("%s", %s),`+LF, prop.Key.Value, prop.Value),
		)
	}

	buf.WriteString(")" + LF)

	return buf.Bytes(), nil
}

func (t *transformer) transformSubroutine(sub *ast.SubroutineDeclaration) ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString(
		fmt.Sprintf("func %s(ctx *fastly.Runtime) (vintage.State, error) {\n", sub.Name.String()),
	)
	inside, err := t.transformBlockStatement(sub.Block.Statements)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf.Write(inside)
	buf.WriteString("}\n")

	return buf.Bytes(), nil

}
