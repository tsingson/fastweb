package utils

import (
	"bytes"
	"errors"
	"io/ioutil"

	"github.com/rs/zerolog"

	"github.com/tsingson/chardet"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func Detect(b []byte) (*chardet.Result, error) {
	textDetector := chardet.NewTextDetector()
	return textDetector.DetectBest(b)

}

//convert GBK to UTF-8
func Decodegbk(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//convert BIG5 to UTF-8
func Decodebig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//convert UTF-8 to BIG5
func Encodebig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewEncoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// Trans  translate  code to simplechinese
func Trans(input []byte, log zerolog.Logger) (output []byte, err error) {

	var code *chardet.Result

	code, err = Detect(input)
	if err != nil {
		log.Error().Err(err).Msg("ioutil ReadAll error")
		return nil, err
	}

	switch code.Charset {
	case "GB-18030":
		output, err = Decodegbk(input)
		if err != nil {
			log.Error().Err(err).Msg("Error")
			//	return output, err
		}
		break
	case "Big5":
		output, err = Decodebig5(input)
		if err != nil {
			log.Error().Err(err).Msg("Error")
			//	return output, err
		}
		break
	case "UTF-8":
		output = input
		//return output, nil
		break
	default:
		err = errors.New("unknow code type")
	}
	return output, err

}
