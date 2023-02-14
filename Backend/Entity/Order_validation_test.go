package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestAllRight(t *testing.T) {
	g := NewGomegaWithT(t)

	video := Video{
		Name: "meow",            //ถูก
		Url:  "www.youtube.com", //ถูก
	}

	ok, err := govalidator.ValidateStruct(video)

	//ขึ้น ok คือ รูปแบบที่ผิดแล้วให้ขึ้น ok
	g.Expect(ok).To(BeTrue())

	//err เป็นค่า nil ไม่ได้ ต้องแสดงออกมา
	g.Expect(err).To(BeNil())
}

func TestNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	video := Video{
		Name: "",            //ผิด -->เช็คตรงนี้
		Url:  "www.youtube.com", //ถูก
	}

	ok, err := govalidator.ValidateStruct(video)

	//ขึ้น ok คือ รูปแบบที่ผิดแล้วให้ขึ้น ok
	g.Expect(ok).NotTo(BeTrue())

	//err เป็นค่า nil ไม่ได้ ต้องแสดงออกมา
	g.Expect(err).NotTo(BeNil())

	//if err pose message error ออกมา
	g.Expect(err.Error()).To(Equal("Name: non zero value required"))

}

func TestTimeOutIsNotPast(t *testing.T) {
	g := NewGomegaWithT(t)

	ordertech := Video{
		Name: "meow",            //ผิด -->เช็คตรงนี้
		Url:  "meoeoe", //ถูก
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(ordertech)

	// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
	g.Expect(ok).NotTo(BeTrue())

	// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
	g.Expect(err).NotTo(BeNil())

	// err.Error() ต้องมี message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Url: meoeoe does not validate as url"))
}