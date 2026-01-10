package service

import (
	"log"
	"secret-manager/pkg/persistence"

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

	log.Printf("Found %d pending rerolls", len(rerolls))

	for _, reroll := range rerolls {
		err := r.Service.FetchAndStoreTemplate(reroll)
		if err != nil {
			log.Println(reroll.FilePath + "can't be rerolled")
			continue
		}
		err = persistence.SaveSecretConfig(r.Db, reroll)
		if err != nil {
			log.Println(reroll.FilePath + "can't be resaved")
		}
	}

}
