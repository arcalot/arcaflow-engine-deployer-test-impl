package testimpl

import (
	"context"
	"go.flow.arcalot.io/pluginsdk/atp"
	"go.flow.arcalot.io/testdeployer/plugin"
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
	reader *io.PipeReader
	writer *io.PipeWriter
	cancel context.CancelFunc
}

func (p *pluginConnection) Read(buf []byte) (n int, err error) {
	return p.reader.Read(buf)
}

func (p *pluginConnection) Write(buf []byte) (n int, err error) {
	return p.writer.Write(buf)
}

func (p *pluginConnection) Close() error {
	// Cancel the context that was sent to the ATP server.
	// This will instruct it to finish up and close its stdin.
	// You need to let it close it instead of closing it here, or else it will panic due to being unable to
	// send the CBOR messages.
	p.cancel()
	return nil
}

func (c *connector) Deploy(ctx context.Context, image string) (deployer.Plugin, error) {
	c.logger.Infof("Mimicking deployment of a plugin with image %s for testing.", image)

	// Simulate how it takes time to start the deployment.
	time.Sleep(time.Duration(c.config.DeployTime) * time.Millisecond)

	// Simulate stdin and stdout with simple pipes.
	stdinSub, stdinWriter := io.Pipe()
	stdoutReader, stdoutSub := io.Pipe()

	// TODO: Allow plugin crash simulation by terminating the ATP server early.
	// Set up the context to close the server when the deployer is closed.
	subCtx, cancel := context.WithCancel(ctx)
	go func() {
		c.logger.Debugf("Starting ATP server in test deployer impl\n")
		// Just run the ATP server until the context is cancelled, or it completes. Whatever comes first.
		err := atp.RunATPServer(subCtx, stdinSub, stdoutSub, plugin.WaitSchema)
		if err != nil {
			c.logger.Errorf("Error while running ATP server %e", err)
		}
		c.logger.Debugf("ATP server execution finished in test deployer impl\n")
	}()

	pluginIO := &pluginConnection{
		writer: stdinWriter,
		reader: stdoutReader,
		cancel: cancel,
	}

	c.logger.Infof("Plugin initialized.")

	return pluginIO, nil
}
