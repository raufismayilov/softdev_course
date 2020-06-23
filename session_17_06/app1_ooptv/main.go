package main

import "fmt"

// tv can   be switch on/off
// volume   up/down
// channel   up/down

// func showTV

type tv struct {
	isOn    bool
	volume  int8
	channel int8
	mute    bool
}

func (m *tv) tvSwitchOnOff() string {
	m.isOn = !m.isOn
	return "ok"
}

func (m *tv) tvVolumeUp() string {
	if !m.isOn {
		return "failed: TV is off"
	}
	if m.mute {
		m.mute = false
		return "ok"
	}
	m.volume++
	if m.volume > 100 {
		m.volume = 100
	}
	return "ok"
}

func (m *tv) tvVolumeDown() string {
	if !m.isOn {
		return "failed: TV is off"
	}
	if m.mute {
		m.mute = false
		return "ok"
	}
	m.volume--
	if m.volume < 0 {
		m.volume = 0
	}
	return "ok"
}

func (m *tv) tvChannelUp() string {
	if !m.isOn {
		return "failed: TV is off"
	}
	m.channel++
	if m.channel > 100 {
		m.channel = 1
	}

	return "ok"
}

func (m *tv) tvChannelDown() string {
	if !m.isOn {
		return "failed: TV is off"
	}
	m.channel--
	if m.channel < 1 {
		m.channel = 100
	}

	return "ok"
}

func (m *tv) tvMute() string {
	if !m.isOn {
		return "failed: TV is off"
	}
	m.mute = !m.mute
	return "ok"
}

func newTV() *tv {
	return &tv{
		isOn:    false,
		volume:  10,
		channel: 1,
	}
}
func (m *tv) tvShow() {
	if !m.isOn {
		fmt.Println("TV is off")
		return
	}
	volume := m.volume
	if m.mute {
		volume = 0
	}

	fmt.Printf("TV is on, volume %d, channel %d\n",
		volume, m.channel)
}

func main() {

	m := newTV()

	m.tvShow()

	m.tvSwitchOnOff()

	m.tvShow()

	m.tvMute()

	m.tvShow()

	m.tvVolumeUp()

	m.tvShow()

	m.tvVolumeUp()

	m.tvShow()

}
