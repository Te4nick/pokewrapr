package pokewrapr

import "github.com/Te4nick/pokewrapr/entity"

func LocationArea(id string) (result entity.LocationArea, err error) {
	err = fetch("/location-area/"+id+"/", &result)
	return
}
