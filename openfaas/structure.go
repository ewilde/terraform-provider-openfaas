package openfaas

func expandStringList(configured []interface{}) []string {
	list := make([]string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			list = append(list, v.(string))
		}
	}
	return list
}

func expandStringMap(m map[string]interface{}) map[string]string {
	list := make(map[string]string, len(m))
	for i, v := range m {
		list[i] = v.(string)
	}
	return list
}

func pointersMapToStringList(pointers *map[string]string) map[string]string {
	if pointers != nil{
		return *pointers
	}

	return nil
}