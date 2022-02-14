package apputils

import (
	"database/sql"
	"hash/crc32"
	"strings"
)

/*
*
 */
func StringArrToStringPointersArr(inputArr []string) []*string {

	var retArr []*string
	for i := 0; i < len(inputArr); i++ {
		retArr = append(retArr, &inputArr[i])
	}
	return retArr
}

/*
*
 */
func StringsMapToJsonString(inputMap map[string]string) string {

	var jsonStringBuilder strings.Builder
	jsonStringBuilder.WriteString(`{`)
	for key, element := range inputMap {

		jsonStringBuilder.WriteString(`"`)
		jsonStringBuilder.WriteString(key)
		jsonStringBuilder.WriteString(`":`)
		jsonStringBuilder.WriteString(`"`)
		jsonStringBuilder.WriteString(element)
		jsonStringBuilder.WriteString(`",`)

	}
	tempString := jsonStringBuilder.String()
	return tempString[:len(tempString)-1] + "}"
}

/*
*
 */
func StringArrToCsvString(inputArr []string) string {

	var csvBuilder strings.Builder
	for _, val := range inputArr {
		csvBuilder.WriteString(`"`)
		csvBuilder.WriteString(val)
		csvBuilder.WriteString(`",`)
	}
	retString := csvBuilder.String()
	retString = retString[:len(retString)-1]
	return retString
}

/*
*
 */
func NullColumnValue(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

/*
*
 */
func Crc32OfString(of string) uint32 {
	return crc32.ChecksumIEEE([]byte(of))
}
