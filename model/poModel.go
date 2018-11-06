package model

//Po is a model for dao
type Po struct {
	PoID    string `bson:"poId"`
	Name    string `bson:"name"`
	Address string `bson:"alamat"`
}
