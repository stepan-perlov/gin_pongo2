package gin_pongo2

import (
	"encoding/json"
	"github.com/flosch/pongo2"
	"path"
)

func init() {
	pongo2.RegisterFilter("abspath", filterAbspath)
	pongo2.RegisterFilter("json", filterJson)
}

func filterAbspath(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return pongo2.AsValue(path.Join(rootPath, in.String())), nil
}

func filterJson(in *pongo2.Value, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
	var value []byte

	mapValue, ok := in.Interface().(map[string]interface{})
	if ok {
		value, _ = json.Marshal(mapValue)
	} else {
		var arrMapValue []map[string]interface{}
		in.Iterate(func(idx int, count int, key *pongo2.Value, value *pongo2.Value) bool {
			mapItem, ok := key.Interface().(map[string]interface{})
			if ok {
				arrMapValue = append(arrMapValue, mapItem)
			} else {
				panic("[gin_pongo2.filterJson] Unknow format #1")
			}
			return true
		}, func() {})
		value, _ = json.Marshal(arrMapValue)
	}
	return pongo2.AsValue(string(value)), nil
}
