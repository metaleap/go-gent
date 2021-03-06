package genttest

import (
	"sort"
	"time"
)

type simWorld struct {
	Cities [123]city
}

type town (**(**([12]**(**(sort.Float64Slice)))))

type city struct {
	Name      string
	ClosestTo *****city
	Companies []company
	Families  **[]family
	Schools   **[]**school
}

type company struct {
	Suppliers []*company
	Clients   []*company
	Staff     []*person
}

type school struct {
	Teachers []*person
	Pupils   []*person
}

type family struct {
	LastName string
	Pets     map[string]*petAnimal `ggd:" *petCat *petDog *petHamster *petPiranha  "`
}

type person struct {
	FirstName   string
	Family      *family
	DateOfBirth time.Time
	Parents     [2]*person
	FavPet      petAnimal `ggd:" *petPiranha *petHamster   *petDog *petCat  "`
	Top5Hobbies [5]hobby
}

type hobby struct {
	Name            string
	PopularityScore float64 // starting here, naive byte-count until Description string would be 32 but eg. for gc/amd64 is a size of 40 bytes incl padding
	AvgPerDay       struct {
		TimeNeededMinMax                        [2]time.Duration
		JustNeedToCheckPaddingAndAlignmentsHere bool
		CostInCentsMinMax                       [2]uint16
	}
	GroupSizeMinMax [2]uint
	Outdoorsy       bool
	Description     *string
}

type petAnimal interface {
	carnivore() bool
	mammal() bool
}

type pet struct {
	DailyFoodBill  float32
	AgeWhenAdopted time.Duration
	LastIllness    struct {
		Days          time.Duration
		Date          *time.Time
		NotSerialized sort.Interface   `ggd:"-"`
		Notes         sort.StringSlice `ggd:"[]string"`
	}
	OrigCostIfKnown *complex128
}

func (me *pet) carnivore() bool { return true }
func (me *pet) mammal() bool    { return true }

type petPiranha struct {
	pet
	Weird map[****[1234]byte][]fixedSize
}

func (me *petPiranha) mammal() bool { return false }

type petHamster struct {
	pet
}

func (me *petHamster) carnivore() bool { return false }

type petCat struct {
	pet
	RabbitsSlaynPerDayOnAvg *uint8
}

type petDog struct {
	pet
	WalkLog *map[*time.Time][7]time.Duration
}

type fixedSize struct {
	eight1 float64
	eight2 [1]uint64
	eight3 [2][3]int64
	eight4 [4][5]complex64
	four1  [6][7]float32
	four2  [8][9]int32
	four3  [8][7]uint32
	four4  [6][5]rune
	one1   [4][3]uint8
	one2   [2][1]int8
	one3   [2][3]byte
	sixt1  [4][5]complex128
	sixt2  [6][7]complex384
}

type Num int

const (
	zero Num = iota
	one
	two
	three
)

type complex384 [three]complex128
