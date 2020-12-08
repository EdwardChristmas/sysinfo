// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

// Package sysinfo is a Go library providing Linux OS / kernel / hardware system information.
package sysinfo

// SysInfo struct encapsulates all other information structs.
type SysInfo struct {
	OS      OS              `json:"os"`
	Kernel  Kernel          `json:"kernel"`
}

// GetSysInfo gathers all available system information.
func (si *SysInfo) GetSysInfo() {
	si.getOSInfo()
	si.getKernelInfo()
}
