package core

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/vintage/transformer/value"
)

const lineFeed = "\n"

func (tf *CoreTransformer) transformAcl(acl *ast.AclDeclaration) ([]byte, error) {
	var buf bytes.Buffer

	name := acl.Name.String()
	tf.acls[name] = value.NewValue(value.ACL, "acl__"+name)

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

	buf.WriteString(
		fmt.Sprintf(`var backend__%s = vintage.NewBackend("%s",`+lineFeed, name, name),
	)
	// We will use first found backend as default
	if len(tf.backends) == 0 {
		buf.WriteString("vintage.BackendDefault()," + lineFeed)
	}
	for _, prop := range backend.Properties {
		switch prop.Key.Value {
		case "port":
			buf.WriteString(fmt.Sprintf(`vintage.BackendPort(%s),`+lineFeed, toString(prop.Value)))
		case "host":
			buf.WriteString(fmt.Sprintf(`vintage.BackendHost(%s),`+lineFeed, toString(prop.Value)))
		case "ssl":
			buf.WriteString(fmt.Sprintf(`vintage.BackendSSL(%s),`+lineFeed, toString(prop.Value)))
		case "connect_timeout":
			buf.WriteString(fmt.Sprintf(`vintage.BackendConnectTimeout(%s),`+lineFeed, toString(prop.Value)))
		case "first_byte_timeout":
			buf.WriteString(fmt.Sprintf(`vintage.BackendFirstByteTimeout(%s),`+lineFeed, toString(prop.Value)))
		case "between_bytes_timeout":
			buf.WriteString(fmt.Sprintf(`vintage.BackendBetweenBytesTimeout(%s),`+lineFeed, toString(prop.Value)))
		}
	}
	buf.WriteString(")" + lineFeed)

	tf.backends[name] = value.NewValue(value.BACKEND, "backend__"+name)
	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformDirector(director *ast.DirectorDeclaration) ([]byte, error) {
	var buf bytes.Buffer

	name := director.Name.String()
	tf.backends[name] = value.NewValue(value.BACKEND, "director__"+name)

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
	tf.tables[name] = value.NewValue(value.IDENT, "table__"+name)

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
		tf.functionSubroutines[name] = value.NewValue(value.VCLType(sub.ReturnType.Value), name)
		buf.WriteString(fmt.Sprintf(
			"func %s(ctx *%s) (%s, error) {"+lineFeed,
			name,
			tf.runtimeName,
			value.GoTypeString(value.VCLType(sub.ReturnType.Value)),
		))
		buf.WriteString("re := vintage.RegexpMatchedGroup{}" + lineFeed)
		buf.WriteString("defer re.Release()" + lineFeed)
		buf.WriteString(lineFeed)
		inside, _, err := tf.transformBlockStatement(sub.Block.Statements)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.Write(inside)
		buf.WriteString("}" + lineFeed)
		return buf.Bytes(), nil
	}

	tf.subroutines[name] = value.NewValue(value.IDENT, name)

	buf.WriteString(fmt.Sprintf(
		"func %s(ctx *%s) (vintage.State, error) {"+lineFeed,
		name,
		tf.runtimeName,
	))
	buf.WriteString("re := vintage.RegexpMatchedGroup{}" + lineFeed)
	buf.WriteString("defer re.Release()" + lineFeed)
	buf.WriteString(lineFeed)
	inside, rs, err := tf.transformBlockStatement(sub.Block.Statements)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf.Write(inside)
	if !rs {
		buf.WriteString(`return vintage.NONE, nil` + lineFeed)
	}
	buf.WriteString("}" + lineFeed)

	return buf.Bytes(), nil

}

func (tf *CoreTransformer) transformLoggingEndpoint(name string) ([]byte, error) {
	var buf bytes.Buffer

	// Need to replace from "-" to "_" due to name is used on program variable
	tf.loggingEndpoints[name] = "logging__" + strings.ReplaceAll(name, "-", "_")
	buf.WriteString(
		fmt.Sprintf(`var %s = vintage.NewLoggingEndpoint("%s")`, tf.loggingEndpoints[name], name),
	)

	return buf.Bytes(), nil
}
