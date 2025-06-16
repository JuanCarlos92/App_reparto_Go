package rutas

import (
	"jornada-backend/handlers"
	"net/http"
)

// ConfigurarRutas define todas las rutas de la aplicación
func ConfigurarRutas() {
	// Ruta para login
	http.HandleFunc("/login", handlers.LoginHandler) // Ruta para hacer login

	// Ruta para iniciar sesión de trabajo
	http.HandleFunc("/work-sessions/start", handlers.StartWorkSessionHandler)
	http.HandleFunc("/work-sessions/end", handlers.EndWorkSessionHandler)
	http.HandleFunc("/work-sessions/update", handlers.UpdateWorkSessionHandler)

	// Ruta para obtener eventos de agenda
	http.HandleFunc("/event", handlers.GetEventsHandler()) // Ruta para obtener eventos de agenda

	// Ruta para obtener información básica del cliente
	http.HandleFunc("/client/basic-info", handlers.GetBasicClientInfoHandler()) // Ruta para obtener información básica del cliente

	//Ruta para obtener tickets
	http.HandleFunc("/tickets", handlers.GetBasicTicketInfoHandler()) // Ruta para obtener tickets

	// Ruta para obtener información básica de productos
	http.HandleFunc("/products", handlers.GetBasicProductInfoHandler()) // Ruta para obtener información básica de productos

}
