package provider

import (
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func convertSetToStringSlice(s *schema.Set) []string {
	strings := make([]string, 0, s.Len())

	for _, val := range s.List() {
		str, ok := val.(string)

		if ok && str != "" {
			strings = append(strings, str)
		}
	}
	return strings
}
func convertIdsToTypeIdRelationship(jsonType string, ids []string) []epcc.TypeIdRelationship {
	relationships := make([]epcc.TypeIdRelationship, len(ids))

	for idx, fileId := range ids {
		relationships[idx] = epcc.TypeIdRelationship{
			Id:   fileId,
			Type: jsonType,
		}
	}
	return relationships
}

func convertJsonTypesToIds(productFiles *[]epcc.TypeIdRelationship) []string {
	fileIds := make([]string, 0, len(*productFiles))

	for _, file := range *productFiles {
		fileIds = append(fileIds, file.Id)
	}
	return fileIds
}
