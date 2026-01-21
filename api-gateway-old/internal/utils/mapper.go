package utils

// Mapper provides utilities for mapping data transformations
type Mapper struct{}

// NewMapper creates a new mapper utility
func NewMapper() *Mapper {
	return &Mapper{}
}

// MapInterfaceToOrder converts an interface{} to an Order struct
func (m *Mapper) MapInterfaceToOrder(data interface{}) (map[string]interface{}, bool) {
	if orderMap, ok := data.(map[string]interface{}); ok {
		return orderMap, true
	}
	return nil, false
}

// MapInterfaceToUser converts an interface{} to a User struct
func (m *Mapper) MapInterfaceToUser(data interface{}) (map[string]interface{}, bool) {
	if userMap, ok := data.(map[string]interface{}); ok {
		return userMap, true
	}
	return nil, false
}

// MapInterfaceToProduct converts an interface{} to a Product struct
func (m *Mapper) MapInterfaceToProduct(data interface{}) (map[string]interface{}, bool) {
	if productMap, ok := data.(map[string]interface{}); ok {
		return productMap, true
	}
	return nil, false
}

// MapInterfaceToCartItem converts an interface{} to a CartItem struct
func (m *Mapper) MapInterfaceToCartItem(data interface{}) (map[string]interface{}, bool) {
	if itemMap, ok := data.(map[string]interface{}); ok {
		return itemMap, true
	}
	return nil, false
}

// MapStringSliceToString concatenates a slice of strings with a separator
func (m *Mapper) MapStringSliceToString(slice []string, separator string) string {
	if len(slice) == 0 {
		return ""
	}
	
	result := slice[0]
	for i := 1; i < len(slice); i++ {
		result += separator + slice[i]
	}
	return result
}

// MapInterfaceSliceToStringSlice converts a slice of interface{} to a slice of strings
func (m *Mapper) MapInterfaceSliceToStringSlice(slice []interface{}) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		if str, ok := v.(string); ok {
			result[i] = str
		}
	}
	return result
}

// MapFilterMapByKey filters a map of maps by a specific key value
func (m *Mapper) MapFilterMapByKey(data []interface{}, key string, value interface{}) []interface{} {
	var result []interface{}
	for _, item := range data {
		if itemMap, ok := item.(map[string]interface{}); ok {
			if itemValue, exists := itemMap[key]; exists {
				if itemValue == value {
					result = append(result, item)
				}
			}
		}
	}
	return result
}

// MapTransform applies a transformation function to each element in a slice
func (m *Mapper) MapTransform(data []interface{}, transformFunc func(interface{}) interface{}) []interface{} {
	result := make([]interface{}, len(data))
	for i, item := range data {
		result[i] = transformFunc(item)
	}
	return result
}

// MapExtractField extracts a specific field from each element in a slice
func (m *Mapper) MapExtractField(data []interface{}, field string) []interface{} {
	result := make([]interface{}, len(data))
	for i, item := range data {
		if itemMap, ok := item.(map[string]interface{}); ok {
			if value, exists := itemMap[field]; exists {
				result[i] = value
			}
		}
	}
	return result
}