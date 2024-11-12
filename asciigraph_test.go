package asciigraph

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestPlot(t *testing.T) {
	cases := []struct {
		data     []float64
		opts     []Option
		expected string
	}{

		{
			[]float64{1, 1, 1, 1, 1},
			[]Option{AlwaysY(false)},
			` 1.00 ┼────`},
		{
			[]float64{0, 0, 0, 0, 0},
			[]Option{AlwaysY(false)},
			` 0.00 ┼────`},
		{
			[]float64{2, 1, 1, 2, -2, 5, 7, 11, 3, 7, 1},
			[]Option{AlwaysY(false)},
			`
 11.00 ┤      ╭╮
 10.00 ┤      ││
  9.00 ┤      ││
  8.00 ┤      ││
  7.00 ┤     ╭╯│╭╮
  6.00 ┤     │ │││
  5.00 ┤    ╭╯ │││
  4.00 ┤    │  │││
  3.00 ┤    │  ╰╯│
  2.00 ┼╮ ╭╮│    │
  1.00 ┤╰─╯││    ╰
  0.00 ┤   ││
 -1.00 ┤   ││
 -2.00 ┤   ╰╯`},
		{
			[]float64{2, 1, 1, 2, -2, 5, 7, 11, 3, 7, 4, 5, 6, 9, 4, 0, 6, 1, 5, 3, 6, 2},
			[]Option{Caption("Plot using asciigraph."), AlwaysY(false)},
			`
 11.00 ┤      ╭╮
 10.00 ┤      ││
  9.00 ┤      ││    ╭╮
  8.00 ┤      ││    ││
  7.00 ┤     ╭╯│╭╮  ││
  6.00 ┤     │ │││ ╭╯│ ╭╮  ╭╮
  5.00 ┤    ╭╯ │││╭╯ │ ││╭╮││
  4.00 ┤    │  ││╰╯  ╰╮││││││
  3.00 ┤    │  ╰╯     ││││╰╯│
  2.00 ┼╮ ╭╮│         ││││  ╰
  1.00 ┤╰─╯││         ││╰╯
  0.00 ┤   ││         ╰╯
 -1.00 ┤   ││
 -2.00 ┤   ╰╯
        Plot using asciigraph.`},
		{
			[]float64{.2, .1, .2, 2, -.9, .7, .91, .3, .7, .4, .5},
			[]Option{Caption("Plot using asciigraph."), AlwaysY(false)},
			`
  2.00 ┤  ╭╮ ╭╮
  0.55 ┼──╯│╭╯╰───
 -0.90 ┤   ╰╯
        Plot using asciigraph.`},
		{
			[]float64{2, 1, 1, 2, -2, 5, 7, 11, 3, 7, 1},
			[]Option{Height(4), Offset(3), AlwaysY(false)},
			`
 11.00 ┤      ╭╮
  7.75 ┤    ╭─╯│╭╮
  4.50 ┼╮ ╭╮│  ╰╯│
  1.25 ┤╰─╯││    ╰
 -2.00 ┤   ╰╯`},
		{
			[]float64{.453, .141, .951, .251, .223, .581, .771, .191, .393, .617, .478},
			[]Option{AlwaysY(false)},
			`
 0.95 ┤ ╭╮
 0.85 ┤ ││  ╭╮
 0.75 ┤ ││  ││
 0.65 ┤ ││ ╭╯│ ╭╮
 0.55 ┤ ││ │ │ │╰
 0.44 ┼╮││ │ │╭╯
 0.34 ┤│││ │ ││
 0.24 ┤││╰─╯ ╰╯
 0.14 ┤╰╯`},

		{
			[]float64{.01, .004, .003, .0042, .0083, .0033, 0.0079},
			[]Option{AlwaysY(false)},
			`
 0.010 ┼╮
 0.009 ┤│
 0.008 ┤│  ╭╮╭
 0.007 ┤│  │││
 0.006 ┤│  │││
 0.005 ┤│  │││
 0.004 ┤╰╮╭╯││
 0.003 ┤ ╰╯ ╰╯`},

		{
			[]float64{192, 431, 112, 449, -122, 375, 782, 123, 911, 1711, 172},
			[]Option{Height(10), AlwaysY(false)},
			`
 1711 ┤        ╭╮
 1528 ┤        ││
 1344 ┤        ││
 1161 ┤        ││
  978 ┤       ╭╯│
  794 ┤     ╭╮│ │
  611 ┤     │││ │
  428 ┤╭╮╭╮╭╯││ │
  245 ┼╯╰╯││ ╰╯ ╰
   61 ┤   ││
 -122 ┤   ╰╯`},
		{
			[]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[]Option{Height(10), AlwaysY(true)},
			`
 0.00 ┤
 0.00 ┤
 0.00 ┤
 0.00 ┤
 0.00 ┤
 0.00 ┤
 0.00 ┤
 0.00 ┤
 0.00 ┤
 0.00 ┤
 0.00 ┼──────────`},
		{
			[]float64{0.3189989805, 0.149949026, 0.30142492354, 0.195129182935, 0.3142492354, 0.1674974513, 0.3142492354, 0.1474974513, 0.3047974513},
			[]Option{Width(30), Height(5), Caption("Plot with custom height & width."), AlwaysY(false)},
			`
 0.32 ┼╮            ╭─╮     ╭╮     ╭
 0.29 ┤╰╮    ╭─╮   ╭╯ │    ╭╯│     │
 0.26 ┤ │   ╭╯ ╰╮ ╭╯  ╰╮  ╭╯ ╰╮   ╭╯
 0.23 ┤ ╰╮ ╭╯   ╰╮│    ╰╮╭╯   ╰╮ ╭╯
 0.20 ┤  ╰╮│     ╰╯     ╰╯     │╭╯
 0.16 ┤   ╰╯                   ╰╯
       Plot with custom height & width.`},
		{
			[]float64{
				0, 0, 0, 0, 1.5, 0, 0, -0.5, 9, -3, 0, 0, 1, 2, 1, 0, 0, 0, 0,
				0, 0, 0, 0, 1.5, 0, 0, -0.5, 8, -3, 0, 0, 1, 2, 1, 0, 0, 0, 0,
				0, 0, 0, 0, 1.5, 0, 0, -0.5, 10, -3, 0, 0, 1, 2, 1, 0, 0, 0, 0,
			},
			[]Option{Offset(10), Height(10), Caption("I'm a doctor, not an engineer."), AlwaysY(false)},
			`
     10.00    ┤                                             ╭╮
      8.70    ┤       ╭╮                                    ││
      7.40    ┤       ││                 ╭╮                 ││
      6.10    ┤       ││                 ││                 ││
      4.80    ┤       ││                 ││                 ││
      3.50    ┤       ││                 ││                 ││
      2.20    ┤       ││   ╭╮            ││   ╭╮            ││   ╭╮
      0.90    ┤   ╭╮  ││  ╭╯╰╮       ╭╮  ││  ╭╯╰╮       ╭╮  ││  ╭╯╰╮
     -0.40    ┼───╯╰──╯│╭─╯  ╰───────╯╰──╯│╭─╯  ╰───────╯╰──╯│╭─╯  ╰───
     -1.70    ┤        ││                 ││                 ││
     -3.00    ┤        ╰╯                 ╰╯                 ╰╯
                            I'm a doctor, not an engineer.`},
		{
			[]float64{-5, -2, -3, -4, 0, -5, -6, -7, -8, 0, -9, -3, -5, -2, -9, -3, -1},
			[]Option{AlwaysY(false)},
			`
  0.00 ┤   ╭╮   ╭╮
 -1.00 ┤   ││   ││     ╭
 -2.00 ┤╭╮ ││   ││  ╭╮ │
 -3.00 ┤│╰╮││   ││╭╮││╭╯
 -4.00 ┤│ ╰╯│   │││││││
 -5.00 ┼╯   ╰╮  │││╰╯││
 -6.00 ┤     ╰╮ │││  ││
 -7.00 ┤      ╰╮│││  ││
 -8.00 ┤       ╰╯││  ││
 -9.00 ┤         ╰╯  ╰╯`},
		{
			[]float64{-0.000018527, -0.021, -.00123, .00000021312, -.0434321234, -.032413241234, .0000234234},
			[]Option{Height(5), Width(45), AlwaysY(false)},
			`
  0.000 ┼─╮           ╭────────╮                    ╭
 -0.008 ┤ ╰──╮     ╭──╯        ╰─╮                ╭─╯
 -0.017 ┤    ╰─────╯             ╰╮             ╭─╯
 -0.025 ┤                         ╰─╮         ╭─╯
 -0.034 ┤                           ╰╮   ╭────╯
 -0.042 ┤                            ╰───╯`},
		{
			[]float64{57.76, 54.04, 56.31, 57.02, 59.5, 52.63, 52.97, 56.44, 56.75, 52.96, 55.54, 55.09, 58.22, 56.85, 60.61, 59.62, 59.73, 59.93, 56.3, 54.69, 55.32, 54.03, 50.98, 50.48, 54.55, 47.49, 55.3, 46.74, 46, 45.8, 49.6, 48.83, 47.64, 46.61, 54.72, 42.77, 50.3, 42.79, 41.84, 44.19, 43.36, 45.62, 45.09, 44.95, 50.36, 47.21, 47.77, 52.04, 47.46, 44.19, 47.22, 45.55, 40.65, 39.64, 37.26, 40.71, 42.15, 36.45, 39.14, 36.62},
			[]Option{Width(-10), Height(-10), Offset(-1), AlwaysY(false)},
			`
 60.61 ┤             ╭╮ ╭╮
 59.60 ┤   ╭╮        │╰─╯│
 58.60 ┤   ││      ╭╮│   │
 57.59 ┼╮ ╭╯│      │││   │
 56.58 ┤│╭╯ │ ╭─╮  │╰╯   ╰╮
 55.58 ┤││  │ │ │╭─╯      │╭╮    ╭╮
 54.57 ┤╰╯  │ │ ││        ╰╯╰╮ ╭╮││      ╭╮
 53.56 ┤    │╭╯ ╰╯           │ ││││      ││
 52.56 ┤    ╰╯               │ ││││      ││           ╭╮
 51.55 ┤                     ╰╮││││      ││           ││
 50.54 ┤                      ╰╯│││      ││╭╮      ╭╮ ││
 49.54 ┤                        │││  ╭─╮ ││││      ││ ││
 48.53 ┤                        │││  │ │ ││││      ││ ││
 47.52 ┤                        ╰╯│  │ ╰╮││││      │╰─╯╰╮╭╮
 46.52 ┤                          ╰─╮│  ╰╯│││      │    │││
 45.51 ┤                            ╰╯    │││   ╭──╯    ││╰╮
 44.50 ┤                                  │││ ╭╮│       ╰╯ │
 43.50 ┤                                  ││╰╮│╰╯          │
 42.49 ┤                                  ╰╯ ╰╯            │   ╭╮
 41.48 ┤                                                   │   ││
 40.48 ┤                                                   ╰╮ ╭╯│
 39.47 ┤                                                    ╰╮│ │╭╮
 38.46 ┤                                                     ││ │││
 37.46 ┤                                                     ╰╯ │││
 36.45 ┤                                                        ╰╯╰`},
		{
			[]float64{2, 1, 1, 2, -2, 5, 7, 11, 3, 7, 4, 5, 6, 9, 4, 0, 6, 1, 5, 3, 6, 2},
			[]Option{LowerBound(-3), UpperBound(13), AlwaysY(false)},
			` 13.00 ┤
 12.00 ┤
 11.00 ┤      ╭╮
 10.00 ┤      ││
  9.00 ┤      ││    ╭╮
  8.00 ┤      ││    ││
  7.00 ┤     ╭╯│╭╮  ││
  6.00 ┤     │ │││ ╭╯│ ╭╮  ╭╮
  5.00 ┤    ╭╯ │││╭╯ │ ││╭╮││
  4.00 ┤    │  ││╰╯  ╰╮││││││
  3.00 ┤    │  ╰╯     ││││╰╯│
  2.00 ┼╮ ╭╮│         ││││  ╰
  1.00 ┤╰─╯││         ││╰╯
  0.00 ┤   ││         ╰╯
 -1.00 ┤   ││
 -2.00 ┤   ╰╯
 -3.00 ┤`},
		{
			[]float64{2, 1, 1, 2, -2, 5, 7, 11, 3, 7, 4, 5, 6, 9, 4, 0, 6, 1, 5, 3, 6, 2},
			[]Option{LowerBound(0), UpperBound(3), AlwaysY(false)},
			` 11.00 ┤      ╭╮
 10.00 ┤      ││
  9.00 ┤      ││    ╭╮
  8.00 ┤      ││    ││
  7.00 ┤     ╭╯│╭╮  ││
  6.00 ┤     │ │││ ╭╯│ ╭╮  ╭╮
  5.00 ┤    ╭╯ │││╭╯ │ ││╭╮││
  4.00 ┤    │  ││╰╯  ╰╮││││││
  3.00 ┤    │  ╰╯     ││││╰╯│
  2.00 ┼╮ ╭╮│         ││││  ╰
  1.00 ┤╰─╯││         ││╰╯
  0.00 ┤   ││         ╰╯
 -1.00 ┤   ││
 -2.00 ┤   ╰╯`},

		{
			[]float64{1, 1, math.NaN(), 1, 1},
			[]Option{AlwaysY(false)},
			` 1.00 ┼─╴╶─`},
		{
			[]float64{math.NaN(), 1},
			[]Option{AlwaysY(false)},
			` 1.00 ┤╶`},
		{
			[]float64{0, 0, 1, 1, math.NaN(), math.NaN(), 3, 3, 4},
			[]Option{AlwaysY(false)},
			`
 4.00 ┤       ╭
 3.00 ┤     ╶─╯
 2.00 ┤
 1.00 ┤ ╭─╴
 0.00 ┼─╯`},
		{
			[]float64{.1, .2, .3, math.NaN(), .5, .6, .7, math.NaN(), math.NaN(), .9, 1},
			[]Option{AlwaysY(false)},
			`
 1.00 ┤         ╭
 0.90 ┤        ╶╯
 0.80 ┤
 0.70 ┤     ╭╴
 0.60 ┤    ╭╯
 0.50 ┤   ╶╯
 0.40 ┤
 0.30 ┤ ╭╴
 0.20 ┤╭╯
 0.10 ┼╯`},
		{
			[]float64{-0.000018527, -0.021, -.00123, .00000021312, -.0434321234, -.032413241234, .0000234234},
			[]Option{Height(5), Width(45), Precision(5), AlwaysY(false)},
			`
  0.000023 ┼─╮           ╭────────╮                    ╭
 -0.008467 ┤ ╰──╮     ╭──╯        ╰─╮                ╭─╯
 -0.016958 ┤    ╰─────╯             ╰╮             ╭─╯
 -0.025449 ┤                         ╰─╮         ╭─╯
 -0.033940 ┤                           ╰╮   ╭────╯
 -0.042430 ┤                            ╰───╯`},

		{
			[]float64{math.NaN(), 1},
			[]Option{Caption("color test"), CaptionColor(Red), AxisColor(Green), LabelColor(Blue), AlwaysY(false)},
			`
\x1b[94m 1.00\x1b[0m \x1b[32m┤\x1b[0m╶
       \x1b[91mcolor test\x1b[0m`},
		{
			[]float64{.02, .03, .02},
			[]Option{AlwaysY(false)},
			`
 0.030 ┤╭╮
 0.020 ┼╯╰`},
		{
			[]float64{.2, .3, .1, .3},
			[]Option{AlwaysY(false)},
			`
 0.30 ┤╭╮╭
 0.20 ┼╯││
 0.10 ┤ ╰╯`},
		{
			[]float64{70 * 1024 * 1024 * 1024, 90 * 1024 * 1024 * 1024, 80 * 1024 * 1024 * 1024, 2 * 1024 * 1024 * 1024},
			[]Option{Height(5), Width(45), ValueFormatter(func(v float64) string {
				return fmt.Sprintf("%.2f Foo", v/1024/1024/1024)
			}), AlwaysY(false)},
			` 89.77 Foo ┤      ╭──────────────────────╮
 72.22 Foo ┼──────╯                      ╰──╮
 54.66 Foo ┤                                ╰───╮
 37.11 Foo ┤                                    ╰──╮
 19.55 Foo ┤                                       ╰──╮
  2.00 Foo ┤                                          ╰─`,
		},
	}

	for i := range cases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			c := cases[i]
			expected := strings.Replace(strings.TrimPrefix(c.expected, "\n"), `\x1b`, "\x1b", -1)
			actual := Plot(c.data, c.opts...)
			if actual != expected {
				conf := configure(config{AlwaysY: false}, c.opts)
				t.Errorf("Plot(%f, %#v)", c.data, conf)
				t.Logf("expected:\n%s\n", expected)
			}
			t.Logf("actual:\n%s\n", actual)
		})
	}
}

func TestPlotMany(t *testing.T) {
	cases := []struct {
		data     [][]float64
		opts     []Option
		expected string
	}{
		{
			[][]float64{{0}, {1}, {2}},
			nil,
			`
 2.00 ┼
 1.00 ┼
 0.00 ┼`},
		{
			[][]float64{{0, 0, 2, 2, math.NaN()}, {1, 1, 1, 1, 1, 1, 1}, {math.NaN(), math.NaN(), math.NaN(), 0, 0, 2, 2}},
			nil,
			`
 2.00 ┤ ╭─╴╭─
 1.00 ┼────│─
 0.00 ┼─╯╶─╯`},
		{
			[][]float64{{0, 0, 0}, {math.NaN(), 0, 0}, {math.NaN(), math.NaN(), 0}},
			nil,
			` 0.00 ┼╶╶`},
		{
			[][]float64{{0, 1, 0}, {2, 3, 4, 3, 2}, {4, 5, 6, 7, 6, 5, 4}},
			[]Option{Width(21), Caption("interpolation test"), AlwaysY(false)},
			`
 7.00 ┤        ╭──╮
 6.00 ┤    ╭───╯  ╰───╮
 5.00 ┤ ╭──╯          ╰──╮
 4.00 ┼─╯  ╭───╮         ╰─
 3.00 ┤ ╭──╯   ╰──╮
 2.00 ┼─╯         ╰─╴
 1.00 ┤ ╭───╮
 0.00 ┼─╯   ╰╴
        interpolation test`},

		{
			[][]float64{{0, 0}, {math.NaN(), 0}},
			[]Option{SeriesColors(Red), AlwaysY(false)},
			" 0.00 ┼╶"},
		{
			[][]float64{{0, 0}, {math.NaN(), 0}},
			[]Option{SeriesColors(Default, Red), AlwaysY(false)},
			" 0.00 ┼\x1b[91m╶\x1b[0m"},
		{
			[][]float64{{math.NaN(), 0, 2}, {0, 2}},
			[]Option{SeriesColors(Red, Red), AlwaysY(false)},
			`
 2.00 ┤\x1b[91m╭╭\x1b[0m
 1.00 ┤\x1b[91m││\x1b[0m
 0.00 ┼\x1b[91m╯╯\x1b[0m`},
		{
			[][]float64{{0, 1, 0}, {2, 3, 4, 3, 2}},
			[]Option{SeriesColors(Red, Blue), SeriesLegends("Red", "Blue"),
				Caption("legends with caption test"), AlwaysY(false)},
			`
 4.00 ┤ [94m╭╮[0m
 3.00 ┤[94m╭╯╰╮[0m
 2.00 ┼[94m╯[0m  [94m╰[0m
 1.00 ┤[91m╭╮[0m
 0.00 ┼[91m╯╰[0m
       legends with caption test

       [91m■[0m Red   [94m■[0m Blue`},
	}

	for i := range cases {
		name := fmt.Sprintf("%d", i)
		t.Run(name, func(t *testing.T) {
			c := cases[i]
			expected := strings.Replace(strings.TrimPrefix(c.expected, "\n"), `\x1b`, "\x1b", -1)
			actual := PlotMany(c.data, c.opts...)
			if actual != expected {
				conf := configure(config{}, c.opts)
				t.Errorf("Plot(%f, %#v)", c.data, conf)
				t.Logf("expected:\n%s\n", expected)
			}
			t.Logf("actual:\n%s\n", actual)
		})
	}
}

func BenchmarkPlot(b *testing.B) {
	data := []float64{2, 1, 1, 2, -2, 5, 7, 11, 3, 7, 1}
	opts := []Option{Height(4), Offset(3)}

	for i := 0; i < b.N; i++ {
		Plot(data, opts...)
	}
}

func BenchmarkPlotMany(b *testing.B) {
	data1 := []float64{2, 1, 1, 2, -2, 5, 7, 11, 3, 7, 1}
	data2 := []float64{5, 3, 2, 7, 1, -2, 9, 4, 3, 2, 1}
	opts := []Option{Height(4), Offset(3)}
	datasets := [][]float64{data1, data2}

	for i := 0; i < b.N; i++ {
		PlotMany(datasets, opts...)
	}
}
