package sqlite

func tables() []string {
	createScheduleRequestTableQuery := `
		CREATE TABLE IF NOT EXISTS scheduled_requests (
		    id TEXT PRIMARY KEY,
		    title TEXT NOT NULL,
		    description TEXT,
		    status NUMERIC NOT NULL,
		    url TEXT NOT NULL,
		    payload JSON,
		    header JSON,
		    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    scheduled_at DATETIME NOT NULL,
		    error TEXT,
			deleted_at DATETIME,
			executed_at DATETIME
		);
		CREATE INDEX IF NOT EXISTS idx_scheduled_requests_status ON scheduled_requests(status);
		CREATE INDEX IF NOT EXISTS idx_scheduled_requests_scheduled_at ON scheduled_requests(scheduled_at);`

	return []string{
		createScheduleRequestTableQuery,
	}
}
