
package pipe

// DomainGrouper groups the records by domain.
// It keeps the other fields intact.
// For example: It returns the page field as well.
// Exercise: Write a solution that removes the unnecessary data.
func DomainGrouper(r Record) string {
	return r.domain
}

// Page groups records by page.
func Page(r Record) string {
	return r.domain + r.page
}
