package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// ตรวจสอบค่าว่างของชื่อแล้วต้องเจอ Error
func ReviewCommentSystemNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	review := Review{
		CheckedPayment_ID:1,
		Satisfaction_System_ID:5,
		Review_Comment_System:"",
		Satisfaction_Technician_ID: 5,
		Review_Comment_Technician: "BBBB",
		TimestampReview:time.Now(),
		StatusReview: false, //ผิด
		Customer_ID:1,
		CheckSucceed: false,

	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(review)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Review Comment System is Blank"))
}

// func TestEmailMustBeValid(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	user := User{
// 		Name:     "Abc",
// 		Email:    "qwe#123", // ผิด
// 		Password: "111",
// 		Role:     "employee",
// 	}

// 	ok, err := govalidator.ValidateStruct(user)

// 	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
// 	g.Expect(ok).ToNot(BeTrue())

// 	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
// 	g.Expect(err).ToNot(BeNil())

// 	// err.Error ต้องมี error message แสดงออกมา
// 	g.Expect(err.Error()).To(Equal("Email: qwe#123 does not validate as email"))
// }
