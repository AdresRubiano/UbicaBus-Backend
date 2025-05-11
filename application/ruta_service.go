package application

import (
	"context"
	"log"

	"UbicaBus/UbicaBusBackend/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RutaService struct {
	DB *mongo.Database
}

func NewRutaService(db *mongo.Database) *RutaService {
	return &RutaService{DB: db}
}

func (s *RutaService) GetAllRutas(ctx context.Context) ([]domain.Ruta, error) {
	collection := s.DB.Collection("rutas")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("error al obtener rutas:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var rutas []domain.Ruta
	if err := cursor.All(ctx, &rutas); err != nil {
		log.Println("error al decodificar rutas:", err)
		return nil, err
	}

	return rutas, nil
}
