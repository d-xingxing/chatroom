package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"chatroom/logic"
	"html/template"
)

func homeHandleFunc(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("template/home.html")
	if err != nil {
		fmt.Fprint(w, "模板解析错误！")
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Fprint(w, "模板执行错误！")
		return
	}
}

func userListHandleFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	userList := logic.Broadcaster.GetUserList()
	b, err := json.Marshal(userList)

	if err != nil {
		fmt.Fprint(w, `[]`)
	} else {
		fmt.Fprint(w, string(b))
	}
}
