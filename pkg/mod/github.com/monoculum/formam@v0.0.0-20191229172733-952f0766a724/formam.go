// Package formam decodes HTTP form and query parameters.
package formam

import (
	"encoding"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const tagName = "formam"

// pathMap holds the values of a map with its key and values correspondent
type pathMap struct {
	ma    reflect.Value
	key   string
	value reflect.Value

	path string
}

// pathMaps holds the values for each key
type pathMaps []*pathMap

// find finds and gets the value by the given key
func (m pathMaps) find(id reflect.Value, key string) *pathMap {
	for i := range m {
		if m[i].ma == id && m[i].key == key {
			return m[i]
		}
	}
	return nil
}

// DecodeCustomTypeFunc for decoding a custom type.
type DecodeCustomTypeFunc func([]string) (interface{}, error)

// decodeCustomTypeField is registered for a specific field.
type decodeCustomTypeField struct {
	field reflect.Value
	fun   DecodeCustomTypeFunc
}

// decodeCustomType fields for custom types.
type decodeCustomType struct {
	fun    DecodeCustomTypeFunc
	fields []*decodeCustomTypeField
}

// Decoder to decode a form.
type Decoder struct {
	main       reflect.Value
	formValues url.Values
	opts       *DecoderOptions

	curr   reflect.Value
	values []string

	path    string
	field   string
	bracket string
	//isKey   bool

	maps pathMaps

	customTypes map[reflect.Type]*decodeCustomType
}

// DecoderOptions options for decoding the values.
type DecoderOptions struct {
	// Struct field tag name; default is "formam".
	TagName string

	// Prefer UnmarshalText over custom types.
	PrefUnmarshalText bool

	// Ignore unknown form fields. By default unknown fields are an error
	// (although all valid keys will still be decoded).
	IgnoreUnknownKeys bool
}

// RegisterCustomType registers a functions for decoding custom types.
func (dec *Decoder) RegisterCustomType(fn DecodeCustomTypeFunc, types []interface{}, fields []interface{}) *Decoder {
	if dec.customTypes == nil {
		dec.customTypes = make(map[reflect.Type]*decodeCustomType, 100)
	}
	lenFields := len(fields)
	for i := range types {
		typ := reflect.TypeOf(types[i])
		if dec.customTypes[typ] == nil {
			dec.customTypes[typ] = &decodeCustomType{fun: fn, fields: make([]*decodeCustomTypeField, 0, lenFields)}
		}
		if lenFields > 0 {
			for j := range fields {
				val := reflect.ValueOf(fields[j])
				field := &decodeCustomTypeField{field: val, fun: fn}
				dec.customTypes[typ].fields = append(dec.customTypes[typ].fields, field)
			}
		}
	}
	return dec
}

// NewDecoder creates a new instance of Decoder.
func NewDecoder(opts *DecoderOptions) *Decoder {
	dec := &Decoder{opts: opts}
	if dec.opts == nil {
		dec.opts = &DecoderOptions{}
	}
	if dec.opts.TagName == "" {
		dec.opts.TagName = tagName
	}
	return dec
}

// Decode the url.Values and populate the destination dst, which must be a
// pointer.
func (dec Decoder) Decode(vs url.Values, dst interface{}) error {
	main := reflect.ValueOf(dst)
	if main.Kind() != reflect.Ptr {
		return newError(ErrCodeNotAPointer, "", "", "dst %q is not a pointer", main.Kind())
	}
	dec.main = main.Elem()
	dec.formValues = vs
	return dec.init()
}

// Decode the url.Values and populate the destination dst, which must be a
// pointer.
func Decode(vs url.Values, dst interface{}) error {
	main := reflect.ValueOf(dst)
	if main.Kind() != reflect.Ptr {
		return newError(ErrCodeNotAPointer, "", "", "dst %q is not a pointer", main.Kind())
	}
	dec := &Decoder{
		main:       main.Elem(),
		formValues: vs,
		opts: &DecoderOptions{
			TagName: tagName,
		},
	}
	return dec.init()
}

// init initializes the decoding
func (dec Decoder) init() error {
	// iterate over the form's values and decode it
	for k, v := range dec.formValues {
		dec.path = k
		dec.values = v
		dec.curr = dec.main
		if err := dec.analyzePath(); err != nil {
			if dec.curr.Kind() == reflect.Struct && dec.opts.IgnoreUnknownKeys {
				continue
			}
			return err
		}
	}
	// set values of maps
	for _, v := range dec.maps {
		key := v.ma.Type().Key()
		ptr := false
		// check if the key implements the UnmarshalText interface
		var val reflect.Value
		if key.Kind() == reflect.Ptr {
			ptr = true
			val = reflect.New(key.Elem())
		} else {
			val = reflect.New(key).Elem()
		}
		// decode key
		dec.path = v.path
		dec.field = v.path
		dec.values = []string{v.key}
		dec.curr = val
		//dec.isKey = true
		if err := dec.decode(); err != nil {
			return err
		}
		// check if the key is a pointer or not. And if it is, then get its address
		if ptr && dec.curr.Kind() != reflect.Ptr {
			dec.curr = dec.curr.Addr()
		}
		// set key with its value
		v.ma.SetMapIndex(dec.curr, v.value)
	}
	return nil
}

// analyzePath analyzes the current path to walk through it
func (dec *Decoder) analyzePath() (err error) {
	inBracket := false
	bracketClosed := false
	lastPos := 0
	endPos := 0

	// parse path
	for i, char := range []byte(dec.path) {
		if char == '[' && inBracket == false {
			// found an opening bracket
			bracketClosed = false
			inBracket = true
			dec.field = dec.path[lastPos:i]
			lastPos = i + 1
			continue
		} else if inBracket {
			// it is inside of bracket, so get its value
			if char == ']' {
				// found an closing bracket, so it will be recently close, so put as true the bracketClosed
				// and put as false inBracket and pass the value of bracket to dec.key
				inBracket = false
				bracketClosed = true
				if endPos == 0 { // foo[] without number.
					dec.bracket = dec.path[lastPos:i]
				} else {
					dec.bracket = dec.path[lastPos:endPos]
				}

				lastPos = i + 1
				if err = dec.traverse(); err != nil {
					return
				}
			} else {
				// still inside the bracket, so to save the end position
				endPos = i + 1
			}
			continue
		} else if !inBracket {
			// not found any bracket, so try found a field
			if char == '.' {
				// found a field, we need to know if the field is next of a closing bracket,
				// for example: [0].Field
				if bracketClosed {
					bracketClosed = false
					lastPos = i + 1
					continue
				}
				// found a field, but is not next of a closing bracket, for example: Field1.Field2
				dec.field = dec.path[lastPos:i]
				//dec.field = tmp[:i]
				lastPos = i + 1
				if err = dec.traverse(); err != nil {
					return
				}
			}
			continue
		}
	}
	// last field of path
	dec.field = dec.path[lastPos:]

	return dec.end()
}

// walk traverses the current path until to the last field
func (dec *Decoder) traverse() error {
	// check if there is field, if is so, then it should be struct or map (access by .)
	if dec.field != "" {
		// check if is a struct or map
		switch dec.curr.Kind() {
		case reflect.Struct:
			if err := dec.findStructField(); err != nil {
				return err
			}
		case reflect.Map:
			dec.traverseInMap(true)
		}
		dec.field = ""
	}
	// check if is a interface and it is not nil. This mean that the interface
	// has a struct, map or slice as value
	if dec.curr.Kind() == reflect.Interface && !dec.curr.IsNil() {
		dec.curr = dec.curr.Elem()
	}
	// check if it is a pointer
	if dec.curr.Kind() == reflect.Ptr {
		if dec.curr.IsNil() {
			dec.curr.Set(reflect.New(dec.curr.Type().Elem()))
		}
		dec.curr = dec.curr.Elem()
	}
	// check if there is access to slice/array or map (access by [])
	if dec.bracket != "" {
		switch dec.curr.Kind() {
		case reflect.Array:
			index, err := strconv.Atoi(dec.bracket)
			if err != nil {
				return newError(ErrCodeArrayIndex, dec.field, dec.path, "array index is not a number: %s", err)
			}
			dec.curr = dec.curr.Index(index)
		case reflect.Slice:
			index, err := strconv.Atoi(dec.bracket)
			if err != nil {
				return newError(ErrCodeArrayIndex, dec.field, dec.path, "slice index is not a number: %s", err)
			}
			if dec.curr.Len() <= index {
				dec.expandSlice(index + 1)
			}
			dec.curr = dec.curr.Index(index)
		case reflect.Map:
			dec.traverseInMap(false)
		default:
			return newError(ErrCodeArrayIndex, dec.field, dec.path, "has an array index but it is a %v", dec.curr.Kind())
		}
		dec.bracket = ""
	}
	return nil
}

// walkMap puts in d.curr the map concrete for decode the current value
func (dec *Decoder) traverseInMap(byField bool) {
	n := dec.curr.Type()
	makeAndAppend := func() {
		if dec.maps == nil {
			dec.maps = make(pathMaps, 0, 500)
		}
		m := reflect.New(n.Elem()).Elem()
		if byField {
			dec.maps = append(dec.maps, &pathMap{dec.curr, dec.field, m, dec.path})
		} else {
			dec.maps = append(dec.maps, &pathMap{dec.curr, dec.bracket, m, dec.path})
		}
		dec.curr = m
	}
	if dec.curr.IsNil() {
		// map is nil
		dec.curr.Set(reflect.MakeMap(n))
		makeAndAppend()
	} else {
		// map is not nil, so try find value by the key
		var a *pathMap
		if byField {
			a = dec.maps.find(dec.curr, dec.field)
		} else {
			a = dec.maps.find(dec.curr, dec.bracket)
		}
		if a == nil {
			// the key not exists
			makeAndAppend()
		} else {
			dec.curr = a.value
		}
	}
}

// end finds the last field for decode its value correspondent
func (dec *Decoder) end() error {
	switch dec.curr.Kind() {
	case reflect.Struct:
		if err := dec.findStructField(); err != nil {
			return err
		}
	case reflect.Map:
		// leave backward compatibility for access to maps by .
		dec.traverseInMap(true)
	}
	return dec.decode()
}

// decode sets the value in the field
func (dec *Decoder) decode() error {
	// check if has UnmarshalText method or a custom function to decode it
	if dec.opts.PrefUnmarshalText {
		if ok, err := dec.isUnmarshalText(dec.curr); ok || err != nil {
			return err
		}
		if ok, err := dec.isCustomType(); ok || err != nil {
			return err
		}
	} else {
		if ok, err := dec.isCustomType(); ok || err != nil {
			return err
		}
		if ok, err := dec.isUnmarshalText(dec.curr); ok || err != nil {
			return err
		}
	}

	switch dec.curr.Kind() {
	case reflect.Array:
		if dec.bracket == "" {
			// not has index, so to decode all values in the slice
			if err := dec.setValues(); err != nil {
				return err
			}
		} else {
			// has index, so to decode value by index indicated
			index, err := strconv.Atoi(dec.bracket)
			if err != nil {
				return newError(ErrCodeArrayIndex, dec.field, dec.path, "array index is not a number: %s", err)
			}
			dec.curr = dec.curr.Index(index)
			return dec.decode()
		}
	case reflect.Slice:
		if dec.bracket == "" {
			// not has index, so to decode all values in the slice
			// only for slices
			dec.expandSlice(len(dec.values))
			if err := dec.setValues(); err != nil {
				return err
			}
		} else {
			// has index, so to decode value by index indicated
			index, err := strconv.Atoi(dec.bracket)
			if err != nil {
				return newError(ErrCodeArrayIndex, dec.field, dec.path, "slice index is not a number: %s", err)
			}
			// only for slices
			if dec.curr.Len() <= index {
				dec.expandSlice(index + 1)
			}
			dec.curr = dec.curr.Index(index)
			return dec.decode()
		}
	case reflect.String:
		dec.curr.SetString(dec.values[0])
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if num, err := strconv.ParseInt(dec.values[0], 10, dec.curr.Type().Bits()); err != nil {
			code := ErrCodeConversion
			if err, ok := err.(*strconv.NumError); ok && err.Err == strconv.ErrRange {
				code = ErrCodeRange
			}
			return newError(code, dec.field, dec.path, "could not parse number: %s", err)
		} else {
			dec.curr.SetInt(num)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if num, err := strconv.ParseUint(dec.values[0], 10, dec.curr.Type().Bits()); err != nil {
			code := ErrCodeConversion
			if err, ok := err.(*strconv.NumError); ok && err.Err == strconv.ErrRange {
				code = ErrCodeRange
			}
			return newError(code, dec.field, dec.path, "could not parse number: %s", err)
		} else {
			dec.curr.SetUint(num)
		}
	case reflect.Float32, reflect.Float64:
		if num, err := strconv.ParseFloat(dec.values[0], dec.curr.Type().Bits()); err != nil {
			code := ErrCodeConversion
			if err, ok := err.(*strconv.NumError); ok && err.Err == strconv.ErrRange {
				code = ErrCodeRange
			}
			return newError(code, dec.field, dec.path, "could not parse float: %s", err)
		} else {
			dec.curr.SetFloat(num)
		}
	case reflect.Bool:
		switch dec.values[0] {
		case "true", "on", "1", "checked":
			dec.curr.SetBool(true)
		default:
			dec.curr.SetBool(false)
			return nil
		}
	case reflect.Interface:
		dec.curr.Set(reflect.ValueOf(dec.values[0]))
	case reflect.Ptr:
		n := reflect.New(dec.curr.Type().Elem())
		if dec.curr.CanSet() {
			dec.curr.Set(n)
		} else {
			dec.curr.Elem().Set(n.Elem())
		}
		dec.curr = dec.curr.Elem()
		return dec.decode()
	case reflect.Struct:
		switch dec.curr.Interface().(type) {
		case time.Time:
			var t time.Time
			// if the value is empty then no to try to parse it and leave "t" as a zero value to set it in the field
			if dec.values[0] != "" {
				var err error
				t, err = time.Parse("2006-01-02", dec.values[0])
				if err != nil {
					return newError(ErrCodeConversion, dec.field, dec.path, "could not parse field: %s", err)
				}
			}
			dec.curr.Set(reflect.ValueOf(t))
		case url.URL:
			u, err := url.Parse(dec.values[0])
			if err != nil {
				return newError(ErrCodeConversion, dec.field, dec.path, "could not parse field: %s", err)
			}
			dec.curr.Set(reflect.ValueOf(*u))
		default:
			if dec.opts.IgnoreUnknownKeys {
				return nil
			}
			num := dec.curr.NumField()
			for i := 0; i < num; i++ {
				field := dec.curr.Type().Field(i)
				tag := field.Tag.Get(dec.opts.TagName)
				if tag == "-" {
					// skip this field
					return nil
				}
			}
			return newError(ErrCodeUnknownType, dec.field, dec.path,
				"unsupported type; maybe include it the UnmarshalText interface or register it using custom type?")
		}
	default:
		if dec.opts.IgnoreUnknownKeys {
			return nil
		}

		return newError(ErrCodeUnknownType, dec.field, dec.path, "unsupported type")
	}

	return nil
}

// findStructField finds a field by its name, if it is not found,
// then retry the search examining the tag "formam" of every field of struct
func (dec *Decoder) findStructField() error {
	var anon reflect.Value

	num := dec.curr.NumField()
	for i := 0; i < num; i++ {
		field := dec.curr.Type().Field(i)

		if field.Name == dec.field {
			tag := field.Tag.Get(dec.opts.TagName)
			if tag == "-" {
				// skip this field
				return nil
			}
			// check if the field's name is equal
			dec.curr = dec.curr.Field(i)
			return nil
		} else if field.Anonymous {
			// if the field is a anonymous struct, then iterate over its fields
			tmp := dec.curr
			dec.curr = dec.curr.FieldByIndex(field.Index)
			if dec.curr.Kind() == reflect.Ptr {
				if dec.curr.IsNil() {
					dec.curr.Set(reflect.New(dec.curr.Type().Elem()))
				}
				dec.curr = dec.curr.Elem()
			}
			if err := dec.findStructField(); err != nil {
				dec.curr = tmp
				continue
			}
			// field in anonymous struct is found,
			// but first it should found the field in the rest of struct
			// (a field with same name in the current struct should have preference over anonymous struct)
			anon = dec.curr
			dec.curr = tmp
		} else if dec.field == getTagName(field.Tag, dec.opts.TagName) {
			// is not found yet, then retry by its tag name "formam"
			dec.curr = dec.curr.Field(i)
			return nil
		}
	}
	if anon.IsValid() {
		dec.curr = anon
		return nil
	}

	if dec.opts.IgnoreUnknownKeys {
		return nil
	}
	return newError(ErrCodeUnknownField, dec.field, dec.path, "unknown field")
}

// expandSlice expands the length and capacity of the current slice
func (dec *Decoder) expandSlice(length int) {
	n := reflect.MakeSlice(dec.curr.Type(), length, length)
	reflect.Copy(n, dec.curr)
	dec.curr.Set(n)
}

// setValues set the values in current slice/array
func (dec *Decoder) setValues() error {
	tmp := dec.curr // hold current field
	for i, v := range dec.values {
		dec.curr = tmp.Index(i)
		dec.values[0] = v
		if err := dec.decode(); err != nil {
			return err
		}
	}
	return nil
}

// isCustomType checks if the field's type to decode has a custom type registered
func (dec *Decoder) isCustomType() (bool, error) {
	if dec.customTypes == nil {
		return false, nil
	}
	if v, ok := dec.customTypes[dec.curr.Type()]; ok {
		if len(v.fields) > 0 {
			for i := range v.fields {
				// check if the current field is registered
				// in the fields of the custom type
				if v.fields[i].field.Elem() == dec.curr {
					va, err := v.fields[i].fun(dec.values)
					if err != nil {
						return true, err
					}
					dec.curr.Set(reflect.ValueOf(va))
					return true, nil
				}
			}
		}
		// check if the default function exists for fields not specific
		if v.fun != nil {
			va, err := v.fun(dec.values)
			if err != nil {
				return true, err
			}
			dec.curr.Set(reflect.ValueOf(va))
			return true, nil
		}
	}
	return false, nil
}

var (
	typeTime    = reflect.TypeOf(time.Time{})
	typeTimePtr = reflect.TypeOf(&time.Time{})
)

// isUnmarshalText returns a boolean and error. The boolean is true if the
// field's type implements TextUnmarshaler, and false if not.
func (dec *Decoder) isUnmarshalText(v reflect.Value) (bool, error) {
	// check if implements the interface
	m, ok := v.Interface().(encoding.TextUnmarshaler)
	addr := v.CanAddr()
	if !ok && !addr {
		return false, nil
	} else if addr {
		return dec.isUnmarshalText(v.Addr())
	}
	// skip if the type is time.Time
	n := v.Type()
	if n.ConvertibleTo(typeTime) || n.ConvertibleTo(typeTimePtr) {
		return false, nil
	}
	// return result
	return true, m.UnmarshalText([]byte(dec.values[0]))
}

func getTagName(t reflect.StructTag, tagName string) string {
	tag := t.Get(tagName)
	if p := strings.Index(tag, ","); p != -1 {
		return tag[:p]
	}
	return tag
}
