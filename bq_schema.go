package main

import (
	bigquery "google.golang.org/api/bigquery/v2"
)

func generateBQSchema() *bigquery.TableSchema {
	return &bigquery.TableSchema{
		Fields: []*bigquery.TableFieldSchema{
			{Mode: "NULLABLE", Name: "id", Type: "STRING"},
			{Mode: "NULLABLE", Name: "type", Type: "STRING"},
			{Mode: "NULLABLE", Name: "create_time", Type: "STRING"},
			{Mode: "NULLABLE", Name: "archive_url", Type: "STRING"},
			{Mode: "NULLABLE", Name: "long_archive_url", Type: "STRING"},
			{Mode: "NULLABLE", Name: "status", Type: "STRING"},
			{Mode: "NULLABLE", Name: "emails_sent", Type: "INTEGER"},
			{Mode: "NULLABLE", Name: "send_time", Type: "STRING"},
			{Mode: "NULLABLE", Name: "content_type", Type: "STRING"},
			{Mode: "NULLABLE", Name: "recipients", Type: "RECORD", Fields: []*bigquery.TableFieldSchema{
				{Mode: "NULLABLE", Name: "list_id", Type: "STRING"},
				{Mode: "NULLABLE", Name: "list_name", Type: "STRING"},
				{Mode: "NULLABLE", Name: "segment_text", Type: "STRING"},
				{Mode: "NULLABLE", Name: "recipient_count", Type: "INTEGER"},
			}},
			{Mode: "NULLABLE", Name: "settings", Type: "RECORD", Fields: []*bigquery.TableFieldSchema{
				{Mode: "NULLABLE", Name: "subject_line", Type: "STRING"},
				{Mode: "NULLABLE", Name: "title", Type: "STRING"},
				{Mode: "NULLABLE", Name: "from_name", Type: "STRING"},
				{Mode: "NULLABLE", Name: "reply_to", Type: "STRING"},
			}},
			{Mode: "NULLABLE", Name: "variate_settings", Type: "RECORD", Fields: []*bigquery.TableFieldSchema{
				{Mode: "NULLABLE", Name: "winning_combination_id", Type: "STRING"},
				{Mode: "NULLABLE", Name: "winning_campaign_id", Type: "STRING"},
				{Mode: "NULLABLE", Name: "winner_criteria", Type: "STRING"},
				{Mode: "NULLABLE", Name: "wait_time", Type: "INTEGER"},
				{Mode: "NULLABLE", Name: "test_size", Type: "INTEGER"},
				{Mode: "REPEATED", Name: "subject_lines", Type: "STRING"},
				{Mode: "REPEATED", Name: "send_times", Type: "STRING"},
				{Mode: "REPEATED", Name: "reply_to_addresses", Type: "STRING"},
				{Mode: "REPEATED", Name: "contents", Type: "STRING"},
				{Mode: "REPEATED", Name: "combinations", Type: "RECORD", Fields: []*bigquery.TableFieldSchema{
					{Mode: "NULLABLE", Name: "id", Type: "STRING"},
					{Mode: "NULLABLE", Name: "subject_line", Type: "INTEGER"},
					{Mode: "NULLABLE", Name: "send_time", Type: "INTEGER"},
					{Mode: "NULLABLE", Name: "from_name", Type: "INTEGER"},
					{Mode: "NULLABLE", Name: "reply_to", Type: "INTEGER"},
					{Mode: "NULLABLE", Name: "content_description", Type: "INTEGER"},
					{Mode: "NULLABLE", Name: "recipients", Type: "INTEGER"},
				}},
			}},
			{Mode: "NULLABLE", Name: "report_summary", Type: "RECORD", Fields: []*bigquery.TableFieldSchema{
				{Mode: "NULLABLE", Name: "opens", Type: "INTEGER"},
				{Mode: "NULLABLE", Name: "unique_opens", Type: "INTEGER"},
				{Mode: "NULLABLE", Name: "open_rate", Type: "FLOAT"},
				{Mode: "NULLABLE", Name: "clicks", Type: "INTEGER"},
				{Mode: "NULLABLE", Name: "subscriber_clicks", Type: "INTEGER"},
				{Mode: "NULLABLE", Name: "click_rate", Type: "FLOAT"},
			}},
			{Mode: "NULLABLE", Name: "delivery_status", Type: "RECORD", Fields: []*bigquery.TableFieldSchema{
				{Mode: "NULLABLE", Name: "enabled", Type: "BOOLEAN"},
				{Mode: "NULLABLE", Name: "can_cancel", Type: "BOOLEAN"},
				{Mode: "NULLABLE", Name: "status", Type: "STRING"},
				{Mode: "NULLABLE", Name: "emails_sent", Type: "INTEGER"},
				{Mode: "NULLABLE", Name: "emails_canceled", Type: "INTEGER"},
			}},
		},
	}
}
