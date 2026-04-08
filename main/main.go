package main

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	db := connect()
	defer db.Close()

	db.Exec("INSERT INTO users (name, email) VALUES ('adam', 'adam.com')")
	db.Exec("INSERT INTO users (name, email) VALUES ('adam2', 'adam2com')")
}
