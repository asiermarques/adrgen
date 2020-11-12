package domain

import (
	"fmt"
	"github.com/gosimple/slug"
	"regexp"
	"strconv"
	"strings"
)

type ADRFilename interface {
	Value() string
}

type privateADRFilename struct {
	ID    int
	value string
}

func (f privateADRFilename) Value() string {
	return f.value
}

func CreateADRFilename(id int, title string, idDigits int) ADRFilename {
	if idDigits < 1 {
		return privateADRFilename{value: fmt.Sprintf("%d-%s.md", id, slug.Make(title))}
	}
	return privateADRFilename{
		value: fmt.Sprintf("%0"+strconv.Itoa(idDigits)+"d-%s.md", id, slug.Make(title)),
	}
}

func CreateADRFilenameFromFilenameString(filename string) (ADRFilename, error) {
	if !ValidateADRFilename(filename) {
		return &privateADRFilename{}, fmt.Errorf("filename not valid %s", filename)
	}

	return &privateADRFilename{value: filename}, nil
}

func ValidateADRFilename(name string) bool {
	pattern := regexp.MustCompile(`(?mi)^\d+-.+\.md`)
	return pattern.MatchString(name)
}

type ADR struct {
	ID       int
	Filename ADRFilename
	Status   string
	Content  string
	SupersededByID int
	SupersedesID int
}

func (adr ADR) getTitleFromContent() (string, error) {
	if adr.Content == "" {
		return "", fmt.Errorf("ADR content not present")
	}

	re := regexp.MustCompile(`(?mi)^# (.+)$`)
	if !re.MatchString(adr.Content) {
		return "", fmt.Errorf("title not present in ADR Content")
	}

	matches := re.FindStringSubmatch(adr.Content)
	if len(matches) < 2 || matches[1] == "" {
		return "", fmt.Errorf("could not possible extracting the title from ADR Content")
	}

	return matches[1], nil
}


type ADRRepository interface {
	FindAll() ([]ADR, error)
	FindById(id int) (ADR, error)
	GetLastId() int
}

type ADRWriter interface {
	Persist(adr ADR) error
}

type RelationsManager interface {
	Supersede(adr ADR, targetADR ADR)  (ADR, ADR, error)
}

type privateRelationsManager struct {
	templateService TemplateService
	statusManager ADRStatusManager
}

func (m privateRelationsManager) Supersede(adr ADR, targetADR ADR) (ADR, ADR, error) {
	re := regexp.MustCompile(`(?mi)^Status:\s?(.+)$`)
	if !re.MatchString(adr.Content) {
		return adr, targetADR, fmt.Errorf("ADR content have not a status field")
	}
	if !re.MatchString(targetADR.Content) {
		return adr, targetADR, fmt.Errorf("target ADR content have not a status field")
	}

	targetADR, _ = m.statusManager.ChangeStatus(targetADR, "superseded")

	matches := re.FindStringSubmatch(targetADR.Content)
	targetADR.Content = strings.Replace(targetADR.Content, matches[0], matches[0] + "\n\n" + m.templateService.CreateSupersededByLink(adr), 1)

	matches = re.FindStringSubmatch(adr.Content)
	adr.Content = strings.Replace(adr.Content, matches[0], matches[0] + "\n\n" + m.templateService.CreateSupersedesLink(targetADR), 1)

	return adr, targetADR, nil
}

func CreateRelationsManager(service TemplateService, manager ADRStatusManager) RelationsManager {
	return privateRelationsManager{service, manager}
}

type ADRStatusManager interface {
	ChangeStatus(adr ADR, newStatus string) (ADR, error)
	ValidateStatus(targetStatus string) bool
}

type privateADRStatusManager struct {
	configuration Config
}

func (manager privateADRStatusManager) ChangeStatus(adr ADR, newStatus string) (ADR, error) {
	if !manager.ValidateStatus(newStatus) {
		return adr, fmt.Errorf(
			"status %s not allowed, please use one of these %s",
			newStatus,
			strings.Join(manager.configuration.Statuses, ", "),
		)
	}

	re := regexp.MustCompile(`(?mi)^Status:\s?(.+)$`)
	if !re.MatchString(adr.Content) {
		return ADR{}, fmt.Errorf("ADR content have not a status field")
	}

	return ADR{
		ID:       adr.ID,
		Filename: adr.Filename,
		Status:   newStatus,
		Content:  re.ReplaceAllString(adr.Content, "Status: "+newStatus),
	}, nil
}

func (manager privateADRStatusManager) ValidateStatus(targetStatus string) bool {
	if len(manager.configuration.Statuses) < 1 {
		return true
	}

	for _, status := range manager.configuration.Statuses {
		if status == targetStatus {
			return true
		}
	}
	return false
}

func CreateADRStatusManager(configuration Config) ADRStatusManager {
	return privateADRStatusManager{
		configuration: configuration,
	}
}
