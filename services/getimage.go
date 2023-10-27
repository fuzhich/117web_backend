package services

import "io/ioutil"

var BaseDir = "D:\\117webimages\\"

func GetImage(img string) (string, error) {
	if file, err := ioutil.ReadFile(BaseDir + img); err != nil {
		return "", err
	} else {
		return string(file), nil
	}
}
