package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestAllRightMeow(t *testing.T) {
	g := NewGomegaWithT(t)

	meow := Meow{
		Test_Matches: "test",       //right
		Phone:        "0885870149", //right
		Cost:         10,           //right
	}

	ok, err := govalidator.ValidateStruct(meow)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}

func TestMatchesNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	meow := Meow{
		Test_Matches: "",           //wrong-->>pls check
		Phone:        "0885870149", //right
		Cost:         10,           //right
	}

	ok, err := govalidator.ValidateStruct(meow)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Test cannot be blank"))
}
func TestCostLessThenZero(t *testing.T) {
	g := NewGomegaWithT(t)

	meow := Meow{
		Test_Matches: "test",           //wrong-->>pls check
		Phone:        "0885870149", //right
		Cost:         -9,           //right
	}

	ok, err := govalidator.ValidateStruct(meow)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Cost: -9 does not validate as range(0|10)"))
}

func TestLengthPhone(t *testing.T) {
	g := NewGomegaWithT(t)

	meow := Meow{
		Test_Matches: "test",           //wrong-->>pls check
		Phone:        "08858701491", //right
		Cost:         9,           //right
	}

	ok, err := govalidator.ValidateStruct(meow)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Phone: 08858701491 does not validate as length(0|10)"))
}
