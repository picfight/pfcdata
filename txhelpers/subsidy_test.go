package txhelpers

import (
	"testing"

	"github.com/picfight/pfcd/chaincfg"
)

func TestBlockSubsidyDecred(t *testing.T) {
	totalSubsidy := UltimateSubsidy(&chaincfg.DecredNetParams)

	if totalSubsidy != 2099999999800912 {
		t.Errorf("Bad total subsidy; want 2099999999800912, got %v", totalSubsidy)
	}
}

func TestBlockSubsidyPicFightCoin(t *testing.T) {
	totalSubsidy := UltimateSubsidy(&chaincfg.PicFightCoinNetParams)

	if totalSubsidy != 2099999999800912 {
		t.Errorf("Bad total subsidy; want 2099999999800912, got %v", totalSubsidy)
	}
}
