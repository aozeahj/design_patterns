package main

import (
	"github.com/aozeahj/design_patterns/chain_of_responsibility"
	"github.com/aozeahj/design_patterns/chain_of_responsibility/demo/doctor"
	"github.com/aozeahj/design_patterns/chain_of_responsibility/demo/medical"
	 "github.com/aozeahj/design_patterns/chain_of_responsibility/demo/pay"
	"github.com/aozeahj/design_patterns/chain_of_responsibility/demo/reception"
)

func main() {
	receptionHandler := &reception.Handler{}
    doctorHandler := &doctor.Handler{}
    payHandler := &pay.Handler{}
    medicalHandler := &medical.Handler{}

    receptionHandler.SetNext(doctorHandler.SetNext(payHandler.SetNext(medicalHandler)))

    receptionHandler.Handle(&chain_of_responsibility.Client{})
}
