package postgres

type PostgresStore struct {
	DB *PGCon
}

func (m *PostgresStore) Init() {
	m.DB = GetPGCon()
}
