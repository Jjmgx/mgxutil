package mgxutil

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
	"unicode/utf8"

	"github.com/axgle/mahonia"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var ErrInvalidUtf8Rune = errors.New("Not Invalid Utf8 runes")

func Utf8ToGbk(src string) string {
	desCoder := mahonia.NewEncoder("gbk")
	return desCoder.ConvertString(src)
}
func GbkToUtf8(src string) string {
	srcCoder := mahonia.NewDecoder("gbk")
	return srcCoder.ConvertString(src)
}

func Ucs2ToUtf8(in string) (string, error) {
	r := bytes.NewReader([]byte(in))
	t := transform.NewReader(r, unicode.All[1].NewDecoder()) //UTF-16 bigendian, no-bom
	out, err := ioutil.ReadAll(t)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func Utf8ToGB18030(in string) (string, error) {
	if !utf8.ValidString(in) {
		return "", ErrInvalidUtf8Rune
	}

	r := bytes.NewReader([]byte(in))
	t := transform.NewReader(r, simplifiedchinese.GB18030.NewEncoder())
	out, err := ioutil.ReadAll(t)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func Utf8ToUcs2Byte(in string) ([]byte, error) {
	if !utf8.ValidString(in) {
		return []byte{}, ErrInvalidUtf8Rune
	}
	r := bytes.NewReader([]byte(in))
	t := transform.NewReader(r, unicode.All[1].NewEncoder())
	return ioutil.ReadAll(t)
}

func Nrzm(s string) string {
	s = GbkToUtf8(s)
	s = strings.Replace(s, "\\", "\\\\", -1)
	s = strings.Replace(s, "\r", "\\\r", -1)
	s = strings.Replace(s, "\n", "\\\n", -1)
	s = strings.Replace(s, "'", "\\'", -1)
	s = strings.Replace(s, "(", "\\(", -1)
	s = strings.Replace(s, ")", "\\)", -1)
	s = strings.Replace(s, "`", "\\`", -1)
	return Utf8ToGbk(s)
}
