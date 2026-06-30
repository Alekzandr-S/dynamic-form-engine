package domain

type TemplateStatus string

const (
	DrafTemplate      TemplateStatus = "DRAFT"
	PublishedTemplate TemplateStatus = "PUBLISHED"
	ArchivedTemplate  TemplateStatus = "ARCHIVED"
)

func (s TemplateStatus) IsValid() bool {
	switch s {
	case DrafTemplate, PublishedTemplate, ArchivedTemplate:
		return true
	default:
		return false
	}
}
