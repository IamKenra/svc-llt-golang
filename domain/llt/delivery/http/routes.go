package http

import (
	"log"
	"svc-llt-golang/domain/llt/repository"
	"svc-llt-golang/domain/llt/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterLltRoutes(api fiber.Router, db *gorm.DB) {
	// Initialize repository and usecase
	lltRepo := repository.NewMysqlLltRepository(db)
	lltUC := usecase.NewLltUsecase(lltRepo)
	
	// Initialize handlers
	lansiaHandler := NewLansiaHandler(lltUC)
	identitasHandler := NewIdentitasHandler(lltUC)
	alamatHandler := NewAlamatHandler(lltUC)

	// Lansia routes under /lansia
	lansia := api.Group("/lansia")
	{
		lansia.Get("/", lansiaHandler.GetAllLansia)
		lansia.Get("/detail", lansiaHandler.GetOneLansia)
		lansia.Post("/", lansiaHandler.StoreLansia)
		lansia.Post("/complete", lansiaHandler.StoreLansiaComplete) // Orchestrated creation
		lansia.Put("/", lansiaHandler.UpdateLansia)
		lansia.Delete("/", lansiaHandler.DeleteLansia)
	}

	// Identitas routes under /identitas
	identitas := api.Group("/identitas")
	{
		identitas.Get("/", identitasHandler.GetAllIdentitas)
		identitas.Get("/detail", identitasHandler.GetOneIdentitas)
		identitas.Post("/", identitasHandler.StoreIdentitas)
		identitas.Put("/", identitasHandler.UpdateIdentitas)
		identitas.Delete("/", identitasHandler.DeleteIdentitas)
	}

	// Alamat routes under /alamat
	alamat := api.Group("/alamat")
	{
		alamat.Get("/", alamatHandler.GetAllAlamat)
		alamat.Get("/detail", alamatHandler.GetOneAlamat)
		alamat.Post("/", alamatHandler.StoreAlamat)
		alamat.Put("/", alamatHandler.UpdateAlamat)
		alamat.Delete("/", alamatHandler.DeleteAlamat)
	}
	
	// Debug: Log that routes are registered
	log.Println("LLT routes registered successfully (Lansia, Identitas, Alamat)")
}
