package database

type Conf struct {
	db string
}

func New(conn string) *Conf {
	if conn == "" {
		panic("conn string is empty")
	}
	return &Conf{db: conn}
}
