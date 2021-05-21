/*
 * Copyright 2021 Nikki Nikkhoui <nnikkhoui@wikimedia.org> and Wikimedia Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"

	log "github.com/eevans/servicelib-golang/logger"
)

var (
	version   = "unknown"
	buildDate = "unknown"
	buildHost = "unknown"
)

// Entrypoint for our service
func main() {
	var confFile = flag.String("config", "./config.yaml", "Path to the configuration file")

	var config *Config
	var err error
	var logger *log.Logger

	flag.Parse()

	if config, err = ReadConfig(*confFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	logger, err = log.NewLogger(
		os.Stdout,
		config.ServiceName,
		config.LogLevel,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to initialize the logger: %s", err)
		os.Exit(1)
	}

	logger.Info("Initializing service %s (Go version: %s, Build host: %s, Timestamp: %s", config.ServiceName, version, buildHost, buildDate)

	http.Handle("/healthz", &HealthzHandler{NewHealthz(version, buildDate, buildHost)})
	http.Handle(path.Join(config.BaseURI, "echo"), &EchoHandler{})
	http.ListenAndServe(fmt.Sprintf("%s:%d", config.Address, config.Port), nil)
}
