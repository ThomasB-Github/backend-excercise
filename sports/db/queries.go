package db

const (
	eventsList = "list"
)

func getEventsQueries() map[string]string {
	return map[string]string{
		eventsList: `
			SELECT 
				id, 
				name, 
				advertised_start_time
			FROM sports
		`,
	}
}
