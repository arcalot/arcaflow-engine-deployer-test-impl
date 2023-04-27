package testimpl

import (
	"context"
	"go.flow.arcalot.io/pluginsdk/atp"
	"io"
	"time"

	log "go.arcalot.io/log/v2"
	"go.flow.arcalot.io/deployer"
)

type connector struct {
	config *Config
	logger log.Logger
}

// pluginConnection holds the IO for a plugin, and fulfills the deployer Plugin interface.
type pluginConnection struct {
	reader io.Reader
	writer io.Writer
}

func (p pluginConnection) Read(buf []byte) (n int, err error) {
	return p.reader.Read(buf)
}

func (p pluginConnection) Write(buf []byte) (n int, err error) {
	return p.writer.Write(buf)
}

func (p pluginConnection) Close() error {
	panic("Not implemented. Careful to prevent goroutine leak")
}

type stdinSubstitute struct {
	reader io.Reader
}

func (s stdinSubstitute) Read(p []byte) (n int, err error) {
	return s.reader.Read(p)
}

func (s stdinSubstitute) Close() error {
	// We don't depend on this being closed.
	// If the output is done, it should be closed.
	panic("Not implemented")
}

func newStdinSubstitute() (stdinSubstitute, io.Writer) {
	stdinReader, stdinWriter := io.Pipe()
	return stdinSubstitute{stdinReader}, stdinWriter
}

type stdoutSubstitute struct {
	writer io.Writer
}

func (s stdoutSubstitute) Write(p []byte) (n int, err error) {
	return s.writer.Write(p)
}

func (s stdoutSubstitute) Close() error {
	// We don't depend on this being closed.
	// If the output is done, it should be closed.
	panic("Not implemented")
}

func newStdoutSubstitute() (stdoutSubstitute, io.Reader) {
	stdoutReader, stdoutWriter := io.Pipe()
	return stdoutSubstitute{stdoutWriter}, stdoutReader
}

func (c *connector) Deploy(ctx context.Context, image string) (deployer.Plugin, error) {
	c.logger.Infof("Mimicking deployment of a plugin with image %s for testing", image)

	// Simulate how it takes time to start the deployment.
	time.Sleep(time.Duration(c.config.DeployTime) * time.Millisecond)

	stdinSub, stdinWriter := newStdinSubstitute()
	stdoutSub, stdoutReader := newStdoutSubstitute()

	// TODO: run in a goroutine.
	go func() {
		err := atp.RunATPServer(ctx, stdinSub, stdoutSub, nil) // TODO: Fulfill schema.
		if err != nil {
			c.logger.Errorf("Error while running ATP server %e", err)
		}
	}()

	pluginIO := pluginConnection{
		writer: stdinWriter,
		reader: stdoutReader,
	}
	// TODO: Close on pluginIO needs to shutdown ATP server.

	c.logger.Infof("Plugin initialized.")

	return pluginIO, nil
}
