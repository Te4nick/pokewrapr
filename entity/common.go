package entity

type APIResource struct {
	URL string `json:"url"`
}

type Description struct {
	Description string             `json:"description"`
	Language    []NamedAPIResource `json:"language"`
}

type Effect struct {
	Effect   string             `json:"effect"`
	Language []NamedAPIResource `json:"language"`
}

type Encounter struct {
	MinLevel        int              `json:"min_level"`
	MaxLevel        int              `json:"max_level"`
	Chance          int              `json:"chance"`
	ConditionValues []any            `json:"condition_values"`
	Method          NamedAPIResource `json:"method"`
}

type FlavorText struct {
	FlavorText string           `json:"flavor_text"`
	Language   NamedAPIResource `json:"language"`
	Version    NamedAPIResource `json:"version"`
}

type GenerationGameIndex struct {
	GameIndex  int              `json:"game_index"`
	Generation NamedAPIResource `json:"generation"`
}

type MachineVersionDetail struct {
	Machine      APIResource      `json:"machine"`
	VersionGroup NamedAPIResource `json:"version_group"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"language"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VerboseEffect struct {
	Effect      string           `json:"effect"`
	ShortEffect string           `json:"short_effect"`
	Language    NamedAPIResource `json:"language"`
}

type VersionEncounterDetail struct {
	EncounterDetails []Encounter      `json:"encounter_details"`
	MaxChance        int              `json:"max_chance"`
	Version          NamedAPIResource `json:"version"`
}

type VersionGameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIResource `json:"version"`
}

type VersionGroupFlavorText struct {
	Text         string           `json:"text"`
	Language     NamedAPIResource `json:"language"`
	VersionGroup NamedAPIResource `json:"version_group"`
}
