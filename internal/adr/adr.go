package adr

import (
	"fmt"
	"github.com/asiermarques/adrgen/internal/config"
	"github.com/asiermarques/adrgen/internal/template"
	"regexp"
	"strconv"
	"strings"

	"github.com/gosimple/slug"
)

// Filename is the domain value object for the ADR filename property
type Filename interface {
	Value() string
}

type privateFilename struct {
	ID    int
	value string
}

func (f privateFilename) Value() string {
	return f.value
}

// CreateFilename creates the Filename value object
func CreateFilename(id int, title string, idDigits int, extension string) Filename {
	if idDigits < 1 {
		return privateFilename{value: fmt.Sprintf("%d-%s%s", id, slug.Make(title), extension)}
	}
	return privateFilename{
		value: fmt.Sprintf("%0"+strconv.Itoa(idDigits)+"d-%s%s", id, slug.Make(title), extension),
	}
}

// CreateFilenameFromFilenameString creates the Filename value object from a filename string
func CreateFilenameFromFilenameString(filename string) (Filename, error) {
	if !ValidateFilename(filename) {
		return &privateFilename{}, fmt.Errorf("filename not valid %s", filename)
	}

	return &privateFilename{value: filename}, nil
}

// ValidateFilename validates if a string is a correct filename for an ADR File
func ValidateFilename(name string) bool {
	pattern := regexp.MustCompile(`(?mi)^\d+-.+\.(md|adoc)`)
	return pattern.MatchString(name)
}

// ADR is the domain entity representing the ADR file
type ADR interface {
	ID() int
	Filename() Filename
	Status() string
	Title() string
	Date() string
	Content() string
}

type privateADR struct {
	id       int
	filename Filename
	content  string
}

func (a privateADR) ID() int {
	return a.id
}

func (a privateADR) Filename() Filename {
	return a.filename
}

func (a privateADR) Status() string {
	status, _ := a.getStatusFromContent()
	return status
}

func (a privateADR) Title() string {
	title, _ := a.getTitleFromContent()
	return title
}

func (a privateADR) Date() string {
	date, _ := a.getDateFromContent()
	return date
}

func (a privateADR) Content() string {
	return a.content
}

func (a privateADR) getTitleFromContent() (string, error) {
	if a.content == "" {
		return "", fmt.Errorf("ADR content not present")
	}

	re := regexp.MustCompile(`(?mi)^[#=] (.+)$`)
	if !re.MatchString(a.content) {
		return "", fmt.Errorf("title not present in ADR Content")
	}

	matches := re.FindStringSubmatch(a.content)
	if len(matches) < 2 || matches[1] == "" {
		return "", fmt.Errorf("could not possible extracting the title from ADR Content")
	}

	return matches[1], nil
}

func (a privateADR) getDateFromContent() (string, error) {
	if a.content == "" {
		return "", fmt.Errorf("ADR content not present")
	}

	re := regexp.MustCompile(`(?mi)^Date: (.+)$`)
	if !re.MatchString(a.content) {
		return "", fmt.Errorf("date not present in ADR Content")
	}

	matches := re.FindStringSubmatch(a.content)
	if len(matches) < 2 || matches[1] == "" {
		return "", fmt.Errorf("could not possible extracting the date from ADR Content")
	}

	return matches[1], nil
}

func (a privateADR) getStatusFromContent() (string, error) {
	if a.content == "" {
		return "", fmt.Errorf("ADR content not present")
	}

	re := regexp.MustCompile(`(?mi)^(##|==) Status\n\n?(.+)$`)
	if !re.MatchString(a.content) {
		return "", fmt.Errorf("status not present in ADR Content")
	}

	matches := re.FindStringSubmatch(a.content)
	if len(matches) < 3 || matches[2] == "" {
		return "", fmt.Errorf("could not possible extracting the status from ADR Content")
	}

	return matches[2], nil
}

// CreateADR creates the ADR entity and validates its required properties
func CreateADR(id int, content string, filename Filename) (ADR, error) {

	if id < 1 {
		return nil, fmt.Errorf("id parameter is required")
	}

	if content == "" {
		return nil, fmt.Errorf("content is required")
	}

	if filename == nil {
		return nil, fmt.Errorf("filename is required")
	}

	return privateADR{id, filename, content}, nil
}

// Repository is the repository for the ADR domain entity
type Repository interface {
	FindAll() ([]ADR, error)
	Query(filterParams map[string][]string) ([]ADR, error)
	FindById(id int) (ADR, error)
	GetLastId() int
}

// Writer service that persist the ADR entity
type Writer interface {
	Persist(adr ADR) error
}

// RelationsManager service that manage the relation links between ADR files
type RelationsManager interface {
	AddRelation(adr ADR, targetADR ADR, relation string) (ADR, ADR, error)
	RelationIsValid(relation string) bool
}

type relation struct {
	mainTitle    string
	targetTitle  string
	targetStatus string
}

type privateRelationsManager struct {
	relations       map[string]relation
	templateService template.Service
	statusManager   StatusManager
}

func (m privateRelationsManager) RelationIsValid(relation string) bool {
	_, result := m.relations[relation]
	return result
}

func (m privateRelationsManager) AddRelation(
	adr ADR,
	targetADR ADR,
	relation string,
) (ADR, ADR, error) {
	if !m.RelationIsValid(relation) {
		return adr, targetADR, fmt.Errorf("relation %s is not valid", relation)
	}

	re := regexp.MustCompile(`(?mi)^## Status\n\n?(.+)$`)
	if !re.MatchString(adr.Content()) {
		return adr, targetADR, fmt.Errorf("ADR content have not a status field")
	}
	if !re.MatchString(targetADR.Content()) {
		return adr, targetADR, fmt.Errorf("target ADR content have not a status field")
	}

	targetADR, _ = m.statusManager.ChangeStatus(targetADR, m.relations[relation].targetStatus)

	matches := re.FindStringSubmatch(targetADR.Content())
	if len(matches) < 1 {
		return adr, targetADR, fmt.Errorf("target ADR content have not a status field match")
	}

	targetADRContent := strings.Replace(
		targetADR.Content(),
		matches[0],
		matches[0]+"\n\n"+m.templateService.RenderRelationLink(
			adr.Title(),
			adr.Filename().Value(),
			m.relations[relation].targetTitle,
		),
		1,
	)

	matches = re.FindStringSubmatch(adr.Content())
	adrContent := strings.Replace(
		adr.Content(),
		matches[0],
		matches[0]+"\n\n"+m.templateService.RenderRelationLink(
			targetADR.Title(),
			targetADR.Filename().Value(),
			m.relations[relation].mainTitle,
		),
		1,
	)

	newAdr, _ := CreateADR(adr.ID(), adrContent, adr.Filename())
	newTargetAdr, _ := CreateADR(targetADR.ID(), targetADRContent, targetADR.Filename())

	return newAdr, newTargetAdr, nil
}

// CreateRelationsManager creates the RelationsManager service with its dependencies
func CreateRelationsManager(service template.Service, manager StatusManager) RelationsManager {
	relations := make(map[string]relation)
	relations["supersede"] = relation{
		mainTitle:    "Supersedes",
		targetTitle:  "Superseded by",
		targetStatus: "superseded",
	}
	relations["amend"] = relation{
		mainTitle:    "Amends",
		targetTitle:  "Amended by",
		targetStatus: "amended",
	}

	return privateRelationsManager{
		relations,
		service,
		manager,
	}
}

// StatusManager is the service that manages the status change in a ADR entity
type StatusManager interface {
	ChangeStatus(adr ADR, newStatus string) (ADR, error)
	ValidateStatus(targetStatus string) bool
}

type privateStatusManager struct {
	configuration config.Config
}

func (manager privateStatusManager) ChangeStatus(adr ADR, newStatus string) (ADR, error) {
	if !manager.ValidateStatus(newStatus) {
		return adr, fmt.Errorf(
			"status %s not allowed, please use one of these %s",
			newStatus,
			strings.Join(manager.configuration.Statuses, ", "),
		)
	}

	re := regexp.MustCompile(`(?mi)^## Status\n\n?(.+)$`)
	if !re.MatchString(adr.Content()) {
		return nil, fmt.Errorf("ADR content have not a status field")
	}

	return CreateADR(
		adr.ID(),
		re.ReplaceAllString(adr.Content(), fmt.Sprintf("## Status\n\n%s", newStatus)),
		adr.Filename(),
	)
}

func (manager privateStatusManager) ValidateStatus(targetStatus string) bool {
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

// CreateStatusManager creates the StatusManager service with its dependencies
func CreateStatusManager(configuration config.Config) StatusManager {
	return privateStatusManager{
		configuration: configuration,
	}
}
