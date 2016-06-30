package main

type MCCampaign struct {
	Id              string                   `json:"id"`
	Type            string                   `json:"type"`
	CreateTime      string                   `json:"create_time"`
	ArchiveUrl      string                   `json:"archive_url"`
	LongArchiveUrl  string                   `json:"long_archive_url"`
	Status          string                   `json:"status"`
	EmailsSent      int32                    `json:"emails_sent"`
	SendTime        string                   `json:"send_time"`
	ContentType     string                   `json:"content_type"`
	Recipients      MCCampaignRecipients     `json:"recipients"`
	Settings        MCCampaignSettings       `json:"settings"`
	VariateSettings MCCampaignVariateSetting `json:"variate_settings"`
	ReportSummary   MCCampaignReportSummary  `json:"report_summary"`
	DeliveryStatus  MCCampaignDeliveryStatus `json:"delivery_status"`
}

type MCCampaignRecipients struct {
	ListId         string `json:"list_id"`
	ListName       string `json:"list_name"`
	SegmentText    string `json:"segment_text"`
	RecipientCount int32  `json:"recipient_count"`
}

type MCCampaignSettings struct {
	SubjectLine string `json:"subject_line"`
	Title       string `json:"title"`
	FromName    string `json:"from_name"`
	ReplyTo     string `json:"reply_to"`
}

type MCCampaignVariateSetting struct {
	WinningCombinationId string                                `json:"winning_combination_id"`
	WinningCampaignId    string                                `json:"winning_campaign_id"`
	WinnerCriteria       string                                `json:"winner_criteria"`
	WaitTime             int32                                 `json:"wait_time"`
	TestSize             int32                                 `json:"test_size"`
	SubjectLines         []string                              `json:"subject_lines"`
	SendTimes            []string                              `json:"send_times"`
	ReplyToAddresses     []string                              `json:"reply_to_addresses"`
	Contents             []string                              `json:"contents"`
	Combinations         []MCCampaignVariateSettingCombination `json:"combinations"`
}

type MCCampaignVariateSettingCombination struct {
	Id                 string `json:"id"`
	SubjectLine        int32  `json:"subject_line"`
	SendTime           int32  `json:"send_time"`
	FromName           int32  `json:"from_name"`
	ReplyTo            int32  `json:"reply_to"`
	ContentDescription int32  `json:"content_description"`
	Recipients         int32  `json:"recipients"`
}

type MCCampaignReportSummary struct {
	Opens            int32   `json:"opens"`
	UniqueOpens      int32   `json:"unique_opens"`
	OpenRate         float32 `json:"open_rate"`
	Clicks           int32   `json:"clicks"`
	SubscriberClicks int32   `json:"subscriber_clicks"`
	ClickRate        float32 `json:"click_rate"`
}

type MCCampaignDeliveryStatus struct {
	Enabled        bool   `json:"enabled"`
	CanCancel      bool   `json:"can_cancel"`
	Status         string `json:"status"`
	EmailsSent     int32  `json:"emails_sent"`
	EmailsCanceled int32  `json:"emails_canceled"`
}
