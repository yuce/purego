// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2025 The Ebitengine Authors

package purego

// Source for constants: https://github.com/openbsd/src/blob/master/include/dlfcn.h

const (
	intSize      = 32 << (^uint(0) >> 63) // 32 or 64
	RTLD_DEFAULT = -2
	RTLD_LAZY    = 1
	RTLD_NOW     = 2
	RTLD_LOCAL   = 0x000
	RTLD_GLOBAL  = 0x100
)
