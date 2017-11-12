package mapper

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
)

type InitContainerMapper struct {
}

func (e InitContainerMapper) FromInternal(data map[string]interface{}) {
	containers, _ := data["containers"].([]interface{})

	for _, initContainer := range convert.ToMapSlice(data["initContainers"]) {
		if initContainer == nil {
			continue
		}
		initContainer["initContainer"] = true
		containers = append(containers, initContainer)
	}

	data["containers"] = containers
}

func (e InitContainerMapper) ToInternal(data map[string]interface{}) {
	newContainers := []interface{}{}
	newInitContainers := []interface{}{}

	for _, container := range convert.ToMapSlice(data["container"]) {
		if convert.ToBool(container["initContainer"]) {
			newInitContainers = append(newInitContainers, container)
		} else {
			newContainers = append(newContainers, container)
		}
		delete(container, "initContainer")
	}

	data["containers"] = newContainers
	data["initContainers"] = newInitContainers
}

func (e InitContainerMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	delete(schema.ResourceFields, "initContainers")
	return nil
}
