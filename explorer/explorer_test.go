package explorer

import (
	"testing"

	"github.com/picfight/pfcd/chaincfg"
)

func TestTestNet3Name(t *testing.T) {
	netName := netName(&chaincfg.TestNet3Params)
	if netName != "Testnet" {
		t.Errorf(`Net name not "Testnet": %s`, netName)
	}
}

func TestPicfightcoinNetName(t *testing.T) {
	netName := netName(&chaincfg.PicFightCoinNetParams)
	if netName != "PicFight Coin" {
		t.Errorf(`Net name not "PicFight Coin": %s`, netName)
	}
}

func TestSimNetName(t *testing.T) {
	netName := netName(&chaincfg.SimNetParams)
	if netName != "Simnet" {
		t.Errorf(`Net name not "Simnet": %s`, netName)
	}
}
