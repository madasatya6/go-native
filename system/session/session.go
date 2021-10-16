package session

import (
	"log"
	"os"
	"fmt"
	
	"github.com/gorilla/sessions"
	"github.com/srinathgs/mysqlstore"
	"github.com/antonlindstrom/pgstore"
	"github.com/madasatya6/go-native/applications/config"
	"github.com/madasatya6/go-native/system/conf"
	"github.com/madasatya6/go-native/helpers/utils"
)

var SessionConfig Session
var SessType SessionType

type Session struct{
	ID 				string 
	AuthKey			string 
	Encryption 		string
	Expired			int
	TimeForUpdate	int 
	Path 			string 
	HttpOnly		bool
}

type SessionType struct{
	Config		*Session
	Cookie		*sessions.CookieStore
	MySQL 		*mysqlstore.MySQLStore
	Postgre 	*pgstore.PGStore
}

func (c *SessionType) NewConfig(cfg *Session) {
	c.Config = cfg
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

func (m *SessionType) NewMysqlStore(cfg conf.Configuration) {
	env := cfg.Database.MySQL
	authKey := []byte(m.Config.AuthKey)
	//encryptionKey := []byte(m.Config.Encryption)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Local", 
		env.Username, 
		env.Password,
		env.Host,
		env.Port,
		env.DBName,
	)

	store, err := mysqlstore.NewMySQLStore(dsn, 
		"sessionstore", 
		"/", 
		m.Config.Expired, 
		authKey)

	if err != nil {
		log.Println(err.Error())
	}

	store.Options.Path = m.Config.Path
	store.Options.MaxAge = m.Config.Expired
	store.Options.HttpOnly = m.Config.HttpOnly

	m.MySQL = store
} 

func (p *SessionType) NewPostgresStore(cfg conf.Configuration) {
	env := cfg.Database.Postgre
	var url = fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=%s",
				env.DBName, env.Username, env.Password, env.Host, env.SSLMode)

	authKey := []byte(p.Config.AuthKey)
	encryptionKey := []byte(p.Config.Encryption)

	store, err := pgstore.NewPGStore(url, authKey, encryptionKey)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(0)
	}

	p.Postgre = store
}

func (c *SessionType) Send() {
	utils.SetCookie(c.Cookie)
	utils.SetMySQL(c.MySQL)
	utils.SetPostgre(c.Postgre)
}

func Init() *SessionType {
	
	//set config from applications/config
	SessType.NewConfig(&Session{
		ID: config.SessionID,
		AuthKey: config.SessionAuthKey,
		Encryption: config.SessionEncryption,
		Expired: config.SessionExpired,
		TimeForUpdate: config.SessionTimeForUpdate,
		Path: config.SessionPath,
		HttpOnly: config.HttpOnly,
	})
	SessType.NewCookieStore()
	SessType.Send()

	//var env = conf.Config
	//SessType.NewMysqlStore(env)
	//SessType.NewPostgresStore(env)
	
	return &SessType
}
