package duckdb

import (
	"context"

	"github.com/medama-io/medama/model"
)

// GetWebsiteReferrersSummary returns a summary of the referrers for the given hostname.
func (c *Client) GetWebsiteReferrersSummary(ctx context.Context, hostname string) ([]*model.StatsReferrerSummary, error) {
	var referrers []*model.StatsReferrerSummary

	// Array of referrer summaries
	//
	// Referrer is the referrer of the page.
	//
	// Uniques is the number of uniques for the referrer.
	//
	// UniquePercentage is the percentage the referrer contributes to the total uniques.
	query := `
		SELECT
			referrer,
			COUNT(CASE WHEN is_unique = true THEN 1 END) AS uniques,
			ROUND(COUNT(CASE WHEN is_unique = true THEN 1 END) * 100.0 / (SELECT COUNT(CASE WHEN is_unique = true THEN 1 END) FROM views WHERE hostname = ?), 2) AS unique_percentage
		FROM views
		WHERE hostname = ?
		GROUP BY referrer
		ORDER BY uniques DESC;`

	err := c.SelectContext(ctx, &referrers, query, hostname, hostname)
	if err != nil {
		return nil, err
	}

	return referrers, nil
}

// GetWebsiteReferrers returns the referrers for the given hostname.
func (c *Client) GetWebsiteReferrers(ctx context.Context, hostname string) ([]*model.StatsReferrers, error) {
	var referrers []*model.StatsReferrers

	// Array of referrers
	//
	// Referrer is the referrer of the page.
	//
	// Uniques is the number of uniques for the referrer.
	//
	// UniquePercentage is the percentage the referrer contributes to the total uniques.
	//
	// Bounces is the number of bounces for the referrer.
	//
	// Duration is the median duration the user spent on the page in milliseconds.
	query := `
		SELECT
			referrer,
			COUNT(CASE WHEN is_unique = true THEN 1 END) AS uniques,
			ROUND(COUNT(CASE WHEN is_unique = true THEN 1 END) * 100.0 / (SELECT COUNT(CASE WHEN is_unique = true THEN 1 END) FROM views WHERE hostname = ?), 2) AS unique_percentage,
			COUNT(CASE WHEN is_unique = true AND duration_ms < 5000 THEN 1 END) AS bounces,
			CAST(ifnull(median(duration_ms), 0) AS INTEGER) AS duration
		FROM views
		WHERE hostname = ?
		GROUP BY referrer
		ORDER BY uniques DESC;`

	err := c.SelectContext(ctx, &referrers, query, hostname, hostname)
	if err != nil {
		return nil, err
	}

	return referrers, nil
}
