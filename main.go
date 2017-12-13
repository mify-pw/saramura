// Copyright (c) 2017 Christopher König (Kura)
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"time"

	"github.com/evalphobia/logrus_sentry"
	log "github.com/sirupsen/logrus"
)

func setupLogger() {
	// command line formatting
	log.SetFormatter(&log.TextFormatter{})

	if cfg.SentryDSN != "" {
		levels := []log.Level{log.PanicLevel, log.FatalLevel, log.ErrorLevel, log.WarnLevel}
		if cfg.IsDebug() {
			levels = append(levels, log.InfoLevel)
		}

		hook, err := logrus_sentry.NewSentryHook(cfg.SentryDSN, levels)
		if err != nil {
			log.WithError(err).Error("Failed to create sentry hook!")
		} else {
			hook.SetEnvironment(cfg.Environment)
			hook.Timeout = time.Second
			log.AddHook(hook)
			log.WithField("dsn", cfg.SentryDSN).Info("Enabling sentry output")
		}
	}

	if cfg.IsDebug() {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	setupLogger()
}
