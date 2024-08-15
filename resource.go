package pokewrapr

import "github.com/Te4nick/pokewrapr/entity"

func Resource(endpoint string) (out entity.NamedAPIResourceList, err error) {
	err = fetch(endpoint, &out)
	return
}
