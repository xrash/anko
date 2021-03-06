// Package strings implements strings interface for anko script.
package strings

import (
	"github.com/mattn/anko/vm"
	"reflect"
	pkg "strings"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewPackage("strings")
	m.Define("Contains", reflect.ValueOf(pkg.Contains))
	m.Define("ContainsAny", reflect.ValueOf(pkg.ContainsAny))
	m.Define("ContainsRune", reflect.ValueOf(pkg.ContainsRune))
	m.Define("Count", reflect.ValueOf(pkg.Count))
	m.Define("EqualFold", reflect.ValueOf(pkg.EqualFold))
	m.Define("Fields", reflect.ValueOf(pkg.Fields))
	m.Define("FieldsFunc", reflect.ValueOf(pkg.FieldsFunc))
	m.Define("HasPrefix", reflect.ValueOf(pkg.HasPrefix))
	m.Define("HasSuffix", reflect.ValueOf(pkg.HasSuffix))
	m.Define("Index", reflect.ValueOf(pkg.Index))
	m.Define("IndexAny", reflect.ValueOf(pkg.IndexAny))
	m.Define("IndexByte", reflect.ValueOf(pkg.IndexByte))
	m.Define("IndexFunc", reflect.ValueOf(pkg.IndexFunc))
	m.Define("IndexRune", reflect.ValueOf(pkg.IndexRune))
	m.Define("Join", reflect.ValueOf(pkg.Join))
	m.Define("LastIndex", reflect.ValueOf(pkg.LastIndex))
	m.Define("LastIndexAny", reflect.ValueOf(pkg.LastIndexAny))
	m.Define("LastIndexFunc", reflect.ValueOf(pkg.LastIndexFunc))
	m.Define("Map", reflect.ValueOf(pkg.Map))
	m.Define("NewReader", reflect.ValueOf(pkg.NewReader))
	m.Define("NewReplacer", reflect.ValueOf(pkg.NewReplacer))
	m.Define("Repeat", reflect.ValueOf(pkg.Repeat))
	m.Define("Replace", reflect.ValueOf(pkg.Replace))
	m.Define("Split", reflect.ValueOf(pkg.Split))
	m.Define("SplitAfter", reflect.ValueOf(pkg.SplitAfter))
	m.Define("SplitAfterN", reflect.ValueOf(pkg.SplitAfterN))
	m.Define("SplitN", reflect.ValueOf(pkg.SplitN))
	m.Define("Title", reflect.ValueOf(pkg.Title))
	m.Define("ToLower", reflect.ValueOf(pkg.ToLower))
	m.Define("ToLowerSpecial", reflect.ValueOf(pkg.ToLowerSpecial))
	m.Define("ToTitle", reflect.ValueOf(pkg.ToTitle))
	m.Define("ToTitleSpecial", reflect.ValueOf(pkg.ToTitleSpecial))
	m.Define("ToUpper", reflect.ValueOf(pkg.ToUpper))
	m.Define("ToUpperSpecial", reflect.ValueOf(pkg.ToUpperSpecial))
	m.Define("Trim", reflect.ValueOf(pkg.Trim))
	m.Define("TrimFunc", reflect.ValueOf(pkg.TrimFunc))
	m.Define("TrimLeft", reflect.ValueOf(pkg.TrimLeft))
	m.Define("TrimLeftFunc", reflect.ValueOf(pkg.TrimLeftFunc))
	m.Define("TrimPrefix", reflect.ValueOf(pkg.TrimPrefix))
	m.Define("TrimRight", reflect.ValueOf(pkg.TrimRight))
	m.Define("TrimRightFunc", reflect.ValueOf(pkg.TrimRightFunc))
	m.Define("TrimSpace", reflect.ValueOf(pkg.TrimSpace))
	m.Define("TrimSuffix", reflect.ValueOf(pkg.TrimSuffix))
	return m
}
