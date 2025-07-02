package repository

import (
	"database/sql"
	"fmt"
	"log"
	"product_stock/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		" VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	log.Println("Entrou no repository ")
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}
	defer rows.Close()

	var productList []model.Product

	for rows.Next() {
		var p model.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		productList = append(productList, p)
	}

	return productList, nil
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var produto model.Product
	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()
	return &produto, nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) (*model.Product, error) {
	log.Println("Entrou no repository update")

	query := `UPDATE product
	SET product_name = $1, price = $2
	WHERE id = $3
	RETURNING id, product_name, price
	`

	row := pr.connection.QueryRow(query, product.Name, product.Price, product.ID)

	var UpdateProduct model.Product

	err := row.Scan(&UpdateProduct.ID, &UpdateProduct.Name, &UpdateProduct.Price)
	if err != nil {
		log.Printf("Erro ao fazer update: %v", err)
		return nil, err
	}
	log.Printf("Update concluido com sucesso")
	return &UpdateProduct, nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	query := `DELETE FROM product WHERE id = $1`

	res, err := pr.connection.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao deletar produto ID %d: %v", id, err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Printf("Erro ao verificar linhas afetadas: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("produto com ID %d não encontrado", id)
	}

	log.Printf("Produto ID %d deletado com sucesso", id)
	return nil

}
