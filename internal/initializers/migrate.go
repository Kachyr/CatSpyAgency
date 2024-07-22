package initializers

import (
	"github.com/Kachyr/SpyCatAgency/pkg/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) {
	log.Info().Msg("Syncing database")

	if err := db.AutoMigrate(&models.Cat{}, &models.Mission{}, &models.Target{}); err != nil {
		log.Fatal().Err(err).Msg("Error to migrate database")
	}
	// if err := db.SetupJoinTable(&models.User{}, "SeenAnimals", &models.SeenAnimal{}); err != nil {
	// 	log.Fatal().Err(err).Msg("Error to setup join table SeenAnimals")
	// }
	// if err := db.SetupJoinTable(&models.User{}, "LikedAnimals", &models.LikedAnimal{}); err != nil {
	// 	log.Fatal().Err(err).Msg("Error to setup join table LikedAnimals")
	// }
}
