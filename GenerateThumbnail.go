package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

const folderPath = "/Users/adrian/Desktop/PIC/"
const pngFormat = ".png"
const gifFormat = ".gif"
const jpgFormat = ".jpg"

func main() {

	var wg sync.WaitGroup

	fileList, err := ioutil.ReadDir(folderPath)

	if err != nil {
		fmt.Printf("Failed to read Folder\n%v", err)
	}

	for _, filename := range fileList {

		if checkPicFormat(filename.Name()) {
			wg.Add(1)
			thumbName := nameThumbnail(filename.Name())
			go generateThumbnail(folderPath+filename.Name(), thumbName, &wg)

		}
	}
	wg.Wait()
}

func checkPicFormat(filename string) bool {

	if strings.HasSuffix(filename, pngFormat) || strings.HasSuffix(filename, gifFormat) ||
		strings.HasSuffix(filename, jpgFormat) {

		return true

	} else {

		return false

	}

}

func generateThumbnail(fullFilePath string, thumbName string, wg *sync.WaitGroup) {
	originalPic, err := os.Open(fullFilePath)

	if err != nil {
		fmt.Printf("Failed to open picture file\n%v", err)

	}
	decodePic, format, err := image.Decode(originalPic)

	if err != nil {
		fmt.Printf("Failed to decode picture file\n%v", err)

	}

	originalPic.Close()

	afterResize := resize.Thumbnail(200, 200, decodePic, resize.Lanczos3)

	thumbFile, err := os.Create(folderPath + thumbName)

	if err != nil {
		fmt.Printf("Failed to create thumbnail\n%v", err)

	}

	defer thumbFile.Close()

	switch format {
	case `png`:
		png.Encode(thumbFile, afterResize)

	case `gif`:
		gif.Encode(thumbFile, afterResize, nil)

	case `jpeg`:
		jpeg.Encode(thumbFile, afterResize, nil)

	}

	wg.Done()

}

func nameThumbnail(filename string) string {

	suffix := path.Ext(filename)
	newFileName := strings.TrimSuffix(filename, suffix) + "_thumb" + suffix
	return newFileName

}
