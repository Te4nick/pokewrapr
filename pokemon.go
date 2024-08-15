package pokewrapr

import (
	"github.com/Te4nick/pokewrapr/entity"
)

func Pokemon[T Integer | string](id T) (result entity.Pokemon, err error) {
	err = fetch("pokemon/"+convertIdOrName(id)+"/", &result)
	return
}
