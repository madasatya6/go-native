package session

import (
	"log"
	"fmt"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/srinathgs/mysqlstore"
	"github.com/antonlindstrom/pgstore"
	"github.com/madasatya6/go-native/applications/config"
)

type Session struct{
	ID 				string 
	Type 			string 
	AuthKey			string 
	Encryption 		string
	Expired			int
	TimeForUpdate	int 
	Path 			string 
	HttpOnly		bool
}

type SessionType struct{
	Config		Session
	Cookie		*sessions.CookieStore
	MySQL 		*mysqlstore.MySQLStore
	Postgre 	*pgstore.PGStore
}

func (c *SessionType) NewCookieStore() {
	authKey := []byte(c.Config.AuthKey)
	encryptionKey := []byte(c.Config.Encryption)

	store := sessions.NewCookieStore(authKey,encryptionKey)
	store.Options.Path = c.Config.Path
	store.Options.MaxAge = c.Config.Expired
	store.Options.HttpOnly = c.Config.HttpOnly

	c.Cookie = store
}

func (m *SessionType) NewMysqlStore() {
	authKey := []byte(m.Config.AuthKey)
	//encryptionKey := []byte("my-encryption-key-very-secret123")
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

	store.Options.Path = m.Config.Path
	store.Options.MaxAge = m.Config.Expired
	store.Options.HttpOnly = m.Config.HttpOnly

	m.MySQL = store
} 

func (p *SessionType) NewPostgresStore() *pgstore.PGStore {
	

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