package output

import (
	"fmt"
	"reflect"
	"time"
)

type ColumnSpec struct {
	Header string
	Format string // hint for renderer
}

func DefaultColumnsFromType(root any) ([]ColumnSpec, error) {
	rt := deref(reflect.TypeOf(root))
	seen := map[string]bool{}
	var cols []ColumnSpec

	if err := walkColumns(rt, "", seen, &cols); err != nil {
		return nil, err
	}
	return cols, nil
}

func walkColumns(t reflect.Type, prefix string, seen map[string]bool, cols *[]ColumnSpec) error {
	t = deref(t)
	if t.Kind() != reflect.Struct {
		return nil
	}

	timeType := reflect.TypeOf(time.Time{})
	childSeen := false

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}
		meta := parseTableTag(&f)
		ft := deref(f.Type)

		switch {
		case meta.Skip:
			continue

		case meta.Children:
			if childSeen {
				return fmt.Errorf("multiple table_children tags on %s", t.Name())
			}
			childSeen = true

			sliceType := deref(f.Type)
			if sliceType.Kind() != reflect.Slice && sliceType.Kind() != reflect.Array {
				return fmt.Errorf("field %s.%s tagged table_children must be a slice", t.Name(), f.Name)
			}

			elemType := deref(sliceType.Elem())
			if elemType.Kind() != reflect.Struct {
				return fmt.Errorf("field %s.%s tagged table_children must be a slice of structs", t.Name(), f.Name)
			}

			childPrefix := meta.ChildLabel
			if childPrefix == "" {
				childPrefix = fieldName(&f)
			}

			nextPrefix := childPrefix
			if prefix != "" {
				nextPrefix = prefix + "." + childPrefix
			}

			if err := walkColumns(elemType, nextPrefix, seen, cols); err != nil {
				return err
			}

		case isSlice(ft):
			if meta.Header != "" && meta.Default {
				addColumn(prefix, meta.Header, meta.Format, seen, cols)
			}

		case ft.Kind() == reflect.Struct && ft != timeType:
			if err := walkColumns(ft, prefix, seen, cols); err != nil {
				return err
			}

		default:
			if meta.Header != "" && meta.Default {
				addColumn(prefix, meta.Header, meta.Format, seen, cols)
			}
		}
	}
	return nil
}

func addColumn(prefix, header, format string, seen map[string]bool, cols *[]ColumnSpec) {
	full := header
	if prefix != "" {
		full = prefix + "." + header
	}
	if seen[full] {
		return
	}
	seen[full] = true
	*cols = append(*cols, ColumnSpec{
		Header: full,
		Format: format,
	})
}
