package architecture

//Db interface- accessor to access saved users
type Db interface {
	Save(p Person, n int)
	Retrieve(n int) Person
}

type personService struct {
	accesor Db
}

//Person type
type Person struct {
	First string
}

//Get Person from db
func Get(db Db, n int) Person {

	return db.Retrieve(n)
}

//Put person to db
func Put(d Db, n int, p Person) {
	d.Save(p, n)
}
