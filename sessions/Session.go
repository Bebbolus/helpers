package helpers

import (
	"net/http"
	h "github.com/Bebbolus/helpers/crypt"
	"encoding/json"
	b64 "encoding/base64"
	"errors"
)

type Session struct{
	Content map[string]interface{}
	Crypt string
	Key string
}

func (s *Session) SetItem(key string, value interface{}){
	s.Decrypt()
	s.Content[key] = value
	s.Encrypt()
}

func (s *Session) Encrypt(){
	marshalled, _ := json.Marshal(s.Content)					
	s.Crypt = b64.StdEncoding.EncodeToString(h.Encrypt(marshalled, s.Key))
}

func (s *Session) Decrypt(){
	b64tobyte, _ := b64.StdEncoding.DecodeString(s.Crypt)
	b := h.Decrypt(b64tobyte, s.Key)
	s.Content = make(map[string]interface{})
	_ = json.Unmarshal(b, &s.Content)
}

func (s *Session) Init() error{
	if len(s.Content) == 0{
		s.Content = make(map[string]interface{})
		s.Content["sessionid"] =  h.MakeUUID()
		s.Encrypt()
		return nil
	} else {
		return errors.New("TRYING TO INIT A NO EMPTY CONTENT")
	}
}

func (s *Session) SetHeader(r *http.Request){
	s.Encrypt()
	r.Header.Set("alphasession",s.Crypt )
}

func (s *Session) GetFromHeader(r *http.Request) error{
	if r.Header.Get("alphasession") == "" {
		return errors.New("No session stored")
	} else {
		s.Crypt = r.Header.Get("alphasession")
		s.Decrypt()
		return nil
	}		
}

func (s *Session) StoreForm(r *http.Request){
	s.Decrypt()
	for key, values := range r.Form {   // range over map
		for _, value := range values {    // range over []string
			s.Content[key] =value
		}
	}
	s.SetHeader(r)
}