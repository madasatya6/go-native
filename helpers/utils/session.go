package utils

import (
	"net/http"
	
    "github.com/gorilla/sessions"
	"github.com/srinathgs/mysqlstore"
	"github.com/antonlindstrom/pgstore"
)

/****
* Default session is cookie
*/
var SessionCookie *sessions.CookieStore
var SessionMySQL *mysqlstore.MySQLStore
var SessionPostgre *pgstore.PGStore

//set type session
var SessionStore = SessionCookie

func SetFlashdata(w http.ResponseWriter, r *http.Request, name, value string){
	session, _ := SessionStore.Get(r, "fmessages")
	session.AddFlash(value, name)

	session.Save(r, w)
}

func GetFlashdata(w http.ResponseWriter, r *http.Request, name string) []string {
	session, _ := SessionStore.Get(r, "fmessages")
	fm := session.Flashes(name)
	//IF we have some message

	if len(fm) > 0 {
		session.Save(r, w)
		//initiate a strings slice to return messages
		var flashes []string 
		for _, fl := range fm {
			//Add message to the slice
			flashes = append(flashes, fl.(string))
		}
		
		return flashes
	}

	return nil
}



