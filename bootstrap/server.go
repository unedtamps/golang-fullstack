package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	e "github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	c "github.com/unedtamps/chat-app/api/controllers"
	r "github.com/unedtamps/chat-app/api/routes"
	s "github.com/unedtamps/chat-app/api/service"
	cfg "github.com/unedtamps/chat-app/config"
)

type App struct {
	route   *e.Echo
	dBClose func() string
	addr    string
}

func LoadServer() *App {
	dbQuery, dbClose := cfg.NewRepo()
	addr := cfg.GetSeverEnv().Port
	services := s.NewService(dbQuery)
	controlers := c.NewController(services)
	routes := r.LoadRoute(controlers)
	return &App{
		route:   routes,
		dBClose: dbClose,
		addr:    addr,
	}
}

func (app *App) Start() {
	serve := http.Server{
		Addr:         fmt.Sprintf(":%s", app.addr),
		Handler:      app.route,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	shotdown := make(chan struct{})
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := serve.Shutdown(ctx); err != nil {
			log.Fatal().Msg(err.Error())
		}
		shotdown <- struct{}{}
	}()
	log.Info().Msgf("Start Server %s", app.addr)
	if err := serve.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal().Msgf("Shutting Down Server: %v", err.Error())
	}
	<-shotdown
	log.Info().Msg(app.dBClose())
	log.Info().Msg("Server Down")
}
