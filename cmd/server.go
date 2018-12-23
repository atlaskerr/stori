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

package cmd

import (
	//	"github.com/atlaskerr/stori/stori"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the Stori registry server.",
	Run: func(cmd *cobra.Command, args []string) {
		//		configFile := cmd.Flag("config").Value.String()
		logLevel := cmd.Flag("log-level").Value.String()
		devMode, _ := cmd.Flags().GetBool("dev")

		logger, _ := newLogger(logLevel, devMode)
		logger.Info("Logger construction succeeded.")

	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringP(
		"config",
		"c",
		"/etc/stori/config.json",
		"Path to config file.",
	)

	serverCmd.Flags().StringP(
		"log-level",
		"l",
		"info",
		"Set the server's log level.",
	)

	serverCmd.Flags().Bool(
		"dev",
		false,
		"Start server in development mode.",
	)
}

// newLogger initializes a new logger. When dev is set to true, structured
// logging is disabled and all logs will be output to the console.
func newLogger(lvl string, dev bool) (l *zap.Logger, err error) {
	var ll zapcore.Level

	switch lvl {
	case "debug":
		ll = zapcore.DebugLevel
	case "info":
		ll = zapcore.InfoLevel
	case "warn":
		ll = zapcore.WarnLevel
	case "error":
		ll = zapcore.ErrorLevel
	case "dpanic":
		ll = zapcore.DPanicLevel
	case "panic":
		ll = zapcore.PanicLevel
	case "fatal":
		ll = zapcore.FatalLevel
	}

	levelEncoder := func() zapcore.LevelEncoder {
		if dev {
			return zapcore.CapitalColorLevelEncoder
		}
		return zapcore.CapitalLevelEncoder
	}()

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    levelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	logEncoding := func() string {
		if dev {
			return "console"
		}
		return "json"
	}()

	logConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(ll),
		Development:      dev,
		Encoding:         logEncoding,
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{os.Stdout.Name()},
		ErrorOutputPaths: []string{os.Stderr.Name()},
	}

	l, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
	return
}
