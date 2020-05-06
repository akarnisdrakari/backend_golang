package db

import (
	"context"
	"time"

	"github.com/akarnisdrakari/backend_golang/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertoRegistro se usa para insertar datos de la tabla Usuario*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("gobd")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.string(), true, nil
}
