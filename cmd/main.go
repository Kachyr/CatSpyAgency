package main

import (
	"context"
	"time"

	"github.com/Kachyr/SpyCatAgency/internal/handlers"
	"github.com/Kachyr/SpyCatAgency/internal/initializers"
	"github.com/Kachyr/SpyCatAgency/internal/services"
	"github.com/Kachyr/SpyCatAgency/internal/store"
	"github.com/Kachyr/SpyCatAgency/logger"
	"github.com/Kachyr/SpyCatAgency/pkg/db"
	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

const appName = "SpyCatAgency"

var gormDB *gorm.DB
var configuration *config

func init() {
	// prepare config
	var err error
	configuration, err = loadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to prepare config")
	}

	ctx := context.Background()

	initLogger(ctx, configuration.LogLevel)
	gormDB, err = connectDB(ctx, configuration.Database)

	if err != nil {
		log.Ctx(ctx).Fatal().Err(err).Msg("unable to connect to database")
	}

	initializers.SyncDatabase(gormDB)
}

func main() {
	catStore := store.NewCatStore(gormDB)
	missionStore := store.NewMissionStore(gormDB)

	catService := services.NewCatService(catStore, missionStore)
	missionService := services.NewMissionService(missionStore)

	catHandler := handlers.NewCatHandler(catService)
	missionHandler := handlers.NewMissionHandler(missionService)

	router := initializers.NewRouter(catHandler, missionHandler)

	ginEngine := gin.Default()
	router.SetupAPIs(ginEngine)

	ginEngine.Run(configuration.GinPort)
}

func initLogger(ctx context.Context, logLevel zerolog.Level) {
	logger.Init(logLevel, appName)
	log.Ctx(ctx).Info().Msg("logger initialized")
}

func connectDB(ctx context.Context, dbConfig *dbConfig) (*gorm.DB, error) {
	gormDB, err := db.Connect(ctx, dbConfig.ReadURL, dbConfig.WriteURL, 3, time.Second*2)

	return gormDB, err
}
