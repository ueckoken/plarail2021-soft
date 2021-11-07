package stationNameId

import (
	"fmt"
	"ueckoken/plarail2021-soft-external/spec"
)

func Id2Name(id int32) (string, error) {
	name, ok := spec.Stations_StationId_name[id]
	if !ok {
		return "", fmt.Errorf("station ID `%d` not found\n", id)
	}
	return name, nil
}
func Name2Id(name string) (int32, error) {
	id, ok := spec.Stations_StationId_value[name]
	if !ok {
		return 0, fmt.Errorf("station `%s` not found", name)
	}
	return id, nil
}
