package utils

import (
    "html/template"
	"fmt"
	"time"
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
	"log"
	"reflect"

	"github.com/gorilla/mux"
)

var Response http.ResponseWriter
var Request *http.Request

func SetContext(w http.ResponseWriter, r *http.Request){
	Response = w 
	Request = r
}

var FuncMap = template.FuncMap{
    "sumi": func(a,b int) int {
		return a+b
	},
	"sumf": func(a,b float64) float64 {
		return a+b
	},
	"mini": func(a,b int) int {
		return a-b
	},
	"minf": func(a,b float64) float64 {
		return a-b
	},
	"xi": func(a,b int) int {
		return a*b
	},
	"xf": func(a,b float64) float64 {
		return a*b
	},
	"di": func(a,b int) int {
		return a/b
	},
	"df": func(a,b float64) float64 {
		return a/b
	},
	"strrmvarray": func(data []string) string {
		var result string
		for i, kata := range data{
			if i == 0 {
				result = fmt.Sprintf("%v", kata)
			} else {
				result = fmt.Sprintf("%v, %v", result, kata)
			}
		}
		return result
	},
	"avg": func(n ...int) int {
		var total = 0
		for _, angka := range n {
			total += angka
		}
		return total/len(n)
	},
	"unescape": func(s string) template.HTML{
		return template.HTML(s)
	},
	"TimetoYmd": func(tm time.Time) string {
		if tm.IsZero() {
			return ""
		} else {
			return fmt.Sprintf("%v-%v-%v %v:%v:%v", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
		}
		
	},
	"Iterate" : func(start uint, sign string, count uint) []uint {
		var i uint 
		var Items []uint 
		if sign == "<" {
			for i=start; i < count; i++ {
				Items = append(Items, i)
			}
		} else if sign == "<=" {
			for i=start; i <= count; i++ {
				Items = append(Items, i)
			}
		} else if sign == ">" {
			for i=start; i > count; i++ {
				Items = append(Items, i)
			}
		} else if sign == ">=" {
			for i=start; i >= count; i++ {
				Items = append(Items, i)
			}
		}
		return Items
	},
	"IndexString" : func(data []string, i uint) string {
		if len(data) > 0 {
			var index = int(i)
			if len(data) >= index { 
				return data[index]
			} else {
				return ""
			}
		} else { 
			return ""
		}
	},
	"Date" : func(date time.Time) string {
		format := fmt.Sprintf("%v %v %v", date.Day(), date.Month(), date.Year())
		return format
	},
	"Time" : func(date time.Time) string {
		format := date.Format("15:04")
		return format
	},
	"DateTime" : func(date time.Time) string {
		format := date.Format("2006-01-02 15:04")
		return format
	},
	"GetFlashdata" : func(w http.ResponseWriter, r *http.Request, key string) string {
		var sess = GetFlashdata(w, r, key)
		var data = ArraytoString(sess)
		return data
	},
	"UnmarshalJSONSession" : func(w http.ResponseWriter, r *http.Request, data string) (*map[string]interface{}) {
		var sess = ArraytoString(GetFlashdata(w, r, data))
		var temp = new(map[string]interface{})
		err := json.Unmarshal([]byte(sess), &temp)
		if err != nil {
			log.Println(err.Error())
		}
		return temp
	},
	"FormError" : func(key string) string {
		var keyMsg = key+"-msg"
		var sess = GetFlashdata(Response, Request, keyMsg)
		var data = ArraytoString(sess)
		return data
	},
	"SetFlashdata" : func(w http.ResponseWriter, r *http.Request, key string, value string) bool {
		SetFlashdata(w, r, key, value)
		return true
	},
	"SliceQueryParam" : func(r *http.Request, key string) []string {
		var param, ok = r.URL.Query()[key]
		if ok {
			return param
		}
		return nil
	},
	"QueryParam" : func(r *http.Request, key string) string {
		var param = r.URL.Query().Get(key)
		return param
	},
	"Param" : func(r *http.Request, key string) string {
		var queryParam = mux.Vars(r)
		return queryParam[key]
	},
	"FindString" : func(sentence, key string) bool {
		var status = strings.Index(sentence, key)
		if status == -1 {
			return false
		} else {
			return true
		}
	},
	"IntToString" : func(num int) string {
		return fmt.Sprintf("%d", num)
	},
	"PaginatorIsActive": func(page int) bool {
		urlpath := ""
		u, _ := url.Parse(urlpath)
		values, _ := url.ParseQuery(u.RawQuery)
		p := values.Get("p")
		pageNow := fmt.Sprintf("%d", page)
		if pageNow == p {
			return true
		} else {
			return false
		}
	},
	"equal" : func(param1 string, param2 string) bool {
		if param1 != param2 {
			return false
		} else {
			return true
		}
	},
	"in_array": func(val interface{}, array interface{}) (exists bool) {
		exists = false
		switch reflect.TypeOf(array).Kind() {
			case reflect.Slice :
				s := reflect.ValueOf(array)
				for i:=0; i < s.Len(); i++ {
					if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
						exists = true 
						return 
					}
				}
		}
		return
	},
}	

func ArraytoString(array []string) string {
	var result string
	for i, kata := range array{
		if i == 0 {
			result = fmt.Sprintf("%v", kata)
		} else {
			result = fmt.Sprintf("%v, %v", result, kata)
		}
	}
	return result
}
