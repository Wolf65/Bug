package main

import (
	imgui "github.com/AllenDang/cimgui-go"
)

var (
	baseFontSize        float32 = 13
	fontAwesomeFontSize float32 = baseFontSize * 4 / 3
)

func main() {
	backend := imgui.CreateBackend(imgui.NewGLFWBackend())
	backend.CreateWindow("bug", 1200, 900, 0)

	// Base Font

	io := imgui.CurrentIO()

	baseConfig := *imgui.NewFontConfig()
	baseConfig.SetPixelSnapH(true)

	baseRange := imgui.NewGlyphRange()
	baseBuilder := imgui.NewFontGlyphRangesBuilder()
	baseBuilder.AddRanges(io.Fonts().GlyphRangesCyrillic())
	baseBuilder.BuildRanges(baseRange)

	io.Fonts().AddFontFromFileTTFV("fonts/ttf/JetBrainsMono-Medium.ttf",
		baseFontSize,
		&baseConfig,
		baseRange.Data())

	//File: D:/a/cimgui-go/cimgui-go/cimgui/imgui/imgui_draw.cpp, Line: 2208
	//exit status 1

	// FontAwesome
	fontAwesomeConfig := *imgui.NewFontConfig()
	fontAwesomeConfig.SetMergeMode(true)
	fontAwesomeConfig.SetGlyphMinAdvanceX(fontAwesomeFontSize)
	fontAwesomeConfig.SetGlyphMaxAdvanceX(fontAwesomeFontSize)
	fontAwesomeConfig.SetPixelSnapH(true)
	fontAwesomeConfig.SetGlyphOffset(imgui.Vec2{X: 0, Y: 2})

	////Whoa, I caught a bug while I was doing the test. If you uncomment the code below
	//IconRange := []uint16{uint16(IconsFontAwesome6.Min), uint16(IconsFontAwesome6.Max16), 0}
	//io.Fonts().AddFontFromFileTTFV(fmt.Sprintf("fonts/ttf/%s",
	//	IconsFontAwesome6.Filenames[1][1]),
	//	fontAwesomeFontSize,
	//	&fontAwesomeConfig,
	//	imgui.SliceToPtr(IconRange))
	////.\main.go:49:3: cannot use imgui.SliceToPtr(IconRange) (value of type *uint16) as *imgui.Wchar value in argument to io.Fonts().AddFontFromFileTTFV

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
