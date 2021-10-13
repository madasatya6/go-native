package config

import "database/sql"
import "time"

//database
var MySQL *sql.DB
var Postgree *sql.DB

//time 
var DefaultLocation = "Asia/Jakarta"
var TimeZone *time.Location

