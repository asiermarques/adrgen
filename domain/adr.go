package domain

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gosimple/slug"
)

// ADRFilename is the domain value object for the ADR filename property
//
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

// CreateADRFilename creates the ADRFilename value object
//
func CreateADRFilename(id int, title string, idDigits int) ADRFilename {
	if idDigits < 1 {
		return privateADRFilename{value: fmt.Sprintf("%d-%s.md", id, slug.Make(title))}
	}
	return privateADRFilename{
		value: fmt.Sprintf("%0"+strconv.Itoa(idDigits)+"d-%s.md", id, slug.Make(title)),
	}
}

// CreateADRFilenameFromFilenameString creates the ADRFilename value object from a filename string
//
func CreateADRFilenameFromFilenameString(filename string) (ADRFilename, error) {
	if !ValidateADRFilename(filename) {
		return &privateADRFilename{}, fmt.Errorf("filename not valid %s", filename)
	}

	return &privateADRFilename{value: filename}, nil
}

// ValidateADRFilename validates if a string is a correct filename for an ADR File
//
func ValidateADRFilename(name string) bool {
	pattern := regexp.MustCompile(`(?mi)^\d+-.+\.md`)
	return pattern.MatchString(name)
}

// ADR is the domain entity representing the ADR file
//
type ADR interface {
	ID() int
	Filename() ADRFilename
	Status() string
	Title() string
	Date() string
	Content() string
}

type privateADR struct {
	id       int
	filename ADRFilename
	content  string
}

func (a privateADR) ID() int {
	return a.id
}

func (a privateADR) Filename() ADRFilename {
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

	re := regexp.MustCompile(`(?mi)^# (.+)$`)
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
		return "", fmt.Errorf("title not present in ADR Content")
	}

	matches := re.FindStringSubmatch(a.content)
	if len(matches) < 2 || matches[1] == "" {
		return "", fmt.Errorf("could not possible extracting the title from ADR Content")
	}

	return matches[1], nil
}

func (a privateADR) getStatusFromContent() (string, error) {
	if a.content == "" {
		return "", fmt.Errorf("ADR content not present")
	}

	re := regexp.MustCompile(`(?mi)^## Status\n\n?(.+)$`)
	if !re.MatchString(a.content) {
		return "", fmt.Errorf("status not present in ADR Content")
	}

	matches := re.FindStringSubmatch(a.content)
	if len(matches) < 2 || matches[1] == "" {
		return "", fmt.Errorf("could not possible extracting the status from ADR Content")
	}

	return matches[1], nil
}

// CreateADR creates the ADR entity and validates its required properties
//
func CreateADR(id int, content string, filename ADRFilename) (ADR, error) {

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

// ADRRepository is the repository for the ADR domain entity
//
type ADRRepository interface {
	FindAll() ([]ADR, error)
	FindById(id int) (ADR, error)
	GetLastId() int
}

// ADRWriter service that persist the ADR entity
//
type ADRWriter interface {
	Persist(adr ADR) error
}

// RelationsManager service that manage the relation links between ADR files
//
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
	templateService TemplateService
	statusManager   ADRStatusManager
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
			adr,
			m.relations[relation].targetTitle,
		),
		1,
	)

	matches = re.FindStringSubmatch(adr.Content())
	adrContent := strings.Replace(
		adr.Content(),
		matches[0],
		matches[0]+"\n\n"+m.templateService.RenderRelationLink(
			targetADR,
			m.relations[relation].mainTitle,
		),
		1,
	)

	newAdr, _ := CreateADR(adr.ID(), adrContent, adr.Filename())
	newTargetAdr, _ := CreateADR(targetADR.ID(), targetADRContent, targetADR.Filename())

	return newAdr, newTargetAdr, nil
}

// CreateRelationsManager creates the RelationsManager service with its dependencies
//
func CreateRelationsManager(service TemplateService, manager ADRStatusManager) RelationsManager {
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

// ADRStatusManager is the service that manages the status change in a ADR entity
//
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

// CreateADRStatusManager creates the ADRStatusManager service with its dependencies
//
func CreateADRStatusManager(configuration Config) ADRStatusManager {
	return privateADRStatusManager{
		configuration: configuration,
	}
}
