package service

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func StartTaskCleanupScheduler(stop chan struct{}) {
	ticker := time.NewTicker(24 * time.Second)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				log.Println("Starting removal of outdated tasks..")
				//DeleteExpiredTasks(db)

			case <-stop:
				log.Println("Stopping background task of removal...")
				return
			}
		}
	}()
}
