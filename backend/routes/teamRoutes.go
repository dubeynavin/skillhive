package routes

import (
	"net/http"

	"freelance-connect-backend/controllers"
)

func TeamRoutes() {
	http.HandleFunc("/api/team", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			controllers.GetTeamMembers(w, r)
		} else if r.Method == "POST" {
			controllers.CreateTeamMember(w, r)
		}
	})
	http.HandleFunc("/api/team/id", controllers.GetTeamMemberByID)
}
