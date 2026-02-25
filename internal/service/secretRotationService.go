package service

import (
	"context"
	"fmt"
	"secret-manager/internal/logging"
	"secret-manager/pkg/persistence"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type RotationJob struct {
	Db      *gorm.DB
	Service SecretService
}

var (
	// RerollSelectedTotal Track how many times an item was picked up by the DB query
	RerollSelectedTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "secret_manager_item_selected_total",
		Help: "Total times a specific secret was selected for rotation",
	}, []string{"file_path"})

	// RerollSuccessTotal Track how many times an item successfully rotated
	RerollSuccessTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "secret_manager_item_success_total",
		Help: "Total times a specific secret was successfully rotated",
	}, []string{"file_path"})

	// RerollFailureTotal Track failures specifically per item
	RerollFailureTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "secret_manager_item_failure_total",
		Help: "Total times a specific secret failed rotation",
	}, []string{"file_path"})
)

func (r RotationJob) Run() {

	rerolls, err := persistence.GetPendingRerolls(r.Db)
	if err != nil {
		logging.L.App.Error(fmt.Sprintf("Pending Reroll query failed %s", err.Error()))
	}

	g, ctx := errgroup.WithContext(context.Background())
	g.SetLimit(10)

	var failedCount int32

	for _, reroll := range rerolls {
		reroll := reroll

		// 1. Increment "Selected" immediately (Item was picked up by the query)
		RerollSelectedTotal.WithLabelValues(reroll.FilePath).Inc()

		g.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			// 2. Perform the work
			if err := r.Service.FetchAndStoreTemplate(reroll); err != nil {
				logging.L.App.Error(fmt.Sprintf("%s fetch failed: %v", reroll.FilePath, err))
				RerollFailureTotal.WithLabelValues(reroll.FilePath).Inc()
				return nil
			}

			if err := persistence.SaveSecretConfig(r.Db, reroll); err != nil {
				logging.L.App.Error(fmt.Sprintf("[ERROR] %s save failed: %v", reroll.FilePath, err))
				RerollFailureTotal.WithLabelValues(reroll.FilePath).Inc()
				return nil
			}

			// 3. Increment "Success" only if both steps passed
			RerollSuccessTotal.WithLabelValues(reroll.FilePath).Inc()
			return nil
		})
	}

	_ = g.Wait()
	logging.L.App.Debug(fmt.Sprintf("Batch complete. Total: %d, Failures: %d", len(rerolls), failedCount))

}
