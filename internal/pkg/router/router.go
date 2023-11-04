package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"gitlab.ozon.dev/kavkazov/homework-8/internal/pkg/server"
)

func RootRouter(svr *server.Server, mwf ...mux.MiddlewareFunc) *mux.Router {
	router := mux.NewRouter()
	router.Use(mwf...)

	postRouter := router.PathPrefix("/post/").Subrouter()
	initPostRouter(postRouter, svr)

	commentRouter := router.PathPrefix("/comment/").Subrouter()
	initCommentRouter(commentRouter, svr)
	return router
}

func initPostRouter(router *mux.Router, svr *server.Server) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

			id, status := svr.ParseGetID(r)
			if status != http.StatusOK {
				w.WriteHeader(status)
			}
			data, status := svr.GetPost(r.Context(), id)
			dataJson, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(status)
			w.Write(dataJson)

		case http.MethodPost:

			post, status := svr.ParsePostBody(r)
			if status != http.StatusOK {
				w.WriteHeader(status)
			}
			data, status := svr.AddPost(r.Context(), post)

			dataJson, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(status)
			w.Write(dataJson)

		case http.MethodPut:
			post, status := svr.ParsePostBodyUpdate(r)
			if status != http.StatusOK {
				w.WriteHeader(status)
			}
			status = svr.UpdatePost(r.Context(), post)
			w.WriteHeader(status)
		default:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "unsupported method")
		}
	})

	router.HandleFunc(
		fmt.Sprintf("/{%s:[0-9]+}", server.PostIDKey),
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodDelete:
				id, status := svr.ParsePathID(r)
				if status != http.StatusOK {
					w.WriteHeader(status)
				}
				status = svr.RemovePost(r.Context(), id)
				w.WriteHeader(status)
			default:
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println(w, "unsupported method")
			}
		},
	)
}

func initCommentRouter(router *mux.Router, svr *server.Server) {
	router.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				comment, status := svr.ParseCommentReq(r)
				if status != http.StatusOK {
					w.WriteHeader(status)
				}
				data, status := svr.AddComment(r.Context(), comment)

				dataJson, err := json.Marshal(data)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				w.WriteHeader(status)
				w.Write(dataJson)

			case http.MethodDelete:
				id, status := svr.ParseCommentID(r)
				if status != http.StatusOK {
					w.WriteHeader(status)
				}
				status = svr.RemoveComment(r.Context(), id)
				w.WriteHeader(status)
			default:
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "unsupported method")
			}
		},
	)
}
