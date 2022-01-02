package puzzle8

import "strings"

type Display struct {
	Top         string
	TopLeft     string
	TopRight    string
	Middle      string
	BottomLeft  string
	BottomRight string
	Bottom      string
}

func (d Display) String() string {
	template := [][]string{
		{" ", "X", "X", "X", "X", " "},
		{"X", " ", " ", " ", " ", "Y"},
		{"X", " ", " ", " ", " ", "Y"},
		{" ", "X", "X", "X", "X", " "},
		{"X", " ", " ", " ", " ", "Y"},
		{"X", " ", " ", " ", " ", "Y"},
		{" ", "X", "X", "X", "X", " "},
	}

	sb := strings.Builder{}
	for i, line := range template {
		var key map[string]string
		switch i {
		case 0:
			key = map[string]string{"X": d.Top}
		case 1, 2:
			key = map[string]string{"X": d.TopLeft, "Y": d.TopRight}
		case 3:
			key = map[string]string{"X": d.Middle}
		case 4, 5:
			key = map[string]string{"X": d.BottomLeft, "Y": d.BottomRight}
		case 7:
			key = map[string]string{"X": d.Bottom}
		}

		if i != 0 {
			sb.WriteString("\n")
		}
		for _, k := range line {
			if v, found := key[k]; found {
				sb.WriteString(v)
			} else {
				sb.WriteString(k)
			}
		}
	}

	return sb.String()
}
