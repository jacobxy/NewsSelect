package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	TokenUrl     = "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s"
	CliendId     = "M5qShRNYdPw9pvF49Fjr3MG9"
	ClientSecret = "SVuTYUyxiTGawGiPcQmGCY9v0invT06Q"
)

var Token = "24.d80e3e3aa32976ea5ab2e600cfa047d8.2592000.1588301461.282335-18015470"

func GetTokenURL() string {
	return fmt.Sprintf(TokenUrl, CliendId, ClientSecret)
}

type TokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func GetAccessToken() string {
	if Token != "" {
		return Token
	}
	temp := TokenRes{}
	res, _ := http.Post(GetTokenURL(), "json", nil)
	b, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return "error"
	}
	Token = temp.AccessToken
	return temp.AccessToken
}

const (
	NewsSummaryUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/news_summary?access_token=%s&charset=UTF-8"
)

func GetNewsSummaryUrl() string {
	return fmt.Sprintf(NewsSummaryUrl, Token)
}

type NewsSummaryPostRequest struct {
	Title         string `json:"title"`
	Content       string `json:"content"`
	MaxSummaryLen int    `json:"max_summary_len"`
}

type NewsSummaryResponse struct {
	LogId   int64  `json:"log_id"`
	Summary string `json:"summary"`
}

func AskNewsSummary(title, content string) string {

	content = strings.Replace(content, " ", "", -1)	

	temp := NewsSummaryPostRequest{
		Title:          title,
		Content:        content,
		MaxSummaryLen: 200,
	}

	b, _ := json.Marshal(temp)

	client := http.Client{}

	body := bytes.NewBuffer(b)
	req, err := http.NewRequest("POST", GetNewsSummaryUrl(), body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	resTemp := NewsSummaryResponse{}

	err = json.Unmarshal(r, &resTemp)
	if err != nil {
		panic(err)
	}

	fmt.Println(resTemp)
	fmt.Println(string(r))

	return resTemp.Summary
}

func main() {
	title := "第二语言习得的领域关系"
	content := `有许多领域与第二语言习得有关，有些在第一章中已经提过。本章将简要提及几个“相邻”学科并向读者介绍这些领域，指出它们相同和不同之处。虽然 二语习得已是一个自足的研究领域，但它的根扎在其他领域，如语言教学领域, 也是在其他领域最早证明自身存在的价值，并且一直受到诸如语言学和心理学 等其他领域的影响。但是，它与儿童语言习得(child language acquisition)有一 种特殊的关系，原因在于儿童语言习得构成了第二语言习得研究的基础，第二语 言习得研究的许多原始问题都源于儿童语言习得研究中的相同问题。还有一些 领域，如第三语言习得(third language acquisition )或继承语习得(heritage language acquisition),是第二语言习得的特例，近年来也得到了发展，特别是继 承语学习。最后，双语习得(bilingual acquisition)融合了有关第二语言习得和第 一语言习得的问题。
作为本章的开始，我们首先简单浏览一下这些相关领域研究的一些问题。 我们只能粗略地介绍，因为如果不这样，我们就会偏离本书的中心——第二语言 习得。但是，我们觉得介绍这些相关领域很重要，因为它们能够帮助我们了解二 语习得的复杂性。这些学科都有各自完备的发展史，大多数还有专门研究该学 科问题的期刊。在本章中，我们在总结这些领域的研究范围之余,还要稍微多做 一点介绍。
每一个领域都与第二语言习得构成了不同的关系。第三语言习得和继承语习得这两个领域与第二语言习得是派生关系，从相关但更具体的问题发展而来。 双语研究与那些在一定程度上同第二语言习得相异的问题平行发展，它关注的 是同时开始学习两种语言之类的问题。对习得种类进行划分，正如我们在本章 中所做的那样，有些不自然，但是出于讲解的目的还是必要的。下面我们分领域
单独介绍。`
	r := (AskNewsSummary(title, content))
	fmt.Println(r)
}

