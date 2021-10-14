package config

import "fmt"

/****
* Jenis penyimpanan session :
* cookie dismpan dalam cookie
* mysql disimpan dalam database mysql
* postgres disimpan dalam database postgres 
*/
var TypeSession = "cookie" 

var SessionEncryption = "my-auth-key-very-secret"
var SessionExpired = 86400 //sehari
var SessionTimeForUpdate = 14400

