package util

import ()

func Contains(strs []string, target string) bool {
	for _, str := range strs {
		if str == target {
			return true
		}
	}
	return false
}

func Remove(strs []string, target string) []string {
	res := []string{}
	for _, str := range strs {
		if str != target {
			res = append(res, str)
		}
	}
	return res
}
