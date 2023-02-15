package entity

import {
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
	
}


type Employee struct {
	gorm.Model
	Name 	string 			`valid: "require~Name cannot be blank"`
	Email	string			
	EmployeeID string		`valid: "matches(^[JMS]\\d{8}$)"`
}


func TestNameconnotbeBlank(t *testing.T){
	g := gomega.NewGomegaWithT(t)

	emp := Employee {
		Name: "",
		Email: "abc@hotmail.com",
		EmployeeID: "M12345678",
	}

	ok, err := govalidator.validatestruct(emp)

	g.expect(ok).Tonot(Betrue())
	g.expect(err).Tonot(BeNill())
	g.expect(err.Error()).To(equal("Name can not be blank"))
	


}
