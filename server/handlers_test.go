package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var s *http.Server

func TestMain(m *testing.M){
	s = New(":8060")

	code := m.Run()
	os.Exit(code)
}


func Test_index(t *testing.T) {
        type args struct {
                w http.ResponseWriter
                r *http.Request
        }
        tests := []struct {
                name string
                args args
        }{
                {
					name: "index",
					args: args{
						httptest.NewRecorder(),
						r, _:= http.NewRequest("GET","/"),
					},

				},
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        index(tt.args.w, tt.args.r)
                })
        }
}

func Test_info(t *testing.T) {
        type args struct {
                w http.ResponseWriter
                r *http.Request
        }
        tests := []struct {
                name string
                args args
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        info(tt.args.w, tt.args.r)
                })
        }
}

func Test_listInfo(t *testing.T) {
        type args struct {
                w http.ResponseWriter
                r *http.Request
        }
        tests := []struct {
                name string
                args args
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        listInfo(tt.args.w, tt.args.r)
                })
        }
}

func Test_sendInfo(t *testing.T) {
        type args struct {
                w http.ResponseWriter
                r *http.Request
        }
        tests := []struct {
                name string
                args args
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        sendInfo(tt.args.w, tt.args.r)
                })
        }
}

func Test_getDB(t *testing.T) {
        tests := []struct {
                name string
                want []*Info
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        if got := getDB(); !reflect.DeepEqual(got, tt.want) {
                                t.Errorf("getDB() = %v, want %v", got, tt.want)
                        }
                })
        }
}

func Test_saveDB(t *testing.T) {
        type args struct {
                i *Info
        }
        tests := []struct {
                name string
                args args
        }{
                // TODO: Add test cases.
        }
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        saveDB(tt.args.i)
                })
        }
}