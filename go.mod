module github.com/dsauerbrun/cragcast

go 1.18

replace github.com/dsauerbrun/cragcast/noaaclient => ./pkg/noaaClient

replace github.com/dsauerbrun/cragcast/api => ./api

require (
	github.com/dsauerbrun/cragcast/noaaclient v0.0.0-00010101000000-000000000000
	github.com/dsauerbrun/cragcast/api v0.0.0-00010101000000-000000000000
)
