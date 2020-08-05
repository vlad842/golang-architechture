package mongo
import (
	"github.com/golang-architechture"
)

// Db map
type Db map[int] Person

//Save in MongoDB
func (mongo Db) save(p Person, n int) {
	mongo[n] = p
}

func (mongo Db) retrieve(n int) Person {
	return mongo[n]
}
