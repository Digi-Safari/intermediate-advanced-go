package main

func main() {
	c := Conf{
		Host: "localhost",
		Port: 3000,
	}
	CreatesConnection(c)
}

type Conf struct {
	Host     string
	Port     int
	User     string
	Password string
}

func CreatesConnection(c Conf) {

	if c.User == "" {
		c.User = "root"
	}
	if c.Port == 0 {
		c.Port = 3306
	}

}
