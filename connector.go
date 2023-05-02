package testimpl

import (
	"context"
	"fmt"
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
}

func (p pluginConnection) Read(buf []byte) (n int, err error) {
	return p.reader.Read(buf)
}

func (p pluginConnection) Write(buf []byte) (n int, err error) {
	return p.writer.Write(buf)
}

func (p pluginConnection) Close() error {
	//panic("Not implemented. Careful to prevent goroutine leak")
	err := p.reader.Close()
	if err != nil {
		return err
	}
	err = p.writer.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *connector) Deploy(ctx context.Context, image string) (deployer.Plugin, error) {
	c.logger.Infof("Mimicking deployment of a plugin with image %s for testing.", image)

	// Simulate how it takes time to start the deployment.
	time.Sleep(time.Duration(c.config.DeployTime) * time.Millisecond)

	stdinSub, stdinWriter := io.Pipe()
	stdoutReader, stdoutSub := io.Pipe()

	// TODO: Allow plugin crash simulation by terminating the ATP server early.
	s, err := plugin.WaitSchema.SelfSerialize()
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s\n", s)
	go func() {
		c.logger.Debugf("Starting ATP server in test deployer impl\n")
		err := atp.RunATPServer(ctx, stdinSub, stdoutSub, plugin.WaitSchema)
		if err != nil {
			c.logger.Errorf("Error while running ATP server %e", err)
		}
		c.logger.Debugf("ATP server execution finished in test deployer impl\n")
	}()

	pluginIO := &pluginConnection{
		writer: stdinWriter,
		reader: stdoutReader,
	}

	c.logger.Infof("Plugin initialized.")

	return pluginIO, nil
}
