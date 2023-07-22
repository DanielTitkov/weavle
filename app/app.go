package app

import (
	"time"

	"github.com/DanielTitkov/weavle/configs"
	"github.com/DanielTitkov/weavle/domain"
	"github.com/DanielTitkov/weavle/logger"
	"github.com/DanielTitkov/weavle/repository"
	"github.com/DanielTitkov/weavle/util"
)

type (
	App struct {
		Cfg           configs.Config
		log           *logger.Logger
		repo          *repository.Repository
		locales       []string // locale count is not very big so no need to have map
		Errors        []error
		systemSummary *domain.SystemSummary
		Events        []domain.Event
	}
)

func New(
	cfg configs.Config,
	logger *logger.Logger,
	repo *repository.Repository,
	// store sessions.Store,
) (*App, error) {
	start := time.Now()
	defer util.InfoExecutionTime(start, "app.New", logger)
	app := App{
		Cfg:  cfg,
		log:  logger,
		repo: repo,
		// Store: store,
		// locales: domain.Locales(),
	}

	app.AddEvent("app.New", start)

	return &app, nil
}

func (a *App) IsDev() bool {
	return a.Cfg.Env == domain.EnvDev
}

func (a *App) AddError(err error) {
	a.Errors = append(a.Errors, err)
}

func (a *App) AddEvent(name string, start time.Time) {
	if len(a.Events) >= domain.AppMaxEvents {
		a.Events = a.Events[:len(a.Events)-1]
	}
	// prepend event to show in reverse order
	a.Events = append([]domain.Event{domain.Event{
		Name:      name,
		StartTime: start,
		EndTime:   time.Now(),
		Elapsed:   time.Since(start),
	}}, a.Events...)
}
