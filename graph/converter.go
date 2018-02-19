package graph

func MakeInterfaceAndIntConverters(vertexes *[]interface{}) (*map[int]interface{}, *map[interface{}]int) {
	intToInterface := make(map[int]interface{}, len(*vertexes))
	interfaceToInt := make(map[interface{}]int, len(*vertexes))
	for k, v := range *vertexes {
		intToInterface[k] = v
		interfaceToInt[v] = k
	}
	return &intToInterface, &interfaceToInt
}

func GetCommonPath(paths *[][]int, intToInterface *map[int]interface{}) [][]interface{} {
	res := make([][]interface{}, len(*paths))
	for k, v := range *paths {
		for i := 0; i < len(v); i++ {
			res[k] = append(res[k], (*intToInterface)[v[i]])
		}
	}
	return res
}
