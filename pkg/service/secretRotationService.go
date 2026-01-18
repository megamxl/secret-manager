package service

import (
	"context"
	"log"
	"secret-manager/pkg/persistence"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type RotationJob struct {
	Db      *gorm.DB
	Service TemplateService
}

func (r RotationJob) Run() {

	rerolls, err := persistence.GetPendingRerolls(r.Db)
	if err != nil {
		log.Println(err)
	}

	g, ctx := errgroup.WithContext(context.Background())
	g.SetLimit(10)

	var failedCount int32

	for _, reroll := range rerolls {
		reroll := reroll
		g.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			if err := r.Service.FetchAndStoreTemplate(reroll); err != nil {
				log.Printf("[ERROR] %s fetch failed: %v", reroll.FilePath, err)
				atomic.AddInt32(&failedCount, 1)
				return nil // Return nil to keep the group running
			}

			if err := persistence.SaveSecretConfig(r.Db, reroll); err != nil {
				log.Printf("[ERROR] %s save failed: %v", reroll.FilePath, err)
				atomic.AddInt32(&failedCount, 1)
			}

			return nil
		})
	}

	_ = g.Wait()
	log.Printf("Batch complete. Total: %d, Failures: %d", len(rerolls), failedCount)

}
