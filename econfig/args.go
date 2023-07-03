// Copyright (c) 2023, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// note: parsing code adapted from pflag package https://github.com/spf13/pflag
// Copyright (c) 2012 Alex Ogier. All rights reserved.
// Copyright (c) 2012 The Go Authors. All rights reserved.

package econfig

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/goki/ki/kit"
	"github.com/iancoleman/strcase"
)

// SetFromArgs sets Config values from command-line args,
// based on the field names in the Config struct.
// Returns any args that did not start with a `-` flag indicator.
// For more robust error processing, it is assumed that all flagged args (-)
// must refer to fields in the config, so any that fail to match trigger
// an error.  Errors can also result from parsing.
// Errors are automatically logged because these are user-facing.
func SetFromArgs(cfg any, args []string) (leftovers []string, err error) {
	leftovers, err = parseArgs(cfg, args)
	if err != nil {
		fmt.Println(Usage(cfg))
	}
	return
}

// parseArgs does the actual arg parsing
func parseArgs(cfg any, args []string) ([]string, error) {
	allArgs := FieldArgNames(cfg)
	var leftovers []string
	var err error
	for len(args) > 0 {
		s := args[0]
		args = args[1:]
		if len(s) == 0 || s[0] != '-' || len(s) == 1 {
			leftovers = append(leftovers, s)
			continue
		}

		if s[1] == '-' && len(s) == 2 { // "--" terminates the flags
			// f.argsLenAtDash = len(f.args)
			leftovers = append(leftovers, args...)
			break
		}
		args, err = parseArg(s, args, allArgs)
		if err != nil {
			return leftovers, err
		}
	}
	return leftovers, nil
}

func parseArg(s string, args []string, allArgs map[string]reflect.Value) (a []string, err error) {
	a = args
	name := s[1:]
	if name[0] == '-' {
		name = name[1:]
	}
	if len(name) == 0 || name[0] == '-' || name[0] == '=' {
		err = fmt.Errorf("SetFromArgs: bad flag syntax: %s", s)
		log.Println(err)
		return
	}

	split := strings.SplitN(name, "=", 2)
	name = split[0]
	fval, exists := allArgs[name]
	if !exists {
		err = fmt.Errorf("SetFromArgs: flag name not recognized: %s", name)
		log.Println(err)
		return
	}

	isbool := kit.NonPtrValue(fval).Kind() == reflect.Bool

	var value string
	switch {
	case len(split) == 2:
		// '--flag=arg'
		value = split[1]
	case isbool:
		// '--flag' bare
		lcnm := strings.ToLower(name)
		negate := false
		if len(lcnm) > 3 {
			if lcnm[:3] == "no_" || lcnm[:3] == "no-" {
				negate = true
			} else if lcnm[:2] == "no" {
				if _, has := allArgs[lcnm[2:]]; has { // e.g., nogui and gui is on list
					negate = true
				}
			}
		}
		if negate {
			value = "false"
		} else {
			value = "true"
		}
	case len(a) > 0:
		// '--flag arg'
		value = a[0]
		a = a[1:]
	default:
		// '--flag' (arg was required)
		err = fmt.Errorf("SetFromArgs: flag needs an argument: %s", s)
		log.Println(err)
		return
	}

	err = setArgValue(name, fval, value)
	return
}

func setArgValue(name string, fval reflect.Value, value string) error {
	nptyp := kit.NonPtrType(fval.Type())
	vk := nptyp.Kind()
	switch {
	case vk >= reflect.Int && vk <= reflect.Uint64 && kit.Enums.TypeRegistered(nptyp):
		return kit.SetEnumValueFromString(fval, value)
	case vk == reflect.Map:
		mval := make(map[string]any)
		err := ReadBytes(&mval, []byte("tmp="+value)) // use toml decoder
		if err != nil {
			log.Println(err)
			return err
		}
		ok := kit.SetRobust(fval.Interface(), mval["tmp"])
		if !ok {
			err := fmt.Errorf("SetFromArgs: not able to set field from arg: %s val: %s", name, value)
			log.Println(err)
			return err
		}
	default:
		ok := kit.SetRobust(fval.Interface(), value) // overkill but whatever
		if !ok {
			err := fmt.Errorf("SetFromArgs: not able to set field from arg: %s val: %s", name, value)
			log.Println(err)
			return err
		}
	}
	return nil
}

// FieldArgNames returns map of all the different ways the field names
// can be specified as arg flags, mapping to the reflect.Value
func FieldArgNames(obj any) map[string]reflect.Value {
	allArgs := make(map[string]reflect.Value)
	fieldArgNamesStruct(obj, "", allArgs)
	return allArgs
}

func addAllCases(nm, path string, pval reflect.Value, allArgs map[string]reflect.Value) {
	if nm == "Includes" {
		return // skip
	}
	if path != "" {
		nm = path + "." + nm
	}
	allArgs[nm] = pval
	allArgs[strings.ToLower(nm)] = pval
	allArgs[strcase.ToKebab(nm)] = pval
	allArgs[strcase.ToSnake(nm)] = pval
	allArgs[strcase.ToScreamingSnake(nm)] = pval
}

// fieldArgNamesStruct returns map of all the different ways the field names
// can be specified as arg flags, mapping to the reflect.Value
func fieldArgNamesStruct(obj any, path string, allArgs map[string]reflect.Value) {
	typ := kit.NonPtrType(reflect.TypeOf(obj))
	val := kit.NonPtrValue(reflect.ValueOf(obj))
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		fv := val.Field(i)
		if kit.NonPtrType(f.Type).Kind() == reflect.Struct {
			nwPath := f.Name
			if path != "" {
				nwPath = path + "." + nwPath
			}
			fieldArgNamesStruct(kit.PtrValue(fv).Interface(), nwPath, allArgs)
			continue
		}
		pval := kit.PtrValue(fv)
		addAllCases(f.Name, path, pval, allArgs)
		if f.Type.Kind() == reflect.Bool {
			addAllCases("No"+f.Name, path, pval, allArgs)
		}
	}
}
