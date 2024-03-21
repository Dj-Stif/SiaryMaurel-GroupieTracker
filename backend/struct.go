package backend

type Armures []struct {
	Rank   string     `json:"rank"`
	Name   string     `json:"name"` //nom d'armure
	Pieces []struct { //liste des éléments d'une armure
		Type    string `json:"type"` // head, chest..
		Rank    string `json:"rank"`
		Rarity  int    `json:"rarity"`
		Defense struct {
			Base      int `json:"base"`
			Max       int `json:"max"`
			Augmented int `json:"augmented"`
		} `json:"defense"`
		Resistances struct {
			Fire    int `json:"fire"`
			Water   int `json:"water"`
			Ice     int `json:"ice"`
			Thunder int `json:"thunder"`
			Dragon  int `json:"dragon"`
		} `json:"resistances"`
		Name   string `json:"name"` // nom de la pièce d'armure
		Assets struct {
			ImageMale   string `json:"imageMale"`
			ImageFemale string `json:"imageFemale"`
		} `json:"assets"`
	} `json:"pieces"`
}

type Weapons []struct {
	Type       string     `json:"type"`       //type d'arme
	Rarity     int        `json:"rarity"`     //rareté 1 à 12
	DamageType string     `json:"damageType"` //type de dégats
	Name       string     `json:"name"`       //nom de l'arme
	Durability []struct { //durabilité de l'arme
		Red    int `json:"red"`
		Orange int `json:"orange"`
		Yellow int `json:"yellow"`
		Green  int `json:"green"`
		Blue   int `json:"blue"`
		White  int `json:"white"`
		Purple int `json:"purple"`
	} `json:"durability,omitempty"`
	Assets struct {
		Icon  string `json:"icon"`
		Image string `json:"image"`
	} `json:"assets"`
}

type Monsters []struct {
	Type        string `json:"type"`
	Species     string `json:"species"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Weaknesses  []struct {
		Element string `json:"element"`
		Stars   int    `json:"stars"`
	} `json:"weaknesses"`
}
