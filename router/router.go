package router

import "github.com/gin-gonic/gin"

/* Essa função servirá para inicializar o servidor. */

/* Qualquer função que comece com letra maiúscula está sendo exportada. */
func Initialize() {

	/* Inicializa o router com as configurações default do Gin. */
	router := gin.Default()

	/* Define uma rota. O primeiro parâmetro é o nome da rota, e o segundo parâmetro é um handler. Ele
	serve para serializarmos o código como JSON e devolvendo uma resposta. */

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/* Por padrão, estamos rodando na porta 8080. */
	router.Run(":8080")
	print("Estou testando!\n")
}
