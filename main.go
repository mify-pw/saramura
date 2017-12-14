// Copyright (c) 2017 Christopher König (Kura)
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"time"

	"github.com/evalphobia/logrus_sentry"
	log "github.com/sirupsen/logrus"
	suture "gopkg.in/thejerf/suture.v2"
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

	// setup supervisor system
	spec := suture.Spec{
		Log: func(l string) {
			log.Infoln(l)
		},
	}

	sv := suture.New("mify", spec)
	sv.ServeBackground()
	log.Info("suture running")

	// now start up all needed clients
	// this is done via ssh, even for localhost instances.
	// make sure to
}
