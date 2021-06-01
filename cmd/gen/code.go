package main

type Code struct {
	Data struct {
		Question struct {
			Questionid            string        `json:"questionId"`
			Questionfrontendid    string        `json:"questionFrontendId"`
			Categorytitle         string        `json:"categoryTitle"`
			Boundtopicid          int           `json:"boundTopicId"`
			Title                 string        `json:"title"`
			Titleslug             string        `json:"titleSlug"`
			Content               string        `json:"content"`
			Translatedtitle       string        `json:"translatedTitle"`
			Translatedcontent     string        `json:"translatedContent"`
			Ispaidonly            bool          `json:"isPaidOnly"`
			Difficulty            string        `json:"difficulty"`
			Likes                 int           `json:"likes"`
			Dislikes              int           `json:"dislikes"`
			Isliked               interface{}   `json:"isLiked"`
			Similarquestions      string        `json:"similarQuestions"`
			Contributors          []interface{} `json:"contributors"`
			Langtovalidplayground string        `json:"langToValidPlayground"`
			Topictags             []struct {
				Name           string `json:"name"`
				Slug           string `json:"slug"`
				Translatedname string `json:"translatedName"`
				Typename       string `json:"__typename"`
			} `json:"topicTags"`
			Companytagstats interface{} `json:"companyTagStats"`
			Codesnippets    []struct {
				Lang     string `json:"lang"`
				Langslug string `json:"langSlug"`
				Code     string `json:"code"`
				Typename string `json:"__typename"`
			} `json:"codeSnippets"`
			Stats             string        `json:"stats"`
			Hints             []string      `json:"hints"`
			Solution          interface{}   `json:"solution"`
			Status            interface{}   `json:"status"`
			Sampletestcase    string        `json:"sampleTestCase"`
			Metadata          string        `json:"metaData"`
			Judgeravailable   bool          `json:"judgerAvailable"`
			Judgetype         string        `json:"judgeType"`
			Mysqlschemas      []interface{} `json:"mysqlSchemas"`
			Enableruncode     bool          `json:"enableRunCode"`
			Envinfo           string        `json:"envInfo"`
			Book              interface{}   `json:"book"`
			Issubscribed      bool          `json:"isSubscribed"`
			Isdailyquestion   bool          `json:"isDailyQuestion"`
			Dailyrecordstatus string        `json:"dailyRecordStatus"`
			Editortype        string        `json:"editorType"`
			Ugcquestionid     interface{}   `json:"ugcQuestionId"`
			Style             string        `json:"style"`
			Exampletestcases  string        `json:"exampleTestcases"`
			Typename          string        `json:"__typename"`
		} `json:"question"`
	} `json:"data"`
}
