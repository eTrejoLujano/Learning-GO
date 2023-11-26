package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
const (
  host     = "localhost"
  port     = 5432
  user     = "eriktrejolujano"
  password = ""
  dbname   = "recordings"
)

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}
  var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	var db, err = sql.Open("postgres", psqlInfo)

func main() {
  if err != nil {
    panic(err)
  }
  defer db.Close()
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  fmt.Println("Established a successful connection!")
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []Album
    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}