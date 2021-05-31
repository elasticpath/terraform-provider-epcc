package provider

import (
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func convertArrayToStringSlice(arr []interface{}) []string {
	var result []string
	for _, param := range arr {
		result = append(result, param.(string))
	}
	return result
}

func convertArrayToIntSlice(arr []interface{}) []int {
	var result []int
	for _, param := range arr {
		result = append(result, param.(int))
	}
	return result
}

func convertMapToIntMap(m map[string]interface{}) map[string]int {
	result := make(map[string]int, len(m))
	for k, v := range m {
		result[k] = v.(int)
	}
	return result
}

func convertArrayToIntMaps(arr []interface{}) []map[string]int {
	var result []map[string]int
	for _, param := range arr {
		intMap := convertMapToIntMap(param.(map[string]interface{}))
		result = append(result, intMap)
	}
	return result
}

func convertArrayToFloatSlice(arr []interface{}) []float64 {
	var result []float64
	for _, param := range arr {
		result = append(result, param.(float64))
	}
	return result
}

func convertMapToFloatMap(m map[string]interface{}) map[string]float64 {
	result := make(map[string]float64, len(m))
	for k, v := range m {
		result[k] = v.(float64)
	}
	return result
}

func convertArrayToFloatMaps(arr []interface{}) []map[string]float64 {
	var result []map[string]float64
	for _, param := range arr {
		floatMap := convertMapToFloatMap(param.(map[string]interface{}))
		result = append(result, floatMap)
	}
	return result
}

func convertMapToStringMap(m map[string]interface{}) map[string]string {
	result := make(map[string]string, len(m))
	for k, v := range m {
		result[k] = v.(string)
	}
	return result
}

func convertMapToBooleanMap(m map[string]interface{}) map[string]bool {
	result := make(map[string]bool, len(m))
	for k, v := range m {
		result[k] = v.(bool)
	}
	return result
}

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
