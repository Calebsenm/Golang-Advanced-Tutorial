package main

type college struct {
	departments []department
}

func (c *college) addDepartament(departmentName string, numOfProfessors int) {
	if departmentName == "computerscience" {
		computerScienceDepartment := &computerscience{numberOfProfessors: numOfProfessors}
		c.departments = append(c.departments, computerScienceDepartment)
	}

	if departmentName == "mechanical" {
		mechanicalDepartment := &mechanical{numberOfProfessors: numOfProfessors}
		c.departments = append(c.departments, mechanicalDepartment)
	}
	return
}

func (c *college) getDepartment(departmentName string) department {

	for _, department := range c.departments {
		if department.getName() == departmentName {
			return department
		}
	}

	// return a nill department if the deparment doesn´t exists
	return &nullDepartament{}
}
