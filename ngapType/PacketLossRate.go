package ngapType

// Need to import "free5gc/lib/aper" if it uses "aper"

type PacketLossRate struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:1000"`
}
