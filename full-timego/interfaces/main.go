package main

import "fmt"

type NumberStorer interface {
	GetAll() ([]int, error)
	Put() error
}

type MongoDBNumberStore struct {
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4}, nil
}
func (m MongoDBNumberStore) Put() error {
	fmt.Println("Printling ")
	return nil
}

type ApiServer struct {
	numberStore NumberStorer
}

func main() {

	apiServer := ApiServer{
		numberStore: MongoDBNumberStore{},
	}

	number, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(number)

}
