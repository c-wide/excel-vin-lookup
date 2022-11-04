package excel

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/c-wide/go-nhtsa"
	"github.com/xuri/excelize/v2"
)

func ProcessFile(filePath string) (*[]nhtsa.VinRequest, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file.\npath: %s\nerror: %s", filePath, err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("error closing file. %s\n", err)
		}
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, fmt.Errorf("error reading rows. %s", err)
	}

	pattern, err := regexp.Compile(`^(?:19|20)\d{2}$`)
	if err != nil {
		return nil, fmt.Errorf("error creating regex. %s", err)
	}

	vList := make([]nhtsa.VinRequest, len(rows))

	for rowIdx, row := range rows {
		rowLen := len(row)

		if rowLen == 0 {
			return nil, fmt.Errorf("row %d is blank", rowIdx+1)
		}

		if row[0] == "" {
			return nil, fmt.Errorf("missing vin in row %d detected", rowIdx+1)
		}

		year := ""
		if rowLen > 1 {
			s := strings.Trim(row[1], " ")

			if pattern.MatchString(s) {
				year = s
			} else {
				fmt.Printf("ignoring year for VIN %s in row %d, wrong format.\n", row[0], rowIdx+1)
			}
		}

		vList[rowIdx] = nhtsa.VinRequest{Vin: strings.Trim(row[0], " "), Year: year}
	}

	return &vList, nil
}

func WriteFile(vReqs *[]nhtsa.VinRequest, vData *[]nhtsa.DecodeFlatResult, filePath string) error {
	f := excelize.NewFile()

	index := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)

	style, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "gradient", Color: []string{"#FF0000", "#FF0000"}, Pattern: 1},
	})
	if err != nil {
		return fmt.Errorf("error creating cell style. %s", err)
	}

	reqs := *vReqs

	for vIdx, vin := range *vData {
		tgtIdx := vIdx + 1

		rowTgt := map[string]string{
			"a": fmt.Sprintf("A%d", tgtIdx),
			"b": fmt.Sprintf("B%d", tgtIdx),
			"c": fmt.Sprintf("C%d", tgtIdx),
		}

		f.SetCellValue("Sheet1", rowTgt["a"], reqs[vIdx].Vin)
		f.SetCellValue("Sheet1", rowTgt["b"], reqs[vIdx].Year)

		if vin.ErrorCode != "0" {
			f.SetCellValue("Sheet1", rowTgt["c"], vin.ErrorText)

			for _, row := range rowTgt {
				if err := f.SetCellStyle("Sheet1", row, row, style); err != nil {
					return fmt.Errorf("error setting cell style. %s", err)
				}
			}

			continue
		}

		f.SetCellValue("Sheet1", rowTgt["c"], vin.GVWR)
	}

	basePath := filepath.Base(filePath)
	fileName := fmt.Sprintf("%s-complete.xlsx", basePath[:len(basePath)-len(filepath.Ext(filePath))])

	if err := f.SaveAs(fileName); err != nil {
		return fmt.Errorf("error saving excel file. %s", err)
	}

	return nil
}
