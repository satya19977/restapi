package models

import (
	"fmt"
	"github.com/satya19977/restapi/db"
	"time"

	// "github.com/pelletier/go-toml/query"
	// "go.mongodb.org/mongo-driver/event"
)

type Event struct {
	ID          int64
	Name        string `binding:required`
	Description string `binding:required`
	Location    string `binding:required`
	DateTime    time.Time `binding:required`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `INSERT INTO events(name,description,location,dateTime,user_id)
	VALUES(?,?,?,?,?)` // SQL injection sage way, the resaon ? this is used
	stmt,err := db.DB.Prepare(query) // prepares a sql statement, can use exec but this is better performant
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserID)
	if err != nil{  // Exec is used when you want to insert/update data
		return err
	}
	id, err := result.LastInsertId() // Generates an ID which is unique
	e.ID = id // Assign that generated ID to the id of struct
	return err
}

func GetAllEvents() ([]Event,error){
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query) // Query is used when you want to fetch data
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next(){
		var event Event
		err := rows.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)
		if err != nil{
			return nil,err
		}
		events = append(events, event)
	}
	fmt.Println(events)

	return events, nil
	
}

// This query fetches data by ID
func GetEventByID(id int64) (*Event, error){
	query := "SELECT * FROM events WHERE id = ?" // ? is used to protect against sql injection
	row := db.DB.QueryRow(query,id)
	var event Event
	err := row.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)
	if err != nil {
		return nil,err
	}
	return &event,nil
}
// This method works to update a request
func (event Event)  Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil{
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name,event.Description,event.Location,event.DateTime,event.ID)
	return err
}

//This method workds to delete a request
func(event Event) Delete() error {
	query := `
	DELETE FROM events
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}
