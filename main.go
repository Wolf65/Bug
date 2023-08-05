package main

import (
	"unsafe"

	imgui "github.com/AllenDang/cimgui-go"
)

func main() {
	backend := imgui.CreateBackend(imgui.NewGLFWBackend())
	backend.CreateWindow("bug", 1200, 900, 0)

	RusRanges := []uint16{0x0020, 0xA69F, 0}
	IconRange := []uint16{uint16(IconsFontAwesome6.Min), uint16(IconsFontAwesome6.Max16), 0}

	jetBrainsConfig := imgui.NewFontConfig()
	fontAwesomeConfig := imgui.NewFontConfig()
	//v0.0.0-20230801075634-218b71149299 crach or SetMergeMode(false)
	fontAwesomeConfig.SetMergeMode(true) //v0.0.0-20230619023324-e4dae85333e0 good
	fontAwesomeConfig.SetPixelSnapH(true)
	io := imgui.CurrentIO()

	io.Fonts().AddFontFromFileTTFV("JetBrainsMono-Medium.ttf",
		15,
		jetBrainsConfig,
		(*imgui.Wchar)(unsafe.Pointer(&RusRanges[0])))

	io.Fonts().AddFontFromFileTTFV("fa-solid-900.ttf",
		40,
		fontAwesomeConfig,
		(*imgui.Wchar)(unsafe.Pointer(&IconRange[0])))

	backend.Run(loop)
}

func loop() {
	imgui.Text("Test")
	imgui.Text("Тест")
	imgui.SameLine()
	imgui.Text(IconsFontAwesome6.Icons["AddressBook"])
}

type Font struct {
	// The filenames of the associated TTF files are provided in Filenames,
	// where each entry stores first an abbreviated name for the fot and
	// then the actual filename.
	Filenames [][2]string
	// The range of Unicode code points is given by [Min, Max). The largest
	// 16-bit code point is stored in Max16.
	Min, Max16, Max int
	// Icons stores the mapping from user-friendly names to code points.
	Icons map[string]string
}
