package main

import (
    "errors"
    "flag"
    "fmt"
    "io"
    "log"
    "net/http"
    "net/url"
    "os"
    "path/filepath"

    "github.com/PuerkitoBio/goquery"
)

const bingURL = "https://www.bing.com"

//从Bing获取壁纸链接.
func fetchWallpaperLink() (string, error) {
    resp, err := http.Get(bingURL)

    if err != nil {
        return "", err
    }

    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)

    if err != nil {
        return "", err
    }

    //找到包含壁纸链接的#preloadBg 元素.
    sel := doc.Find("#preloadBg").First()
    link, exists := sel.Attr("href")

    if !exists {
        return "", errors.New("Could not find #preloadBg element on Bing. Cannot fetch wallpaper link")
    }

    return fmt.Sprintf("%s%s", bingURL, link), nil
}

//返回壁纸链接的文件名.
//在 id GET 参数中找到.
//如果 overrideName 不为空,则使用它作为名称.
func getWallpaperName(link string, overrideName string) (string, error) {
    // 将壁纸链接解析为 *URL.
    u, err := url.Parse(link)

    if err != nil {
        return "", err
    }

    //从URL中提取GET参数.
    getParams, err := url.ParseQuery(u.RawQuery)
    idParam, ok := getParams["id"]

    if !ok {
        return "", fmt.Errorf("Could not find id GET parameter in link: %s. Cannot resolve wallpaper filename", link)
    }

    if len(idParam) != 1 {
        return "", fmt.Errorf("id GET parameter is not valid in link: %s. Cannot resolve wallpaper filename", link)
    }

    filename := idParam[0]

    if overrideName != "" {
        ext := filepath.Ext(filename)
        filename = fmt.Sprintf("%s%s", overrideName, ext)
    }

    return filename, nil
}

//保存壁纸从链接到目的地.
//返回最终目标路径.
func saveWallpaper(link string, dest string, filename string) (string, error) {
    outputDest := filepath.Join(dest, filename)

    resp, err := http.Get(link)

    if err != nil {
        return "", err
    }

    defer resp.Body.Close()

    f, err := os.Create(outputDest)

    if err != nil {
        return "", err
    }

    defer f.Close()

    //复制数据到文件.
    _, err = io.Copy(f, resp.Body)

    if err != nil {
        return "", err
    }

    return outputDest, nil
}

//检查图片是否存在;
func Exist(FileName string) bool {
    _, err := os.Stat(FileName)
    return err == nil || os.IsExist(err)
}

func main() {
    //图片保存路径: *dest
    //图片名称:filename
    //图片地址:link
    //dest := flag.String("output-dir", "", "Output directory to save wallpaper to")
    dest := flag.String("o", "", "Output directory to save wallpaper to")
    overrideFilename := flag.String("filename", "", "Name to give the wallpaper picture. Extension is automatically added.")
    flag.Parse()

    if *dest == "" {
        log.Fatal("You must provide an output directory using the -output-dir flag")
    }

    link, err := fetchWallpaperLink()

    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Found wallpaper link: %s\n", link)

    filename, err := getWallpaperName(link, *overrideFilename)

    if err != nil {
        log.Fatal(err)
    }

    //检查图片是否存在;
    if Exist(*dest+filename) == true {
        os.Exit(0)
    }

    finalDest, err := saveWallpaper(link, *dest, filename)          //下载图片;

    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Saved wallpaper to: %s\n", finalDest)
}
