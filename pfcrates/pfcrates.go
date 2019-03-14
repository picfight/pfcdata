// Copyright (c) 2019, The Decred developers
// See LICENSE for details.
package dcrrates

import "github.com/picfight/pfcd/pfcutil"

const (
	DefaultKeyName  = "rpc.key"
	DefaultCertName = "rpc.cert"
)

var (
	DefaultAppDirectory = pfcutil.AppDataDir("dcrrates", false)
)
