package testimpl

import (
	"context"
	"errors"
	"fmt"
	"go.flow.arcalot.io/pluginsdk/atp"
	testplugin "go.flow.arcalot.io/testplugin"
	"io"
	"sync"
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
	wg     *sync.WaitGroup
	id     string
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
	readerCloseErr := p.reader.Close()
	writerCloseErr := p.writer.Close()
	if readerCloseErr != nil || writerCloseErr != nil {
		return fmt.Errorf("error while closing pipes (%w)", errors.Join(readerCloseErr, writerCloseErr))
	}
	p.wg.Wait()
	return nil
}

func (p *pluginConnection) ID() string {
	return p.id
}

// badConnection holds the IO for a plugin, and fulfills the deployer Plugin interface.
type badConnection struct {
	reader *io.PipeReader
	writer *io.PipeWriter
	cancel context.CancelFunc
	id     string
}

func (p *badConnection) Read(buf []byte) (n int, err error) {
	return p.reader.Read(buf)
}

func (p *badConnection) Write(buf []byte) (n int, err error) {
	return 0, fmt.Errorf("bad connection to writer")
}

func (p *badConnection) Close() error {
	// Cancel the context that was sent to the ATP server.
	// This will instruct it to finish up and close its stdin.
	// You need to let it close it instead of closing it here, or else it will panic due to being unable to
	// send the CBOR messages.
	p.cancel()
	return nil
}

func (p *badConnection) ID() string {
	return p.id
}

func (c *connector) Deploy(ctx context.Context, image string) (deployer.Plugin, error) {
	c.logger.Infof("Mimicking deployment of a plugin with image %s for testing.", image)

	// Simulate how it takes time to start the deployment.
	time.Sleep(time.Duration(c.config.DeployTime) * time.Millisecond)

	if !c.config.DeploySucceed {
		return nil, fmt.Errorf("intentional deployment fail after %d ms", c.config.DeployTime)
	}

	// Simulate stdin and stdout with simple pipes.
	stdinSub, stdinWriter := io.Pipe()
	stdoutReader, stdoutSub := io.Pipe()

	// Give the plugin an independent context, so it can handle itself.
	pluginCtx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1) // For the goroutine
	go func() {
		c.logger.Debugf("Starting ATP server in test deployer impl\n")
		// Just run the ATP server until the context is cancelled, or it completes. Whatever comes first.
		schemaClone := *testplugin.GetSchema()
		errs := atp.RunATPServer(pluginCtx, stdinSub, stdoutSub, &schemaClone)
		if len(errs) > 0 {
			c.logger.Errorf("%d errors while running ATP server", len(errs))
			for _, err := range errs {
				c.logger.Errorf("error while running ATP server: %w", err)
			}
		}
		c.logger.Debugf("ATP server execution finished in test deployer impl\n")
		wg.Done()
	}()

	if c.config.DisablePluginWrites {
		return &badConnection{
			writer: stdinWriter,
			reader: stdoutReader,
			cancel: cancel,
			id:     image,
		}, nil
	}

	pluginIO := &pluginConnection{
		writer: stdinWriter,
		reader: stdoutReader,
		cancel: cancel,
		wg:     &wg,
		id:     image,
	}

	c.logger.Infof("Plugin initialized.")

	return pluginIO, nil
}
