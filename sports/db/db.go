package db

import (
	"time"

	"syreclabs.com/go/faker"
)

func (r *eventsRepo) seed() error {
	// create new table if one does not exist and populate with dummy data"
	statement, err := r.db.Prepare(`CREATE TABLE IF NOT EXISTS sports (id INTEGER PRIMARY KEY, name TEXT, advertised_start_time DATETIME)`)
	if err == nil {
		_, err = statement.Exec()
	}

	// Create array that contains valid event names
	eventArray := []string{"AMERICAN FOOTBALL", "AUSTRALIAN RULES", "BASEBALL", "BASKETBALL", "BOXING",
		"CRICKET", "CYCLING", "DARTS", "ESPORTS", "GAELIC SPORTS", "GOLF", "HANDBALL", "ICE HOCKEY",
		"MIXED MARTIAL ARTS", "MOTOR SPORT", "NETBALL", "NOVELTY", "POLITICS", "POOL", "RUGBY LEAGUE",
		"RUGBY UNION", "SNOOKER", "SOCCER", "SURFING", "TABLE TENNIS", "TENNIS", "VOLLEYBALL"}

	for i := 1; i <= len(eventArray); i++ {
		statement, err = r.db.Prepare(`INSERT OR IGNORE INTO sports(id, name, advertised_start_time) VALUES (?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,
				eventArray[i-1],
				faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2)).Format(time.RFC3339),
			)
		}
	}

	return err
}
