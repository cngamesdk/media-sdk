package model

// ReportReq 报表请求
type ReportReq struct {
	AdvertiserID string   `json:"advertiser_id"`
	StartDate    string   `json:"start_date"`
	EndDate      string   `json:"end_date"`
	Level        string   `json:"level"` // ACCOUNT, CAMPAIGN, UNIT, CREATIVE
	GroupBy      []string `json:"group_by,omitempty"`
	Page         int      `json:"page"`
	PageSize     int      `json:"page_size"`
}

// ReportResp 报表响应
type ReportResp struct {
	List     []*ReportData `json:"list"`
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	Summary  *ReportData   `json:"summary,omitempty"`
}

// ReportData 报表数据
type ReportData struct {
	Date           string  `json:"date"`
	CampaignID     string  `json:"campaign_id,omitempty"`
	UnitID         string  `json:"unit_id,omitempty"`
	CreativeID     string  `json:"creative_id,omitempty"`
	Impressions    int64   `json:"impressions"`
	Clicks         int64   `json:"clicks"`
	Cost           float64 `json:"cost"`
	Ctr            float64 `json:"ctr"`
	Cpm            float64 `json:"cpm"`
	Cpc            float64 `json:"cpc"`
	Conversions    int64   `json:"conversions"`
	ConversionRate float64 `json:"conversion_rate"`
	ConversionCost float64 `json:"conversion_cost"`
}
