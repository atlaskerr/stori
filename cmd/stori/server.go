// Copyright © 2018 Atlas Kerr atlaskerr@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/atlaskerr/stori/cmd/stori/server"
	storihttp "github.com/atlaskerr/stori/http"
	"github.com/atlaskerr/stori/stori"

	"github.com/mitchellh/colorstring"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the Stori registry server.",
	Run:   startServer,
}

type serverCommand struct {
	config   string
	dev      bool
	logLevel string
}

var sc serverCommand

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(
		&sc.config,
		"config",
		"c",
		"/etc/stori/config.json",
		"Path to config file.",
	)

	serverCmd.Flags().StringVarP(
		&sc.logLevel,
		"log-level",
		"l",
		"info",
		"Set the server's log level.",
	)

	serverCmd.Flags().BoolVar(
		&sc.dev,
		"dev",
		false,
		"Start server in development mode.",
	)
}

func startServer(cmd *cobra.Command, args []string) {
	if sc.dev {
		lineOne := "Starting Stori Registry in development mode."
		lineTwo := "DO NOT USE IN PRODUCTION!"
		colorstring.Printf("[light_yellow][bold]%s %s\n", lineOne, lineTwo)
	}

	logger := initLogger(sc.logLevel, sc.dev)

	c := initServerConfig(sc.config, sc.dev)

	confDir := path.Dir(sc.config)
	_ = os.Chdir(confDir)

	registryConf := initRegistryConfig(logger, sc.dev)
	registry, _ := stori.NewRegistry(registryConf)

	handler := initHandler(registry)

	addr := c.Address
	tlsOpts := c.TLS
	ln := initListener(addr, tlsOpts, sc.dev)

	serve(ln, handler)
}

func initLogger(logLevel string, dev bool) *zap.Logger {
	if dev {
		return server.Logger("debug", true)
	}
	return server.Logger(logLevel, false)
}

func initServerConfig(path string, dev bool) *server.Config {
	if dev {
		return server.DevConfig
	}
	conf, err := server.LoadConfigFile(path)
	if err != nil {
		colorstring.Printf("[light_red][bold]%v\n", err)
		os.Exit(1)
	}
	return conf
}

func initRegistryConfig(logger *zap.Logger, dev bool) *stori.RegistryConfig {
	if dev {
		return &stori.RegistryConfig{
			Logger: logger,
		}
	}
	return &stori.RegistryConfig{
		Logger: logger,
	}
}

func initListener(addr string, tlsOpts server.TLS, dev bool) net.Listener {
	if dev || !tlsOpts.Enabled {
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			fmt.Printf("unable to start listener address %v: %v\n", addr, err)
		}
		return ln
	}
	cert, err := tls.LoadX509KeyPair(tlsOpts.CertFile, tlsOpts.KeyFile)
	if err != nil {
		fmt.Printf("unable to load server certificates: %v\n", err)
		os.Exit(1)
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	ln, err := tls.Listen("tcp", addr, tlsConfig)
	if err != nil {
		fmt.Printf("unable to start listener address %v: %v\n", addr, err)
	}
	return ln
}

func initHandler(reg *stori.Registry) http.Handler {
	handlerProps := &stori.HandlerProperties{
		Registry: reg,
	}
	handlerOpts := &storihttp.HandlerOpts{
		Metrics: true,
	}
	return storihttp.Handler(handlerProps, handlerOpts)
}

func serve(ln net.Listener, handler http.Handler) {
	srv := http.Server{
		Handler: handler,
	}
	go srv.Serve(ln)
	colorstring.Printf("[bold]Server listening on: %v\n", ln.Addr().String())

	// Wait for a signal to stop the server.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	for {
		select {
		case sig := <-sigs:
			if sig == syscall.SIGINT {
				colorstring.Printf("\n[light_yellow][bold]SIGINT received: Initiating graceful shutdown.\n")
				srv.Shutdown(context.Background())
				colorstring.Println("[light_green][bold]Shudown complete.")
				os.Exit(0)
			}
		}
	}
}
