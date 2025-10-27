package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"

	"bookit.com/model"
	db "bookit.com/utils/database"
)

// nextWeekendDates returns the next Saturday and Sunday (00:00) in the provided location.
// If today is Saturday or Sunday, it returns the Saturday/Sunday of the next week.
func nextWeekendDates(loc *time.Location) (time.Time, time.Time) {
	now := time.Now().In(loc)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	weekday := int(now.Weekday())

	// days until next Saturday
	daysToSat := (int(time.Saturday) - weekday + 7) % 7
	if weekday == int(time.Saturday) || weekday == int(time.Sunday) {
		daysToSat += 7 // move to next week's weekend
	}

	sat := today.AddDate(0, 0, daysToSat)
	sun := sat.AddDate(0, 0, 1)
	return sat, sun
}

func main() {
	// load env if present
	_ = godotenv.Load()

	dbConn := db.DBconnect()
	fmt.Println("🚀 Initiating seeder...")

	// migrate models
	if err := dbConn.AutoMigrate(&model.User{}, &model.Facility{}, &model.Booking{}, &model.BookingSlot{}); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	fmt.Println("✅ Database migrated.")

	// timezone: Kuala Lumpur
	loc, err := time.LoadLocation("Asia/Kuala_Lumpur")
	if err != nil {
		log.Printf("⚠️ Failed to load Asia/Kuala_Lumpur timezone, fallback to local: %v", err)
		loc = time.Local
	}

	// compute next weekend (Saturday & Sunday)
	sat, sun := nextWeekendDates(loc)

	// helper to create hourly slots between startHour (inclusive) and endHour (exclusive)
	makeHourlySlots := func(facilityID uint, day time.Time, startHour, endHour int) []model.BookingSlot {
		var slots []model.BookingSlot
		for h := startHour; h < endHour; h++ {
			start := time.Date(day.Year(), day.Month(), day.Day(), h, 0, 0, 0, loc)
			end := time.Date(day.Year(), day.Month(), day.Day(), h+1, 0, 0, 0, loc)
			s := model.BookingSlot{
				FacilityID:  facilityID,
				StartTime:   start,
				EndTime:     end,
				IsAvailable: true,
			}
			slots = append(slots, s)
		}
		return slots
	}

	// create facilities
	facilities := []model.Facility{
		{Name: "Football Field A", Price: 80.0, Capacity: 22, Available: true},
		{Name: "Football Field B", Price: 80.0, Capacity: 22, Available: true},
		{Name: "Swimming Pool", Price: 120.0, Capacity: 30, Available: true},
	}

	for i := range facilities {
		if err := dbConn.Create(&facilities[i]).Error; err != nil {
			log.Fatalf("❌ Failed to create facility %s: %v", facilities[i].Name, err)
		}
		fmt.Printf("🏟️ Created facility: %s (ID=%d)\n", facilities[i].Name, facilities[i].ID)
	}

	// create hourly slots 08:00-12:00 (8-9,9-10,10-11,11-12) for both days
	totalSlots := 0
	for _, f := range facilities {
		// Saturday slots
		satSlots := makeHourlySlots(f.ID, sat, 8, 12)
		if len(satSlots) > 0 {
			if err := dbConn.Create(&satSlots).Error; err != nil {
				log.Fatalf("❌ Failed to create saturday slots for facility %d: %v", f.ID, err)
			}
			totalSlots += len(satSlots)
		}

		// Sunday slots
		sunSlots := makeHourlySlots(f.ID, sun, 8, 12)
		if len(sunSlots) > 0 {
			if err := dbConn.Create(&sunSlots).Error; err != nil {
				log.Fatalf("❌ Failed to create sunday slots for facility %d: %v", f.ID, err)
			}
			totalSlots += len(sunSlots)
		}

		fmt.Printf("🕒 Created %d hourly slots for facility %s\n", 4*2, f.Name) // 4 hours * 2 days
	}

	fmt.Printf("🎉 Seeding complete: %d facilities and %d booking slots created.\n", len(facilities), totalSlots)
}
