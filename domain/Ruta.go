package domain

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ruta struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre         string             `bson:"nombre" json:"nombre"`
	Descripcion    string             `bson:"descripcion" json:"descripcion"`
	Origen         bson.M             `bson:"origen" json:"origen"` // o crea una struct si quieres validación
	Destino        bson.M             `bson:"destino" json:"destino"`
	ModoTransporte string             `bson:"modo_transporte" json:"modo_transporte"`
	Waypoints      []bson.M           `bson:"waypoints" json:"waypoints"` // o []map[string]interface{}
}
