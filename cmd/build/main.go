package main

import (
	"os"
	"time"

	"github.com/dmikusa/rust-dist-cnb/rust"
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/cargo"
	"github.com/paketo-buildpacks/packit/pexec"
	"github.com/paketo-buildpacks/packit/postal"
	"github.com/paketo-buildpacks/packit/scribe"
)

func main() {
	transport := cargo.NewTransport()
	dependencyService := postal.NewService(transport)
	clock := rust.NewClock(time.Now)
	logEmitter := rust.NewLogEmitter(os.Stdout)
	runner := rust.NewInstallRunner(
		pexec.NewExecutable("install.sh"),
		scribe.NewWriter(os.Stdout, scribe.WithIndent(4)),
		scribe.NewWriter(os.Stderr, scribe.WithIndent(4)),
	)

	packit.Build(rust.Build(dependencyService, &runner, clock, logEmitter))
}
