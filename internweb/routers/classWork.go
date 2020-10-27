
package routers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Data struct {
	ID primitive.ObjectID `json:"_id,omitempty bson:"_id,omitempty"`
	Details *user `json:"details" bson:"details"`
}


