package graph

type cEdge [2]interface{}

type CEdges []cEdge

func (es *CEdges) GetUniqueVertexes() []interface{} {
	var uniqueVertexes = make(map[interface{}]bool)
	for _, e := range *es {
		for _, v := range e {
			if !uniqueVertexes[v] {
				uniqueVertexes[v] = true
			}
		}
	}
	var vertexes = make([]interface{}, 0, len(uniqueVertexes))
	for uv := range uniqueVertexes {
		vertexes = append(vertexes, uv)
	}
	return vertexes
}

func MakeIntM(vertexes *[]interface{}) (map[int]interface{}, map[interface{}]int) {
	res1 := make(map[int]interface{}, len(*vertexes))
	res2 := make(map[interface{}]int, len(*vertexes))
	for k, v := range *vertexes {
		res1[k] = v
		res2[v] = k
	}
	return res1, res2
}

func (es *CEdges) FromCEtoE(fromInter map[int]interface{}, toInt map[interface{}]int) Edges {
	e := make(Edges, 0, len(*es))
	for _, v := range *es {
		e = append(e, edge{toInt[v[0]], toInt[v[1]]})
	}
	return e
}

func FromIntPaths(paths [][]int, toInter map[int]interface{}) [][]interface{} {
	res := make([][]interface{}, len(paths))
	for k, v := range paths {
		for i := 0; i < len(v); i++ {
			res[k] = append(res[k], toInter[v[i]])
		}
	}
	return res
}

func Dfs(ed *Edges, from interface{}, to interface{}, toInt map[interface{}]int, toInter map[int]interface{}) [][]interface{} {
	ff := toInt[from]
	tt := toInt[to]
	res := ed.Dfs(ff, tt)
	return FromIntPaths(res, toInter)
}