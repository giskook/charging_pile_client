package base

type Price struct {
	Start_hour      uint8
	Start_min       uint8
	End_hour        uint8
	End_min         uint8
	Elec_unit_price uint16
	Service_price   uint16
}

type Charging_Pile struct {
	ID              uint64
	BoxVersion      byte
	ProtocolVersion byte
	MaxCurrent      uint8
	Mode            uint8
	Prices          []Price
	Status          uint8

	ConnectWay uint8
	Wifi       string
	Passwd     string
	Interface  uint8
}