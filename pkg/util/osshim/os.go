// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package osshim

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate -header ../../../license_prefix.txt

import (
	"os"
)

// Os is shim for methods from os package
//counterfeiter:generate . Os
type Os interface {
	ReadFile(name string) ([]byte, error)
	IsNotExist(err error) bool
	Lstat(name string) (os.FileInfo, error)
}

// OsShim is default Os implementation
type OsShim struct{}

// ReadFile see os.ReadFile
func (sh *OsShim) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

// IsNotExist see os.IsNotExist
func (sh *OsShim) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

// Lstat see os.Lstat
func (sh *OsShim) Lstat(name string) (os.FileInfo, error) {
	return os.Lstat(name)
}
