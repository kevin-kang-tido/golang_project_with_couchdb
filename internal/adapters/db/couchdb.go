package db

import (
    "encoding/json"
    "errors"
    "fmt"
    "golang_with_couchdb2/internal/domain/entities"
    "golang_with_couchdb2/internal/usecases/ports"
    "net/http"
    "bytes"
    // "io/ioutil"
)

// add new 

type CouchDBProductRepository struct {
    CouchDBURL string
}

func NewCouchDBProductRepository(url string) ports.ProductRepository {
    return &CouchDBProductRepository{
        CouchDBURL: url,
    }
}

// Function to get all products
func (repo *CouchDBProductRepository) GetAllProducts() ([]entities.Product, error) {
    // Construct CouchDB _all_docs URL
    url := fmt.Sprintf("%s/products/_all_docs?include_docs=true", repo.CouchDBURL)

    // Make the HTTP request
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Define a structure to capture CouchDB response
    type Row struct {
        Doc entities.Product `json:"doc"`
    }
    var result struct {
        Rows []Row `json:"rows"`
    }

    // Parse the response
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    // Extract products from the rows
    var products []entities.Product
    for _, row := range result.Rows {
        products = append(products, row.Doc)
    }

    return products, nil
}

func (r *CouchDBProductRepository) Create(product *entities.Product) error {
    productJSON, err := json.Marshal(product)
    if err != nil {
        return err
    }

    resp, err := http.Post(fmt.Sprintf("%s/products", r.CouchDBURL), "application/json", bytes.NewBuffer(productJSON))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        return errors.New("failed to create product")
    }

    return nil
}

func (r *CouchDBProductRepository) GetByID(id string) (*entities.Product, error) {
    resp, err := http.Get(fmt.Sprintf("%s/products/%s", r.CouchDBURL, id))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusNotFound {
        return nil, errors.New("product not found")
    }

    var product entities.Product
    err = json.NewDecoder(resp.Body).Decode(&product)
    if err != nil {
        return nil, err
    }

    return &product, nil
}

// Update method implementation
func (repo *CouchDBProductRepository) Update(product *entities.Product) error {
    // Ensure the product has an ID and Rev
    if product.ID == "" || product.Rev == "" {
        return fmt.Errorf("product ID and Rev must not be empty")
    }

    // Prepare the URL for the update
    url := fmt.Sprintf("%s/products/%s", repo.CouchDBURL, product.ID)

    // Create the JSON payload for the update
    jsonProduct, err := json.Marshal(product)
    if err != nil {
        return err
    }

    // Make the PUT request to update the product in CouchDB
    req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonProduct))
    if err != nil {
        return err
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Check the response status
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("failed to update product, status code: %d", resp.StatusCode)
    }

    return nil
}

func (r *CouchDBProductRepository) Delete(id string) error {
    req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/products/%s", r.CouchDBURL, id), nil)
    if err != nil {
        return err
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return errors.New("failed to delete product")
    }

    return nil
}
