package postgres


//PostgresDB like a map implemitatation
type PostgresDB map[int]Person

func (postgres PostgresDB) save(p Person, n int) {
	postgres[n] = p
}

func (postgres PostgresDB) retrieve(n int) Person {
	return postgres[n]
}