package config

import (
	"log"

	"github.com/JalalA984/apptrak/internal/models"
)

type ApplicationConfig struct {
	ErrorLog *log.Logger
	InfoLog *log.Logger

	Applications *models.ApplicationModel
}
