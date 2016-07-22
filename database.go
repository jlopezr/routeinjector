package routeinjector

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//StartDatabase connects to the Mongo Database
func (m *RouteInjector) StartDatabase() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	m.Db = session
}

//StopDatabase closes connections with Mongo Database
func (m *RouteInjector) StopDatabase() {
	m.Db.Close()
}

//Person is an example schema
type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

// The DB should be droped at startup
var (
	IsDrop = true
)

//TestMongo is an initial example of MGO usage
func TestMongo() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Drop Database
	if IsDrop {
		err = session.DB("test").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// Collection People
	c := session.DB("test").C("people")

	// Index
	index := mgo.Index{
		Key:        []string{"name", "phone"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	// Insert Datas
	err = c.Insert(
		&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
		&Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})

	if err != nil {
		panic(err)
	}

	// Query One
	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"_id": 0, "phone": 1}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Phone %+v\n", result.Phone)

	result2 := bson.M{}
	err = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"_id": 0, "phone": 1}).One(&result2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Phone %+v\n", result2)

	// Query All
	var results []Person
	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	// Update
	colQuerier := bson.M{"name": "Ale"}
	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	err = c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}

	// Query All
	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)
}
