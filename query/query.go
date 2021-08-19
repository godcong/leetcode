package query

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Query struct {
	cookie string
	//token string
}

func NewQuery(cookie string) Query {
	return Query{
		cookie: cookie,
		//token:  token,
	}
}

func (q Query) questionData(code *Code, name string) error {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl 'https://leetcode-cn.com/graphql/' \
	//   -H 'authority: leetcode-cn.com' \
	//   -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="90", "Google Chrome";v="90"' \
	//   -H 'x-timezone: Asia/Shanghai' \
	//   -H 'x-operation-name: questionData' \
	//   -H 'accept-language: zh-CN' \
	//   -H 'sec-ch-ua-mobile: ?0' \
	//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36' \
	//   -H 'content-type: application/json' \
	//   -H 'accept: */*' \
	//   -H 'x-csrftoken: Jknx9LMqHv5VCPwDVDUXnBpH0q77RSR3ItVyzffyBmkJHIK8pDsiZ7OFbPZo9Edg' \
	//   -H 'x-definition-name: question' \
	//   -H 'origin: https://leetcode-cn.com' \
	//   -H 'sec-fetch-site: same-origin' \
	//   -H 'sec-fetch-mode: cors' \
	//   -H 'sec-fetch-dest: empty' \
	//   -H 'referer: https://leetcode-cn.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day/' \
	//   -H 'cookie: _uab_collina=158494287144877725711875; __auc=197e35ef1798ce16c8387fd8db5; gr_user_id=43d5f89f-a476-4f4e-a2cf-324c63dc0b5d; a2873925c34ecbd2_gr_last_sent_cs1=cong-shen-bu-shi-shen; _ga=GA1.2.560771403.1621566256; Hm_lvt_fa218a3ff7179639febdb15e372f411c=1621566254,1622117541; csrftoken=Jknx9LMqHv5VCPwDVDUXnBpH0q77RSR3ItVyzffyBmkJHIK8pDsiZ7OFbPZo9Edg; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjQ2MzI3IiwiX2F1dGhfdXNlcl9iYWNrZW5kIjoiYWxsYXV0aC5hY2NvdW50LmF1dGhfYmFja2VuZHMuQXV0aGVudGljYXRpb25CYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiYjFkNTUyMTQ5NmViNTFhYTdiOTdiMTlmMDhhNTNiN2E0ZDU0YzBlZTFhMGUzNTk0YTZhNjFkY2Q4YWMzYzkxYyIsImlkIjoyNDYzMjcsImVtYWlsIjoianVtYnljY0AxNjMuY29tIiwidXNlcm5hbWUiOiJjb25nLXNoZW4tYnUtc2hpLXNoZW4iLCJ1c2VyX3NsdWciOiJjb25nLXNoZW4tYnUtc2hpLXNoZW4iLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS1jbi5jb20vYWxpeXVuLWxjLXVwbG9hZC91c2Vycy9jb25nLXNoZW4tYnUtc2hpLXNoZW4vYXZhdGFyXzE1OTY3ODg3OTEucG5nIiwicGhvbmVfdmVyaWZpZWQiOnRydWUsIl90aW1lc3RhbXAiOjE2MjIxMTc1NDYuNTcyMjU1LCJleHBpcmVkX3RpbWVfIjoxNjIyNjYwNDAwfQ.zFmv2ZoXn0-xa2GwlQquFByUvaAcSPsXbfhlJG4Q0MQ; __asc=c03719f4179c7aa1b58e18f261f; a2873925c34ecbd2_gr_session_id=43e5189c-43bc-4706-9de7-d467d9e8fbd3; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=43e5189c-43bc-4706-9de7-d467d9e8fbd3; a2873925c34ecbd2_gr_session_id_43e5189c-43bc-4706-9de7-d467d9e8fbd3=true; _gid=GA1.2.1864074649.1622552486; _gat_gtag_UA_131851415_1=1; Hm_lpvt_fa218a3ff7179639febdb15e372f411c=1622552495; a2873925c34ecbd2_gr_cs1=cong-shen-bu-shi-shen' \
	//   --data-raw $'{"operationName":"questionData","variables":{"titleSlug":"can-you-eat-your-favorite-candy-on-your-favorite-day"},"query":"query questionData($titleSlug: String\u0021) {\\n  question(titleSlug: $titleSlug) {\\n    questionId\\n    questionFrontendId\\n    categoryTitle\\n    boundTopicId\\n    title\\n    titleSlug\\n    content\\n    translatedTitle\\n    translatedContent\\n    isPaidOnly\\n    difficulty\\n    likes\\n    dislikes\\n    isLiked\\n    similarQuestions\\n    contributors {\\n      username\\n      profileUrl\\n      avatarUrl\\n      __typename\\n    }\\n    langToValidPlayground\\n    topicTags {\\n      name\\n      slug\\n      translatedName\\n      __typename\\n    }\\n    companyTagStats\\n    codeSnippets {\\n      lang\\n      langSlug\\n      code\\n      __typename\\n    }\\n    stats\\n    hints\\n    solution {\\n      id\\n      canSeeDetail\\n      __typename\\n    }\\n    status\\n    sampleTestCase\\n    metaData\\n    judgerAvailable\\n    judgeType\\n    mysqlSchemas\\n    enableRunCode\\n    envInfo\\n    book {\\n      id\\n      bookName\\n      pressName\\n      source\\n      shortDescription\\n      fullDescription\\n      bookImgUrl\\n      pressImgUrl\\n      productUrl\\n      __typename\\n    }\\n    isSubscribed\\n    isDailyQuestion\\n    dailyRecordStatus\\n    editorType\\n    ugcQuestionId\\n    style\\n    exampleTestcases\\n    __typename\\n  }\\n}\\n"}' \
	//   --compressed
	if DEBUG {
		fmt.Println("code name", name)
	}

	current := name
	if current == "" {
		for i, record := range code.Data.DailyQuestionRecords {
			if record.Date == code.Data.TodayRecord[0].Date {
				current = code.Data.DailyQuestionRecords[i].Question.QuestionTitleSlug
				break
			}
		}

		if current == "" {
			current = code.Data.TodayRecord[0].Question.QuestionTitleSlug
		}
	}

	data := Payload{
		OperationName: "questionData",
		Variables: Variables{
			TitleSlug: current,
		},
		Query: `query questionData($titleSlug: String) {
  question(titleSlug: $titleSlug) {
    questionId
    questionFrontendId
    categoryTitle
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
    style
    exampleTestcases
    __typename
  }
}`,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://leetcode-cn.com/graphql/", body)
	if err != nil {
		return err
	}
	req.Header.Set("Authority", "leetcode-cn.com")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Set("X-Timezone", "Asia/Shanghai")
	req.Header.Set("X-Operation-Number", "questionData")
	req.Header.Set("Accept-Language", "zh-CN")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.20 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")

	req.Header.Set("X-Definition-Number", "question")
	req.Header.Set("Origin", "https://leetcode-cn.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://leetcode-cn.com/problems/contiguous-array/")
	req.Header.Set("X-Csrftoken", "Q0GQIqMVCZipkRvFVh6YMZG00NIY1DwX4BzekgruSJjHceUAzSYZDqufv5dXbAx2")
	req.Header.Set("Cookie", q.cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(all, &code)
	if err != nil {
		return err
	}
	//fmt.Println("every day code", string(all))
	return err
}

func (q Query) problemSetQuestionList(skip, limit int) (Code, error) {
	var code Code
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl 'https://leetcode-cn.com/graphql/' \
	//   -H 'authority: leetcode-cn.com' \
	//   -H 'sec-ch-ua: "Chromium";v="94", "Google Chrome";v="94", ";Not A Brand";v="99"' \
	//   -H 'authorization: ' \
	//   -H 'x-csrftoken: IXVisBq53IUHsBo5OBeCPLl1DFL5zk2sEPb8i6MgZeI7UEXYbFtkXpcsshyLdGbp' \
	//   -H 'sec-ch-ua-mobile: ?0' \
	//   -H 'content-type: application/json' \
	//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4590.0 Safari/537.36' \
	//   -H 'sec-ch-ua-platform: "Windows"' \
	//   -H 'accept: */*' \
	//   -H 'origin: https://leetcode-cn.com' \
	//   -H 'sec-fetch-site: same-origin' \
	//   -H 'sec-fetch-mode: cors' \
	//   -H 'sec-fetch-dest: empty' \
	//   -H 'referer: https://leetcode-cn.com/problemset/all/?page=2' \
	//   -H 'accept-language: zh-CN,zh;q=0.9' \
	//   -H 'cookie: __auc=be3552e4179647261bb747fe3d6; gr_user_id=af2ef593-93e0-4d57-987f-c3793b8a60fd; _ga=GA1.2.2006997331.1620887891; a2873925c34ecbd2_gr_last_sent_cs1=cong-shen-bu-shi-shen; Hm_lvt_fa218a3ff7179639febdb15e372f411c=1628171536; __asc=9d75d71f17b169609ec3ea1cf99; a2873925c34ecbd2_gr_session_id=54602608-2025-4194-9cb4-818a1c097c11; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=54602608-2025-4194-9cb4-818a1c097c11; a2873925c34ecbd2_gr_session_id_54602608-2025-4194-9cb4-818a1c097c11=true; _gid=GA1.2.470683429.1628171538; csrftoken=IXVisBq53IUHsBo5OBeCPLl1DFL5zk2sEPb8i6MgZeI7UEXYbFtkXpcsshyLdGbp; NEW_PROBLEMLIST_PAGE=1; Hm_lpvt_fa218a3ff7179639febdb15e372f411c=1628171675; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjQ2MzI3IiwiX2F1dGhfdXNlcl9iYWNrZW5kIjoiYWxsYXV0aC5hY2NvdW50LmF1dGhfYmFja2VuZHMuQXV0aGVudGljYXRpb25CYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiYjFkNTUyMTQ5NmViNTFhYTdiOTdiMTlmMDhhNTNiN2E0ZDU0YzBlZTFhMGUzNTk0YTZhNjFkY2Q4YWMzYzkxYyIsImlkIjoyNDYzMjcsImVtYWlsIjoianVtYnljY0AxNjMuY29tIiwidXNlcm5hbWUiOiJjb25nLXNoZW4tYnUtc2hpLXNoZW4iLCJ1c2VyX3NsdWciOiJjb25nLXNoZW4tYnUtc2hpLXNoZW4iLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS1jbi5jb20vYWxpeXVuLWxjLXVwbG9hZC91c2Vycy9jb25nLXNoZW4tYnUtc2hpLXNoZW4vYXZhdGFyXzE1OTY3ODg3OTEucG5nIiwicGhvbmVfdmVyaWZpZWQiOnRydWUsIl90aW1lc3RhbXAiOjE2MjgxNzE1NDUuNjM5MTAyMiwiZXhwaXJlZF90aW1lXyI6MTYzMDY5NTYwMCwibGF0ZXN0X3RpbWVzdGFtcF8iOjE2MjgxNzE5OTV9.q5LGYOfyQoXS667fQUPhKp1EmMczuTlv9_3mlspfM7c; a2873925c34ecbd2_gr_cs1=cong-shen-bu-shi-shen' \
	//   --data-raw '{"query":"\n    query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {\n  problemsetQuestionList(\n    categorySlug: $categorySlug\n    limit: $limit\n    skip: $skip\n    filters: $filters\n  ) {\n    hasMore\n    total\n    questions {\n      acRate\n      difficulty\n      freqBar\n      frontendQuestionId\n      isFavor\n      paidOnly\n      solutionNum\n      status\n      title\n      titleCn\n      titleSlug\n      topicTags {\n        name\n        nameTranslated\n        id\n        slug\n      }\n      extra {\n        hasVideoSolution\n        topCompanyTags {\n          imgUrl\n          slug\n          numSubscribed\n        }\n      }\n    }\n  }\n}\n    ","variables":{"categorySlug":"","skip":50,"limit":50,"filters":{}},"operationName":"problemsetQuestionList"}' \
	//   --compressed

	data := Payload{
		OperationName: "problemsetQuestionList",
		Variables: Variables{
			CategorySlug: "",
			TitleSlug:    "",
			Year:         0,
			Month:        0,
			Skip:         skip,
			Limit:        limit,
			Filters:      Filters{},
		},
		Query: `query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {
  problemsetQuestionList(
    categorySlug: $categorySlug
    limit: $limit
    skip: $skip
    filters: $filters
  ) {
    hasMore
    total
    questions {
      acRate
      difficulty
      freqBar
      frontendQuestionId
      isFavor
      paidOnly
      solutionNum
      status
      title
      titleCn
      titleSlug
      topicTags {
        name
        nameTranslated
        id
        slug
      }
      extra {
        hasVideoSolution
        topCompanyTags {
          imgUrl
          slug
          numSubscribed
        }
      }
    }
  }
}`,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return code, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://leetcode-cn.com/graphql/", body)
	if err != nil {
		return code, err
	}
	req.Header.Set("Authority", "leetcode-cn.com")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"94\", \"Google Chrome\";v=\"94\", \";Not A Brand\";v=\"99\"")
	req.Header.Set("Authorization", "")
	req.Header.Set("X-Csrftoken", "IXVisBq53IUHsBo5OBeCPLl1DFL5zk2sEPb8i6MgZeI7UEXYbFtkXpcsshyLdGbp")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4590.0 Safari/537.36")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://leetcode-cn.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://leetcode-cn.com/problemset/all/?page=2")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", q.cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return code, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return code, err
	}
	//fmt.Println("every day", string(all))

	err = json.Unmarshal(all, &code)
	return code, err
}

func (q Query) dailyQuestionRecords(now time.Time) (Code, error) {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl 'https://leetcode-cn.com/graphql' \
	//   -H 'authority: leetcode-cn.com' \
	//   -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="90", "Google Chrome";v="90"' \
	//   -H 'accept: */*' \
	//   -H 'x-csrftoken: Jknx9LMqHv5VCPwDVDUXnBpH0q77RSR3ItVyzffyBmkJHIK8pDsiZ7OFbPZo9Edg' \
	//   -H 'sec-ch-ua-mobile: ?0' \
	//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36' \
	//   -H 'content-type: application/json' \
	//   -H 'origin: https://leetcode-cn.com' \
	//   -H 'sec-fetch-site: same-origin' \
	//   -H 'sec-fetch-mode: cors' \
	//   -H 'sec-fetch-dest: empty' \
	//   -H 'referer: https://leetcode-cn.com/problemset/all/' \
	//   -H 'accept-language: zh-CN,zh;q=0.9' \
	//   -H 'cookie: _uab_collina=158494287144877725711875; __auc=197e35ef1798ce16c8387fd8db5; gr_user_id=43d5f89f-a476-4f4e-a2cf-324c63dc0b5d; a2873925c34ecbd2_gr_last_sent_cs1=cong-shen-bu-shi-shen; _ga=GA1.2.560771403.1621566256; Hm_lvt_fa218a3ff7179639febdb15e372f411c=1621566254,1622117541; csrftoken=Jknx9LMqHv5VCPwDVDUXnBpH0q77RSR3ItVyzffyBmkJHIK8pDsiZ7OFbPZo9Edg; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjQ2MzI3IiwiX2F1dGhfdXNlcl9iYWNrZW5kIjoiYWxsYXV0aC5hY2NvdW50LmF1dGhfYmFja2VuZHMuQXV0aGVudGljYXRpb25CYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiYjFkNTUyMTQ5NmViNTFhYTdiOTdiMTlmMDhhNTNiN2E0ZDU0YzBlZTFhMGUzNTk0YTZhNjFkY2Q4YWMzYzkxYyIsImlkIjoyNDYzMjcsImVtYWlsIjoianVtYnljY0AxNjMuY29tIiwidXNlcm5hbWUiOiJjb25nLXNoZW4tYnUtc2hpLXNoZW4iLCJ1c2VyX3NsdWciOiJjb25nLXNoZW4tYnUtc2hpLXNoZW4iLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS1jbi5jb20vYWxpeXVuLWxjLXVwbG9hZC91c2Vycy9jb25nLXNoZW4tYnUtc2hpLXNoZW4vYXZhdGFyXzE1OTY3ODg3OTEucG5nIiwicGhvbmVfdmVyaWZpZWQiOnRydWUsIl90aW1lc3RhbXAiOjE2MjIxMTc1NDYuNTcyMjU1LCJleHBpcmVkX3RpbWVfIjoxNjIyNjYwNDAwfQ.zFmv2ZoXn0-xa2GwlQquFByUvaAcSPsXbfhlJG4Q0MQ; _gid=GA1.2.1864074649.1622552486; __asc=2e464209179caaa0c3d8afb1ac1; a2873925c34ecbd2_gr_session_id=c42d1559-0c6f-4dda-a0cf-20e057f020bf; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=c42d1559-0c6f-4dda-a0cf-20e057f020bf; a2873925c34ecbd2_gr_session_id_c42d1559-0c6f-4dda-a0cf-20e057f020bf=true; _gat_gtag_UA_131851415_1=1; a2873925c34ecbd2_gr_cs1=cong-shen-bu-shi-shen; Hm_lpvt_fa218a3ff7179639febdb15e372f411c=1622603596' \
	//   --data-raw $'{"operationName":"dailyQuestionRecords","variables":{"year":2021,"month":6},"query":"query dailyQuestionRecords($year: Int\u0021, $month: Int\u0021) {\\n  dailyQuestionRecords(year: $year, month: $month) {\\n    date\\n    question {\\n      questionId\\n      questionFrontendId\\n      questionTitle\\n      questionTitleSlug\\n      translatedTitle\\n      __typename\\n    }\\n    userStatus\\n    __typename\\n  }\\n}\\n"}' \
	//   --compressed

	data := Payload{
		OperationName: "dailyQuestionRecords",
		Variables: Variables{
			Year:  now.Year(),
			Month: int(now.Month()),
		},
		Query: `query dailyQuestionRecords($year: Int!, $month: Int!) {
  dailyQuestionRecords(year: $year, month: $month) {
    date
    question {
      questionId
      questionFrontendId
      questionTitle
      questionTitleSlug
      translatedTitle
      __typename
    }
    userStatus
    __typename
  }
}`,
	}
	var code Code
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return code, err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://leetcode-cn.com/graphql", body)
	if err != nil {
		return code, err
	}
	req.Header.Set("Authority", "leetcode-cn.com")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.20 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://leetcode-cn.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://leetcode-cn.com/problemset/all/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("X-Csrftoken", "Q0GQIqMVCZipkRvFVh6YMZG00NIY1DwX4BzekgruSJjHceUAzSYZDqufv5dXbAx2")
	req.Header.Set("Cookie", q.cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return code, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return code, err
	}
	//fmt.Println("every day", string(all))

	err = json.Unmarshal(all, &code)
	return code, err
}

func (q Query) questionOfToday(code *Code) error {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl 'https://leetcode-cn.com/graphql' \
	//   -H 'authority: leetcode-cn.com' \
	//   -H 'sec-ch-ua: "Chromium";v="92", " Not A;Brand";v="99", "Google Chrome";v="92"' \
	//   -H 'accept: */*' \
	//   -H 'x-csrftoken: uEJ9oYMMLB83jahDtJ3FxQIWSA15VkBja5UOMccjH6NLCbwvAekvyFbnqqoTtY6G' \
	//   -H 'sec-ch-ua-mobile: ?0' \
	//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.20 Safari/537.36' \
	//   -H 'content-type: application/json' \
	//   -H 'origin: https://leetcode-cn.com' \
	//   -H 'sec-fetch-site: same-origin' \
	//   -H 'sec-fetch-mode: cors' \
	//   -H 'sec-fetch-dest: empty' \
	//   -H 'referer: https://leetcode-cn.com/problemset/all/' \
	//   -H 'accept-language: zh-CN,zh;q=0.9' \
	//   -H 'cookie: __auc=be3552e4179647261bb747fe3d6; gr_user_id=af2ef593-93e0-4d57-987f-c3793b8a60fd; _ga=GA1.2.2006997331.1620887891; a2873925c34ecbd2_gr_last_sent_cs1=cong-shen-bu-shi-shen; Hm_lvt_fa218a3ff7179639febdb15e372f411c=1620887888,1621005075,1622700730; a2873925c34ecbd2_gr_session_id=9bc707e9-0f16-4412-a6a5-90fe24aec10c; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=9bc707e9-0f16-4412-a6a5-90fe24aec10c; a2873925c34ecbd2_gr_session_id_9bc707e9-0f16-4412-a6a5-90fe24aec10c=true; __asc=579659c8179d080277becfa346f; _gid=GA1.2.1394740010.1622700731; _gat_gtag_UA_131851415_1=1; csrftoken=uEJ9oYMMLB83jahDtJ3FxQIWSA15VkBja5UOMccjH6NLCbwvAekvyFbnqqoTtY6G; a2873925c34ecbd2_gr_cs1=cong-shen-bu-shi-shen; Hm_lpvt_fa218a3ff7179639febdb15e372f411c=1622700769' \
	//   --data-raw '{"operationName":"questionOfToday","variables":{},"query":"query questionOfToday {\n  todayRecord {\n    question {\n      questionFrontendId\n      questionTitleSlug\n      __typename\n    }\n    lastSubmission {\n      id\n      __typename\n    }\n    date\n    userStatus\n    __typename\n  }\n}\n"}' \
	//   --compressed

	//type Payload struct {
	//	Operationname string    `json:"operationName"`
	//	Variables     Variables `json:"variables"`
	//	Query         string    `json:"query"`
	//}

	data := Payload{
		OperationName: "questionOfToday",
		Variables:     Variables{},
		Query: `query questionOfToday {
  todayRecord {
    question {
      questionFrontendId
      questionTitleSlug
      __typename
    }
    lastSubmission {
      id
      __typename
    }
    date
    userStatus
    __typename
  }
}`,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://leetcode-cn.com/graphql", body)
	if err != nil {
		return err
	}
	req.Header.Set("Authority", "leetcode-cn.com")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"92\", \" Not A;Brand\";v=\"99\", \"Google Chrome\";v=\"92\"")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.20 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://leetcode-cn.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://leetcode-cn.com/problemset/all/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("X-Csrftoken", "Q0GQIqMVCZipkRvFVh6YMZG00NIY1DwX4BzekgruSJjHceUAzSYZDqufv5dXbAx2")
	req.Header.Set("Cookie", q.cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//fmt.Println("today records", string(all))
	err = json.Unmarshal(all, code)
	if err != nil {
		return err
	}

	return nil
}

func (q Query) getQuestionTranslation(code *Code) error {

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl 'https://leetcode-cn.com/graphql' \
	//   -H 'authority: leetcode-cn.com' \
	//   -H 'sec-ch-ua: " Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"' \
	//   -H 'accept: */*' \
	//   -H 'x-csrftoken: NkyRIGswJRNb0MgnVCRStJesTnIUHevUyWgihUrwVZIL4Mzmz9NNndWKbOc6TjQ1' \
	//   -H 'sec-ch-ua-mobile: ?0' \
	//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36' \
	//   -H 'content-type: application/json' \
	//   -H 'origin: https://leetcode-cn.com' \
	//   -H 'sec-fetch-site: same-origin' \
	//   -H 'sec-fetch-mode: cors' \
	//   -H 'sec-fetch-dest: empty' \
	//   -H 'referer: https://leetcode-cn.com/problemset/all/' \
	//   -H 'accept-language: zh-CN,zh;q=0.9' \
	//   -H 'cookie: _uab_collina=158494287144877725711875; __auc=197e35ef1798ce16c8387fd8db5; gr_user_id=43d5f89f-a476-4f4e-a2cf-324c63dc0b5d; a2873925c34ecbd2_gr_last_sent_cs1=cong-shen-bu-shi-shen; _ga=GA1.2.560771403.1621566256; Hm_lvt_fa218a3ff7179639febdb15e372f411c=1622117541,1622778812,1623307679,1623392791; csrftoken=NkyRIGswJRNb0MgnVCRStJesTnIUHevUyWgihUrwVZIL4Mzmz9NNndWKbOc6TjQ1; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjQ2MzI3IiwiX2F1dGhfdXNlcl9iYWNrZW5kIjoiYWxsYXV0aC5hY2NvdW50LmF1dGhfYmFja2VuZHMuQXV0aGVudGljYXRpb25CYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiYjFkNTUyMTQ5NmViNTFhYTdiOTdiMTlmMDhhNTNiN2E0ZDU0YzBlZTFhMGUzNTk0YTZhNjFkY2Q4YWMzYzkxYyIsImlkIjoyNDYzMjcsImVtYWlsIjoianVtYnljY0AxNjMuY29tIiwidXNlcm5hbWUiOiJjb25nLXNoZW4tYnUtc2hpLXNoZW4iLCJ1c2VyX3NsdWciOiJjb25nLXNoZW4tYnUtc2hpLXNoZW4iLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS1jbi5jb20vYWxpeXVuLWxjLXVwbG9hZC91c2Vycy9jb25nLXNoZW4tYnUtc2hpLXNoZW4vYXZhdGFyXzE1OTY3ODg3OTEucG5nIiwicGhvbmVfdmVyaWZpZWQiOnRydWUsIl90aW1lc3RhbXAiOjE2MjMzOTM0MTMuMDk4MjU4LCJleHBpcmVkX3RpbWVfIjoxNjIzOTU2NDAwfQ.89xly4S1XWugho3mqJAekpXX68UkLI8XPZHLzIGaCnc; Hm_lpvt_fa218a3ff7179639febdb15e372f411c=1623467310; __asc=05051f8117a0d9e42f787bd3144; a2873925c34ecbd2_gr_session_id=6b1a6526-daf7-4fcd-aae0-b71d65cb8e70; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=6b1a6526-daf7-4fcd-aae0-b71d65cb8e70; a2873925c34ecbd2_gr_session_id_6b1a6526-daf7-4fcd-aae0-b71d65cb8e70=true; _gid=GA1.2.999558998.1623726115; _gat_gtag_UA_131851415_1=1; a2873925c34ecbd2_gr_cs1=cong-shen-bu-shi-shen' \
	//   --data-raw '{"operationName":"getQuestionTranslation","variables":{},"query":"query getQuestionTranslation($lang: String) {\n  translations: allAppliedQuestionTranslations(lang: $lang) {\n    title\n    questionId\n    __typename\n  }\n}\n"}' \
	//   --compressed

	data := Payload{
		OperationName: "getQuestionTranslation",
		Variables:     Variables{},
		Query: `query getQuestionTranslation($lang: String) {
  translations: allAppliedQuestionTranslations(lang: $lang) {
    title
    questionId
    __typename
  }
}`,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://leetcode-cn.com/graphql", body)
	if err != nil {
		return err
	}
	req.Header.Set("Authority", "leetcode-cn.com")
	req.Header.Set("Sec-Ch-Ua", "\" Not;A Brand\";v=\"99\", \"Google Chrome\";v=\"91\", \"Chromium\";v=\"91\"")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("X-Csrftoken", "NkyRIGswJRNb0MgnVCRStJesTnIUHevUyWgihUrwVZIL4Mzmz9NNndWKbOc6TjQ1")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://leetcode-cn.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://leetcode-cn.com/problemset/all/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", q.cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return decodeCode(resp.Body, code)
}

func (q *Query) getNumberCode(codeNum int64) (*Code, error) {
	fmt.Println("Get Code", codeNum)
	skip := int(codeNum / 50 * 50)
	limit := 50

	records, err := q.problemSetQuestionList(skip, limit)
	if err != nil {
		return nil, err
	}

	for _, question := range records.Data.ProblemSetQuestionList.Questions {
		fmt.Printf("question list:%+v\n", question)
		i, err := strconv.ParseInt(question.FrontendQuestionID, 10, 32)
		if err != nil {
			continue
		}
		if codeNum == i {
			records.Result.Number = question.FrontendQuestionID
			records.Result.Slug = question.TitleSlug
			break
		}
	}

	return &records, nil
}

func (q *Query) getTodayCode() (*Code, error) {
	var code Code
	if err := q.questionOfToday(&code); err != nil {
		return nil, fmt.Errorf("questionOfToday:%v", err)
	}

	for i := range code.Data.TodayRecord {
		fmt.Printf("Today:%+v\n", code.Data.TodayRecord[i].Question)
		code.Result.Number = code.Data.TodayRecord[i].Question.QuestionFrontendID
		code.Result.Slug = code.Data.TodayRecord[i].Question.QuestionTitleSlug
	}
	return &code, nil
}

func (q *Query) getNameCode(codeName string) (*Code, error) {
	fmt.Println("Get Code", codeName)
	t := time.Now()

	records, err := q.dailyQuestionRecords(t)
	if err != nil {
		return nil, fmt.Errorf("dailyQuestionRecords:%v", err)
	}

	for _, record := range records.Data.DailyQuestionRecords {
		if DEBUG {
			fmt.Printf("Records:%+v\n", record.Question)
		}
		//fmt.Println("compare:", record.Question.QuestionTitleSlug == strings.TrimSpace(codeName))
		if record.Question.QuestionTitleSlug == strings.TrimSpace(codeName) {
			records.Result.Number = record.Question.QuestionFrontendID
			records.Result.Slug = record.Question.QuestionTitleSlug
			if DEBUG {
				fmt.Printf("result data:%+v\n", records.Result)
			}
			return &records, nil
		}
	}
	return nil, errors.New("failed to get name code")
}

func GetCode(cookie string, codeName string) (*Code, error) {
	var code *Code
	var err error
	q := NewQuery(cookie)
	if codeName == "" {
		code, err = q.getTodayCode()
		if err != nil {
			return nil, err
		}
	} else {

		parseInt, err := strconv.ParseInt(codeName, 10, 32)
		if err != nil {
			code, err = q.getNameCode(codeName)
			if err != nil {
				fmt.Println("GetCode error", err)
				return code, err
			}

		} else {
			code, err = q.getNumberCode(parseInt)
			if err != nil {
				fmt.Println("GetNumberCode error", err)
				return code, err
			}
		}
	}
	if code == nil {
		return nil, errors.New("empty code")
	}
	if err = q.questionData(code, code.Result.Slug); err != nil {
		return nil, fmt.Errorf("questionData:%v", err)
	}

	fmt.Println("Data:", code.Data.Question.Title)

	return code, nil
}
