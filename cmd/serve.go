package cmd

import (
	"context"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/loggingx"
	"go.infratographer.com/x/otelx"
	"go.infratographer.com/x/versionx"
	"go.uber.org/zap"

	"go.infratographer.com/location-api/internal/config"
	ent "go.infratographer.com/location-api/internal/ent/generated"
	"go.infratographer.com/location-api/internal/graphapi"
)

var defaultListenAddr = ":7906"

var (
	enablePlayground bool
	serveDevMode     bool
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the location API",
	Run: func(cmd *cobra.Command, args []string) {
		serve(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	echox.MustViperFlags(viper.GetViper(), serveCmd.Flags(), defaultListenAddr)
	echojwtx.MustViperFlags(viper.GetViper(), serveCmd.Flags())
	events.MustViperFlagsForPublisher(viper.GetViper(), serveCmd.Flags(), appName)
	permissions.MustViperFlags(viper.GetViper(), serveCmd.Flags())

	// only available as a CLI arg because it shouldn't be something that could accidentially end up in a config file or env var
	serveCmd.Flags().BoolVar(&serveDevMode, "dev", false, "dev mode: enables playground, disables all auth checks, sets CORS to allow all, pretty logging, etc.")
	serveCmd.Flags().BoolVar(&enablePlayground, "playground", false, "enable the graph playground")
}

func serve(ctx context.Context) {
	if serveDevMode {
		enablePlayground = true
		config.AppConfig.Logging.Debug = true
		config.AppConfig.Logging.Pretty = true
		config.AppConfig.Server.WithMiddleware(middleware.CORS())
		// reinit the logger
		logger = loggingx.InitLogger(appName, config.AppConfig.Logging)
		// this is a hack, echojwt needs to be updated to go into AppConfig
		viper.Set("oidc.enabled", false)
	}

	publisher, err := events.NewPublisher(config.AppConfig.Events.Publisher)
	if err != nil {
		logger.Fatal("unable to initialize event publisher", zap.Error(err))
	}

	err = otelx.InitTracer(config.AppConfig.Tracing, appName, logger)
	if err != nil {
		logger.Fatal("unable to initialize tracing system", zap.Error(err))
	}

	db, err := crdbx.NewDB(config.AppConfig.CRDB, config.AppConfig.Tracing.Enabled)
	if err != nil {
		logger.Fatal("unable to initialize crdb client", zap.Error(err))
	}

	defer db.Close()

	entDB := entsql.OpenDB(dialect.Postgres, db)

	cOpts := []ent.Option{ent.Driver(entDB), ent.EventsPublisher(publisher)}

	if config.AppConfig.Logging.Debug {
		cOpts = append(cOpts,
			ent.Log(logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	client := ent.NewClient(cOpts...)
	defer client.Close()

	var middleware []echo.MiddlewareFunc

	perms, err := permissions.New(config.AppConfig.Permissions,
		permissions.WithLogger(logger),
		permissions.WithDefaultChecker(permissions.DefaultAllowChecker),
	)
	if err != nil {
		logger.Fatal("failed to initialize permissions", zap.Error(err))
	}

	middleware = append(middleware, perms.Middleware())

	// jwt auth middleware
	if viper.GetBool("oidc.enabled") {
		auth, err := echojwtx.NewAuth(ctx, config.AppConfig.OIDC)
		if err != nil {
			logger.Fatal("failed to initialize jwt authentication", zap.Error(err))
		}

		auth.JWTConfig.Skipper = echox.SkipDefaultEndpoints

		middleware = append(middleware, auth.Middleware())
	}

	config.AppConfig.Server = config.AppConfig.Server.WithMiddleware(middleware...)

	srv, err := echox.NewServer(logger.Desugar(), echox.ConfigFromViper(viper.GetViper()), versionx.BuildDetails())
	if err != nil {
		logger.Fatal("failed to initialize new server", zap.Error(err))
	}

	r := graphapi.NewResolver(client, logger.Named("resolvers"))
	handler := r.Handler(enablePlayground)

	srv.AddHandler(handler)

	// TODO: we should have a database check
	// srv.AddReadinessCheck("database", r.DatabaseCheck)

	if err := srv.RunWithContext(ctx); err != nil {
		logger.Fatal("failed to run server", zap.Error(err))
	}
}
