// Copyright (c) 2019, The Fonero developers
// See LICENSE for details.

package fnorates

import "github.com/fonero-project/fnod/fnoutil"

const (
	DefaultKeyName  = "rpc.key"
	DefaultCertName = "rpc.cert"
)

var (
	DefaultAppDirectory = fnoutil.AppDataDir("fnorates", false)
)
