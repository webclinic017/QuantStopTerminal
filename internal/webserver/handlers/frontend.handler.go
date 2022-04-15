package handlers

import (
	"github.com/quantstop/quantstopterminal/internal"
	"github.com/quantstop/quantstopterminal/internal/database/models"
	"github.com/quantstop/quantstopterminal/internal/webserver/write"
	"net/http"
)

func GetVersion(bot internal.IEngine, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {

	return write.JSON(bot.GetVersion())
}

func GetSubsystemStatus(bot internal.IEngine, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {

	return write.JSON(bot.GetSubsystemsStatus())
}

func GetUptime(bot internal.IEngine, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {

	return write.JSON(bot.GetUptime())
}
