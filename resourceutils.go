package main

func Load(path string) string {
	resource, err := Asset(path)
	if err != nil {
		panic(err)
	}

	return str(resource)
}
