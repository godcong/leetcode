package types

// DEBUG flag for development
var DEBUG = false

// Code represents the complete problem data structure
type Code struct {
	Data   Data   `json:"data"`
	Result Result `json:"-"`
}

// UserStatus represents the completion status
type UserStatus string

const (
	Finish   UserStatus = "FINISH"
	NotStart UserStatus = "NOT_START"
)

// Result contains problem metadata
type Result struct {
	Number string `json:"-"`
	Slug   string `json:"-"`
}

// Data is the main data container
type Data struct {
	ProblemSetQuestionList ProblemSetQuestionList `json:"problemsetQuestionList"`
	DailyQuestionRecords   []DailyQuestionRecord  `json:"dailyQuestionRecords,omitempty"`
	Question               DataQuestion           `json:"question,omitempty"`
	TodayRecord            []TodayRecord          `json:"todayRecord,omitempty"`
	Translations           []Translation          `json:"translations,omitempty"`
}

// ProblemSetQuestionList contains list of questions
type ProblemSetQuestionList struct {
	Typename  string     `json:"__typename,omitempty"`
	Questions []Question `json:"questions,omitempty"`
	HasMore   bool       `json:"hasMore,omitempty"`
	Total     int64      `json:"total,omitempty"`
}

// Question represents a single problem
type Question struct {
	TranslatedTitle    string     `json:"translatedTitle"`
	QuestionFrontendID string     `json:"questionFrontendId,omitempty"`
	QuestionTitleSlug  string     `json:"questionTitleSlug,omitempty"`
	TypeName           string     `json:"__typename,omitempty"`
	ACRate             float64    `json:"acRate,omitempty"`
	Difficulty         Difficulty `json:"difficulty,omitempty"`
	FreqBar            int64      `json:"freqBar,omitempty"`
	PaidOnly           bool       `json:"paidOnly,omitempty"`
	Status             Status     `json:"status,omitempty"`
	FrontendQuestionID string     `json:"frontendQuestionId,omitempty"`
	IsFavor            bool       `json:"isFavor,omitempty"`
	SolutionNum        int64      `json:"solutionNum,omitempty"`
	Title              string     `json:"title,omitempty"`
	TitleCN            string     `json:"titleCn,omitempty"`
	TitleSlug          string     `json:"titleSlug,omitempty"`
	TopicTags          []TopicTag `json:"topicTags,omitempty"`
	Extra              Extra      `json:"extra,omitempty"`
}

// Extra contains additional problem information
type Extra struct {
	CompanyTagNum    int64           `json:"companyTagNum"`
	HasVideoSolution bool            `json:"hasVideoSolution"`
	TopCompanyTags   []TopCompanyTag `json:"topCompanyTags"`
	Typename         ExtraTypename   `json:"__typename"`
}

// TopCompanyTag represents company tags
type TopCompanyTag struct {
	ImgURL   string `json:"imgUrl,omitempty"`
	Slug     Slug   `json:"slug,omitempty"`
	Typename string `json:"__typename,omitempty"`
}

// TopicTag represents topic categories
type TopicTag struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Slug           string `json:"slug,omitempty"`
	NameTranslated string `json:"nameTranslated,omitempty"`
	TranslatedName string `json:"translatedName,omitempty"`
	Typename       string `json:"__typename,omitempty"`
}

// Difficulty level
type Difficulty string

const (
	Easy   Difficulty = "EASY"
	Hard   Difficulty = "HARD"
	Medium Difficulty = "MEDIUM"
)

// Slug for company tags
type Slug string

const (
	Adobe        Slug = "adobe"
	Amazon       Slug = "amazon"
	Apple        Slug = "apple"
	Bytedance    Slug = "bytedance"
	Facebook     Slug = "facebook"
	GoldmanSachs Slug = "goldman-sachs"
	Google       Slug = "google"
	Intuit       Slug = "intuit"
	Jpmorgan     Slug = "jpmorgan"
	Linkedin     Slug = "linkedin"
	Microsoft    Slug = "microsoft"
	Tencent      Slug = "tencent"
)

// TopCompanyTagTypename type
type TopCompanyTagTypename string

const (
	CommonTagNode TopCompanyTagTypename = "CommonTagNode"
)

// ExtraTypename type
type ExtraTypename string

const (
	QuestionExtraInfoNode ExtraTypename = "QuestionExtraInfoNode"
)

// Status represents submission status
type Status string

const (
	AC         Status = "AC"
	NotStarted Status = "NOT_STARTED"
)

// QuestionTypename type
type QuestionTypename string

const (
	QuestionLightNode QuestionTypename = "QuestionLightNode"
)

// TodayRecord for daily records
type TodayRecord struct {
	Question       Question    `json:"question"`
	LastSubmission interface{} `json:"lastSubmission"`
	Date           string      `json:"date"`
	UserStatus     interface{} `json:"userStatus"`
	Typename       string      `json:"__typename"`
}

// DailyQuestionRecord for daily question tracking
type DailyQuestionRecord struct {
	Date       string                      `json:"date"`
	Question   DailyQuestionRecordQuestion `json:"question"`
	UserStatus string                      `json:"userStatus"`
	Typename   string                      `json:"__typename"`
}

// DailyQuestionRecordQuestion structure
type DailyQuestionRecordQuestion struct {
	QuestionID         string `json:"questionId"`
	QuestionFrontendID string `json:"questionFrontendId"`
	QuestionTitle      string `json:"questionTitle"`
	QuestionTitleSlug  string `json:"questionTitleSlug"`
	TranslatedTitle    string `json:"translatedTitle"`
	Title              string `json:"title"`
	TitleSlug          string `json:"titleSlug"`
	Typename           string `json:"__typename"`
}

// DataQuestion detailed question data
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
}

// CodeSnippet for different languages
type CodeSnippet struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
}

// Translation for multilingual support
type Translation struct {
	Title       string `json:"title"`
	TitleCn     string `json:"titleCn"`
	Content     string `json:"content"`
	ContentCn   string `json:"contentCn"`
	Description string `json:"description"`
	DescCn      string `json:"descCn"`
}
