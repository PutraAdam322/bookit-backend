package service

import (
	"fmt"
	"log"
	"time"

	"bookit.com/model"
	"gorm.io/gorm"

	"github.com/robfig/cron/v3"
)

func StartBookingWorker(db *gorm.DB) {
	loc, err := time.LoadLocation("Asia/Kuala_Lumpur")
	if err != nil {
		log.Printf("failed to load timezone, using local: %v", err)
		loc = time.Local
	}

	c := cron.New(cron.WithLocation(loc))

	spec := "0 0 * * 1"
	_, err = c.AddFunc(spec, func() {
		log.Println("[worker] triggered createWeekendSlots:", time.Now().In(loc).Format(time.RFC3339))
		if err := CreateWeekendSlots(db, loc); err != nil {
			log.Printf("[worker] CreateWeekendSlots error: %v\n", err)
		} else {
			log.Println("[worker] CreateWeekendSlots done")
		}
	})
	if err != nil {
		log.Printf("[worker] failed to add cron job: %v\n", err)
		return
	}

	c.Start()
	log.Println("[worker] cron started")
}

func CreateWeekendSlots(db *gorm.DB, loc *time.Location) error {
	if db == nil {
		return fmt.Errorf("db is nil")
	}

	now := time.Now().In(loc)
	weekday := now.Weekday()
	daysToSat := (int(time.Saturday) - int(weekday) + 7) % 7
	sat := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc).AddDate(0, 0, daysToSat)

	dates := []time.Time{
		time.Date(sat.Year(), sat.Month(), sat.Day(), 0, 0, 0, 0, loc),
		time.Date(sat.Year(), sat.Month(), sat.Day()+1, 0, 0, 0, 0, loc),
	}

	startHours := []int{8, 9, 10, 11}

	var facilities []model.Facility
	if err := db.Find(&facilities).Error; err != nil {
		return fmt.Errorf("failed to fetch facilities: %w", err)
	}
	if len(facilities) == 0 {
		log.Println("[worker] no facilities found")
		return nil
	}

	created := 0
	skipped := 0

	for _, f := range facilities {
		for _, d := range dates {
			for _, h := range startHours {
				start := time.Date(d.Year(), d.Month(), d.Day(), h, 0, 0, 0, loc)
				end := start.Add(time.Hour)

				var exist model.BookingSlot
				err := db.Where("facility_id = ? AND start_time = ?", f.ID, start).First(&exist).Error
				if err == nil {
					skipped++
					continue
				}
				if err != gorm.ErrRecordNotFound {
					log.Printf("[worker] lookup error facility %d start %s: %v", f.ID, start, err)
					continue
				}

				slot := model.BookingSlot{
					FacilityID:  f.ID,
					StartTime:   start,
					EndTime:     end,
					IsAvailable: true,
				}
				if err := db.Create(&slot).Error; err != nil {
					log.Printf("[worker] failed to create slot facility=%d start=%s err=%v", f.ID, start, err)
					continue
				}
				created++
			}
		}
	}

	log.Printf("[worker] weekend slots created=%d skipped=%d for weekend starting %s\n", created, skipped, sat.Format("2006-01-02"))
	return nil
}
