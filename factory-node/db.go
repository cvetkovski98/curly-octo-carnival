package main

import (
	"github.com/cvetkovski98/factory-shared/schemas"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var resistorProductionMachine = schemas.Machine{
	Id:             0,
	Name:           "Resistor production machine",
	SerialNumber:   "C001-331",
	LineId:         0,
	LastServicedAt: &timestamppb.Timestamp{Seconds: 1646229356306},
	InstalledAt:    &timestamppb.Timestamp{Seconds: 1646229356306},
}

var transistorProductionMachine = schemas.Machine{
	Id:             1,
	Name:           "Transistor production machine",
	SerialNumber:   "C301-511",
	LineId:         0,
	LastServicedAt: &timestamppb.Timestamp{Seconds: 1646229356306},
	InstalledAt:    &timestamppb.Timestamp{Seconds: 1646229356306},
}

var pcbAssemblerMachine = schemas.Machine{
	Id:             2,
	Name:           "PCB assembling machine",
	SerialNumber:   "C042-031",
	LineId:         0,
	LastServicedAt: &timestamppb.Timestamp{Seconds: 1646229356306},
	InstalledAt:    &timestamppb.Timestamp{Seconds: 1646229356306},
}

var PCBProductionMachines = []*schemas.Machine{
	&resistorProductionMachine,
	&transistorProductionMachine,
	&pcbAssemblerMachine,
}

var rimProductionMachine = schemas.Machine{
	Id:             0,
	Name:           "Rim production machine",
	SerialNumber:   "A412442",
	LineId:         1,
	LastServicedAt: &timestamppb.Timestamp{Seconds: 1646229356306},
	InstalledAt:    &timestamppb.Timestamp{Seconds: 1646229356306},
}

var tireProductionMachine = schemas.Machine{
	Id:             1,
	Name:           "Tire production machine",
	SerialNumber:   "A003142",
	LineId:         1,
	LastServicedAt: &timestamppb.Timestamp{Seconds: 1646229356306},
	InstalledAt:    &timestamppb.Timestamp{Seconds: 1646229356306},
}

var wheelAssemblerMachine = schemas.Machine{
	Id:             2,
	Name:           "Wheel assembling machine",
	SerialNumber:   "A800213",
	LineId:         1,
	LastServicedAt: &timestamppb.Timestamp{Seconds: 1646229356306},
	InstalledAt:    &timestamppb.Timestamp{Seconds: 1646229356306},
}

var WheelProductionMachines = []*schemas.Machine{
	&rimProductionMachine,
	&tireProductionMachine,
	&wheelAssemblerMachine,
}

var wheelProductionLine = schemas.ProductionLine{
	Name:     "Wheel Production Line",
	Factory:  "The Wheel Factory ltd.",
	Zone:     "Zurich Industrial Zone 3",
	Machines: WheelProductionMachines,
}

var pcbProductionLine = schemas.ProductionLine{
	Name:     "PCB Production Line",
	Factory:  "The PCB Factory ltd.",
	Zone:     "Boston Industrial Zone 1",
	Machines: PCBProductionMachines,
}

var ProductionLines = []*schemas.ProductionLine{
	&wheelProductionLine,
	&pcbProductionLine,
}
