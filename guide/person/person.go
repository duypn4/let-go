package person

type Person struct {
	name string
}

func (person *Person) SetName(name string) {
	person.name = name
}

func New(name string) Person {
	return Person{
		name: name,
	}
}
