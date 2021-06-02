package main

type Code struct {
	Data Data `json:"data"`
}

type Data struct {
	DailyQuestionRecords []DailyQuestionRecord `json:"dailyQuestionRecords,omitempty"`
	Question             DataQuestion          `json:"question,omitempty"`
}

type DailyQuestionRecord struct {
	Date       string                      `json:"date"`
	Question   DailyQuestionRecordQuestion `json:"question"`
	UserStatus string                      `json:"userStatus"`
	Typename   string                      `json:"__typename"`
}

type DailyQuestionRecordQuestion struct {
	QuestionID         string `json:"questionId"`
	QuestionFrontendID string `json:"questionFrontendId"`
	QuestionTitle      string `json:"questionTitle"`
	QuestionTitleSlug  string `json:"questionTitleSlug"`
	TranslatedTitle    string `json:"translatedTitle"`
	Typename           string `json:"__typename"`
}

type DataQuestion struct {
	QuestionID            string        `json:"questionId"`
	QuestionFrontendID    string        `json:"questionFrontendId"`
	CategoryTitle         string        `json:"categoryTitle"`
	BoundTopicID          int64         `json:"boundTopicId"`
	Title                 string        `json:"title"`
	TitleSlug             string        `json:"titleSlug"`
	Content               string        `json:"content"`
	TranslatedTitle       string        `json:"translatedTitle"`
	TranslatedContent     string        `json:"translatedContent"`
	IsPaidOnly            bool          `json:"isPaidOnly"`
	Difficulty            string        `json:"difficulty"`
	Likes                 int64         `json:"likes"`
	Dislikes              int64         `json:"dislikes"`
	IsLiked               interface{}   `json:"isLiked"`
	SimilarQuestions      string        `json:"similarQuestions"`
	Contributors          []interface{} `json:"contributors"`
	LangToValidPlayground string        `json:"langToValidPlayground"`
	TopicTags             []TopicTag    `json:"topicTags"`
	CompanyTagStats       interface{}   `json:"companyTagStats"`
	CodeSnippets          []CodeSnippet `json:"codeSnippets"`
	Stats                 string        `json:"stats"`
	Hints                 []interface{} `json:"hints"`
	Solution              interface{}   `json:"solution"`
	Status                string        `json:"status"`
	SampleTestCase        string        `json:"sampleTestCase"`
	MetaData              string        `json:"metaData"`
	JudgerAvailable       bool          `json:"judgerAvailable"`
	JudgeType             string        `json:"judgeType"`
	MysqlSchemas          []interface{} `json:"mysqlSchemas"`
	EnableRunCode         bool          `json:"enableRunCode"`
	EnvInfo               string        `json:"envInfo"`
	Book                  interface{}   `json:"book"`
	IsSubscribed          bool          `json:"isSubscribed"`
	IsDailyQuestion       bool          `json:"isDailyQuestion"`
	DailyRecordStatus     string        `json:"dailyRecordStatus"`
	EditorType            string        `json:"editorType"`
	UgcQuestionID         interface{}   `json:"ugcQuestionId"`
	Style                 string        `json:"style"`
	ExampleTestcases      string        `json:"exampleTestcases"`
	Typename              string        `json:"__typename"`
}

type CodeSnippet struct {
	Lang     string   `json:"lang"`
	LangSlug string   `json:"langSlug"`
	Code     string   `json:"code"`
	Typename Typename `json:"__typename"`
}

type TopicTag struct {
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	TranslatedName string `json:"translatedName"`
	Typename       string `json:"__typename"`
}

type Typename string

const (
	CodeSnippetNode Typename = "CodeSnippetNode"
)
