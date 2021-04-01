package query

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const MainPage = "https://leetcode-cn.com"

type Variable struct {
	TitleSlug string
}

type RequestJSON struct {
	OperationName string
	Variables     []Variable
	Query         string
}

const query = `query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionId
    questionFrontendId
    boundTopicId
    title
    titleSlug
    content
    translatedTitle
    translatedContent
    isPaidOnly
    difficulty
    likes
    dislikes
    isLiked
    similarQuestions
    contributors {
      username
      profileUrl
      avatarUrl
      __typename
    }
    langToValidPlayground
    topicTags {
      name
      slug
      translatedName
      __typename
    }
    companyTagStats
    codeSnippets {
      lang
      langSlug
      code
      __typename
    }
    stats
    hints
    solution {
      id
      canSeeDetail
      __typename
    }
    status
    sampleTestCase
    metaData
    judgerAvailable
    judgeType
    mysqlSchemas
    enableRunCode
    envInfo
    book {
      id
      bookName
      pressName
      source
      shortDescription
      fullDescription
      bookImgUrl
      pressImgUrl
      productUrl
      __typename
    }
    isSubscribed
    isDailyQuestion
    dailyRecordStatus
    editorType
    ugcQuestionId
    __typename
  }
}`

func Request() {
	request, err := http.NewRequest(http.MethodPost, MainPage+"/graphql", QueryJSON())
	if err != nil {
		return
	}
	do, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	all, err := ioutil.ReadAll(do.Body)
	if err != nil {
		return
	}
	fmt.Println(string(all))
}

func QueryJSON() io.Reader {
	j := RequestJSON{
		OperationName: "getQuestionData",
		Variables: []Variable{
			{
				TitleSlug: "two-sum",
			},
		},
		Query: query,
	}
	marshal, _ := json.Marshal(j)
	return bytes.NewReader(marshal)
}
