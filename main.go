package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)


func main() {
    a := App{} 
    
    a.Initialize("root", "root", "gorestapidb")

    a.Run(":5555")
}

type App struct {
    Router *mux.Router
    DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
    connectionString := fmt.Sprintf(user+":"+password+"@tcp(localhost:8889)/"+dbname)
    
    var err error
    a.DB, err = sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatal(err)
    }
    a.Router = mux.NewRouter()
    a.initializeRoutes()
}

func (a *App) Run(addr string) {
    log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/songs", a.getSongs).Methods("GET")
    a.Router.HandleFunc("/song", a.createSong).Methods("POST")
    a.Router.HandleFunc("/song/{id:[0-9]+}", a.getSong).Methods("GET")
    a.Router.HandleFunc("/song/{id:[0-9]+}", a.updateSong).Methods("PUT")
    a.Router.HandleFunc("/song/{id:[0-9]+}", a.deleteSong).Methods("DELETE")
}


func (a *App) getSongs(w http.ResponseWriter, r *http.Request) {
    count, _ := strconv.Atoi(r.FormValue("count"))
    start, _ := strconv.Atoi(r.FormValue("start"))
    if count > 10 || count < 1 {
        count = 10
    }
    if start < 0 {
        start = 0
    }
    songs, err := getSongs(a.DB, start, count)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, songs)
}

func (a *App) getSong(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid song ID")
        return
    }
    s := song{ID: id}
    if err := s.getSong(a.DB); err != nil {
        switch err {
        case sql.ErrNoRows:
            respondWithError(w, http.StatusNotFound, "Song not found")
        default:
            respondWithError(w, http.StatusInternalServerError, err.Error())
        }
        return
    }
    respondWithJSON(w, http.StatusOK, s)
}



func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}


func (a *App) createSong(w http.ResponseWriter, r *http.Request) {
    var s song
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&s); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()
    if err := s.createSong(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusCreated, s)
}

func (a *App) updateSong(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid song ID")
        return
    }
    var s song
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&s); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
        return
    }
    defer r.Body.Close()
    s.ID = id
    if err := s.updateSong(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, s)
}

func (a *App) deleteSong(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid Song ID")
        return
    }
    s := song{ID: id}
    if err := s.deleteSong(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}


type song struct {
    ID    int     `json:"id"`
    Title  string `json:"title"`
    Singer string `json:"singer"`
}

func (s *song) getSong(db *sql.DB) error {
    statement := fmt.Sprintf("SELECT title, singer FROM songs WHERE id=%d", s.ID)
    return db.QueryRow(statement).Scan(&s.Title, &s.Singer)
}
func (s *song) updateSong(db *sql.DB) error {
    statement := fmt.Sprintf("UPDATE songs SET title='%s', singer='%s' WHERE id=%d", s.Title, s.Singer, s.ID)
    _, err := db.Exec(statement)
    return err
}
func (s *song) deleteSong(db *sql.DB) error {
    statement := fmt.Sprintf("DELETE FROM songs WHERE id=%d", s.ID)
    _, err := db.Exec(statement)
    return err
}
func (s *song) createSong(db *sql.DB) error {
    statement := fmt.Sprintf("INSERT INTO songs(title, singer) VALUES('%s', '%s')", s.Title, s.Singer)
    _, err := db.Exec(statement)
    if err != nil {
        return err
    }
    err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&s.ID)
    if err != nil {
        return err
    }
    return nil
}
func getSongs(db *sql.DB, start, count int) ([]song, error) {
    statement := fmt.Sprintf("SELECT id, title, singer FROM songs LIMIT %d OFFSET %d", count, start)
    rows, err := db.Query(statement)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    songs := []song{}
    for rows.Next() {
        var s song
        if err := rows.Scan(&s.ID, &s.Title, &s.Singer); err != nil {
            return nil, err
        }
        songs = append(songs, s)
    }
    return songs, nil
}
