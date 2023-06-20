package logstore

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/akxcix/log-router/pkg/database"
	"github.com/akxcix/log-router/pkg/database/models"
)

const (
	delemiter    string        = ":|:|:|:|:|:"
	timeInterval time.Duration = 10 * time.Second
	bufferLength int           = 10_000
)

type Store struct {
	database *database.Database
	ticker   <-chan time.Time
	buffer   chan string
	mutex    sync.Mutex
}

func NewStore() *Store {
	db := database.NewDatabase()
	ticker := time.NewTicker(timeInterval)
	buffer := make(chan string, bufferLength)

	store := &Store{
		database: db,
		ticker:   ticker.C,
		buffer:   buffer,
	}

	return store
}

func (s *Store) StartProcessing() {
	go s.process()
}

func (s *Store) process() {
	file, err := os.OpenFile("logEvents.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	for {
		select {
		case <-s.ticker:
			fmt.Println("received ticker")
			s.mutex.Lock()
			data, err := ioutil.ReadFile("logEvents.txt")
			if err != nil {
				fmt.Println("Error reading the file:", err)
				s.mutex.Unlock()
				continue
			}
			events := strings.Split(string(data), delemiter)
			if len(events) > 0 {
				events = events[:len(events)-1]
			}

			if len(events) > 0 {
				go s.SaveToDB(events)
				s.clearFile(file)
			}
			s.mutex.Unlock()
		case msg := <-s.buffer:
			s.mutex.Lock()
			_, err := file.WriteString(fmt.Sprintf("%s%s", msg, delemiter))
			if err != nil {
				fmt.Println("Error writing to the file:", err)
			}
			file.Sync()
			s.mutex.Unlock()
		}
	}
}

func (s *Store) Save(logEvent string) error {
	s.buffer <- logEvent
	return nil
}

func (s *Store) clearFile(file *os.File) {
	err := file.Truncate(0)
	if err != nil {
		fmt.Println("Error truncating the file:", err)
		return
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking the file:", err)
	}
}

func (s *Store) SaveToDB(logevents []string) {
	dbEvents := make([]models.Event, 0)
	for _, logEvent := range logevents {
		dbEvent := models.Event{Log: logEvent}
		dbEvents = append(dbEvents, dbEvent)
	}

	if len(dbEvents) > 0 {
		s.database.Save(dbEvents)
	}
}
