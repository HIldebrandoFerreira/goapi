package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HIldebrandoFerreira/goapi/db"
	"github.com/HIldebrandoFerreira/goapi/entities"
	"github.com/dimfeld/httptreemux"
)

type MainHandler struct{}

func (main *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Página Principal")
}

type MateriaHandler struct {
	Repository *db.MateriaRepository
}

func (materia *MateriaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())

	article, err := materia.Repository.Get(params["id"])
	encoder := json.NewEncoder(w)
	if err == nil {
		erro := encoder.Encode(article)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
		}
	} else {
		log.Println("Failed to findById ", err)
		fmt.Println("not found")
	}
}

type MateriasHandler struct {
	Repository *db.MateriaRepository
}

func (materias *MateriasHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ch1 := make(chan []*entities.Materia)

	go materias.Repository.GetAll(ch1)

	select {
	case data := <-ch1:
		encode := json.NewEncoder(w)
		erro := encode.Encode(data)
		if erro != nil {
			http.Error(w, erro.Error(), http.StatusInternalServerError)
		}
	}
}

type PostMateriaHandler struct {
	Repository *db.MateriaRepository
}

func (materia *PostMateriaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//pegar no contexto
	article := &entities.Materia{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(article)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	error := materia.Repository.Insert(article)

	if error == db.ErrDuplicatedMateria {
		log.Printf("%s is already Created \n", article.Titulo)
	} else if err != nil {
		log.Printf("Fail to create team:", err)
	}
	fmt.Fprintln(w, "OK")
}
