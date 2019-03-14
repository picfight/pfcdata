module github.com/picfight/pfcdata/dcrrates/rateserver

replace github.com/picfight/pfcdata/v4 => ../..

require (
	github.com/btcsuite/go-flags v0.0.0-20150116065318-6c288d648c1c
	github.com/picfight/pfcd/certgen v1.0.2
	github.com/picfight/pfcd/pfcutil v1.2.1-0.20190118223730-3a5281156b73
	github.com/picfight/pfcdata/v4 v4.0.0-20190211084703-a009a10db389
	github.com/decred/slog v1.0.0
	github.com/jrick/logrotate v1.0.0
	google.golang.org/grpc v1.18.0
)
