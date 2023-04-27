package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

var s *http.Server

func TestMain(m *testing.M){
	s = New(":8060")

	code := m.Run()
	os.Exit(code)
}


func Test_index(t *testing.T) {
		w :=  httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet,"/", nil)
		t.Run("indexRegTest", func(t *testing.T) {
			err := indexReg(w, r)
			if(err!=nil) {t.Errorf("Unexpected error: %s", err)}
		})
}

func Test_listInfo(t *testing.T) {
	godotenv.Load("../.env")
	w :=  httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet,"/info", nil)
	t.Run("listInfoTest", func(t *testing.T) {
		err := listInfo(w, r)
		if(err!=nil) {t.Errorf("Unexpected error: %s", err)}
	})
}

func Test_sendInfo(t *testing.T) {
	godotenv.Load("../.env")
	payloads := &Info{Name: "testFile",
	Value: "17",}
	body, err := json.Marshal(payloads)
	if(err != nil){
		t.Errorf("Unexpected error: %s", err)
	}

	w :=  httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodPost,"/info", strings.NewReader(string(body)))
	t.Run("sendInfoTest", func(t *testing.T) {
		err := sendInfo(w, r)
		if(err!=nil) {t.Errorf("Unexpected error: %s", err)}
	})
}


