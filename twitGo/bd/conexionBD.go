package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Todo lo que se exporte debe tener un comentario
// MongoCN, es el objeto de conexión a la DB
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://wjbr13:0123456789@cluster0.tv7jd.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConectarBD, es la func que me permite conectarme a la DB (este comentario es obligatorio y debe comenzar
  con el nombre de la función, lo exige GO)*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("conexión exitosa con la BD")
	return client
}

/*ChequeoConection, es el ping a la DB*/
func ChequeoConection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
