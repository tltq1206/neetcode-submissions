type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		 return ""
	}
	sizes := []string{}
	for _, str := range strs {
		sizes = append(sizes, strconv.Itoa(len(str)))
	}
	return strings.Join(sizes,",") + "#" + strings.Join(strs, "")

}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}

	parts := strings.SplitN(encoded, "#", 2)
	sizes := strings.Split(parts[0], ",")

	res := []string{}
	i := 0
	for _, sizeStr := range sizes {
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			break
		}

      res = append(res, parts[1][i : i + size])
	  i += size
	}

	return res
}
