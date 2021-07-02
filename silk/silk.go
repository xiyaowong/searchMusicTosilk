package silk

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"time"

	"searchMusicTosilk/asset"
	"searchMusicTosilk/util"
)

var encoder string

func init() {
	fmt.Println("====请确保已下载 FFmpeg 并已设置环境变量!====")
	// 检查依赖
	var err error
	switch runtime.GOOS {
	case "windows":
		encoder = "encoder.exe"
		if !util.FileExist(encoder) {
			err = ioutil.WriteFile(encoder, asset.EncoderExe, os.ModePerm)
		}
	case "linux":
		encoder = "encoder"
		if !util.FileExist(encoder) {
			err = ioutil.WriteFile(encoder, asset.Encoder, os.ModePerm)
		}
	default:
		log.Fatal("only support windows or linux - (x86)")
	}
	if err != nil {
		log.Fatal("导出依赖文件失败，请尝试重启!")
	}
	encoder = "./" + encoder
}

// File2SilkBase64 音频文件转成silk base64
func File2SilkBase64(filePath string) (b64 string, err error) {
	// p := fmt.Println
	// 1. 转pcm
	pcmPath := path.Join(path.Dir(filePath), "file.wav")
	defer os.Remove(pcmPath)

	cmd := exec.Command("ffmpeg", "-i", filePath, "-f", "s16le", "-ar", "24000", "-ac", "1", "-acodec", "pcm_s16le", pcmPath)
	cmd.Run()

	// 2. 转silk
	silkPath := path.Join(path.Dir(pcmPath), "file.silk")
	defer os.Remove(silkPath)

	cmd = exec.Command(encoder, pcmPath, silkPath, "-quiet", "-tencent")
	cmd.Run()
	// 3. 转base64
	content, err := ioutil.ReadFile(silkPath)
	if err != nil {
		log.Println("打开文件[silk]失败")
		return
	}

	b64 = base64.StdEncoding.EncodeToString(content)

	return
}

// URL2SilkBase64 音频下载链接转base64
func URL2SilkBase64(l string) (b64 string, err error) {
	tempDir, _ := ioutil.TempDir("", "silk")
	filePath := path.Join(tempDir, "file.wav")
	defer os.RemoveAll(tempDir)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, _ := http.NewRequest("GET", l, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("网络请求错误, err: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取网络数据失败, err: %v", err)
		return
	}
	err = ioutil.WriteFile(filePath, body, os.ModePerm)
	if err != nil {
		log.Printf("写入文件失败, err: %v", err)
		return
	}
	return File2SilkBase64(filePath)
}
