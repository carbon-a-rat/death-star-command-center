package main

import (
	"log"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
)

// LauncherHeartbeatJob starts a ticker that checks heartbeats every 10 seconds for a minute.
func LauncherHeartbeatJob(app *pocketbase.PocketBase) {

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// Counter to track the number of iterations
	counter := 0
	maxIterations := 6

	for range ticker.C {
		//log.Println("Running launcher heartbeat checker...")

		// Increment the counter
		counter++
		if counter >= maxIterations {
			//log.Println("Max iterations reached. Stopping ticker.")
			break
		}

		// Define the timeout duration
		timeout := time.Second * 10
		cutoff := time.Now().Add(-timeout)

		// Fetch all launchers
		records, err := app.FindAllRecords("launchers", dbx.NewExp("online = true OR current_user != \"\""))

		if err != nil {
			log.Printf("Error fetching launchers: %v", err)
			continue
		}

		for _, record := range records {
			lastPingAt := record.GetDateTime("last_ping_at")

			online := lastPingAt.Time().After(cutoff)

			currentUser := record.GetString("current_user")

			if !online {
				record.Set("online", false)

				record.Set("current_user", nil)

				err = app.Save(record)
				if err != nil {
					log.Printf("Error updating launcher %s: %v", record.GetString("id"), err)
					continue
				}
				//log.Printf("Updated launcher %s online status to %v", record.GetString("id"), online)
			} else if currentUser != "" {
				userLastPingAt := record.GetDateTime("last_user_ping_at")
				userOnline := userLastPingAt.Time().After(cutoff)
				if !userOnline {
					record.Set("current_user", nil)
					err = app.Save(record)
					if err != nil {
						log.Printf("Error updating launcher %s current user status: %v", record.GetString("id"), err)
						continue
					}
					//log.Printf("Updated launcher %s current user status to nil", record.GetString("id"))
				}
			}
		}
	}
}

func RegisterLauncherHeartbeatJob(app *pocketbase.PocketBase) {

	// Register the heartbeat job
	app.Cron().MustAdd("heartbeat", "*/1 * * * *", func() {
		// This function will be called every minute
		//log.Println("Running launcher heartbeat job...")
		// Call the heartbeat checker function
		LauncherHeartbeatJob(app)
		//log.Println("Launcher heartbeat job completed.")
	})

}
