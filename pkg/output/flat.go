package output

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"
)

type Options struct {
	// If empty: use the field tagged table_children:"true" in the root struct
	// (and in any nested struct we enter). If multiple are tagged => error.
	RowsetPath string

	// Bubble nested field headers without prefix (OwnerID, Owner next to State).
	PrefixNested bool // include parent field prefixes; set false to inline headers

	// Summaries for non-children slices on the same struct level.
	SliceSummary string // "join" | "count" | "first"
	JoinSep      string
	TimeFormat   string
}

func DefaultOptions() Options {
	return Options{
		RowsetPath:   "",  // auto from table_children
		PrefixNested: true,
		SliceSummary: "join",
		JoinSep:      ",",
		TimeFormat:   time.RFC3339,
	}
}

type tableMeta struct {
	Header     string
	Default    bool
	Format     string // "age","datetime","string"
	Skip       bool
	Children   bool
	ChildLabel string
}

func parseTableTag(f *reflect.StructField) tableMeta {
	m := tableMeta{}
	if v := f.Tag.Get("table"); v != "" {
		parts := strings.Split(v, ",")
		if len(parts) > 0 && parts[0] != "" {
			if parts[0] == "-" {
				m.Skip = true
			} else {
				m.Header = parts[0]
			}
		}
		for _, p := range parts[1:] {
			switch {
			case p == "default":
				m.Default = true
			case strings.HasPrefix(p, "format="):
				m.Format = strings.TrimPrefix(p, "format=")
			}
		}
	}
	if v := f.Tag.Get("table_children"); v != "" {
		m.Children = true
		if v != "true" {
			m.ChildLabel = v
		}
	}
	return m
}

func Flatten(v any, opts Options) ([]map[string]any, error) {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return nil, nil
	}
	// If top is slice, flatten each element
	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		var all []map[string]any
		for i := 0; i < rv.Len(); i++ {
			rows, err := Flatten(rv.Index(i).Interface(), opts)
			if err != nil {
				return nil, err
			}
			all = append(all, rows...)
		}
		return all, nil
	}

	out := []map[string]any{}
	if err := flattenInto(rv, opts /*path*/, "" /*cur*/, map[string]any{}, &out, true); err != nil {
		return nil, err
	}
	return out, nil
}

func flattenInto(val reflect.Value, opts Options, path string, cur map[string]any, out *[]map[string]any, emit bool) error {
	val = derefValue(val)
	switch val.Kind() {
	case reflect.Struct:
		rt := val.Type()
		timeType := reflect.TypeOf(time.Time{})

		// Find at most one children slice here (by table_children tag).
		var childFieldIdx = -1
		var childPath string

		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			if f.PkgPath != "" {
				continue
			}
			meta := parseTableTag(&f)
			if meta.Children {
				if childFieldIdx != -1 {
					return fmt.Errorf("multiple table_children on %s; exactly one allowed", rt.Name())
				}
				if !isSlice(f.Type) {
					return fmt.Errorf("table_children must be a slice: %s.%s", rt.Name(), f.Name)
				}
				childFieldIdx = i

				label := meta.ChildLabel
				if label == "" {
					label = fieldName(&f)
				}
				if path != "" {
					childPath = joinPath(path, label)
				} else {
					childPath = label
				}
			}
		}

		// First, bubble up all scalar/nested struct/table-tagged fields (NOT slices).
		for i := 0; i < rt.NumField(); i++ {
			f := rt.Field(i)
			if f.PkgPath != "" {
				continue
			}
			fv := val.Field(i)
			meta := parseTableTag(&f)

			// Skip children slice here; we expand it later.
			if i == childFieldIdx {
				continue
			}

			ft := deref(f.Type)

			switch {
			case meta.Skip:
				continue

			case isSlice(ft):
				// Non-children slice at this level → summarize
				if err := summarizeSliceInto(cur, fv, opts, path, meta); err != nil {
					return err
				}

			case ft.Kind() == reflect.Struct && ft != timeType:
				nestedPath := path
				if meta.Header != "" {
					if nestedPath != "" {
						nestedPath = joinPath(nestedPath, meta.Header)
					} else {
						nestedPath = meta.Header
					}
				} else if opts.PrefixNested {
					nestedPath = joinPath(nestedPath, fieldName(&f))
				}
				// Nested struct: bubble its table-tagged fields; prevent premature row emission.
				if err := flattenInto(fv, opts /*nested path*/, nestedPath, cur, out, false); err != nil {
					return err
				}

			default:
				// Scalar/map/time → include if it has a header (table tag)
				if meta.Header != "" {
					key := qualify(path, meta.Header, opts)
					cur[key] = scalarize(fv, opts)
				}
			}
		}

		// Then, expand the one children slice (if any); each element becomes a row branch.
		if childFieldIdx >= 0 {
			cv := val.Field(childFieldIdx)
			for j := 0; j < cv.Len(); j++ {
				branch := cloneMap(cur)
				if err := flattenInto(cv.Index(j), opts /*path*/, childPath, branch, out, true); err != nil {
					return err
				}
			}
			return nil
		}

		// No children slice here → current cur becomes a row.
		if emit {
			*out = append(*out, cloneMap(cur))
		}
		return nil

	case reflect.Slice, reflect.Array:
		// Slices are only expected when we're inside children expansion; each element → row
		for i := 0; i < val.Len(); i++ {
			branch := cloneMap(cur)
			if err := flattenInto(val.Index(i), opts, path, branch, out, emit); err != nil {
				return err
			}
		}
		return nil

	default:
		// Scalar leaf: just close a row
		if emit {
			*out = append(*out, cloneMap(cur))
		}
		return nil
	}
}

func summarizeSliceInto(cur map[string]any, fv reflect.Value, opts Options, path string, meta tableMeta) error {
	if meta.Skip || meta.Header == "" {
		return nil
	}
	key := qualify(path, meta.Header, opts)
	switch opts.SliceSummary {
	case "count":
		cur[key] = fv.Len()
	case "first":
		if fv.Len() > 0 {
			cur[key] = scalarize(fv.Index(0), opts)
		}
	default: // join
		parts := make([]string, 0, fv.Len())
		for i := 0; i < fv.Len(); i++ {
			parts = append(parts, stringify(fv.Index(i), opts))
		}
		cur[key] = strings.Join(parts, opts.JoinSep)
	}
	return nil
}

// ——— helpers (mostly identical to previous snippet) ———

func deref(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}
func derefValue(v reflect.Value) reflect.Value {
	for v.IsValid() && v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	return v
}
func isSlice(t reflect.Type) bool {
	t = deref(t)
	return t.Kind() == reflect.Slice || t.Kind() == reflect.Array
}

func fieldName(f *reflect.StructField) string {
	if j := f.Tag.Get("json"); j != "" {
		if v := strings.Split(j, ",")[0]; v != "" && v != "-" {
			return v
		}
	}
	return f.Name
}
func joinPath(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	return a + "." + b
}
func qualify(path, hdr string, opts Options) string {
	if !opts.PrefixNested || path == "" {
		return hdr
	}
	return path + "." + hdr
}
func cloneMap(m map[string]any) map[string]any {
	n := make(map[string]any, len(m))
	for k, v := range m {
		n[k] = v
	}
	return n
}

func scalarize(v reflect.Value, opts Options) any {
	v = derefValue(v)
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Map:
		iter := v.MapRange()
		m := map[string]string{}
		for iter.Next() {
			m[fmt.Sprint(iter.Key().Interface())] = fmt.Sprint(iter.Value().Interface())
		}
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		parts := make([]string, 0, len(keys))
		for _, k := range keys {
			parts = append(parts, k+"="+m[k])
		}
		return strings.Join(parts, opts.JoinSep)
	default:
		if v.IsValid() && v.Type() == reflect.TypeOf(time.Time{}) {
			return v.Interface().(time.Time).Format(opts.TimeFormat)
		}
		b, _ := json.Marshal(v.Interface())
		return string(b)
	}
}
func stringify(v reflect.Value, opts Options) string { return fmt.Sprint(scalarize(v, opts)) }
