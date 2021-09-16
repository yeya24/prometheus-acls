package config

import (
	"github.com/prometheus/prometheus/pkg/labels"
	"regexp"
)

type (
	// OidcRole is a string to identify a OIDC role from the OAUTH provider
	OidcRole string

	// MetricName is a string to identify a Prometheus metric
	MetricName string

	// NamedACL hold the LabelMatchers for a specific MetricName
	NamedACL map[MetricName][]*labels.Matcher

	// RegexACL holds the LabelMatchers for all MetricNames that match Regexp
	RegexACL struct {
		Regexp        *regexp.Regexp
		LabelMatchers []*labels.Matcher
	}

	// ACL holds the parsed Named and Regex metricName to LabelMatchers
	ACL struct {
		Named NamedACL
		Regex []RegexACL
	}

	// ACLMap is used to look up OidcRole for its configures ACL
	ACLMap map[OidcRole]*ACL
)