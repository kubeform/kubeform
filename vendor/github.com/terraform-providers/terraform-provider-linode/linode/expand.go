package linode

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func expandStringList(list []interface{}) []string {
	slice := make([]string, 0, len(list))
	for _, s := range list {
		if val, ok := s.(string); ok && val != "" {
			slice = append(slice, val)
		}
	}
	return slice
}

func expandStringSet(set *schema.Set) []string {
	return expandStringList(set.List())
}

func expandIntList(list []interface{}) []int {
	slice := make([]int, 0, len(list))
	for _, n := range list {
		if val, ok := n.(int); ok {
			slice = append(slice, val)
		}
	}
	return slice
}

func expandIntSet(set *schema.Set) []int {
	return expandIntList(set.List())
}
