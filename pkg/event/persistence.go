package event

import (
	"fmt"
	"log"
	"organizer/db"
	"time"
)

// CreateEvent func
func (e *Event) CreateEvent(uid interface{}) (int, error) {

	fmt.Println("UID IS NOW:", uid)

	time := time.Now().Format("_2-Jan-2006 15:04")

	stmt, err := db.Get().Prepare(insertEvent)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(e.Title, e.Description, time)
	if err != nil {
		return 0, err
	}
	r, err := res.LastInsertId()

	if err != nil {
		log.Printf("%s", err)
		return 0, err
	}

	stmt2, err := db.Get().Prepare(userEventPiv)
	defer stmt.Close()
	if err != nil {
		log.Printf("%s", err)
		return 0, err
	}

	_, err2 := stmt2.Exec(uid, r)
	if err2 != nil {
		log.Printf("%s", err2)
		return 0, err2
	}

	return int(r), nil

}

const insertEvent = `INSERT INTO events (title, description, date) VALUES (?, ?, ?)`
const userEventPiv = `INSERT INTO user_event_pivot (user_id, event_id) VALUES (?, ?)`
