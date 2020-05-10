package db

import (
	"errors"

	"github.com/HIldebrandoFerreira/goapi/entities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const MateriaCollection = "materia"

var ErrDuplicatedMateria = errors.New("Duplicated materia")

type MateriaRepository struct {
	session *mgo.Session
}

// enjoy Repository instance to make easy to interact and test
func NewMateriaRepository(session *mgo.Session) *MateriaRepository {
	return &MateriaRepository{session}
}

/*
	OPERATIONS CREATE UPDATE REMOVE FINDALL FINDBYID
*/

func (r *MateriaRepository) Insert(materia *entities.Materia) error {
	session := r.session.Clone()
	//wait til functions return
	defer session.Close()

	mongoCollection := session.DB("").C(MateriaCollection)

	err := mongoCollection.Insert(materia)
	if mongoErr, ok := err.(*mgo.LastError); ok {
		if mongoErr.Code == 11000 {
			return ErrDuplicatedMateria
		}
	}
	return err
}

func (r *MateriaRepository) Update(materia *entities.Materia) error {
	session := r.session.Clone()
	defer session.Close()

	mongoCollection := session.DB("").C(MateriaCollection)
	return mongoCollection.Update(bson.M{"_id": materia.Id}, materia)
}

func (r *MateriaRepository) Delete(id string) error {
	session := r.session.Clone()
	defer session.Close()

	mongoCollection := session.DB("").C(MateriaCollection)
	return mongoCollection.Remove(bson.M{"_id": id})
}

func (r *MateriaRepository) GetAll(c1 chan []*entities.Materia) ([]*entities.Materia, error) {
	session := r.session.Clone()
	defer session.Close()

	mongoCollection := session.DB("").C(MateriaCollection)

	query := bson.M{}

	docs := make([]*entities.Materia, 0)

	err := mongoCollection.Find(query).All(&docs)

	c1 <- docs
	return docs, err
}

func (r *MateriaRepository) Get(id string) (*entities.Materia, error) {
	session := r.session.Clone()
	defer session.Close()

	mongoCollection := session.DB("").C(MateriaCollection)

	query := bson.M{"_id": id}

	materia := &entities.Materia{}

	err := mongoCollection.Find(query).One(materia)

	return materia, err
}
