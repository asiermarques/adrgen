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
	ID int
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

func CreateADRFilenameFromFilenameString(filename string) (ADRFilename,error) {
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
}

type ADRRepository interface {
	FindAll() ([]ADR, error)
	FindById(id int) (ADR, error)
	GetLastId() int
}

type ADRWriter interface {
	Persist(adr ADR) error
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
		return ADR{}, fmt.Errorf(
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
		ID: adr.ID,
		Filename: adr.Filename,
		Status: newStatus,
		Content: re.ReplaceAllString(adr.Content, "Status: " + newStatus),
	}, nil
}

func (manager privateADRStatusManager) ValidateStatus(targetStatus string) bool {
	if len(manager.configuration.Statuses)<1 {
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