package nasdaqProvider

type HistoricalPriceDto struct {
	Data    SymbolData `json:"data"`
	Message string     `json:"message"`
	Status  Status     `json:"status"`
}

type Headers struct {
	Date   string `json:"date"`
	Close  string `json:"close"`
	Volume string `json:"volume"`
	Open   string `json:"open"`
	High   string `json:"high"`
	Low    string `json:"low"`
}

type Rows struct {
	Date   string `json:"date"`
	Close  string `json:"close"`
	Volume string `json:"volume"`
	Open   string `json:"open"`
	High   string `json:"high"`
	Low    string `json:"low"`
}

type TradesTable struct {
	Headers Headers `json:"headers"`
	Rows    []Rows  `json:"rows"`
}

type SymbolData struct {
	Symbol       string      `json:"symbol"`
	TotalRecords int         `json:"totalRecords"`
	TradesTable  TradesTable `json:"tradesTable"`
}

type Status struct {
	RCode            int         `json:"rCode"`
	BCodeMessage     []BCodeMessage `json:"bCodeMessage"`
	DeveloperMessage interface{} `json:"developerMessage"`
}

type BCodeMessage struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"errorMessage"`
}

type CompanyInfoDto struct {
	Data    CompanyData `json:"data"`
	Message string      `json:"message"`
	Status  Status      `json:"status"`
}

type KeyValue struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type CompanyData struct {
	ModuleTitle        KeyValue `json:"ModuleTitle"`
	CompanyName        KeyValue `json:"CompanyName"`
	Symbol             KeyValue `json:"Symbol"`
	Address            KeyValue `json:"Address"`
	Phone              KeyValue `json:"Phone"`
	Industry           KeyValue `json:"Industry"`
	Sector             KeyValue `json:"Sector"`
	Region             KeyValue `json:"Region"`
	CompanyDescription KeyValue `json:"CompanyDescription"`
}

type InsiderActivityDto struct {
	Data struct {
		TransactionTable struct {
			Rows []struct {
				Insider         string `json:"insider"`
				Relation        string `json:"relation"`
				LastDate        string `json:"lastDate"`
				TransactionType string `json:"transactionType"`
				OwnType         string `json:"ownType"`
				SharesTraded    string `json:"sharesTraded"`
				LastPrice       string `json:"lastPrice"`
				SharesHeld      string `json:"sharesHeld"`
			} `json:"rows"`
		} `json:"transactionTable"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Status  Status `json:"status"`
}

// ===========================
// FTP Data Transfer Objects
// ===========================

type nasdaqListedSymbolsFtpDto struct {
	Symbol          string
	SecurityName    string
	MarketCategory  string
	TestIssue       string
	FinancialStatus string
	RoundLotSize    string
	Etf             string
	NextShares      string
}

type nasdaqUnlistedSymbolsFtpDto struct {
	ActSymbol    string
	SecurityName string
	Exchange     string
	CqsSymbol    string
	Etf          string
	RoundLotSize string
	TestIssue    string
	NasdaqSymbol string
}
