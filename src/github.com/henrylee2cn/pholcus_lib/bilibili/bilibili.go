package pholcus_lib

// 基础包
import (
	"log"

	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	. "github.com/henrylee2cn/pholcus/app/spider"           //必需
	"github.com/henrylee2cn/pholcus/common/goquery"         //DOM解析
	"github.com/henrylee2cn/pholcus/logs"                   //信息输出
	// . "github.com/henrylee2cn/pholcus/app/spider/common" //选用
	"github.com/astaxie/beego/httplib"
	// net包
	// "net/http" //设置http.Header
	// "net/url"

	// 编码包
	// "encoding/xml"
	// "encoding/json"

	// 字符串处理包
	//"regexp"
	"strconv"
	"strings"
	// 其他包
	//	"fmt"
	// "math"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"time"
)

func init() {
	BILIBILIVideo.Register()
}

var (
	ents               string
	bets               string
	page               int
	type_BilibiliVideo map[string]string
)

const (
	timeOutPoint = 10 * time.Second
	getPage      = 2
)

func arrType(pStr, bStr, eStr string) {
	type_BilibiliVideo = make(map[string]string)
	//type_BilibiliVideo["bangumi"] = "http://www.bilibili.com/list/hot-33-" + pStr + "-" + bStr + "~" + eStr + ".html" //影视                                // 番剧
	type_BilibiliVideo["douga"] = "https://www.bilibili.com/list/rank-25.html" //动画 （MAD）
	//	type_BilibiliVideo["music"] = "http://www.bilibili.com/list/hot-31-" + pStr + "-" + bStr + "~" + eStr + ".html"   //音乐 (翻唱)
	//	type_BilibiliVideo["kichiku"] = "http://www.bilibili.com/list/hot-22-" + pStr + "-" + bStr + "~" + eStr + ".html" //鬼畜 （鬼畜调教）
	//	type_BilibiliVideo["kichiku"] = "http://www.bilibili.com/list/hot-26-" + pStr + "-" + bStr + "~" + eStr + ".html" //鬼畜 （音MAD）
	//	type_BilibiliVideo["dance"] = "http://www.bilibili.com/list/hot-20-" + pStr + "-" + bStr + "~" + eStr + ".html"   //舞蹈 （宅舞）
	//	type_BilibiliVideo["game"] = "http://www.bilibili.com/list/hot-65-" + pStr + "-" + bStr + "~" + eStr + ".html"    //游戏 （网络竞技）
	//type_BilibiliVideo["technology"] = "http://www.bilibili.com/list/hot-37-" + pStr + "-" + bStr + "~" + eStr + ".html" //科技
	//1 type_BilibiliVideo["life"] = "http://www.bilibili.com/list/hot-138-" + pStr + "-" + bStr + "~" + eStr + ".html" //生活
	//type_BilibiliVideo["fashion"] = "http://www.bilibili.com/list/hot-157-" + pStr + "-" + bStr + "~" + eStr + ".html"   //时尚
	//type_BilibiliVideo["ent"] = "http://www.bilibili.com/list/hot-71-" + pStr + "-" + bStr + "~" + eStr + ".html"        //娱乐
}
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

var BILIBILIVideo = &Spider{
	Name:        "bilibili",
	Description: "哔哩哔哩视频 [Auto Page] [http://www.bilibili.com/]",
	// Pausetime: 300,
	// Keyword:     KEYWORD,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			timeStamp := time.Now().Unix()
			ents = time.Unix(timeStamp, 0).Format("2006-01-02")
			bets = time.Unix(timeStamp-86400, 0).Format("2006-01-02")
			page = 1
			log.Println("bets:", bets, "ents", ents)
			arrType(strconv.Itoa(page), bets, ents)
			for k := range type_BilibiliVideo {
				//ctx.SetTimer(k, time.Second*30, nil)
				ctx.Aid(map[string]interface{}{"loop": k, "p": page}, "LOOP")
			}
		},
		Trunk: map[string]*Rule{
			"LOOP": {
				AidFunc: func(ctx *Context, aid map[string]interface{}) interface{} {
					page = aid["p"].(int)
					arrType(strconv.Itoa(page), bets, ents)
					k := aid["loop"].(string)
					v := type_BilibiliVideo[k]
					log.Println("96 Line:", v)
					ctx.AddQueue(&request.Request{
						Url:        v,
						Rule:       "requestList",
						Temp:       map[string]interface{}{"qtype": k, "p": page},
						Reloadable: true,
					})
					return nil
				},
			},
			"requestList": {
				ParseFunc: func(ctx *Context) {
					var queryType = ctx.GetTemp("qtype", string("")).(string)
					var curr = ctx.GetTemp("p", int(0)).(int)
					if curr > getPage {
						return
					}
					log.Println("113 Line ---", curr)
					defer func() {
						//ctx.RunTimer(queryType)
						ctx.Aid(map[string]interface{}{"loop": queryType, "p": curr + 1}, "LOOP")
					}()
					query := ctx.GetDom()
					sh, _ := query.Html()
					log.Println("query----", sh)
					query.Find(".container-body").Each(func(i int, s *goquery.Selection) {
						sg, _ := s.Html()
						log.Println("------121------", sg)
						log.Println("-----------------")
						log.Println("-----------------", s.Text())
						vidStr, _ := s.Find("a.title").Attr("href")
						log.Println("vidStr--,", vidStr)
						vidStr = strings.Replace(vidStr, `/`, ``, -1)
						vidStr = strings.Replace(vidStr, `video`, ``, -1)
						log.Println("vidStr--,", vidStr)
						titleStr := s.Find("a.title").Text()
						log.Println("titleStr--", titleStr)
						timeStr := s.Find("span.v-date").Text()
						djStr, _ := s.Find(".l-item .v-info .gk span").Attr("number")
						dmStr, _ := s.Find(".l-item .v-info .dm span").Attr("number")
						scStr, _ := s.Find(".l-item .v-info .sc span").Attr("number")
						log.Println("djStr:", djStr, " dmStr:", dmStr, " scStr:", scStr)
						ctx.AddQueue(&request.Request{
							Url:  "http://www.bilibili.com/mobile/video/" + vidStr + ".html",
							Rule: "video",
							Temp: map[string]interface{}{
								"vid":      vidStr,
								"qtype":    queryType,
								"utitle":   titleStr,
								"utime":    timeStr,
								"dianji":   djStr,
								"danmu":    dmStr,
								"shoucang": scStr},
							ConnTimeout: -1,
						})
					})
				},
			},
			"video": {
				//注意：有无字段语义和是否输出数据必须保持一致
				ItemFields: []string{
					"title",
					"time",
					"video_url",
					"state",
					"SID",
					"img_url",
					"postid",
					"videotype",
					"dianjicount",
					"dmcount",
					"stowcount",
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					var video_url, SID, img_url, videotype string
					var state, postid, djnum, dmnum, stownum int
					var title = ctx.GetTemp("utitle", string("")).(string)
					var time1 = ctx.GetTemp("utime", string("")).(string)
					var dianjicount = ctx.GetTemp("dianji", string("")).(string)
					var dmcount = ctx.GetTemp("danmu", string("")).(string)
					var stowcount = ctx.GetTemp("shoucang", string("")).(string)
					query.Find(".wrapper").Each(func(i int, s *goquery.Selection) {
						type backU struct {
							dd int `json:"0"`
							ee int `json:"1"`
						}
						type bilibiliDurl struct {
							Order      int    `json:"order"`
							Length     int    `json:"length"`
							Size       int    `json:"size"`
							Url        string `json:"url"`
							Backup_url backU  `json:"backup_url"`
						}

						type quality struct {
							aa int `json:"0"`
							bb int `json:"1"`
						}
						type Durls struct {
							cc bilibiliDurl `json:"0"`
						}
						type outResult struct {
							From           string  `json:"from"`
							Result         string  `json:"result"`
							Format         string  `json:"format"`
							Timelength     int64   `json:"timelength"`
							Accept_format  string  `json:"accept_format"`
							Accept_quality quality `json:"accept_quality"`
							Seek_param     string  `json:"seek_param"`
							Seek_type      string  `json:"seek_type"`
							Durl           Durls   `json:"durl"`
							Img            string  `json:"img"`
							Cid            string  `json:"cid"`
							Fromview       string  `json:"fromview"`
						}
						var (
							vidStr = ctx.GetTemp("vid", string("")).(string)
							qType  = ctx.GetTemp("qtype", string("")).(string)
							oM     outResult
						)
						djnum, _ = strconv.Atoi(dianjicount)
						dmnum, _ = strconv.Atoi(dmcount)
						stownum, _ = strconv.Atoi(stowcount)

						tysb := time.Now().Unix()
						timeStr := strconv.FormatInt(tysb, 10)
						strToken := GetMd5String("bilibili_" + timeStr)
						vidStr = strings.Replace(vidStr, `av`, ``, -1)
						req := httplib.Get("http://api.bilibili.com/playurl?aid=" + vidStr + "&page=1&platform=html5&quality=1&vtype=mp4&type=jsonp&token=" + strToken + "&_=" + timeStr)
						req.Header("Accept-Encoding", "gzip, deflate")
						req.Header("Host", "api.bilibili.com")
						req.Header("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:49.0) Gecko/20100101 Firefox/49.0")
						req.Header("Cookie", "purl_token=bilibili_"+timeStr)
						//req.SetTimeout(timeOutPoint, timeOutPoint)
						str, err := req.String()
						if err != nil {
							logs.Log.Error("bilibili http connection error:", err)
							return
						}
						err = json.Unmarshal([]byte(str), &oM)
						if err != nil {
							logs.Log.Error("bilibili json data unmarshal error:", err)
							return
						}
						img_url = oM.Img
						video_url = vidStr
						state = 0
						postid = 0
						SID = "bilibili"
						videotype = qType
					})
					// 结果输出
					ctx.Output(map[int]interface{}{
						0:  title,
						1:  time1,
						2:  video_url,
						3:  state,
						4:  SID,
						5:  img_url,
						6:  postid,
						7:  videotype,
						8:  djnum,
						9:  dmnum,
						10: stownum,
					})
				},
			},
		},
	},
}
