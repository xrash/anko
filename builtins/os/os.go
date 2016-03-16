// Package os implements os interface for anko script.
package os

import (
	"github.com/mattn/anko/vm"
	pkg "os"
	"reflect"
)

func Import(env *vm.Env) *vm.Env {
	m := env.NewEnv()
	m.Define("Args", reflect.ValueOf(pkg.Args))
	m.Define("Chdir", reflect.ValueOf(pkg.Chdir))
	m.Define("Chmod", reflect.ValueOf(pkg.Chmod))
	m.Define("Chown", reflect.ValueOf(pkg.Chown))
	m.Define("Chtimes", reflect.ValueOf(pkg.Chtimes))
	m.Define("Clearenv", reflect.ValueOf(pkg.Clearenv))
	m.Define("Create", reflect.ValueOf(pkg.Create))
	m.Define("DevNull", reflect.ValueOf(pkg.DevNull))
	m.Define("Environ", reflect.ValueOf(pkg.Environ))
	m.Define("ErrExist", reflect.ValueOf(pkg.ErrExist))
	m.Define("ErrInvalid", reflect.ValueOf(pkg.ErrInvalid))
	m.Define("ErrNotExist", reflect.ValueOf(pkg.ErrNotExist))
	m.Define("ErrPermission", reflect.ValueOf(pkg.ErrPermission))
	m.Define("Exit", reflect.ValueOf(pkg.Exit))
	m.Define("Expand", reflect.ValueOf(pkg.Expand))
	m.Define("ExpandEnv", reflect.ValueOf(pkg.ExpandEnv))
	m.Define("FindProcess", reflect.ValueOf(pkg.FindProcess))
	m.Define("Getegid", reflect.ValueOf(pkg.Getegid))
	m.Define("Getenv", reflect.ValueOf(pkg.Getenv))
	m.Define("Geteuid", reflect.ValueOf(pkg.Geteuid))
	m.Define("Getgid", reflect.ValueOf(pkg.Getgid))
	m.Define("Getgroups", reflect.ValueOf(pkg.Getgroups))
	m.Define("Getpagesize", reflect.ValueOf(pkg.Getpagesize))
	m.Define("Getpid", reflect.ValueOf(pkg.Getpid))
	handleAppEngine(m)
	m.Define("Getuid", reflect.ValueOf(pkg.Getuid))
	m.Define("Getwd", reflect.ValueOf(pkg.Getwd))
	m.Define("Hostname", reflect.ValueOf(pkg.Hostname))
	m.Define("Interrupt", reflect.ValueOf(pkg.Interrupt))
	m.Define("IsExist", reflect.ValueOf(pkg.IsExist))
	m.Define("IsNotExist", reflect.ValueOf(pkg.IsNotExist))
	m.Define("IsPathSeparator", reflect.ValueOf(pkg.IsPathSeparator))
	m.Define("IsPermission", reflect.ValueOf(pkg.IsPermission))
	m.Define("Kill", reflect.ValueOf(pkg.Kill))
	m.Define("Lchown", reflect.ValueOf(pkg.Lchown))
	m.Define("Link", reflect.ValueOf(pkg.Link))
	m.Define("Lstat", reflect.ValueOf(pkg.Lstat))
	m.Define("Mkdir", reflect.ValueOf(pkg.Mkdir))
	m.Define("MkdirAll", reflect.ValueOf(pkg.MkdirAll))
	m.Define("ModeAppend", reflect.ValueOf(pkg.ModeAppend))
	m.Define("ModeCharDevice", reflect.ValueOf(pkg.ModeCharDevice))
	m.Define("ModeDevice", reflect.ValueOf(pkg.ModeDevice))
	m.Define("ModeDir", reflect.ValueOf(pkg.ModeDir))
	m.Define("ModeExclusive", reflect.ValueOf(pkg.ModeExclusive))
	m.Define("ModeNamedPipe", reflect.ValueOf(pkg.ModeNamedPipe))
	m.Define("ModePerm", reflect.ValueOf(pkg.ModePerm))
	m.Define("ModeSetgid", reflect.ValueOf(pkg.ModeSetgid))
	m.Define("ModeSetuid", reflect.ValueOf(pkg.ModeSetuid))
	m.Define("ModeSocket", reflect.ValueOf(pkg.ModeSocket))
	m.Define("ModeSticky", reflect.ValueOf(pkg.ModeSticky))
	m.Define("ModeSymlink", reflect.ValueOf(pkg.ModeSymlink))
	m.Define("ModeTemporary", reflect.ValueOf(pkg.ModeTemporary))
	m.Define("ModeType", reflect.ValueOf(pkg.ModeType))
	m.Define("NewFile", reflect.ValueOf(pkg.NewFile))
	m.Define("NewSyscallError", reflect.ValueOf(pkg.NewSyscallError))
	m.Define("O_APPEND", reflect.ValueOf(pkg.O_APPEND))
	m.Define("O_CREATE", reflect.ValueOf(pkg.O_CREATE))
	m.Define("O_EXCL", reflect.ValueOf(pkg.O_EXCL))
	m.Define("O_RDONLY", reflect.ValueOf(pkg.O_RDONLY))
	m.Define("O_RDWR", reflect.ValueOf(pkg.O_RDWR))
	m.Define("O_SYNC", reflect.ValueOf(pkg.O_SYNC))
	m.Define("O_TRUNC", reflect.ValueOf(pkg.O_TRUNC))
	m.Define("O_WRONLY", reflect.ValueOf(pkg.O_WRONLY))
	m.Define("Open", reflect.ValueOf(pkg.Open))
	m.Define("OpenFile", reflect.ValueOf(pkg.OpenFile))
	m.Define("PathListSeparator", reflect.ValueOf(pkg.PathListSeparator))
	m.Define("PathSeparator", reflect.ValueOf(pkg.PathSeparator))
	m.Define("Pipe", reflect.ValueOf(pkg.Pipe))
	m.Define("Readlink", reflect.ValueOf(pkg.Readlink))
	m.Define("Remove", reflect.ValueOf(pkg.Remove))
	m.Define("RemoveAll", reflect.ValueOf(pkg.RemoveAll))
	m.Define("Rename", reflect.ValueOf(pkg.Rename))
	m.Define("SEEK_CUR", reflect.ValueOf(pkg.SEEK_CUR))
	m.Define("SEEK_END", reflect.ValueOf(pkg.SEEK_END))
	m.Define("SEEK_SET", reflect.ValueOf(pkg.SEEK_SET))
	m.Define("SameFile", reflect.ValueOf(pkg.SameFile))
	m.Define("Setenv", reflect.ValueOf(pkg.Setenv))
	m.Define("StartProcess", reflect.ValueOf(pkg.StartProcess))
	m.Define("Stat", reflect.ValueOf(pkg.Stat))
	m.Define("Stderr", reflect.ValueOf(pkg.Stderr))
	m.Define("Stdin", reflect.ValueOf(pkg.Stdin))
	m.Define("Stdout", reflect.ValueOf(pkg.Stdout))
	m.Define("Symlink", reflect.ValueOf(pkg.Symlink))
	m.Define("TempDir", reflect.ValueOf(pkg.TempDir))
	m.Define("Truncate", reflect.ValueOf(pkg.Truncate))
	m.DefineType("Signal", reflect.TypeOf(pkg.Signal(nil)))
	return m
}
