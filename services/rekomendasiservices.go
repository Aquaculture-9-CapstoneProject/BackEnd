// services/product_service.go
package services

import (
	"fmt"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
	"gorgonia.org/gorgonia"
)

// ProductService interface
type ProductService interface {
	RecommendProducts(userID int) ([]entities.Product, error)
}

// ProductServiceImpl implementation
type ProductServiceImpl struct {
	productRepo repositories.ProductRepository
}

// NewProductService creates a new ProductService
func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepo: productRepo}
}

// RecommendProducts generates product recommendations based on collaborative filtering
func (s *ProductServiceImpl) RecommendProducts(userID int) ([]entities.Product, error) {
	// Step 1: Get all products from the database
	products, err := s.productRepo.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %v", err)
	}

	// Step 2: Prepare the rating matrix based on product ratings (you may adjust this depending on your rating data)
	// Assuming the rating matrix is created based on user-product relationships
	// The matrix is a 2D array where rows are users and columns are products
	ratingMatrix := [][]float32{
		{5, 3, 0, 1}, // User 1 ratings for products
		{4, 0, 4, 2}, // User 2 ratings for products
		{0, 5, 4, 3}, // User 3 ratings for products
		{2, 2, 3, 0}, // User 4 ratings for products
	}

	// Step 3: Create a tensor for the rating matrix using Gorgonia
	ratingTensor := gorgonia.NewTensor(gorgonia.WithShape(len(ratingMatrix), len(ratingMatrix[0])), gorgonia.WithName("rating"), gorgonia.WithBacking(ratingMatrix))

	// Step 4: Initialize user and item latent factor matrices
	numUsers := len(ratingMatrix)
	numItems := len(ratingMatrix[0])
	numLatentFactors := 2

	// Latent factors for users and items (randomly initialized)
	userMatrix := gorgonia.NewTensor(gorgonia.WithShape(numUsers, numLatentFactors), gorgonia.WithName("userMatrix"), gorgonia.WithShape(numUsers, numLatentFactors), gorgonia.WithBacking(make([]float32, numUsers*numLatentFactors)))
	itemMatrix := gorgonia.NewTensor(gorgonia.WithShape(numItems, numLatentFactors), gorgonia.WithName("itemMatrix"), gorgonia.WithShape(numItems, numLatentFactors), gorgonia.WithBacking(make([]float32, numItems*numLatentFactors)))

	// Step 5: Define the learning rate and number of iterations
	learningRate := 0.01
	numIterations := 1000

	// Step 6: Define the loss function (mean squared error)
	loss := gorgonia.Must(gorgonia.NewScalar(gorgonia.WithShape(1), gorgonia.WithName("loss"), gorgonia.WithBacking(float32(0))))

	// Step 7: Train the model (update the user and item matrices)
	for epoch := 0; epoch < numIterations; epoch++ {
		for i := 0; i < numUsers; i++ {
			for j := 0; j < numItems; j++ {
				if ratingMatrix[i][j] > 0 { // Only consider non-zero ratings
					// Predict the rating by multiplying the user and item matrices
					prediction := gorgonia.Must(gorgonia.Dot(userMatrix, itemMatrix.T()))
					predictedRating := prediction.Get(i, j)

					// Calculate the error
					error := ratingMatrix[i][j] - predictedRating

					// Compute the gradient and update the matrices
					gorgonia.Apply(gorgonia.Add(gorgonia.Mul(userMatrix, itemMatrix), gorgonia.Mul(error, itemMatrix)), userMatrix)
					gorgonia.Apply(gorgonia.Add(gorgonia.Mul(itemMatrix, userMatrix), gorgonia.Mul(error, userMatrix)), itemMatrix)

					// Update loss
					loss.Add(error * error)
				}
			}
		}

		// Step 8: Output the loss at each epoch (for debugging)
		if epoch%100 == 0 {
			fmt.Printf("Epoch %d, Loss: %v\n", epoch, loss.Value())
		}
	}

	// Step 9: Generate product recommendations (return top N products for user)
	// After training, use the userMatrix and itemMatrix to make predictions for the given user
	// For this example, we return dummy products
	recommendedProducts := []entities.Product{
		{ID: 1, Nama: "Product 1"},
		{ID: 2, Nama: "Product 2"},
		{ID: 3, Nama: "Product 3"},
	}

	return recommendedProducts, nil
}
