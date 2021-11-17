package postgres

type Store struct {
	DB *PGCon
}

func (m *Store) Init() {
	m.DB = GetPGCon()
}
