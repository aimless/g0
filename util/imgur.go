package util

import (
	"errors"
	"fmt"
	"log"
	"regexp"
)

func ImgurGetImagesFromAlbum(id string) ([]string, error) {
	var links []string
	var urlregex = regexp.MustCompile(`data\-src="\/\/i\.imgur\.com\/(?P<id>(.*)[^s]\.(jpg|jpeg|png|gif|apng|tiff|bmp))" `)
	urlregex = regexp.MustCompile(`<link rel="image_src" href="(http:\/\/|)i\.imgur\.com\/(?P<id>(.*)[^s]\.(jpg|jpeg|png|gif|apng|tiff|bmp))"\/>`)
	if id == "" {
		err := errors.New("Empty id.")
		log.Printf("Util.ImgurGetImagesFromGallery: %s\n", err.Error())
		return links, err
	}

	src, err := DownloadPage("http://imgur.com/a/" + id)
	if err != nil {
		log.Printf("Util.ImgurGetImagesFromGallery: %s\n", err.Error())
		return links, err
	}

	if urlregex.MatchString(src) {
		raw := urlregex.FindAllStringSubmatch(src, -1)

		for _, n := range raw {
			links = append(links, "http://i.imgur.com/"+n[2])
		}
	} else {
		return nil, errors.New("No links found")
	}
	return links, nil
}

func ImgurGetImageFromSubreddit(id string) (string, error) {
	var links []string
	var urlregex = regexp.MustCompile(`<img src="\/\/i\.imgur\.com\/(?P<id>(.*)\.(jpg|jpeg|png|gif|apng|tiff|bmp))`)

	if id == "" {
		log.Printf("util.ImgurGetImageFromSubreddit: empty id.\n")
		return "", errors.New("No id.")
	}

	src, err := DownloadPage("http://imgur.com/" + id)
	if err != nil {
		log.Printf("util.ImgurGetImageFromSubreddit: %s\n", err.Error())
		return "", err
	}

	if urlregex.MatchString(src) {
		raw := urlregex.FindAllStringSubmatch(src, -1)
		for _, n := range raw {
			links = append(links, "http://i.imgur.com/"+n[1])
		}
	}

	if len(links) == 0 {
		return "", errors.New("No images found")
	}
	return links[0], nil
}

func ImgurGetImagesFromGallery(id string) ([]string, error) {
	var links []string
	var urlregex = regexp.MustCompile(`<img src="\/\/i\.imgur\.com\/(?P<id>(.*)\.(jpg|jpeg|png|gif|apng|tiff|bmp))`)

	if id == "" {
		err := errors.New("Empty id.")
		log.Printf("Util.ImgurGetImagesFromGallery: %s\n", err.Error())
		return links, err
	}

	src, err := DownloadPage("http://imgur.com/gallery/" + id)
	if err != nil {
		log.Printf("Util.ImgurGetImagesFromGallery: %s\n", err.Error())
		return links, err
	}

	if urlregex.MatchString(src) {
		raw := urlregex.FindAllStringSubmatch(src, -1)
		for _, n := range raw {
			links = append(links, "http://i.imgur.com/"+n[1])
		}
	}
	if links == nil {
		err := errors.New("No image found in gallery\n")
		return nil, err
	}
	return links, nil
}

func DropBoxLinkExtractor(url string) (s string, err error) {
	if url == "" {
		err = errors.New("Emtpy url")
		log.Printf("Util.DropBoxLinkExtractor: %s\n", err.Error())
		return "", err
	}

	src, err := DownloadPage(url)
	if err != nil {
		log.Printf("Util.DropBoxLinkExtractor: %s\n", err.Error())
		return "", err
	}
	urlregex := regexp.MustCompile(`", "(?P<url>https:\/\/dl\.dropboxusercontent\.com\/(.*))"\) }\);`)

	if urlregex.MatchString(src) {
		fmt.Println("hi")
		raw := urlregex.FindStringSubmatch(src)

		s = "https://dl.dropboxusercontent.com/" + raw[2]
		fmt.Println(s)
		return s, nil
	}
	panic("unreachable")
}
