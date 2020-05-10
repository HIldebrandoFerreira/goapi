package entities

// MAteria traz os dados que serão trafedados de uma matéria
type Materia struct {
	Titulo          string `json bson:"titulo"`
	Subtitulo       string `json bson: "subtitulo"`
	Cover           string `json bson:"cover"`
	DataCriacao     string `json bson:"datacriacao"`
	DataModificacao string `json bson:"datamodificacao"`
	DataPublicacao  string `json bson:"datapublicacao"`
	Slug            string `json bson:"slug"`
	Id              string `json bson:"id"`
	Corpo           string `json bson:"corpo"`
}
