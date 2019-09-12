// Copyright (c) 2019, The Decred developers
// See LICENSE for details.
package pfcrates

import "github.com/picfight/pfcd/pfcutil"

const (
	DefaultKeyName  = "rpc.key"
	DefaultCertName = "rpc.cert"
)

var (
	DefaultAppDirectory = pfcutil.AppDataDir("pfcrates", false)
)
