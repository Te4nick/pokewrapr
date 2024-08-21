package pokewrapr

import "github.com/Te4nick/pokewrapr/entity"

func ResourceList[E entity.EndpointType | string, T Integer | string](endpoint E, offset, limit T) (out entity.NamedAPIResourceList, err error) {
	err = fetch("/"+string(endpoint)+"/?offset="+convertIdOrName(offset)+"&limit="+convertIdOrName(limit), &out)
	return
}
