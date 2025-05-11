package delivery

import (
	"fmt"
	"net/http"

	"UbicaBus/UbicaBusBackend/application"

	"github.com/gin-gonic/gin"
)

// UserHandler maneja las peticiones relacionadas con usuarios
type UserHandler struct {
	UserService *application.UserService
}

// RutaHandler maneja las peticiones relacionadas con rutas
type RutaHandler struct {
	RutaService *application.RutaService
}

// NewUserHandler crea un nuevo manejador de usuarios
func NewUserHandler(userService *application.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// NewRutaHandler crea un nuevo manejador de rutas
func NewRutaHandler(rutaService *application.RutaService) *RutaHandler {
	return &RutaHandler{RutaService: rutaService}
}

// RegisterUserHandler maneja el registro de un usuario
func (h *UserHandler) RegisterUserHandler(c *gin.Context) {
	var request struct {
		Nombre     string `json:"nombre"`
		Password   string `json:"password"`
		RolID      string `json:"rol_id"`
		CompaniaID string `json:"compania_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de entrada inválidos"})
		return
	}

	userID, err := h.UserService.RegisterUser(request.Nombre, request.Password, request.RolID, request.CompaniaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar usuario: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado correctamente",
		"user_id": userID.Hex(),
	})
}

// GetAllRutasHandler maneja la petición GET para obtener todas las rutas
func (h *RutaHandler) GetAllRutasHandler(c *gin.Context) {
	rutas, err := h.RutaService.GetAllRutas(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las rutas"})
		return
	}
	c.JSON(http.StatusOK, rutas)
}

func StartServer(userService *application.UserService, rutaService *application.RutaService) {
	r := gin.Default()

	userHandler := NewUserHandler(userService)
	rutaHandler := NewRutaHandler(rutaService)

	r.POST("/register", userHandler.RegisterUserHandler)
	r.GET("/rutas", rutaHandler.GetAllRutasHandler)

	fmt.Println("Iniciando servidor en el puerto 8080...")
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
