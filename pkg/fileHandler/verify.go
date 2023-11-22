package fileHandler

import (
	"al-mosso-api/internal/services/types"
	"strings"
)

func CheckFile(input *types.TFile) bool {
	exts := []string{"png", "jpg", "jpeg", "webp", "avif"}

	ext := false
	for _, element := range exts {
		if strings.Contains(strings.ToLower(input.Extension), element) {
			ext = true
		}
	}

	if ext == false {

		return false
	}

	realExt, err := checkFileType(input, input.Extension)
	if err != nil {
		return false
	}

	if !strings.Contains(strings.ToLower(input.Extension), strings.ToLower(realExt)) {

		return false
	}

	return true
}

func checkFileType(input *types.TFile, ext string) (string, error) {

	var magicBytes = map[string][]byte{ //https://en.wikipedia.org/wiki/List_of_file_signatures
		"jpeg": {0xFF, 0xD8, 0xFF},
		"png":  {0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A},
		"gif":  {0x47, 0x49, 0x46},
		"webp": {0x52, 0x49, 0x46, 0x46, 0x00, 0x00, 0x00, 0x00, 0x57, 0x45, 0x42, 0x50},
		"avif": {0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70, 0x61, 0x76, 0x69, 0x66},
		"jpg":  {0xFF, 0xD8, 0xFF},
	}

	//TODO -> CORRIGIR ERRO NO COPY PARA BUFFER,
	/*
		se eu usar input.FileData.Read(buffer) a imagem no final é salva com erro
		se eu uso a linha que está abaixo, ele salva [0,0,0,0,0,0,0,0,0]
		achar um copy seguro
	*/

	var extKey string

	for key := range magicBytes {

		if strings.Contains(ext, key) {
			extKey = key
		}
	}
	magicToTest := magicBytes[extKey]
	result := bytesMatch(input.FileData, magicToTest)
	if result == true {
		return extKey, nil
	}

	return "Unknow", nil
}

func bytesMatch(buffer []byte, magic []byte) bool {

	for i := range magic {
		if buffer[i] != magic[i] { //caso o byte do buffer seja diferente do magic byte

			return false
		}
	}

	return true
}
