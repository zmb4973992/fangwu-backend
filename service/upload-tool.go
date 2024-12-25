package service

import (
	"fangwu-backend/util"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"golang.org/x/image/bmp"
)

// 仿照gin里的c.SaveUploadedFile的写法
func saveFile(fileHeader *multipart.FileHeader, destination string) (resCode int, errDetail *util.ErrDetail) {
	//打开、读取文件
	openedFile, err := fileHeader.Open()
	if err != nil {
		return util.ErrorFailToOpenFile, util.GetErrDetail(err)
	}
	defer openedFile.Close()

	//创建空的新文件
	createdFile, err := os.Create(destination)
	if err != nil {
		return util.ErrorFailToCreateFile, util.GetErrDetail(err)
	}
	defer createdFile.Close()

	//把打开的文件内容复制到新文件中
	_, err = io.Copy(createdFile, openedFile)
	if err != nil {
		return util.ErrorFailToCopyFile, util.GetErrDetail(err)
	}

	return util.Success, nil
}

// 裁剪图片：文件路径，最大边长（像素，另一边将自动调整，0则不修改），图片质量（0-100，建议80。100的话不做压缩）
func jpgResize(destination string, maxSideLength int, imageQuality int) (resCode int, errDetail *util.ErrDetail) {
	// 检查文件后缀是否为.jpg或.jpeg
	if !strings.HasSuffix(destination, ".jpg") &&
		!strings.HasSuffix(destination, ".jpeg") {
		return util.ErrorNonJpgFile, nil
	}

	//打开图片文件
	img, err := imaging.Open(destination)
	if err != nil {
		return util.ErrorFailToOpenFile, util.GetErrDetail(err)
	}

	//获取图片的边界矩形
	bounds := img.Bounds()

	//获取图片的长和宽
	width := bounds.Dx()
	height := bounds.Dy()

	//如果最大边长小于等于0，则直接保存图片
	if maxSideLength <= 0 {
		err = imaging.Save(img, destination, imaging.JPEGQuality(imageQuality))
		if err != nil {
			return util.ErrorFailToSaveFile, util.GetErrDetail(err)
		}
		return util.Success, nil
	}

	//如果图片的长和宽都小于等于最大边长，则直接返回原图
	if width <= maxSideLength && height <= maxSideLength {
		err = imaging.Save(img, destination, imaging.JPEGQuality(imageQuality))
		if err != nil {
			return util.ErrorFailToSaveFile, util.GetErrDetail(err)
		}
		return util.Success, nil
	}

	//获取图片质量
	if imageQuality < 0 || imageQuality > 100 {
		imageQuality = 80
	}

	//如果图片的宽大于高，则按宽度等比例缩放
	if width >= height {
		resizedImage := imaging.Resize(img, maxSideLength, 0, imaging.Lanczos)
		err = imaging.Save(resizedImage, destination, imaging.JPEGQuality(imageQuality))
		if err != nil {
			return util.ErrorFailToSaveFile, util.GetErrDetail(err)
		}
		return util.Success, nil
	} else { // 否则按高度等比例缩放
		resizedImage := imaging.Resize(img, 0, maxSideLength, imaging.Lanczos)
		err = imaging.Save(resizedImage, destination, imaging.JPEGQuality(imageQuality))
		if err != nil {
			return util.ErrorFailToSaveFile, util.GetErrDetail(err)
		}
		return util.Success, nil
	}
}

func pngToJpg(destination string) (newDestination string, resCode int, errDetail *util.ErrDetail) {
	// 打开PNG文件
	pngFile, err := os.Open(destination)
	if err != nil {
		return "", util.ErrorFailToOpenFile, util.GetErrDetail(err)
	}
	defer pngFile.Close()

	// 解码png图像
	pngImg, err := png.Decode(pngFile)
	if err != nil {
		return "", util.ErrorFailToDecodePng, util.GetErrDetail(err)
	}

	// 创建同名的jpg文件
	fileName := strings.TrimSuffix(destination, ".png")
	jpgImg, err := os.Create(fileName + ".jpg")
	if err != nil {
		return "", util.ErrorFailToCreateFile, util.GetErrDetail(err)
	}
	defer jpgImg.Close()

	// 编码jpg图像，写入文件
	err = jpeg.Encode(jpgImg, pngImg, &jpeg.Options{Quality: 100})
	if err != nil {
		return "", util.ErrorFailToEncodeJpg, util.GetErrDetail(err)
	}

	//生成新的文件路径
	newDestination = strings.TrimSuffix(destination, ".png") + ".jpg"

	return newDestination, util.Success, nil
}

func bmpToJpg(destination string) (newDestination string, resCode int, errDetail *util.ErrDetail) {
	// 打开bmp文件
	bmpFile, err := os.Open(destination)
	if err != nil {
		return "", util.ErrorFailToOpenFile, util.GetErrDetail(err)
	}
	defer bmpFile.Close()

	// 解码bmp图像
	bmpImg, err := bmp.Decode(bmpFile)
	if err != nil {
		return "", util.ErrorFailToDecodeBmp, util.GetErrDetail(err)
	}

	// 创建同名的jpg文件
	fileName := strings.TrimSuffix(destination, ".bmp")
	jpgImg, err := os.Create(fileName + ".jpg")
	if err != nil {
		return "", util.ErrorFailToCreateFile, util.GetErrDetail(err)
	}
	defer jpgImg.Close()

	// 编码jpg图像，写入文件
	err = jpeg.Encode(jpgImg, bmpImg, &jpeg.Options{Quality: 100})
	if err != nil {
		return "", util.ErrorFailToEncodeJpg, util.GetErrDetail(err)
	}

	//生成新的文件路径
	newDestination = strings.TrimSuffix(destination, ".bmp") + ".jpg"

	return newDestination, util.Success, nil
}
