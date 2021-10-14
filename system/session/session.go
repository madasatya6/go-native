package session

import (
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/srinathgs/mysqlstore"
	"github.com/antonlindstrom/pgstore"
	"log"
	"fmt"
	"os"
)

const SESSION_ID = "id"

//pilih session yang digunakan
var SessionStore = newCookieStore()

func newCookieStore() *sessions.CookieStore {
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store := sessions.NewCookieStore(authKey,encryptionKey)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7 //expired dalam seminggu
	store.Options.HttpOnly = true

	return store
}

func newMysqlStore() *mysqlstore.MySQLStore {
	authKey := []byte("my-auth-key-very-secret")
	//encryptionKey := []byte("my-encryption-key-very-secret123")

	//store, err := mysqlstore.NewMySQLStore("UN:PASS@tcp(<IP>:<PORT>)/<DB>?parseTime=true&loc=Local", <tablename>, "/", 3600, []byte("<SecretKey>"))
	cs := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local", 
		GetJson["mysql"]["username"], 
		GetJson["mysql"]["password"],
		GetJson["mysql"]["host"],
		GetJson["mysql"]["port"],
		GetJson["mysql"]["dbname"],
	)

	store, err := mysqlstore.NewMySQLStore(cs, 
		"sessionstore", 
		"/", 
		3600, 
		authKey)

	if err != nil {
		log.Println(err.Error())
	}

	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7 //expired dalam seminggu
	store.Options.HttpOnly = true

	return store
} 

func newPostgresStore() *pgstore.PGStore {
	

	//url := "postgres://novalagung:@127.0.0.1:5432/novalagung?sslmode=disable"
	url := "dbname=ecommerce user=postgres password=lampupijar77 host=localhost sslmode=disable"

	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store, err := pgstore.NewPGStore(url, authKey, encryptionKey)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(0)
	}

	return store
}

func SetFlashdata(c echo.Context, name, value string){
	session, _ := SessionStore.Get(c.Request(), "fmessages")
	session.AddFlash(value, name)

	session.Save(c.Request(), c.Response())
}

func GetFlashdata(c echo.Context, name string) []string {
	session, _ := SessionStore.Get(c.Request(), "fmessages")
	fm := session.Flashes(name)
	//IF we have some message

	if len(fm) > 0 {
		session.Save(c.Request(), c.Response())
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

func init (){
	E.Use(echo.WrapMiddleware(context.ClearHandler))
}