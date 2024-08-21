package pokewrapr

import "github.com/Te4nick/pokewrapr/entity"

func LocationArea[T Integer | string](id T) (result entity.LocationArea, err error) {
	err = fetch("/location-area/"+convertIdOrName(id)+"/", &result)
	return
}
