/*
Dependency inversion principle
HLM should not depend LLM
Both should depend on abstraction
*/
package main

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Silbling
)

type Person struct {
	name string
	// something else
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// low-level model

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, rel := range r.relations {
		if rel.from.name == name && rel.relationship == Parent {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{
		from:         parent,
		relationship: Parent,
		to:           child,
	})
	r.relations = append(r.relations, Info{
		from:         child,
		relationship: Child,
		to:           parent,
	})
}

// high-level model

type Research struct {
	// break DIP
	//relationships Relationships
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println(p.name)
	}
}

func main() {
	rs := Relationships{}
	rs.AddParentAndChild(&Person{"John"}, &Person{"Piter"})

	r := Research{&rs}
	r.Investigate()
}
