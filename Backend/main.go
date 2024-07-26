package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

func main() {
	// Load environment variables from go.env file
	err := godotenv.Load("go.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	stripe.Key = os.Getenv("STRIPE_KEY")

	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config))

	router.POST("/create-payment-intent", handleCreatePaymentIntent)
	router.GET("/health", handleHealth)

	// Start the server
	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func handleCreatePaymentIntent(c *gin.Context) {
	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}

	log.Println("Incoming request:", c.Request.Method, c.Request.URL.Path)

	// Bind JSON request body to struct
	if err := c.BindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create PaymentIntent parameters
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrderAmount(req.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	// Create PaymentIntent
	paymentIntent, err := paymentintent.New(params)
	if err != nil {
		log.Printf("Error creating PaymentIntent: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("PaymentIntent created: %s", paymentIntent.ClientSecret)
	// Send response back to the frontend
	var response struct {
		ClientSecret string `json:"clientSecret"`
	}
	response.ClientSecret = paymentIntent.ClientSecret

	c.JSON(http.StatusOK, response)
}

func handleHealth(c *gin.Context) {
	c.String(http.StatusOK, "Server is up and running!")
}

// Calculate the order amount based on product ID
func calculateOrderAmount(productId string) int64 {
	switch productId {
	case "Forever Pants":
		return 26000
	case "Forever Shirt":
		return 15500
	case "Forever Shorts":
		return 30000
	default:
		return 0
	}
}
