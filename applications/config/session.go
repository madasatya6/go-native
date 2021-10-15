package config

/****
* Jenis penyimpanan session :
* cookie dismpan dalam cookie
* mysql disimpan dalam database mysql
* postgres disimpan dalam database postgres 
*/
var TypeSession = "cookie" 

var SessionEncryption = "my-encryption-key-very-secret123"
var SessionAuthKey = "my-auth-key-very-secret"
var SessionExpired = 86400 //sehari
var SessionTimeForUpdate = 14400

var SessionID = "ID"
var SessionPath = "/"
var HttpOnly = true

