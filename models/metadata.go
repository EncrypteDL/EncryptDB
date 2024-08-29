package models

import "time"

//Metadata is an interface tahta defines the minimum requirements for [Encrypt.Keeper] metadata implementation
type Metadata interface{
	Type() string
	//Timestamp should return the last time metadata's parent was opened 
	Timestamp() time.Time
}