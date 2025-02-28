package sentryext

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideSentry(log *zap.Logger) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://d0d1db3c8658d0e93f0e8d09912a7805@o4508897586839552.ingest.de.sentry.io/4508897654997072",
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Error("Sentry initialization failed", zap.Error(err))
		return err
	}

	log.Info("✅ Sentry initialized successfully")
	return nil
}

// SentryModule for Fx
var SentryModule = fx.Module("sentry",
	fx.Invoke(func(lc fx.Lifecycle, log *zap.Logger) error {

		if err := ProvideSentry(log); err != nil {
			return err
		}

		lc.Append(fx.Hook{
			OnStop: func(context.Context) error {
				sentry.Flush(2 * time.Second)
				log.Info("🛑 Sentry flushed and shutting down")
				return nil
			},
		})

		return nil
	}),
)
