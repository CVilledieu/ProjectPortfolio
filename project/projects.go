package projects

type Project struct {
	Title       string
	Description string
}

func createProject(name, description string) *Project {
	return &Project{Title: name,
		Description: description,
	}
}

func createListOfProjects(n, d []string) []*Project {
	list := []*Project{}
	for i, _ := range n {
		list = append(list, createProject(n[i], d[i]))
	}
	return list
}

func GetListOfProjects() []*Project {
	names := []string{"Password Generator"}
	descriptions := []string{"A CLI tool to create passwords. Takes in a length between 8-12 and outputs a random series of Uppercase, lowercase, numbers, and symbols"}
	return createListOfProjects(names, descriptions)
}
