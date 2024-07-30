package fazz

import "fmt"

type Bio struct {
	name         string
	age          int
	gender       string
	isSmoker     bool
	cigarVariant []string
	Skillset []Skillset
}
type Skillset struct {
name string
	
}

func Survey() {
	survey := []Bio{
		Bio{

			name:     "ilyas",
			age:      18,
			gender:   "laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Esse",
				"Marlboro",
				"Surya",
			},
				Skillset: []Skillset{
				{
					name: "javaScript",
				},
				{
					name: "Java",
				},
				{
					name: "PHP",
				},
				{
					name: "HTML",
				},
				{
					name: "Python",
				},
			},
		
		},
		Bio{
			name:     "ilyas",
			age:      18,
			gender:   "laki-laki",
			isSmoker: false,
			cigarVariant: []string{
				"Esse",
				"Marlboro",
				"Surya",
			},
		},
		Bio{

			name:     "ilyas",
			age:      18,
			gender:   "laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Esse",
				"Marlboro",
				"Surya",
			},
		},
		Bio{

			name:     "ilyas",
			age:      18,
			gender:   "laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Esse",
				"Marlboro",
				"Surya",
			},
		},
		Bio{

			name:     "ilyas",
			age:      18,
			gender:   "laki-laki",
			isSmoker: false,
			cigarVariant: []string{
				"Esse",
				"Marlboro",
				"Surya",
				
			},
				Skillset: []Skillset{
				{
					name: "javaScript",
				},
				{
					name: "Java",
				},
				{
					name: "PHP",
				},
				{
					name: "HTML",
				},
				{
					name: "Python",
				},
			},
		},
	}
	for _, j := range survey{
		fmt.Printf("%s,%d,%s,%t ", j.name,j.age,j.gender,j.isSmoker)
		for _, y := range j.cigarVariant {
			fmt.Printf(" %s ", y)
		}
		fmt.Println("")
	}
	fmt.Println(survey[4].Skillset[2].name)
	fmt.Println(survey[0].Skillset[0].name)
}