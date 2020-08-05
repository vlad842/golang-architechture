package mongo
import (
	"github.com/golang-architechture"
)

// Db map
type Db map[int] architecture.Person

//Save in MongoDB
func (mongo Db) Save(p architecture.Person, n int) {
	mongo[n] = p
}

//Retrieve from db
func (mongo Db) Retrieve(n int) architecture.Person {
	return mongo[n]
}
